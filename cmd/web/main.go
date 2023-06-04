package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/whoajitpatil/reflecton/pkg/config"
	"github.com/whoajitpatil/reflecton/pkg/handlers"
	"github.com/whoajitpatil/reflecton/pkg/render"
)

var portNumber = ":8080"
//  app variable set to type AppConfig in global configuration
var app config.AppConfig
var session *scs.SessionManager

func main() {
	// Change this true when in Production
	app.InProduction = false

	// Creating a session and setting a timeout 
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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
