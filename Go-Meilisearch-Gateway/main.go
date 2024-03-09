package main

import (
	"fmt"
	"github.com/meilisearch/meilisearch-go"
	"net/http"
	"os"
)

var meilisearchMasterKey = os.Getenv("MEILISEARCH_V_1_MASTER_KEY")
var meilisearchUrl = os.Getenv("MEILISEARCH_URL")

var client = meilisearch.NewClient(meilisearch.ClientConfig{
	Host:   "http://127.0.0.1:7700",
	APIKey: meilisearchMasterKey,
})
var index = client.Index("speeds")

func main() {
	address := "127.0.0.1:8080"

	http.HandleFunc("/meilisearch/", meilisearchHandler)
	server := &http.Server{
		Addr:    address,
		Handler: http.DefaultServeMux,
	}

	fmt.Printf("HTTP Server listening on %s...\n", address)

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func meilisearchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n\n------")
	queryParams := r.URL.Query()
	q := queryParams.Get("q")
	if q == "" {
		err := fmt.Errorf("Missing 'q' parameter")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		fmt.Println("q: ", q)
	}
	fmt.Println("query params:", queryParams)
	fmt.Printf("request type: %T,\nrequest value: %v\n", r, r)
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	searchRes, err := index.Search("test",
		&meilisearch.SearchRequest{
			Limit: 10,
		})
	fmt.Println("\n---------\nmeiliserach results:", searchRes)
	fmt.Println("meiliearch err:", err)
}
