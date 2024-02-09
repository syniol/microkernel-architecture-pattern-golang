package pkg

import "net/http"

type RestAPIPlugin = func() (string, func(http.ResponseWriter, *http.Request))
