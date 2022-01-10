package adapter

import "net/http"

type HTTPMaker interface {
	Do(w http.ResponseWriter, r *http.Request)
}
