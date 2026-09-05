package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NurulIslam17/students-api/http/handler/student"
	"github.com/NurulIslam17/students-api/internal/config"
)

func main() {
	//Load Config
	cfg := config.MustLoad()
	//Database Setup

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("Get /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Student API route working..."))
	})

	router.HandleFunc("GET /api/students", student.New().ServeHTTP)

	//setup server
	sever := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Println("Server started!", cfg.Addr)
	err := sever.ListenAndServe()
	log.Fatal(err)
	if err != nil {
		log.Fatal("failed to start server")
	}

}
