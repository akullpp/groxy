package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/akullpp/gotenv"
)

func handleRequest(w http.ResponseWriter, r *http.Request, dotenv map[string]string) {
	ps := strings.Split(strings.TrimPrefix(r.URL.Path, "/"), "/")
	id := strings.ToUpper(ps[0])

	t := dotenv[id]
	if t == "" {
		t = dotenv["DEFAULT"]
	}

	u, err := url.Parse(t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Proxying %v to %v\n", r.URL.Path, t)
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.Host = u.Host
			req.URL.Scheme = u.Scheme
			req.URL.Host = u.Host
		},
	}
	proxy.ServeHTTP(w, r)
}

func main() {
	dotenv, err := gotenv.Get()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleRequest(w, r, dotenv)
	})

	s := dotenv["SERVER"]
	fmt.Printf("Serving on %v\n", s)
	log.Fatal(http.ListenAndServe(s, nil))
}
