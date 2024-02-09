package restapi

import "net/http"

func Card() (string, func(http.ResponseWriter, *http.Request)) {
	return "", func(wr http.ResponseWriter, rq *http.Request) {

	}
}

func DeleteCard() (string, func(http.ResponseWriter, *http.Request)) {
	return "", func(wr http.ResponseWriter, rq *http.Request) {

	}
}
