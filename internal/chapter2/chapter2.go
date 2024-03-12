package chapter2

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

func Get(w http.ResponseWriter, r *http.Request) {
	apiURL := "http://mock:80/users"
	params := url.Values{}
	params.Set("age", "25")
	apiURL += "?" + params.Encode()
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		http.Error(w, "Creating Request Error", http.StatusInternalServerError)
		return
	}
	req.Header.Add("key", "dip")
	log.Printf("Request URL: %s\n", req.URL.String())
	log.Printf("Request Header: %s\n", req.Header.Get("key"))

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Sending Request Error", http.StatusInternalServerError)
		log.Printf("Error: %+v\n", err)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(http.StatusOK)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: %+v\n", err)
		return
	}
	log.Printf("Response Body: %s\n", string(body))

	w.Write(body)
}
