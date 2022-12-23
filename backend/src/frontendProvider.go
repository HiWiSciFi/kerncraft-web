package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
)

func serve(addr string, handler http.Handler, name string) {
	var e error = nil
	for true {
		log.Println(name + ": Listening on " + addr)
		e = http.ListenAndServe(addr, handler)
		if e != nil {
			log.Println(name + ": Server on " + addr + " crashed.\n" + e.Error() + "\n" + "Restarting...")
			e = nil
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "HOME")
}

func main() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	directory := path.Dir(ex)

	frontendMux := http.NewServeMux()
	backendMux := http.NewServeMux()

	frontendMux.Handle("/", http.FileServer(http.Dir(directory+"/static")))

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	backendMux.Handle("/", router)

	go serve(":"+"3000", frontendMux, "frontend")
	serve(":"+"3001", backendMux, "backend")
}
