package controller

import (
	"app/model"
	"strconv"
	


	"html/template"
	"net/http"
	"path/filepath"
	"fmt"

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
		http.Error(rw, "Invalid age value", 400) // Обработка ошибки при конвертации
		return
	}

	
	user := model.User{
		Name:    name,
		Surname: surname,
		Age:     age,
	}

	// Добавляем пользователя в модель
	model.AddUser(&user)

	// Ответ
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "User added successfully: %v", user)
}


