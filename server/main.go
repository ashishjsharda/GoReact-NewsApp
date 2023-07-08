package main

import (
	"encoding/json"
	"net/http"
)

type Article struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

type Response struct {
	Articles []Article `json:"articles"`
}

func main() {
	http.HandleFunc("/top-articles", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := http.Get("https://gnews.io/api/v4/top-headlines?token=YOUR_API_KEY&lang=en&max=5")

		var response Response
		json.NewDecoder(resp.Body).Decode(&response)

		json.NewEncoder(w).Encode(response.Articles)
	})

	http.ListenAndServe(":8080", nil)
}
