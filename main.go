package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type CryptoResponse struct {
	Bitcoin  PriceData `json:"bitcoin"`
	Ethereum PriceData `json:"ethereum"`
	Ripple   PriceData `json:"ripple"`
	Solana   PriceData `json:"solana"`
	Cardano  PriceData `json:"cardano"`
	Dogecoin PriceData `json:"dogecoin"`
}

type PriceData struct {
	USD float64 `json:"usd"`
}

func pricesHandler(w http.ResponseWriter, r *http.Request) {
	// Set a timeout so the server doesn't hang forever if the API is slow
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum,ripple,solana,cardano,dogecoin&vs_currencies=usd"
	
	resp, err := client.Get(url)
	
	// ERROR 1: Network Failure (e.g., No Internet)
	if err != nil {
		fmt.Println("Error fetching from API:", err)
		http.Error(w, `{"error": "Network Error"}`, http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	// ERROR  2: Bad Status Code (e.g., API Limit Reached)
	if resp.StatusCode != 200 {
		fmt.Println("API returned bad status:", resp.Status)
		http.Error(w, `{"error": "API Error"}`, http.StatusBadGateway)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var prices CryptoResponse
	
	// ERROR 3: Bad JSON (e.g., API changed format)
	err = json.Unmarshal(body, &prices)
	if err != nil {
		http.Error(w, `{"error": "Data Parse Error"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prices)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/api/prices", pricesHandler)

	fmt.Println("🚀 GoTrade Sentinel running at http://localhost:8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if running locally
	}
	http.ListenAndServe(":"+port, nil)
}