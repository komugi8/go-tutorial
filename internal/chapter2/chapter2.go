package chapter2

import "net/http"

func Get(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", "http://localhost:8081/users?age=25", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	req.Header.Set("key", "dip")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(http.StatusOK)
}
