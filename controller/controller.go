package controller

import (
	"buscadorWeb/model"
	"encoding/json"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"

	"fmt"
	"html/template"
	"log"

	"net/http"

	"github.com/gorilla/securecookie"

	"gopkg.in/mgo.v2/bson"
)

const SEARCH = "search.html"
const HOME = "home.html"
const FILE = "modules/file.html"
const IMAGE = "modules/image.html"
const VIDEO = "modules/video.html"
const LOGIN = "modules/login.html"
const REGISTER = "modules/register.html"
const FAV = "modules/fav.html"
const HISTORIAL = "modules/historial.html"

//OCURRENCIAS numero minimo de palabras repetidas por pag
const OCURRENCIAS int = 3

var search model.SearchBing
var visited = make(map[string]bool)
var page model.PageVars
var terms = []string{"INFORMATICA", "JAVASCRIPT", "SCRIPT", "CSS", "HTML", "INFORMATICO", "SISTEMA", "SISTEMAS", "COMPUTADORA",
	"COMPUTADORAS", "COMPUTADOR", "INTERNET", "INTRANET", "ETHERNET", "PROGRAMACION", "PROGRAMAR", "CODIGO", "LENGUAJE DE PROGRAMACION",
	"SOFTWARE", "HARDWARE", "SERVIDOR", "CODIGO BINARIO", "COMPUTER", "DEVELOPER", "COMPUTER", "SYSTEM", "SYSTEMS", "COMPUTER", "COMPUTERS",
	"COMPUTER", "PROGRAMMING", "CODE", " PROGRAMMING LANGUAGE ", " SERVER ", " BINARY CODE "}

//Controller estructura de datos que apunta a repository
type Controller struct {
	Repository model.Repository
}

//Render Redireccionar las paginas y agregarle valor dinamico a la misma a traves de estructuras de datos
func Render(w http.ResponseWriter, tmpl string, pageVars interface{}) {

	tmpl = fmt.Sprintf("views/%s", tmpl) // prefix the name passed in with templates/
	t, err := template.ParseFiles(tmpl)  //parse the template file held in the templates folde
	if err != nil {                      // si la variable error es distinto de nulo
		log.Print("error al cargar el archivo: ", err)
	}

	err = t.Execute(w, pageVars) //ejecutar el templete

	if err != nil { // si la variable error es distinto de nulo
		log.Print("template executing error: ", err)
	}

}

/******************************************************
* TEMPLATE
*****************************************************************/

//Home vista principal
func Home(w http.ResponseWriter, req *http.Request) {
	userName := getUserName(req)
	p := model.PageVars{}
	p.Title = "CodGo!"
	if userName != "" {
		p.Estado = "d-block"
		p.User = userName
	} else {
		p.Estado = "d-none"
		p.User = ""
	}
	Render(w, HOME, p)
}

//SearchHome vista secundaria
func SearchHome(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	r.ParseForm()
	data := r.FormValue("search-input")
	p := model.PageVars{}
	p.Title = "CodGo! - " + data
	p.Val = data
	if userName != "" {
		p.Estado = "d-block"
		p.User = userName
	} else {
		p.Estado = "d-none"
		p.User = ""
	}
	Render(w, SEARCH, p)
}

//IniciarSesion vista login
func IniciarSesion(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	a := model.Estado{
		Estado: "",
		Valor:  "",
	}
	Render(w, LOGIN, a)
}

//Registrar vista registro
func Registrar(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	a := model.Estado{
		Estado: "",
		Valor:  "",
	}
	Render(w, REGISTER, a)
}

//CerrarSesion cerrar sesion cookie
func CerrarSesion(w http.ResponseWriter, request *http.Request) {
	clearSession(w)
	p := model.PageVars{}
	p.Title = "CodGo!"
	p.Estado = "d-none"
	p.User = ""
	Render(w, HOME, p)
}

//ReceiveAjax comunicacion cliente servidor
func (c *Controller) ReceiveAjax(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//fmt.Println(r.Form.Get("value"))

	data := r.FormValue("value")
	typed := r.FormValue("type")
	pag := r.FormValue("pag")
	userName := getUserName(r)
	var user string
	if userName != "" {
		user = userName
	} else {
		user = ip()
	}
	l := model.Link{
		User:  user,
		Theme: data,
	}

	if r.Method == "POST" && data != "" && typed != "fav" && typed != "historial" {
		if pag != "" {
			//fmt.Println("Receive ajax post data string ", data, typed)
			if typed == "busquedaInformatica" {
				fmt.Println("Receive ajax post data string ", data, typed)
				link := c.Repository.GetLinkByTheme(l)
				a := SearchApi(data, pag, "+%2B'informatica'", link)
				if a.Error.Message != "" {

					w.Write([]byte(a.Error.Message))
				} else {
					Render(w, FILE, a)
				}

			} else if typed == "img" {
				fmt.Println("Receive ajax post data string ", data, typed)
				link := c.Repository.GetLinkByTheme(l)
				a := SearchApi(data, "1", "+%2B'informatica'&searchtype=image", link)
				b := SearchApi(data, "11", "+%2B'informatica'&searchtype=image", link)
				j := make([]model.Items, 0)
				for _, e := range a.Items {
					j = append(j, e)
				}
				for _, e := range b.Items {
					j = append(j, e)
				}
				api := model.SearchBing{ //resultados filtrados de la api
					Favoritos:         a.Favoritos,
					Context:           a.Context,
					URL:               a.URL,
					Kind:              a.Kind,
					Queries:           a.Queries,
					SearchInformation: a.SearchInformation,
					Items:             j,
					Fav:               a.Fav,
				}
				if a.Error.Message != "" {
					w.Write([]byte("false"))
				} else {
					Render(w, IMAGE, api)
				}
			} else if typed == "video" {
				fmt.Println("Receive ajax post data string ", data, typed)
				link := c.Repository.GetLinkByTheme(l)
				a := SearchApi(data, pag, "+%2B'informatica'&siteSearch=https://www.youtube.com", link)

				if a.Error.Message != "" {
					w.Write([]byte("false"))
				} else {
					Render(w, VIDEO, a)
				}

			} else if typed == "doc" || typed == "pdf" {
				fmt.Println("Receive ajax post data string ", data, typed)
				link := c.Repository.GetLinkByTheme(l)
				a := SearchApi(data+"filetype:"+typed, pag, "", link)
				if a.Error.Message != "" {
					w.Write([]byte("false"))
				} else {
					Render(w, FILE, a)
				}
			} else if typed == "busquedaNormal" {
				link := c.Repository.GetLinkByTheme(l)
				fmt.Println("Receive ajax post data string ", data, typed)
				a := SearchApi(data, pag, "", link)
				if a.Error.Message != "" {
					w.Write([]byte("false"))
				} else {
					Render(w, FILE, a)
				}
			} else if typed == "definicion" {
				link := c.Repository.GetLinkByTheme(l)
				fmt.Println("Receive ajax post data string ", data, typed)
				a := SearchApi("definir:"+data, pag, "+%2B'informatica'", link)
				if a.Error.Message != "" {
					w.Write([]byte("false"))
				} else {
					Render(w, FILE, a)
				}
			} else {
				fmt.Println("Receive ajax post data string ", data, typed)
				link := c.Repository.GetLinkByTheme(l)
				a := SearchApi(data, pag, "+%2B'informatica'", link)
				if a.Error.Message != "" {
					w.Write([]byte("false"))
				} else {
					Render(w, FILE, a)
				}

			}
		} else {
			fmt.Println("Receive ajax post data string ", data, typed)
			link := c.Repository.GetLinkByTheme(l)
			a := SearchApi(data, "1", "+%2B'informatica'", link)
			if a.Error.Message != "" {
				w.Write([]byte("false"))
			} else {
				Render(w, FILE, a)
			}
		}

	} else if typed == "fav" {
		links, _ := c.Repository.GetLinksByUser(l)
		fmt.Println("Receive ajax post data string ", typed)
		Render(w, FAV, links)

	} else if typed == "historial" {
		hist, _ := c.Repository.GetHistsByUser(user)
		fmt.Println("Receive ajax post data string ", typed)
		Render(w, HISTORIAL, hist)

	} else if typed == "borrarHistorial" {
		c.DeleteHist(w, r)
		hist, _ := c.Repository.GetHistsByUser(user)
		fmt.Println("Receive ajax post data string ", typed)
		Render(w, HISTORIAL, hist)

	} else if typed == "borrarFavoritos" {
		c.DeleteFav(w, r)
		links, _ := c.Repository.GetLinksByUser(l)
		fmt.Println("Receive ajax post data string ", typed)
		Render(w, FAV, links)

	}

}

/*********************************************************************************
*Search
*********************************************************************************************/

// SearchApi comunicacion con la api de google
func SearchApi(query, pag, filtro string, fav model.L) model.SearchBing {
	a := strings.Join(strings.Fields(query), "+")
	url := "https://www.googleapis.com/customsearch/v1?key=AIzaSyBXvmQYTiWVlw5BBzZKX4-71eiY_rZ8LA4&cx=001073512781794923538:gelo0yxfxxq&lr=lang_es&mkt=es-ES&start=" + pag + "&setLang=ES&q=" + a + filtro

	//url := "https://www.googleapis.com/customsearch/v1?key=AIzaSyCNrW9VpLEJElzM-OMcCbPOT3eyakC2y2Q&cx=001073512781794923538:gelo0yxfxxq&lr=lang_es&mkt=es-ES&start=" + pag + "&setLang=ES&q=" + a + filtro
	//url := "https://www.googleapis.com/customsearch/v1?key=AIzaSyBWbbDssCq10HUX6dbws1Q42E5jGODLdG0&cx=001073512781794923538:gelo0yxfxxq&lr=lang_es&mkt=es-ES&start=" + pag + "&setLang=ES&q=" + a + filtro
	//url := "https://www.googleapis.com/customsearch/v1?key=AIzaSyAB2CccC9GSHgn_B2TIA1uLSMSTKpmYe3I&cx=001073512781794923538:gelo0yxfxxq&lr=lang_es&mkt=es-ES&start=" + pag + "&setLang=ES&q=" + a + filtro

	req, _ := http.NewRequest("GET", url, nil)
	//req.Header.Set("X-Api-Key", "AIzaSyBWbbDssCq10HUX6dbws1Q42E5jGODLdG0")
	resp, err := http.DefaultClient.Do(req)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	api := model.SearchBing{} //resultados obtenidos de la api
	api2 := model.SearchBing{}
	err = json.Unmarshal(body, &api)
	if err != nil {
		fmt.Println("error", err)
	}
	s := make(map[string]model.Items)

	if len(fav.Link) != 0 {
		//j := make([]model.Items, 0)
		for _, a := range api.Items {
			s[a.Link] = a

		}
		j := make([]model.Items, 0)
		for _, a := range s {
			j = append(j, a)
		}

		api2 = model.SearchBing{ //resultados filtrados de la api
			Favoritos:         "si",
			Context:           api.Context,
			URL:               api.URL,
			Kind:              api.Kind,
			Queries:           api.Queries,
			SearchInformation: api.SearchInformation,
			Items:             j,
			Fav:               fav,
		}
	} else {
		api2 = api
		api2.Favoritos = "no"
	}
	return api2
}

//funcion para comparar los terminos con el contenido de la pagina
func compare(f, g, h string) int {
	a, b, c := 0, 0, 0

	if f != "" {
		a = compareTo(f)
	} else if g != "" {
		b = compareTo(g)
	} else if h != "" {
		c = compareTo(h)
	}
	return a + b + c
}

func compareTo(m string) int {
	var t int
	st := strings.ToUpper(m)
	//fmt.Println("compare to " + st)
	for _, v := range terms {
		if strings.Contains(st, v) == true {
			t = t + strings.Count(st, v)
			//		fmt.Println(v, t)
		}
	}
	return t

}

/*------------------
*Obtener tipo
*------------------------*/

// TypeDoc octiene tipo de documento return string
func TypeDoc(url string) string {
	resp, _ := http.Get(url)
	// Check Content-Type is HTML (e.g., "text/html; charset=utf-8").
	ct := resp.Header.Get("Content-Type")
	return ct
}

/*------------------
*API
*------------------------*/

//getApi que obtiene datos de una pag web
func getApi(uri string) model.Api {
	url := "https://mercury.postlight.com/parser?url=" + uri
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Api-Key", "hcBUpDJdktAHjrzn6Y7pV4ueFaDVSl3Wrx1tkbeh")
	resp, err := http.DefaultClient.Do(req)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	api := model.Api{}
	err = json.Unmarshal(body, &api)
	if err != nil {
		fmt.Println("error func get Title", err)
	}

	return api
}

/*****************************
*Session
****************************************/

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32),
)

func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func refreshSession(req *http.Request, resp http.ResponseWriter) {
	u := getUserName(req)
	clearSession(resp)
	if u != "" {
		setSession(u, resp)
	}
}

/******************************
*CRUD USER
*************************************/

//Login confirma si el usuario existe
func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := r.FormValue("email")
	pass := r.FormValue("password")

	u := model.User{
		User:     user,
		Password: pass,
	}

	if c.Repository.Login(u) == true {
		setSession(user, w)

		a := model.PageVars{
			Title:  "CodGo-" + user,
			User:   user,
			Estado: "d-block",
		}
		Render(w, HOME, a)
	} else {
		a := model.Estado{
			Estado: "Los Datos ingresados son incorrectos",
			Valor:  "danger",
		}
		Render(w, LOGIN, a)
	}

}

//AddUserNew agrega usuario a la bdd
func (c *Controller) AddUserNew(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	users := r.FormValue("email")
	password1 := r.FormValue("password1")
	password2 := r.FormValue("password2")

	if r.Method == "POST" && users != "" && password1 != "" && password2 != "" && password1 == password2 {

		fmt.Println("Receive ajax post user ", users, password1, password2)

		user := model.User{
			ID:       bson.NewObjectId(),
			User:     users,
			Password: password1,
		}

		if c.Repository.GetUserByString(user) == false {

			success := c.Repository.AddUser(user) // adds the product to the DB
			if !success {
				a := model.Estado{
					Estado: "Error al ingresar datos en base de datos",
					Valor:  "danger",
				}

				Render(w, REGISTER, a)
			}
			a := model.Estado{
				Estado: "Usuario creado",
				Valor:  "success",
			}

			Render(w, LOGIN, a)
		} else {
			a := model.Estado{
				Estado: "Usuario Ya Existe, por favor ingrese otro usuario",
				Valor:  "warning",
			}

			Render(w, REGISTER, a)
		}

	} else {
		a := model.Estado{
			Estado: "Las Claves no coinciden",
			Valor:  "danger",
		}
		Render(w, REGISTER, a)
	}

	return
}

/******************************
*CRUD Links
*************************************/

//ReceiveAjaxFavoritos agrega u elimina un link a la base de datos
func (c *Controller) ReceiveAjaxFavoritos(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	theme := r.FormValue("theme")
	link := r.FormValue("link")
	class := r.FormValue("class")
	title := r.FormValue("title")
	user := getUserName(r)
	if user != "" {
		user = getUserName(r)
	} else {
		user = ip()
	}

	println(theme, link, class)
	l := model.Link{
		ID:    bson.NewObjectId(),
		User:  user,
		Theme: theme,
		Link:  link,
		Title: title,
	}

	if r.Method == "POST" {

		if class == "icon-star" {
			if c.Repository.GetLinkByUser(l) == false {
				success := c.Repository.AddLink(l) // adds link
				if !success {
					println("Error al agregar link")
					w.Write([]byte("Error al agregar link"))
				}
				w.Write([]byte("icon-star"))

			} else {
				w.Write([]byte("icon-filled"))
				println(" link ya existe")
			}

		} else if class == "icon-star-filled" {
			if c.Repository.GetLinkByUser(l) != false {
				if err := c.Repository.DeleteLink(l); err != true { // delete a link
					println(err, "error al borrar")

				}
				w.Write([]byte("icon-star-filled"))
			} else {
				w.Write([]byte("icon-star"))
				println("link no existe en la bdd")
			}
		}

	}

}
func (c *Controller) DeleteFav(w http.ResponseWriter, r *http.Request) {
	user := getUserName(r)
	if user != "" {
		user = getUserName(r)
	} else {
		user = ip()
	}

	if err := c.Repository.DeleteLinks(user); err != true { // delete a link
		println(err, "error al borrar")

	} else {
		println("Borrado")
	}
}

/******************************
*CRUD HISTORIAL
*************************************/
func (c *Controller) DeleteHist(w http.ResponseWriter, r *http.Request) {
	user := getUserName(r)
	if user != "" {
		user = getUserName(r)
	} else {
		user = ip()
	}

	if err := c.Repository.DeleteHist(user); err != true { // delete a link
		println(err, "error al borrar")

	} else {
		println("Borrado")
	}
}

//ReceiveAjaxFavoritos agrega u elimina un link a la base de datos
func (c *Controller) ReceiveAjaxHISTORIAL(w http.ResponseWriter, r *http.Request) {

	theme := r.FormValue("ty")
	link := r.FormValue("link")

	user := getUserName(r)
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02d",
		t.Year(), t.Month(), t.Day())
	Hora := fmt.Sprintf("%02d:%02d:%02d",
		t.Hour(), t.Minute(), t.Second())
	hist := model.Hist{
		ID:    bson.NewObjectId(),
		User:  user,
		Theme: theme,
		Link:  link,
		Hours: Hora,
		Date:  fecha,
	}
	if user != "" {
		hist.User = getUserName(r)
	} else {
		hist.User = ip()
	}
	if r.Method == "POST" {
		success := c.Repository.AddHist(hist)
		if !success {
			println("Error al agregar link")
			w.Write([]byte("Error al agregar link"))
		}
		w.Write([]byte("agregado al historial "))
	}

}

/***********************************************
*IP
*******************************/

// ip... obtiene la direccion ip de conexion
func ip() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {

				js, err := json.Marshal(ipnet.IP.String())
				if err != nil {
					log.Fatal(err)
				}
				return string(js)
			}
		}

	}
	return ""
}
