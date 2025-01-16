package server

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
	"github.com/scythe504/skribbler-backend/internals"
	"github.com/scythe504/skribbler-backend/internals/utils"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now().UnixMilli()
	// Create the response
	response := internals.Response{
		StatusCode:    http.StatusOK,
		RespStartTime: startTime,
		Data:          map[string]string{"status": "healthy"},
	}

	// Set response headers

	endTime := time.Now().UnixMilli()
	// Marshal the response to JSON
	response.RespEndTime = endTime

	response.NetRespTime = endTime - startTime

	jsonResp, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	// Set the end time just before writing
	w.Header().Set("Content-Type", "application/json")

	// Write status code
	w.WriteHeader(http.StatusOK)
	// Write the JSON response

	w.Write(jsonResp)
}

func GetRandomWords(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now().UnixMilli()
	// Read words from the CSV file
	response := internals.Response{
		StatusCode:    http.StatusOK,
		RespStartTime: startTime,
	}
	words := utils.ReadCsvFile("./word-list.csv")

	// If no words are available, return an error response
	if len(words) == 0 {
		http.Error(w, "No words available", http.StatusInternalServerError)
		return
	}

	// Select 3 random unique words
	selectedWords := make([]internals.Word, 0, 3)
	seenIndices := make(map[int]bool)

	for len(selectedWords) < 3 && len(seenIndices) < len(words) {
		randomIndex := rand.Intn(len(words))
		if seenIndices[randomIndex] {
			continue
		}
		seenIndices[randomIndex] = true
		selectedWords = append(selectedWords, words[randomIndex])
	}

	response.Data = selectedWords
	endTime := time.Now().UnixMilli()
	response.RespEndTime = endTime
	response.NetRespTime = endTime - startTime

	jsonResp, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}
	// Set response headers and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

