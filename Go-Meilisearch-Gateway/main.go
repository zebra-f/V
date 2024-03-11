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

	_, err := index.UpdateFilterableAttributes(&[]string{"is_public"})
	if err != nil {
		fmt.Printf("Can't update filterable attributes in Meilisearch, %s\n", err)
		panic(err)
	}

	http.HandleFunc("/meilisearch/", meilisearchHandler)
	server := &http.Server{
		Addr:    address,
		Handler: http.DefaultServeMux,
	}

	fmt.Printf("HTTP Server listening on %s...\n", address)
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		panic(err)
	}
}

func meilisearchHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	q := queryParams.Get("q")

	searchRes, err := index.Search(q,
		&meilisearch.SearchRequest{
			Limit:  10,
			Filter: "is_public = true",
		})
	if err != nil {
		http.Error(w, "Internal Server Search Error", http.StatusInternalServerError)
		return
	}
	jsonSearchRes, err := searchRes.MarshalJSON()
	if err != nil {
		http.Error(w, "Internal Server JSON Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonSearchRes)
}
