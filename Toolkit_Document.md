# 📘 Capstone Project: GoTrade Sentinel (Beginner’s Toolkit for Go)

## 1. Title & Objective
**Title:** "GoTrade Sentinel: Building a Real-Time Crypto Dashboard with Go"

**Objective:**
To learn the Go programming language (Golang) by building a high-performance microservice that fetches, processes, and displays real-time cryptocurrency prices. The goal was to move beyond simple scripts and create a scalable "Micro-Frontend" application.

**Why Go?**
I chose Go because it is the industry standard for high-frequency trading and fintech systems. Its ability to handle concurrent tasks (using Goroutines) and its strict type safety makes it ideal for financial data processing where speed and accuracy are critical.

## 2. Quick Summary of the Technology
**What is it?**
Go is a statically typed, compiled programming language designed by Google. It combines the speed of C++ with the readability of Python.

**Key Features:**
**Goroutines:** Lightweight threads for multitasking (ideal for watching thousands of markets).
**Strict Typing:** Prevents data errors (critical for money apps).
**Standard Library:** Includes a production-ready Web Server (`net/http`) without needing heavy frameworks.

**Real-World Example:**
Uber uses Go to handle their massive volume of geofencing and dispatch requests.

## 3. System Requirements
OS:** Windows 10/11 (AMD64)
**Language:** Go Version 1.25.5
**Editor:** VS Code with the official "Go" extension.
**Terminal:** Windows PowerShell / Command Prompt.

## 4. Installation & Setup Instructions
**Step 1: Install Go**
Downloaded the installer from [go.dev/dl](https://go.dev/dl) and ran it.

**Step 2: Initialize Project**
Go requires a module file to manage dependencies. I created a folder and ran:
```bash
mkdir gotrade
cd gotrade
go mod init gotrade
go version
# Output: go version go1.25.5 windows/amd64
Minimal Working Example 
Before building the full dashboard, I created a simple "Hello World" API to test the server configuration.

Code (main.go - Initial Version):
package main

import (
    "fmt"
    "net/http"
)

func priceHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    // Minimal JSON response
    fmt.Fprintf(w, `{"symbol": "BTCUSD", "price": 102500.50}`)
}

func main() {
    http.HandleFunc("/", priceHandler)
    fmt.Println("Server running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
Running go run main.go starts the server. Visiting localhost:8080 in the browser displays the JSON data.
#Advancing my prototype
I scaled the project into a Real-Time Dashboard.
Features: Fetches live data for 6 currency pairs (BTC, ETH, XRP, SOL, ADA, DOGE).
Tech Stack: Go Backend (API Fetching) + HTML/JS Frontend (Dark Mode Dashboard).
Architecture: The Go server acts as both the API Gateway and the Static File Server.
## 6. The "Advanced" Prototype 
After the minimal example worked, I scaled the project into a **Robust Real-Time Dashboard**.
* **Features:** Fetches live data for 6 currency pairs (BTC, ETH, XRP, SOL, ADA, DOGE).
* **Resilience:** Implemented error handling logic. If the API fails or the internet disconnects, the dashboard switches to "OFFLINE" mode instead of crashing.
* **Architecture:** The Go server acts as both the API Gateway and the Static File Server.
### Issue 3: API Rate Limiting (429 Error)
**The Problem:**
After leaving the dashboard running, the terminal started showing `API returned bad status: 429 Too Many Requests` and the prices stopped updating.

**The Diagnosis:**
I realized I was "spamming" the CoinGecko API. I had set the refresh interval to 3 seconds, but the free API tier limits users to roughly 10-15 requests per minute.

**The Fix:**
I modified the frontend (`index.html`) to increase the polling interval from 3 seconds to 20 seconds.
```javascript
// Old Code: setInterval(updatePrices, 3000);
// New Code: setInterval(updatePrices, 15000);
