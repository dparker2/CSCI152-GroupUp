package status

import "net/http"

// Index handles thsdoucndsnc
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("V1 status is live!"))
}
