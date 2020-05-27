package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	url := "https://covid-19-data.p.rapidapi.com/country?format=json&name=nigeria"

	r.HandleFunc("/api/covid", func(w http.ResponseWriter, r *http.Request) {

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("x-rapidapi-host", "covid-19-data.p.rapidapi.com")

		req.Header.Add("x-rapidapi-key", "05289646b4msh86f33f76e1807c5p169453jsnea0e6599ecc1")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()

		body, _ := ioutil.ReadAll(res.Body)

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		w.Write([]byte(body))

	})

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./build")))

	http.ListenAndServe(":8100", r)

}
