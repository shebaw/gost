package main

import (
  "fmt"
	"log"
	"net/http"
)

func logHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func cacheHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    r.Header.Del("If-None-Match")
    r.Header.Del("If-Range")
    r.Header.Del("If-Modified-Since")

    w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
    if r.Proto == "HTTP/1.0" {
      w.Header().Set("Pragma", "no-cache")
    }

		handler.ServeHTTP(w, r)
	})
}

func corsHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    origin := r.Header.Get("Origin")

    if r.Method == "OPTIONS" {
      method := r.Header.Get("Access-Control-Request-Method")
      headers := r.Header.Get("Access-Control-Request-Headers")

      if len(origin) == 0 || len(method) == 0 {
        msg := fmt.Sprintf("%d %s: missing required CORS headers",
                           http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
        http.Error(w, msg, http.StatusBadRequest)
        return;
      }

      w.Header().Add("Vary", "Origin")
    	w.Header().Add("Vary", "Access-Control-Request-Method")
    	w.Header().Add("Vary", "Access-Control-Request-Headers")
      w.Header().Set("Access-Control-Allow-Origin", origin)
      w.Header().Set("Access-Control-Allow-Credentials", "true")
      w.Header().Set("Access-Control-Allow-Methods", method)
      w.Header().Set("Access-Control-Allow-Headers", headers)
    } else {
      if len(origin) > 0 {
        w.Header().Add("Vary", "Origin")
        w.Header().Set("Access-Control-Allow-Origin", origin)
        w.Header().Set("Access-Control-Allow-Credentials", "true")
      }
  		handler.ServeHTTP(w, r)
    }
	})
}
