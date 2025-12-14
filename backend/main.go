package main

import (
	"antrianrs/internal/config"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Gagal konek DB:", err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Halo,  Backend Jalan!"))
	})

	log.Println("Server jalan di port :8080")
	http.ListenAndServe(":8080", r)
}
