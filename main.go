package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var templates = template.Must(template.ParseFiles("./templates/base.html", "./templates/body.html"))

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		req := fmt.Sprintf("%s %s", r.Method, r.URL)
		log.Println(req)
		next.ServeHTTP(w, r)
		log.Println(req, "completed in", time.Now().Sub(start))
	})
}

func index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getIP := r.Header.Get("X-Forwarded-For")
		if getIP == "" {
			getIP = r.Header.Get("X-Real-Ip")
		}
		if getIP == "" {
			getIP = r.RemoteAddr
		}
		b := struct {
			Title     template.HTML
			IpAddress string
			Timestamp string
		}{
			Title:     template.HTML("Oracle &verbar; Go Coding Challenge"),
			IpAddress: getIP,
			Timestamp: time.Now().Format(time.RFC850),
		}
		err := templates.ExecuteTemplate(w, "base", &b)
		if err != nil {
			http.Error(w, fmt.Sprintf("index: couldn't parse template: %v", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

// public serves static assets such as CSS and JavaScript to clients.
func public() http.Handler {
	return http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))
}

// func getIP(r *http.Request) string {
//     IPAddress := r.Header.Get("X-Real-Ip")
//     if IPAddress == "" {
//         IPAddress = r.Header.Get("X-Forwarded-For")
//     }
//     if IPAddress == "" {
//         IPAddress = r.RemoteAddr
//     }
//     return IPAddress
// }

func main() {
	mux := http.NewServeMux()
	mux.Handle("/public/", logging(public()))
	mux.Handle("/", logging(index()))

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8083"
	}


	addr := fmt.Sprintf(":%s", port)
	server := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	log.Println("main: running coding challenge solution", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("main: running coding challenge solution: %v\n", err)
	}
}
