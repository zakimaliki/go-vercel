package handler

import (
	"go-vercel/api/src/routes"
	"net/http"

	. "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()
	routes.Router(server)
	server.Handle(w, r)
}
