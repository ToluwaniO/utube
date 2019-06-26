package main

import (
	"context"
	"fmt"
	"github.com/ToluwaniO/utube/controller"
	"github.com/ToluwaniO/utube/global"
	"github.com/ToluwaniO/utube/helper"
	"log"
	"net/http"
	"os"
)

func main() {
	//firebaseClient := initFirebase()
	initGlobal()
	http.HandleFunc("/", indexHandler)
	http.Handle("/addUser", isAuthorized(controller.AddUser))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		_, err := helper.ValidateToken(context.Background(), authorization)

		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "Unauthorized")
			return
		}
		endpoint(w, r)
	})
}

func initGlobal() {
	global.InitVariables()
}

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

