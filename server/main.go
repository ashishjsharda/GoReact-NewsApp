package main

type Article struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type Response struct {
	Articles []Article `json:"articles"`
}

func main() {
	
}
