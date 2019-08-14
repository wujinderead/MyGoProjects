package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	//testServeFile()
	//testServeContent()
	//testFileServer()
	//testStripPrefix()
	testHideDotFile()
}

func testServeFile() {
	http.HandleFunc("/files/txt", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		http.ServeFile(writer, request, "/home/xzy/conn.test7.pipeline.sh")
		return
	})
	http.HandleFunc("/files/pdf", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		http.ServeFile(writer, request, "/home/xzy/图片/schoof algorithm.pdf")
		return
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func testServeContent() {
	// ServeContent deal with modtime and io.ReadSeeker, so that the server can handle cache-control
	// or 'resume from break point' request, such as "If-Modified-Since" and "Range".
	// the responses may be like:
	// 304 not modified (the content hasn't been modified, browser can use cache),
	// 206 partial content (use io.ReadSeeker to randomized read partial data).
	http.HandleFunc("/content", func(w http.ResponseWriter, r *http.Request) {
		// os.File is a io.ReadSeeker that support randomized access
		filename := "/usr/local/go/src/runtime/runtime2.go"
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("open files err")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		stat, err := file.Stat()
		if err != nil {
			fmt.Println("stat files err")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.ServeContent(w, r, "runtime2.go", stat.ModTime(), file)
	})
	http.HandleFunc("/content/somefile", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "src/neter/httper/index.html")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
	// http :8080/content 'Range:bytes=0-100,200-300'
	// http :8080/content 'If-Modified-Since:Tue, 11 Jun 2019 17:27:51 GMT'
}

func testFileServer() {
	// file server can only handle at "/"
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("get user home err:", err)
		return
	}
	fmt.Println("user home:", home)
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(home))))
}

func testStripPrefix() {
	// use StripPrefix to serve file system in alternative path
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("get user home err:", err)
		return
	}
	fmt.Println("user home:", home)
	// on linux, pattern must be "/home/", prefix can be "/home" or "/home/"
	// when request a directory "/home" it will response 301 to redirect to "/home/" instead
	http.Handle("/home/", http.StripPrefix("/home", http.FileServer(http.Dir(home))))
	http.Handle("/goroot/", http.StripPrefix("/goroot/", http.FileServer(http.Dir("/usr/local/go"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func testHideDotFile() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("get user home err:", err)
		return
	}
	fmt.Println("user home:", home)
	// http :8080/expose/golang/.git/info/         200 OK
	// http :8080/hide/golang/.git/info/           403 Forbidden
	http.Handle("/expose/", http.StripPrefix("/expose", http.FileServer(http.Dir(home))))
	http.Handle("/hide/", http.StripPrefix("/hide/", http.FileServer(dotHideFileSystem{http.Dir(home)})))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// use a new struct contains the interface to override some methods the interface
type dotHideFileSystem struct {
	http.FileSystem
}

func (s dotHideFileSystem) Open(name string) (http.File, error) {
	if containsDotFile(name) { // neither file name nor file path can has "."
		return nil, os.ErrPermission
	}
	f, err := s.FileSystem.Open(name)
	if err != nil {
		return nil, os.ErrNotExist
	}
	return dotHideFile{f}, nil
}

type dotHideFile struct {
	http.File
}

func (f dotHideFile) Readdir(n int) ([]os.FileInfo, error) {
	infos, err := f.File.Readdir(n)
	if err != nil {
		return nil, err
	}
	nodot := make([]os.FileInfo, 0)
	for i := range infos {
		if !strings.HasPrefix(infos[i].Name(), ".") {
			nodot = append(nodot, infos[i])
		}
	}
	return nodot, nil
}

func containsDotFile(name string) bool {
	parts := strings.Split(name, "/")
	for _, part := range parts {
		if strings.HasPrefix(part, ".") {
			return true
		}
	}
	return false
}
