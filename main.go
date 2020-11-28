package main

import "net/http"

func main() {
	baseFrontDir := "examples/vanillajs"
	// simple static server
	// https://jeonghwan-kim.github.io/dev/2019/02/07/go-net-http.html
	http.ListenAndServe(":3000", http.FileServer(http.Dir(baseFrontDir)))

}
