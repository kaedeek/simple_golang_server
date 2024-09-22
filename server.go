package main

import (
	"log"
	"net/http"
	"os/exec"
)

func RunScript(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("python3", "script.py")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	w.Write(output)
}

func main() {
	http.HandleFunc("/run", RunScript)

	log.Println("Listening on :3000...")
	http.ListenAndServe(":3000", nil)
}
