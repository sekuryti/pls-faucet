package server

import (
	"fmt"
	"net/http"
)

type Redirect struct{}

func (redirect *Redirect) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	fmt.Printf("proto: %s\n", req.Header.Get("X-Forwarded-Proto"))
	next(w, req)
	// if os.Getenv("ENV") != "production" || req.Header.Get("X-Forwarded-Proto") == "https:" {
	// 	next(w, req)
	// 	return
	// }
	// // remove/add not default ports from req.Host
	// target := "https://" + req.Host + req.URL.Path
	// if len(req.URL.RawQuery) > 0 {
	// 	target += "?" + req.URL.RawQuery
	// }
	// http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}

func NewRedirect() *Redirect {
	return &Redirect{}
}
