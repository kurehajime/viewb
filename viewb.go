package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

var (
	port   int
	user   string
	pass   string
	com    string
	open   bool
	encode string
)

func main() {
	//parse args
	flag.IntVar(&port, "p", 8080, "port /default:8080")
	flag.BoolVar(&open, "o", false, "open web browser")
	flag.StringVar(&user, "user", "", "user (BASIC AUTH)")
	flag.StringVar(&pass, "pass", "", "pass (BASIC AUTH)")
	flag.StringVar(&encode, "e", "utf-8", "input encoding")

	flag.Parse()
	com = strings.Join(flag.Args(), " ")
	url := "http://localhost" + ":" + strconv.Itoa(port)

	//open web browser
	go func() {
		if open == true {
			time.Sleep(500 * time.Millisecond)
			switch runtime.GOOS {
			case "linux":
				exec.Command("xdg-open", url).Start()
			case "windows":
				exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
			case "darwin":
				exec.Command("open", url).Start()
			}
		}
	}()

	//start server
	fmt.Println(url)
	fmt.Println("Stop: Ctrl+C")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

}

//handler: command result
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	if auth(r) == false {
		w.Header().Set("WWW-Authenticate", `Basic realm="MY REALM"`)
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))
		return
	}
	var text string
	text, err := transEnc(cmd(com), encode)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	fmt.Fprint(w, text)
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

//basic auth
func auth(r *http.Request) bool {
	if user == "" || pass == "" {
		return true
	}
	_user, _pass, ok := r.BasicAuth()
	if ok == false {
		return false
	}
	return _user == user && _pass == pass
}

//trans encoding   ->    qiita.com/nobuhito/items/ff782f64e32f7ed95e43
func transEnc(text string, encode string) (string, error) {
	body := []byte(text)
	var f []byte

	encodings := []string{"sjis", "utf-8"}
	if encode != "" {
		encodings = append([]string{encode}, encodings...)
	}
	for _, enc := range encodings {
		if enc != "" {
			ee, _ := charset.Lookup(enc)
			if ee == nil {
				continue
			}
			var buf bytes.Buffer
			ic := transform.NewWriter(&buf, ee.NewDecoder())
			_, err := ic.Write(body)
			if err != nil {
				continue
			}
			err = ic.Close()
			if err != nil {
				continue
			}
			f = buf.Bytes()
			break
		}
	}
	return string(f), nil
}
