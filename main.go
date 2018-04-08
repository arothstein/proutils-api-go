// proutils-api
//
// This is the main package which starts up and runs our REST API service.

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/arothstein/proutils-api/publicip"
	"github.com/arothstein/proutils-api/whois"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

// main launches api web server which runs indefinitely.
func main() {

	// Setup all routes.
	router := httprouter.New()
	router.GET("/whois/:ip", whois.HandleWHOISIP)
	router.GET("/public-ip", publicip.HandlePublicIP)

	// Setup 404 / 405 handlers.
	router.NotFound = http.HandlerFunc(ErrorNotFound)
	router.MethodNotAllowed = http.HandlerFunc(ErrorMethodNotAllowed)

	// Setup middlewares.
	// - Add CORS to restrict request origins.
	origins := []string{"https://proutils.com", "https://www.proutils.com", "http://localhost", "http://localhost:8080"}
	c := cors.New(cors.Options{
		AllowedOrigins: origins,
	})

	handler := c.Handler(router)

	// Start the server.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Starting HTTP server on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
