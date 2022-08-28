package main

import (
	"github.com/gocraft/web"
	"sync"
)

func GetDocument() {

}
func GetAllDocuments() {

}
func GetDocumentById() {

}

type Handler struct {
	repo  int
	cache string
	mutex sync.WaitGroup
}

func (ctx *Handler) InitRoutes() *web.Router {
	router := web.New(ctx)
	router.Post("/api/register", Register)
	return router
}
