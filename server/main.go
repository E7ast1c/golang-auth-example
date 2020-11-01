package main

import (
	"encoding/json"
	"fmt"
	"github.com/E7ast1c/golang-auth-example/api/routes"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {


	if e != nil {

	}
	fmt.Println(e)

	key := "PORT"
	var port string
	val, ok := os.LookupEnv("PORT")
	if ok {
		log.Printf("%s=%s\n", key, val)
		port = val
	} else {
		log.Printf("%s not set\n", key)
		port = "8080"
	}

	port := os.Getenv("PORT")

	// Handle routes
	http.Handle("/", routes.Handlers())

	// serve
	log.Printf("Server up on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

//	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))


	// For dev only - Set up CORS so React client can consume our API
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	// Our application will run on port 8080. Here we declare the port and pass in our router.
	srv := &http.Server{
		Addr:              "127.0.0.1:8080",
		Handler:           corsWrapper.Handler(r),
		TLSConfig:         nil,
		ReadTimeout:       time.Second * 10,
		ReadHeaderTimeout: time.Second * 5,
		WriteTimeout:      time.Second * 10,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	srv.ListenAndServe()
}

func TopFilms(w http.ResponseWriter, r *http.Request) {
	films := []string{"The Shawshank Redemption", "The Godfather", "The Dark Knight",
		"The Godfather: Part II", "The Lord of the Rings: The Return of the King"}

	payload, err := json.Marshal(films)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set(applicationJson, contentType)
	w.Write(payload)
}

const (
	applicationJson string = "application/json"
	contentType     string = "Content-Type"
)





