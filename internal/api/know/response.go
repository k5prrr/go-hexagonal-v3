// Тут мелкие функции, которые часто используются
package know

import (
	"encoding/json"
	"net/http"
)

const contentTypeJSON = "application/json"

func writeJSONReaponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", contentTypeJSON)
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "filed to encode response", http.StatusInternalServerError)
	}
}
