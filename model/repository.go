package model

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

// SERVER servidor de la base de datos
//const SERVER = "mongodb://gautam:gautam@ds157233.mlab.com:57233/dummystore"
const SERVER = "localhost:27017"

// DBNAME instancia del nombre de la base de datos
const DBNAME = "codGo"

// COLLECTION nombre de la coleccion de la BD
const COLLECTION = "users"

// COLLECTION2 nombre de la coleccion de la BD
const COLLECTION2 = "links"

// COLLECTION2 nombre de la coleccion de la BD
const COLLECTION3 = "Hist"

/*-------------------------
*CRUD DE USUARIOS
-------------------------------------------*/

// Login Consulta la existencia del usuario
func (r Repository) Login(user User) bool {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("fallo la conexion con el servidor de MongoDB:", err)
		return false
	}
	defer session.Close()
	c := session.DB(DBNAME).C(COLLECTION)
	results := U{}
	if err := c.Find(bson.M{"user": user.User, "password": user.Password}).One(&results); err != nil {
		fmt.Println("fallo al mostrar los resultados de la la Base de datos:", err)
		return false
	}
	return true
}

// GetUser retornar la lista de usuarios almacenadas en la base de datos
func (r Repository) GetUser() U {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("fallo la conexion con el servidor de MongoDB:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	results := U{}

	if err := c.Find(nil).All(&results.User); err != nil {
		fmt.Println("fallo al mostrar los resultados de la la Base de datos:", err)
	}
	return results
}

// GetUserByString obtiene un string con el nombre de usuario y retorna los datos del usuario
func (r Repository) GetUserByString(user User) bool {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("fallo la conexion con el servidor de MongoDB:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	result := U{}

	if err := c.Find(bson.M{"user": user.User}).One(&result); err != nil {
		fmt.Println("Failed to write result:user exist", err)
		return false
	}

	return true
}

// AddUser agrega un nuevo usuario a la base de datos
func (r Repository) AddUser(user User) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	session.DB(DBNAME).C(COLLECTION).Insert(&user)
	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("nueva Url agregado: ID- ", user.ID)

	return true
}

/*-------------------------
*CRUD DE LINKS
-------------------------------------------*/

//GetLinks obtiene todos los links de la bdd
func (r Repository) GetLinks() L {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("fallo la conexion con el servidor de MongoDB:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION2)
	results := L{}

	if err := c.Find(nil).All(&results.Link); err != nil {
		fmt.Println("fallo al mostrar los resultados de la la Base de datos:", err)
	}
	return results
}

// GetLinkByUser consulta existencia de link
func (r Repository) GetLinkByUser(link Link) bool {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("fallo la conexion con el servidor de MongoDB:", err)
		return false
	}
	defer session.Close()
	c := session.DB(DBNAME).C(COLLECTION2)
	results := U{}
	if err := c.Find(bson.M{"user": link.User, "link": link.Link}).One(&results); err != nil {
		fmt.Println("fallo al mostrar los resultados de la la Base de datos:", err)
		return false
	}
	return true

}

// GetLinkByUser consulta links por Usuario
func (r Repository) GetLinksByUser(link Link) (L, bool) {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("fallo la conexion con el servidor de MongoDB:", err)

	}
	defer session.Close()
	c := session.DB(DBNAME).C(COLLECTION2)
	results := L{}
	if err := c.Find(bson.M{"user": link.User}).All(&results.Link); err != nil {
		fmt.Println("fallo al mostrar los resultados de la la Base de datos:", err)
		return results, false
	}
	return results, true
}

// GetLinkByTheme consulta links por tema
func (r Repository) GetLinkByTheme(link Link) L {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("fallo la conexion con el servidor de MongoDB:", err)

	}
	defer session.Close()
	c := session.DB(DBNAME).C(COLLECTION2)
	results := L{}
	if err := c.Find(bson.M{"theme": link.Theme, "user": link.User}).All(&results.Link); err != nil {
		fmt.Println("fallo al mostrar los resultados de la la Base de datos:", err)
		return results
	}
	return results
}

// AddLink agrega un nuevo Link a la base de datos
func (r Repository) AddLink(link Link) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	session.DB(DBNAME).C(COLLECTION2).Insert(&link)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("nueva Url agregado: ID- ", link.ID)

	return true
}

// DeleteLink elimina las url de la base de datos
func (r Repository) DeleteLink(link Link) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	//borrar url
	if err = session.DB(DBNAME).C(COLLECTION2).Remove(bson.M{"user": link.User, "link": link.Link}); err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Url Borrado - ", link.User)
	// estado de operacion
	return true
}
func (r Repository) DeleteLinks(link string) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	//borrar url
	if _, err = session.DB(DBNAME).C(COLLECTION2).RemoveAll(bson.M{"user": link}); err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Url Borrado - ", link)
	// estado de operacion
	return true
}

/*-------------------------
*CRUD DE HISTORIAL
-------------------------------------------*/

//GetHist obtiene todos los datos del historial de la bdd
func (r Repository) GetHist() H {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("fallo la conexion con el servidor de MongoDB:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION3)
	results := H{}

	if err := c.Find(nil).All(&results.Hist); err != nil {
		fmt.Println("fallo al mostrar los resultados de la la Base de datos:", err)
	}
	return results
}

// GetHistByUser consulta Hist por Usuario
func (r Repository) GetHistsByUser(user string) (H, bool) {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("fallo la conexion con el servidor de MongoDB:", err)

	}
	defer session.Close()
	c := session.DB(DBNAME).C(COLLECTION3)
	results := H{}
	if err := c.Find(bson.M{"user": user}).All(&results.Hist); err != nil {
		fmt.Println("fallo al mostrar los resultados de la la Base de datos:", err)
		return results, false
	}
	return results, true
}

// AddHist agrega un nuevo Link a la base de datos
func (r Repository) AddHist(hist Hist) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	session.DB(DBNAME).C(COLLECTION3).Insert(&hist)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("nueva Url agregado: ID- ", hist.ID)

	return true
}
func (r Repository) DeleteHist(hist string) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	//borrar url
	if _, err = session.DB(DBNAME).C(COLLECTION3).RemoveAll(bson.M{"user": hist}); err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Historial Borrado - ")
	// estado de operacion
	return true
}
