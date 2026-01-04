package bungie

import (
	"io"
	"net/http"
)

func LawlessData(w http.ResponseWriter, r *http.Request, apiKey string) {

	url := "https://www.bungie.net/Platform/Destiny2/1/Account/4611686018430407175/Character/2305843009263048442/Stats/Activities/?mode=93"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Add("x-api-key", apiKey)

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	w.Write(body)

}
