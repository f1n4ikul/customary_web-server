package controller

import (
	"app/model"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GetUsersController(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	users, err := model.GetAllUsers()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	main := filepath.Join("public", "html", "usersDynamicPage.html")
	common := filepath.Join("public", "html", "common.html")

	tmpl, err := template.ParseFiles(main, common)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = tmpl.ExecuteTemplate(rw, "users", users)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func AddUserController(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	name := r.FormValue("name")
	surname := r.FormValue("surname")
	ageStr := r.FormValue("age")

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		http.Error(rw, "Invalid age value", 400)
		return
	}

	user := model.User{
		Name:    name,
		Surname: surname,
		Age:     age,
	}

	
	model.AddUser(&user)

	
	http.Redirect(rw, r, "/users", http.StatusSeeOther)
}