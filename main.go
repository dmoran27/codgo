package main

import (
	"buscadorWeb/controller"
	"buscadorWeb/model"
	"net/http"
)

func main() {

	//solo acceder a los archivos encontrados en la carpeta views
	http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("views"))))
	// cuando se accede  a la carpeta raiz /home llama la funcion Home

	var C = &controller.Controller{Repository: model.Repository{}}
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/searchHome", controller.SearchHome)
	http.HandleFunc("/iniciarSesion", controller.IniciarSesion)
	http.HandleFunc("/CerrarSesion", controller.CerrarSesion)
	http.HandleFunc("/registrar", controller.Registrar)
	http.HandleFunc("/ajax", C.ReceiveAjax)
	http.HandleFunc("/ajaxFavoritos", C.ReceiveAjaxFavoritos)
	http.HandleFunc("/login", C.Login)
	http.HandleFunc("/Historial", C.ReceiveAjaxHISTORIAL)
	http.HandleFunc("/BorrarHistorial", C.DeleteHist)
	http.HandleFunc("/BorrarFavoritos", C.DeleteFav)
	http.HandleFunc("/registro", C.AddUserNew)
	http.ListenAndServe(":8080", nil)

}
