package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	data := `{
        "Athlete": "DEMTSCHENKO, Albert",
        "City": "Turin",
        "Country": "RUS",
        "Discipline": "Luge",
        "Event": "Singles",
        "Gender": "Men",
        "Medal": "Silver",
        "Season": "winter",
        "Sport": "Luge",
        "Year": 2006
    }`
	req, err := http.NewRequest("POST", "http://localhost:4080/api/games3/_doc", strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complex123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func chi() {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))

	})
	http.ListenAndServe(":3333", r)

}
