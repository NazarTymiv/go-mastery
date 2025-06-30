package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var serverStartTime time.Time

type Response interface {
	Send(w http.ResponseWriter)
}

type SuccessResponse struct {
	Message string `json:"message"`
}

func (s SuccessResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (e ErrorResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(e)
}

func sendResponse(w http.ResponseWriter, response Response) {
	response.Send(w)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, SuccessResponse{Message: "pong"})
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		sendResponse(w, ErrorResponse{Error: "Missing name"})
		return
	}

	res := fmt.Sprintf("Hello, %s", name)

	sendResponse(w, SuccessResponse{Message: res})
}

type Data struct {
	Data string `json:"data"`
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if data.Data == "" {
		sendResponse(w, ErrorResponse{Error: "Missing 'data' field"})
		return
	}

	sendJSON(w, http.StatusOK, data)
}

type Status struct {
	Uptime string `json:"uptime"`
	Status string `json:"status"`
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	countedUptime := fmt.Sprintf("%.0fs", time.Since(serverStartTime).Seconds())

	sendJSON(w, http.StatusOK, Status{Uptime: countedUptime, Status: "ok"})
}

func sendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func main() {
	serverStartTime = time.Now()

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/status", statusHandler)

	fmt.Println("Server started at http://localhost:8000")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Error running server: ", err)
	}
}
