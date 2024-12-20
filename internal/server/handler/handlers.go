package handler

import (
	"bytes"
	"dynamic-qris-generator/internal/qris"
	"encoding/json"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
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
		Data   string `json:"data"`
		Amount int    `json:"amount"`
	}
	var req = request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	img, err := qris.DataToQrisDynamic(req.Data, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		log.Println("unable to encode image.")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func ReadQRISHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Data string `json:"data"`
	}

	qrisFile, _, err := r.FormFile("qris")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	qrisData, err := qris.ExtractQris(qrisFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := response{
		Data: qrisData,
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(jsonResp)
}
