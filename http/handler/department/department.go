package department

import "net/http"


func AllStudent() http.HandlerFunc{
	return  func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("fetching all departments"))
	}
	
}