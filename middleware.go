package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checking Authentication")
			if flag {
				f(w, r) //Next middleware
			} else {
				return
			}
		}
	}
}

func loggin() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() { //Anonimous function. Is usefull when you only use it once.
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r) //Next middleware
		}
	}
}
