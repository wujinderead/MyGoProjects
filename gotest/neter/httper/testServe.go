package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	//testHandleServe()
	testNotFoundTimeout()
}

func testHandleServe() {
	// http get :8080/hello?name=aaa
	// http -f post :8080/hello name=aaa
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet && request.Method != http.MethodPost {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		// parse parameter in url, only parse body as form when request is post, put, or patch
		_ = request.ParseForm()
		name := request.Form.Get("name")
		if name == "" {
			writer.WriteHeader(http.StatusBadRequest)
			_, _ = writer.Write([]byte("please provide 'name' parameter."))
			return
		}
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte("hello, "))
		_, _ = writer.Write([]byte(name))
		return
	})
	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("index"))
		return
	}
	// http get :8080/
	http.Handle("/", http.HandlerFunc(indexHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func testNotFoundTimeout() {
	// response '404 Not Found'
	http.HandleFunc("/notfound", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("haha not found. "))
		http.NotFound(w, r)
	}))
	// a time-consuming handler
	timeConsumingHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		workTime := rand.Intn(10)
		log.Println("work time:", workTime)
		time.Sleep(time.Duration(workTime) * time.Second)
		_, _ = w.Write([]byte("done in "))
		_, _ = w.Write([]byte{byte(workTime) + '0'})
		_, _ = w.Write([]byte(" seconds."))
	})
	// timeout handler response '503 Service Unavailable' if handler's process time exceeds the specified timeout
	http.Handle("/timeout", http.TimeoutHandler(timeConsumingHandler, 5*time.Second, "timeout"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func testProxy() {
	http.ProxyFromEnvironment()
}
