package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	Version = "dev"
)

func main() {
	log.Printf("version %s", Version)
	http.HandleFunc("/", open)
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "version %s", Version)
	})
	http.ListenAndServe(":8080", nil)

}

func open(w http.ResponseWriter, r *http.Request) {
	if kontrol(w, r) {
		// ? Authorized
		urlPart := strings.Split(r.URL.Path[1:], "/")
		if urlPart[1] == "lock" {
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
		} else if urlPart[1] == "unlock" {
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
		} else if urlPart[1] == "open" {
			url := r.URL.Query()["url"][0]
			cmd := exec.Command("browse", url)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			if err := cmd.Run(); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fmt.Fprintf(w, "URL OK")
			return
		} else if urlPart[1] == "plus" {
			cmd := exec.Command("amixer", "-D", "pulse", "sset", "Master", "5%+")
			var outb, errb bytes.Buffer
			cmd.Stdout = &outb
			cmd.Stderr = &errb
			if err := cmd.Run(); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			r := regexp.MustCompile("\\[(.*?)\\]")
			fmt.Println(outb.String(), errb.String())
			fmt.Fprintf(w, "Volume UP %s", r.FindStringSubmatch(outb.String())[1])
			return
		} else if urlPart[1] == "minus" {
			cmd := exec.Command("amixer", "-D", "pulse", "sset", "Master", "5%-")
			var outb, errb bytes.Buffer
			cmd.Stdout = &outb
			cmd.Stderr = &errb
			if err := cmd.Run(); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			r := regexp.MustCompile("\\[(.*?)\\]")
			fmt.Println(outb.String(), errb.String())
			fmt.Fprintf(w, "Volume DOWN %s", r.FindStringSubmatch(outb.String())[1])
			return
		}
	}
	w.WriteHeader(401)
}
func kontrol(w http.ResponseWriter, r *http.Request) bool {
	urlPart := strings.Split(r.URL.Path[1:], "/")

	if urlPart[0] == "0b63f3a3-2d28-423e-90f4-da7af27b83f5" {
		return true
	}
	return false
}
