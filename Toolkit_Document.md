# Capstone Project: Beginner’s Toolkit for Go (Golang)

## 1. Title & Objective
**Title:** Building a High-Performance Fintech API with Go
**Objective:** To learn the Go programming language by building a real-time price ticker microservice.
**Why Go?:** It is the industry standard for high-frequency trading because of its speed and efficiency.

## 2. System Requirements
* **OS:** Windows 11
* **Language:** Go (Version 1.25.5)
* **Editor:** VS Code with Go Extension

## 3. Installation & Setup Instructions
1. Downloaded Go from the official website.
2. Verified installation in terminal:
   ```bash
   go version
   # Output: go version go1.25.5 windows/amd64
#  how i initialize  the Project
mkdir gotrade
cd gotrade
go mod init gotrade

### **Summary of Day 1**
1.  **Code:** You have a `main.go` file that runs a server.
2.  **Docs:** You have a `Toolkit_Document.md` with Sections 1, 2, 3, and 4 completed.

**Task:** Once you have pasted that text into your document and taken the screenshot, **Day 1 is 100% complete.** You can stop here for today.

Reply **"Done"** when you are ready to move to **Day 2 (Connecting to Real Live Forex Data)**.

# section 6
Prompt 1 (Theory):

"I am building a trading bot in Go. Explain why we use Structs to parse JSON instead of dynamic maps (like in Python)? How does this help with type safety in financial applications?"

Prompt 2 (The Scalability Logic):

"In Go, I am currently using a for loop to fetch prices. How can I use 'Goroutines' and 'Channels' to make this fetch multiple currency pairs at the same time without waiting for each one? Please show a simple architectural diagram concept."

# section 7
Common Issues & Fixes (or "Iterative Learning").

What to write: "Initially, I only fetched 2 pairs. To scale to 6, I realized I didn't need to write new functions. I simply expanded the Struct definition and the API URL parameters. This showed me how Go's strict typing makes scaling safe—if I forgot to add 'Solana' to the struct, the code wouldn't compile or would just ignore the data, preventing crashes."