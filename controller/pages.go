package controller

import (
	
	"net/http"
	"path/filepath"
	"html/template"
	"github.com/julienschmidt/httprouter"
)

func StartPageController(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	path := filepath.Join("public", "html", "startStaticPage.html")

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = tmpl.Execute(rw, nil)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return 
	}

}


func RenderAddUserForm(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Путь к файлу HTML формы
	formPath := filepath.Join("public", "html", "userAddPage.html")

	// Парсинг HTML-шаблона
	tmpl, err := template.ParseFiles(formPath)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка страницы с формой в браузер
	err = tmpl.Execute(rw, nil)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}