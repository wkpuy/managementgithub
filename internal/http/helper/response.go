package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

func ReturnStatusCode(w http.ResponseWriter, code int) {
	if code != 200 {
		w.WriteHeader(code)
	}
}

func ReturnJsonData(w http.ResponseWriter, code int, d map[string]interface{}) {
	if code != 200 {
		w.WriteHeader(code)
	}

	json.NewEncoder(w).Encode(d)
}

func ReturnFileData(w http.ResponseWriter, r *http.Request, code int, fileName string) {
	if code != 200 {
		w.WriteHeader(code)
	}

	http.ServeFile(w, r, fileName)
}

func DecodeJSONRequest(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return data
	}

	return data
}
