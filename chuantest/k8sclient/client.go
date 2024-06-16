// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 19.
//!+

// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	// each request calls handler
	log.Fatal(http.ListenAndServe("0.0.0.0:31000", nil))

}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("client enter")
	resp, err := http.Get("http://192.168.113.130:32000")
	if err != nil {
		fmt.Fprintf(w, "err = %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	bodystr := string(body)
	fmt.Fprintf(w, "body = %v\n", bodystr)
	fmt.Println("client end")
}
