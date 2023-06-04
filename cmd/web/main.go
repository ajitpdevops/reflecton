package main

import (
	"log"
	"net/http"

	"github.com/whoajitpatil/reflecton/pkg/config"
	"github.com/whoajitpatil/reflecton/pkg/handlers"
	"github.com/whoajitpatil/reflecton/pkg/render"
)

var portNumber = ":8080"

func main() {

	//  app variable set to type AppConfig in global configuration
	var app config.AppConfig

	// calling CreateTemplateCache method from main
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	// Assigning the returned templates to app.TemplateCache global variable
	app.TemplateCache = tc
	// UseCache in Global Reference
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// Call new template chache
	render.NewTemplates(&app)

	// Starting the web server
	log.Printf("Starting application server on %s\n", portNumber)

	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
