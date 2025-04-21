package protocol

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"kitex-multi-protocol/kitex_gen/user"
)

// HTTPHandlerImpl implements the HTTPHandler interface.
type HTTPHandlerImpl struct {
	Service user.UserService
}

// ServeHTTP processes incoming HTTP requests.
func (h *HTTPHandlerImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/UserService/GetUser" && r.Method == "POST" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		var params map[string]interface{}
		err = json.Unmarshal(body, &params)
		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}

		userID, ok := params["userID"].(float64)
		if !ok {
			http.Error(w, "Invalid userID", http.StatusBadRequest)
			return
		}

		result, err := h.Service.GetUser(context.Background(), int64(userID))
		if err != nil {
			http.Error(w, "Error processing request", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"result": result})
	} else {
		http.Error(w, "Route not found", http.StatusNotFound)
	}
}
