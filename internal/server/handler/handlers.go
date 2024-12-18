package handler

import (
	"dynamic-qris-generator/internal/qris"
	"encoding/json"
	"log"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func GenerateQRISHandler(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Data string `json:"data"`
	}
	var req = request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := qris.DataToQrisDynamic(req.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(res)
}

func ReadQRISHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Data string `json:"data"`
	}
}
