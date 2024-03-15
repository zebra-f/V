package main

import (
	"fmt"
	"github.com/meilisearch/meilisearch-go"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var meilisearchMasterKey = os.Getenv("MEILISEARCH_V_1_MASTER_KEY")
var meilisearchUrl = os.Getenv("MEILISEARCH_URL")

var client = meilisearch.NewClient(meilisearch.ClientConfig{
	Host:   meilisearchUrl,
	APIKey: meilisearchMasterKey,
})
var index = client.Index("speeds")

func main() {
	file, err := os.Create("logfile.log")
	if err != nil {
		log.Fatal("Error creating log file:", err)
		panic(err)
	}
	defer file.Close()

	multiWriter := io.MultiWriter(file, os.Stdout)
	criticalLogger := log.New(multiWriter, "[CRITICAL] ", log.Ldate|log.Ltime)
	errorLogger := log.New(multiWriter, "[ERROR] ", log.Ldate|log.Ltime)

	_, err = index.UpdateFilterableAttributes(&[]string{"is_public"})
	if err != nil {
		fmt.Printf("Can't update filterable attributes in Meilisearch, %s\n", err)
		criticalLogger.Printf("Can't update filterable attributes in Meilisearch, %s\n", err)
		panic(err)
	}

	if len(os.Args) < 3 {
		panic("Specify mode [local/container] and port in runtime arguments.")
	}

	port := os.Args[2]
	_, err = strconv.Atoi(port)
	if err != nil {
		panic("Incorrect port argument.")
	}

	mode := os.Args[1]

	var address string
	if mode == "local" {
		address = "127.0.0.1:" + port
	} else if mode == "container" {
		address = "0.0.0.0:" + port
	} else {
		panic("Incorrect mode argument.")
	}

	http.HandleFunc("/meilisearch/", meilisearchHandler(errorLogger))
	server := &http.Server{
		Addr:    address,
		Handler: http.DefaultServeMux,
	}

	fmt.Printf("HTTP Server listening on %s...\n", address)
	if err := server.ListenAndServe(); err != nil {
		criticalLogger.Printf("Error starting server: %s\n", err)
		panic(err)
	}
}

func meilisearchHandler(errorLogger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		q := queryParams.Get("q")

		searchRes, err := index.Search(q,
			&meilisearch.SearchRequest{
				Limit:  10,
				Filter: "is_public = true",
			})
		if err != nil {
			errorLogger.Printf("Internal Server Search Error: %s\n", err)
			http.Error(w, "Internal Server Search Error", http.StatusInternalServerError)
			return
		}

		jsonSearchRes, err := searchRes.MarshalJSON()
		if err != nil {
			errorLogger.Printf("Internal Server JSON Error: %s\n", err)
			http.Error(w, "Internal Server JSON Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonSearchRes)

		currentTime := time.Now()
		fmt.Printf("[%s] Response sent successfully\n", currentTime.Format("2006-01-02 15:04:05"))
	}
}
