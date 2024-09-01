package main

import (
	"net/http"
)

func redirectTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		switch true {
		case r.URL.Path == "/":
			next.ServeHTTP(w, r.WithContext(ctx))
		case r.URL.Path[len(r.URL.Path)-1:] == "/":
			http.Redirect(w, r.WithContext(ctx), r.URL.Path[:len(r.URL.Path)-1], http.StatusSeeOther)
		default:
			next.ServeHTTP(w, r)
		}
	})

}
