package restapi

import "net/http"

func Account() (string, func(http.ResponseWriter, *http.Request)) {
	return "", func(wr http.ResponseWriter, rq *http.Request) {

	}
}

func DeleteAccount() (string, func(wr http.ResponseWriter, rq *http.Request)) {
	return "", func(wr http.ResponseWriter, rq *http.Request) {

	}
}
