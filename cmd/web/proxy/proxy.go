package proxy

import (
	"net/http"
)

func StartProxy(port, ginPort, fiberPort string) {
	// ginProxy := newProxy("http://localhost:" + ginPort)
	// fiberProxy := newProxy("http://localhost:" + fiberPort)

	mux := http.NewServeMux()

	// ------------------------------
	// MIGRATION ROUTING RULES: EDIT HERE
	// ------------------------------

	// // New or migrated routes → Fiber
	// mux.Handle("/fiber/", fiberProxy)

	// // Everything else remains in Gin
	// mux.Handle("/", ginProxy)

	// mux.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
	// 	println("[PROXY] → FIBER:", r.URL.Path)
	// 	fiberProxy.ServeHTTP(w, r)
	// })

	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	println("[PROXY] → GIN:", r.URL.Path)
	// 	ginProxy.ServeHTTP(w, r)
	// })

	// ------------------------------

	println("[Proxy] Listening on :" + port)
	http.ListenAndServe(":"+port, mux)
}
