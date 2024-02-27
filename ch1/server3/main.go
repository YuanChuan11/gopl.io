// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// !+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	// 打印请求信息到浏览器
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	// 打印URL的所有类型到浏览器
	fmt.Fprintf(w, "Url\n")
	fmt.Fprintf(w, "%s\n", r.URL)
	fmt.Fprintf(w, "Url type:%T\n", r.URL)
	fmt.Fprintf(w, "Url String:%q\n", r.URL.String())
	fmt.Fprintf(w, "Url.Path:%q\n", r.URL.Path)
	fmt.Fprintf(w, "Url.Scheme:%q\n", r.URL.Scheme)
	fmt.Fprintf(w, "Url.RawQuery:%q\n", r.URL.RawQuery)
	fmt.Fprintf(w, "Url.Fragment:%q\n", r.URL.Fragment)
	fmt.Fprintf(w, "Url.Host:%q\n", r.URL.Host)
	fmt.Fprintf(w, "Url.User:%q\n", r.URL.User)

	fmt.Fprintf(w, "Body\n")
	data := make([]byte, r.ContentLength)
	r.Body.Read(data)
	defer r.Body.Close()
	fmt.Fprintf(w, "Body = %q\n", string(data))

	fmt.Fprintf(w, "Header\n")
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	fmt.Fprintf(w, "Form\n")
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

//!-handler
