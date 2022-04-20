package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {

	daemon := flag.Bool("daemon", false, "daemon flag")
	flag.Parse()

	if *daemon {
		fmt.Println("daemon")
		cmd := exec.Command("/bin/bash", "-c", os.Args[0])
		err := cmd.Start()
		if err != nil {
			log.Fatalf("daemon http: %v", err)
		}
		return
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("home"))
	})

	server := http.Server{
		Addr:    ":9900",
		Handler: mux,
	}
	log.Printf("start http at %v", ":9900")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("start http: %v", err)
	}
}
