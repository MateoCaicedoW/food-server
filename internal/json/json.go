package json

import (
	"encoding/json"
	"log"
	"net/http"
)

type Data struct {
	Content interface{} `json:"content"`
	Status  int         `json:"status"`
	OK      bool        `json:"ok"`
}

type Error struct {
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, data Data) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data.Status)

	if !data.OK {
		log.Println(data.Content.(Error).Message)
	}

	json.NewEncoder(w).Encode(data)
}

func Decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
