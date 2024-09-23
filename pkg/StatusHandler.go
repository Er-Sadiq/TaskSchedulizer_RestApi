package pkg

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// StatusHandler handles the /status endpoint
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method. Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Assume you pass the task ID as a query parameter
	taskIDStr := r.URL.Query().Get("id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	status, err := CheckStatus(taskID)
	if err != nil {
		http.Error(w, "Error fetching task status: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]bool{
		"status": status,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
