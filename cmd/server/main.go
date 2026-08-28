package main

import (
	"chaincenter/internal/api"
	"chaincenter/internal/service"
	"chaincenter/internal/storage"
	"chaincenter/internal/workflow"
	"log"
	"net/http"
	"os"
)

func main() {
	path := os.Getenv("CHAINCENTER_DB")
	if path == "" {
		path = "chaincenter.db"
	}
	s, e := storage.Open(path)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	a := service.NewAccountService(s)
	r := service.NewRecordService(s, a)
	h := api.NewHandler(a, r, workflow.NewEngine(r))
	log.Fatal(http.ListenAndServe(":8080", h.Routes()))
}
