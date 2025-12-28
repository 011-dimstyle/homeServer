package main

import(
	"html/template"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main(){
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*",http.StripPrefix("/static/",fs))

	templates := template.Must(template.ParseGlob("templates/*.html"))
	router.Get("/",func(w http.ResponseWriter, r *http.Request){
		templates.ExecuteTemplate(w,"index.html", make(map[string]string))
	})

	http.ListenAndServe(":8080",router)
}