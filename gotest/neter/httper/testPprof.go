package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	testHttpPprof()
}

/*
go build -o /tmp/aaa neter/httper/testPprof.go
/tmp/aaa

# sample 10 seconds for cpu profile and then analyze it by pprof
go tool pprof -http=:8081 http://localhost:8080/debug/pprof/profile?seconds=10
# in the meantime, generate requests in another shell
ab -n 100000 -c 1000 http://localhost:8080/

# analysis heap and goroutine profiles
go tool pprof -http=:8081 http://localhost:8080/debug/pprof/heap
go tool pprof -http=:8081 http://localhost:8080/debug/pprof/goroutine

# sample 5 seconds for trace
curl -o /tmp/trace http://localhost:8080/debug/pprof/trace?seconds=5
# in the meantime, generate requests in another shell
ab -n 100000 -c 1000 http://localhost:8080/
# display trace information
go tool trace /tmp/trace
*/

func testHttpPprof() {
	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("i am index"))
		return
	}
	// http get :8080/
	http.Handle("/", http.HandlerFunc(indexHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
