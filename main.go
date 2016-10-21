package main

import (
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	n := negroni.Classic()
	n.UseHandler(http.DefaultServeMux)

	http.HandleFunc("/oauth/authorize", authorizeHandler)

	n.Run(":54321")
}

func authorizeHandler(w http.ResponseWriter, req *http.Request) {
	viewsManager.MustRenderWithLayout(w, "oauth/authorize.html", "layout.html", nil)
}
