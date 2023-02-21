package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"lenslocked/lenslocked/controllers"
	"lenslocked/lenslocked/templates"
	"lenslocked/lenslocked/views"
	"net/http"
)

//func executeTemplate(w http.ResponseWriter, filepath string) {
//	tpl, err := views.Parse(filepath)
//	if err != nil {
//		log.Printf("parsing template: %v", err)
//		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
//		return
//	}
//	tpl.Execute(w, nil)
//}

//func homeHandler(w http.ResponseWriter, r *http.Request) {
//	executeTemplate(w, "templates/home.gohtml")
//}
//
//func contactHandler(w http.ResponseWriter, r *http.Request) {
//	executeTemplate(w, "templates/contact.gohtml")
//}
//
//type Router struct{}

//func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/":
//		homeHandler(w, r)
//	case "/contact":
//		contactHandler(w, r)
//	default:
//		//errorHandling(w, http.StatusNotFound)
//		http.Error(w, "Page not found", http.StatusNotFound)
//	}
//}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	http.ListenAndServe(":3000", r)
}
