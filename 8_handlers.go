package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	app.render(w, r, "home.page.tmpl")
}
func (app *application) ver(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "yandex_f4e9350effdb4931.page.tmpl")
}

func (app *application) sitemap(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/static/other/sitemap.xml")
}

// func find_file(w http.ResponseWriter, r *http.Request)
func (app *application) save_info(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	mail := r.FormValue("mail")
	phone := r.FormValue("phone")
	problem := r.FormValue("problem")
	sendUser := User{name, mail, phone, problem}
	//Отпрака сообщений в телеграм.
	err := SendToAllUser(botUrl, sendUser)
	if err != nil {
		app.serverError(w, err)
	}
	//Отправка сообщений на почту.
	err = SendMailSimple(name, mail, phone, problem)
	if err != nil {
		app.serverError(w, err)
	}

	type form1 struct {
		Name    string
		Mail    string
		Phone   string
		Problem string
	}
	data := form1{Name: name, Mail: mail, Phone: phone, Problem: problem}
	ts, ok := app.templateCache["succes.page.tmpl"]
	if !ok {
		fmt.Println(app.templateCache)
		app.serverError(w, fmt.Errorf("Шаблон %s не существует!", name))
		return
	}

	err = ts.Execute(w, &data)
	if err != nil {
		app.serverError(w, err)
	}

	//http.Redirect(w, r, "/successful", http.StatusSeeOther)
	//http.Redirect(w, r, "/about", http.StatusSeeOther)
}
func (app *application) successful(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "succes.page.tmpl")
}
