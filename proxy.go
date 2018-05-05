package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "New function")
	})

	u1, _ := url.Parse("http://localhost:8085/")
	http.Handle("gospot.mchampaneri.in//", httputil.NewSingleHostReverseProxy(u1))

	u2, _ := url.Parse("http://localhost:8081/")
	http.Handle("temp.mchampaneri.in/", httputil.NewSingleHostReverseProxy(u2))

	u3, _ := url.Parse("http://localhost:8089/")
	http.Handle("www.ytdownload.xyz/", httputil.NewSingleHostReverseProxy(u3))

	// Start the server
	http.ListenAndServe(":80", nil)
}
