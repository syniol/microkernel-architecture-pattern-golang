package restapi

import "net/http"

func Customer() (string, func(http.ResponseWriter, *http.Request)) {
	return "", func(wr http.ResponseWriter, rq *http.Request) {

	}
}

func DeleteCustomer() (string, func(http.ResponseWriter, *http.Request)) {
	return "", func(wr http.ResponseWriter, rq *http.Request) {

	}
}
