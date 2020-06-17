package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/", unlock)
	http.ListenAndServe(":8080", nil)

}

func unlock(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[1:] == "0b63f3a3-2d28-423e-90f4-da7af27b83f5" {
		cmd := exec.Command("loginctl", "lock-session")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		if err := cmd.Run(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Fprintf(w, "LOCK OK")
		return
	} else if r.URL.Path[1:] == "fe910f0f-0ed3-4ad7-ac71-944c96aba71c" {
		cmd := exec.Command("loginctl", "unlock-session")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		if err := cmd.Run(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Fprintf(w, "Unlock OK")
		return
	}

}
