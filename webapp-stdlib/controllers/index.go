package controllers

import (
	"github.com/wintondeshong/Go-Sandbox/webapp-stdlib/lib"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	lib.GenerateHTML(w, nil, "layout", "index")
	return
}
