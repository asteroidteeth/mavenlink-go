package webapp

import (
	"bytes"
	"fmt"
	"github.com/asteroidteeth/mavenlink-go/mavenlink"
	"log"
	"net/http"
	"net/url"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.String())
	log.Println(r.ContentLength)
	log.Println(r.RemoteAddr)
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	s := buf.String()
	log.Println(s)
	userCode := r.URL.Query().Get("code")
	if userCode != "" {
		token := mavenlink.ExchangeForToken(userCode)
		if token != "" {
			fmt.Fprintf(w, "Now I have your very soul, fool! And its name is %s", mavenlink.GetMavenlinkUserName(token))
		}
	}
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Location", fmt.Sprintf(mavenlink.AuthUrl, *mavenlink.AppId, url.QueryEscape(*mavenlink.RedirectUri)))
	w.WriteHeader(303)
}
