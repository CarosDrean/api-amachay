package main

import (
	"flag"
	"fmt"
	"github.com/CarosDrean/api-amachay/constants"
	"github.com/CarosDrean/api-amachay/storage"
	"github.com/CarosDrean/api-amachay/helper"
	"github.com/CarosDrean/api-amachay/middleware"
	routes "github.com/CarosDrean/api-amachay/router"
	"github.com/gonutz/w32"
	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main() {
	f := flag.Bool("execTerminal", false, "to exec cmd")
	flag.Parse()
	if !*f {
		hideConsole()
	}
	// api()
	apiEcho()
}

func apiEcho() {
	storage.DB = helper.Get()
	e := echo.New()

	myLog, err := os.OpenFile(
		"logs.log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Fatalf("No se pudo crear el archivo de logs: %v", err)
	}
	defer myLog.Close()
	logConfig := mid.LoggerConfig{
		Output: myLog,
	}

	// solo en prod
	// e.Use(mid.Recover())
	e.Use(mid.LoggerWithConfig(logConfig))

	e.GET("/", indexEcho)
	api := e.Group("/api")
	routes.RoutesEcho(api)

	e.Logger.Fatal(e.Start(":" + constants.PORT))
}

func indexEcho(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome Inventory Echo!")
}

func api() {
	r := mux.NewRouter()

	storage.DB = helper.Get()

	r.HandleFunc("/", indexRouter)
	r.HandleFunc("/api/login", middleware.Login)

	s := r.PathPrefix("/api").Subrouter()
	routes.Routes(s)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:4200",
			"https://inventario.holosalud.org",
		},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "x-token"},
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = constants.PORT //localhost
	}

	handler := c.Handler(r)

	fmt.Println(fmt.Sprintf("Server online in %s!", constants.PORT))

	log.Fatal(http.ListenAndServe(":"+port, handler))

}

func indexRouter(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Welcome api inventory holo!")
}

func hideConsole() {
	console := w32.GetConsoleWindow()
	if console != 0 {
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		if w32.GetCurrentProcessId() == consoleProcID {
			w32.ShowWindowAsync(console, w32.SW_HIDE)
		}
	}
}
