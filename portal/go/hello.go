package portal

import (
	"encoding/json"
	"net/http"
)

type helloResponse struct {
	Data string `json:"data"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	res := &helloResponse{Data: "hello"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
