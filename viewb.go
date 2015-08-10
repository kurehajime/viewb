package main

import (
	"flag"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var (
	port int
	com  string
)

func main() {
	//parse args
	flag.IntVar(&port, "p", 8080, "port /default:8080")
	flag.IntVar(&port, "port", 8080, "port /default:8080")
	flag.Parse()
	com = strings.Join(flag.Args(), " ")
	//start server
	fmt.Println("http://localhost" + ":" + strconv.Itoa(port))
	fmt.Println("Stop: Ctrl+C")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

//handler: command result
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	fmt.Fprint(w, cmd(com))
}

//exec command
func cmd(commandString string) string {
	var command string
	var op string
	if runtime.GOOS == "windows" {
		command = "cmd"
		op = "/c"
	} else {
		command = "/bin/sh"
		op = "-c"
	}
	out, err := exec.Command(command, op, commandString).Output()
	if err != nil {
		return string(err.Error())
	}
	return string(out)
}
