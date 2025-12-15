package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SCALABILITY UPGRADE: Added slots for 4 new assets
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
	// SCALABILITY UPGRADE: Added more IDs to the URL query
	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum,ripple,solana,cardano,dogecoin&vs_currencies=usd"
	
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch prices", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var prices CryptoResponse
	json.Unmarshal(body, &prices)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prices)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/api/prices", pricesHandler)

	fmt.Println("🚀 GoTrade Dashboard running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}