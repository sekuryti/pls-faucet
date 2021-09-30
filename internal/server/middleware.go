package server

import (
	"fmt"
	"net/http"
)

type Redirect struct{}

func (redirect *Redirect) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	// env := os.Getenv("ENV")
	fmt.Printf("req.URL.Scheme %s", req.URL.Scheme)
	next(w, req)
	// if env != "production" || req.URL.Scheme == "https" {
	// 	fmt.Printf("scheme %#v\n", req.Header.Get("origin"))
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
