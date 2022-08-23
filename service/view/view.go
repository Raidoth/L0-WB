package view

import (
	"log"
	"net/http"
	"test/service/cache"
	"test/service/config"
	"text/template"
)

func MainPage(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	id := r.Form.Get("ids")
	mod, _ := cache.GetCache(id)
	tmpl, err := template.ParseFiles(config.ViewPathMainPage)
	if err != nil {
		log.Println(err)
	}
	tmpl.Execute(w, nil)
	tmpl.ExecuteTemplate(w, "order", mod)

}
