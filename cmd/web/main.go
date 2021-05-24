package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/mrazen/bookings/pkg/config"
	"github.com/mrazen/bookings/pkg/handlers"
	"github.com/mrazen/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"
// go run .\cmd\web\main.go .\cmd\web\middleware.go .\cmd\web\routes.go
var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s \n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
