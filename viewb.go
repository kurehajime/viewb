package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
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
	fmt.Fprint(w, cmd(com))
}

//exec command
func cmd(commandString string) string {
	var command string
	if runtime.GOOS == "windows" {
		command = "cmd"
	} else {
		command = os.Getenv("SHELL")
	}
	out, err := exec.Command(command, "-c", commandString).Output()
	if err != nil {
		return string(err.Error())
	}
	return string(out)
}
