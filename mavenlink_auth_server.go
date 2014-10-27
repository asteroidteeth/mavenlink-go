package main

import (
	"fmt"
	"github.com/asteroidteeth/mavenlink-go/webapp"
	"github.com/rakyll/globalconf"
	"log"
	"net/http"
	"os"
	"path"
)

var port string = ":" + os.Getenv("PORT")

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HI")
}

func main() {
	// Lazy dev config path... for now
	configPath := path.Join(os.Getenv("GOPATH"), "src/github.com/asteroidteeth/go-mavenlink/config.ini")
	config, configErr := globalconf.NewWithOptions(&globalconf.Options{
		Filename: configPath,
	})
	if configErr != nil {
		log.Fatalf("Error loading config! \"%s\"\n", configErr.Error())
	}
	config.Parse()

	http.HandleFunc("/get-registered/", webapp.RegistrationHandler)
	http.HandleFunc("/auth/callback", webapp.Handler)
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println(err)
	}
}
