package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"app/controller"
)

func main() {
	r := httprouter.New()
	routers(r)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil{
		log.Fatal(err)
	}

}

func routers(r *httprouter.Router) {
	r.ServeFiles("/public/*filepath", http.Dir("public"))

	r.GET("/", controller.StartPageController)
	r.GET("/users", controller.GetUsersController)
	r.GET("/add", controller.RenderAddUserForm)
	r.POST("/addUser", controller.AddUserController)
}

