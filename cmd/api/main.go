package main

import (
	"fmt"
	"net/http"

	"github.com/MouraJr/firstgoapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API Service...")

	fmt.Println(`
┳┳┓┓ ┏┳┓  ┏┓•     ┓    ┏┓┏┓┳
┃┃┃┃┃┃┃┃  ┗┓┓┏┳┓┏┓┃┏┓  ┣┫┃┃┃
┛ ┗┗┻┛┻┛  ┗┛┗┛┗┗┣┛┗┗   ┛┗┣┛┻
                ┛          `)

	err := http.ListenAndServe("localhost:8123", r)

	if err != nil {
		log.Fatal(err)
	}
}
