package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func SendResponse(r *http.Request , w http.ResponseWriter ,  res []byte) {
	contentType := r.Header.Get("Content-type")
	// if contentType == "" {
	// 	contentType == "application/json"
	// }
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}