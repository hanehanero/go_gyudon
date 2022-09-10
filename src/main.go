package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	word := callGyudon()
	t, _ := template.ParseFiles("template.html")
	err := t.Execute(os.Stdout, word)
	if err != nil {
		log.Fatalln(err)
	}
}

func callGyudon() string {
	myshop := NewGyudon()
	gochisousama, err := myshop.Eat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot eat : '$s'\n", err)
	}
	return gochisousama
}

type Gyudon struct {
	menu string
}

func NewGyudon() Gyudon {
	return Gyudon{
		menu: "NegitamaGyudon",
	}
}

func (self *Gyudon) Eat() (string, error) {
	if self.menu == "" {
		return "", fmt.Errorf("name is empty.")
	}

	time.Sleep(time.Second * 10)
	//fmt.Println(self.menu)
	return self.menu, nil
}
