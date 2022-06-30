package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/
var errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	// handlers
	router.HandleFunc("/name/{PARAM}", getParam).Methods(http.MethodGet)
	router.HandleFunc("/bad", errorPage).Methods(http.MethodGet)
	router.HandleFunc("/data", getData).Methods(http.MethodPost)
	router.HandleFunc("/headers", getHeader).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

// main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func getParam(w http.ResponseWriter, r *http.Request) {
	endpoint := mux.Vars(r)["PARAM"]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(endpoint))
}

func errorPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func getData(w http.ResponseWriter, r *http.Request) {
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorLog.Printf("Response body: %s", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I got message:\n" + string(resp)))
}

func getHeader(w http.ResponseWriter, r *http.Request) {
	headParam1, err := strconv.Atoi(r.Header.Get("a"))
	if err != nil {
		errorLog.Printf("Header body: %s", err)
	}
	headParam2, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		errorLog.Printf("Header body: %s", err)
	}
	w.Header().Add("a+b", strconv.Itoa(headParam1+headParam2))

	w.WriteHeader(http.StatusOK)
}
