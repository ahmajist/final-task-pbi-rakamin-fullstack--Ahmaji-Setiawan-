package helpers

import (
	"net/http"
)

func GetHeaderValue(r *http.Request, key string) string {
	return r.Header.Get(key)
}
