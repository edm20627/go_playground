package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var tmpl = template.Must(template.New("msg").Parse("<html><body>{{.Name}}さんの運勢は「<b>{{.Omikuji}}</b>」です</body></html>"))

type Result struct {
	Name    string
	Omikuji string
}

func handler(w http.ResponseWriter, r *http.Request) {
	var result string
	switch rand.Intn(3) {
	case 0:
		result = "大吉"
	case 1:
		result = "中吉"
	case 2:
		result = "吉"
	default:
		result = "凶"
	}
	fmt.Fprint(w, result)
}

func hoge(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("p")
	var result string
	if name == "Gopher" {
		result = "大吉"
	} else {
		fmt.Fprint(w, result)
		result = "凶"
	}

	fmt.Fprint(w, name+"さんの運勢は「"+result+"」です！")
}

func fuga(w http.ResponseWriter, r *http.Request) {
	result := Result{
		Name:    r.FormValue("p"),
		Omikuji: omikuji(),
	}
	tmpl.Execute(w, result)
}

func req(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8080/fuga/?p=Gopher")
	if err != nil {
		fmt.Fprint(w, err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	fmt.Fprint(w, string(b))
}

func main() {
	// seed設定
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", handler)
	http.HandleFunc("/hoge/", hoge)
	http.HandleFunc("/fuga/", fuga)
	http.HandleFunc("/req/", req)
	http.ListenAndServe(":8080", nil)
}

func omikuji() string {
	n := rand.Intn(6) // 0-5
	switch n + 1 {
	case 6:
		return "大吉"
	case 5, 4:
		return "中吉"
	case 3, 2:
		return "小吉"
	default:
		return "凶"
	}
}
