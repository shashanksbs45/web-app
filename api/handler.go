package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
		return
	}

	// Optionally, unmarshal the JSON body into a struct
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Error parsing JSON body",
			http.StatusInternalServerError)
		return
	}
	// You can now use `data` map for your application logic
	fmt.Fprintf(w, "Received: %v", data)
}
