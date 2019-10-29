package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	//testHttpPprof()
	testRuntimePprof()
}

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
	/* *
	# visit http://localhost:8080/debug/pprof/ in browser to view profile
	# fetch profile from web and generate pb file
	# generate pprof.___go_build_testPprof_go.samples.cpu.001.pb.gz
	go tool pprof http://loclahost:8080/debug/pprof/profile
	# analysis profile
	go tool pprof -http=:8081 /home/xzy/pprof/pprof.___go_build_testPprof_go.samples.cpu.001.pb.gz
	*/
}

// $ go build -o /tmp/___go_build_testPprof_go /home/xzy/golang/gotest/neter/httper/testPprof.go
// $ /tmp/___go_build_testPprof_go
// go tool pprof /tmp/___go_build_testPprof_go /tmp/cpuprofile
// go tool pprof /tmp/___go_build_testPprof_go /tmp/memprofile
func testRuntimePprof() {
	cpuF, err := os.OpenFile("/tmp/cpuprofile", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer cpuF.Close()
	if err := pprof.StartCPUProfile(cpuF); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	// recursive calculation of fibonacci number, time costly
	fibonacci(46)

	memF, err := os.OpenFile("/tmp/memprofile", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer memF.Close()
	runtime.GC()
	if err := pprof.WriteHeapProfile(memF); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
	// wait memprofile wrote to disk
	time.Sleep(time.Second)
}

func fibonacci(n int) int {
	if n < 2 {
		return 1
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
