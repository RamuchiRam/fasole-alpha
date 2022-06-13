/* package main
import (
  "fmt"
  "net/http"
  "html/template"
  "github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request){
  tmpl, err := template.ParseFiles("templates/index.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  tmpl.Execute(w, "index")
}
func courses(w http.ResponseWriter, r *http.Request){
  tmpl, err := template.ParseFiles("templates/courses.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  tmpl.Execute(w, "courses")
  }
func login(w http.ResponseWriter, r *http.Request){
  tmpl, err := template.ParseFiles("templates/login.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  tmpl.Execute(w, "login")
  }
func register(w http.ResponseWriter, r *http.Request){
  tmpl, err := template.ParseFiles("templates/register.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  tmpl.Execute(w, "register")
  }
func keyscourse(w http.ResponseWriter, r *http.Request){
  tmpl, err := template.ParseFiles("templates/keyscourse.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  tmpl.Execute(w, "keyscourse")
  }
func links(w http.ResponseWriter, r *http.Request){
    tmpl, err := template.ParseFiles("templates/links.html")

    if err != nil {
      fmt.Fprintf(w, err.Error())
    }

    tmpl.Execute(w, "links")
  }
func about(w http.ResponseWriter, r *http.Request){
    tmpl, err := template.ParseFiles("templates/about.html")

    if err != nil {
      fmt.Fprintf(w, err.Error())
    }

    tmpl.Execute(w, "about")
  }
func handleFunc(){
  rtr := mux.NewRouter()
  rtr.HandleFunc("/", index)
  rtr.HandleFunc("/courses", courses)
  rtr.HandleFunc("/login", login)
  rtr.HandleFunc("/register", register)
  rtr.HandleFunc("/keyscourse", keyscourse)
  rtr.HandleFunc("/links", links)
  rtr.HandleFunc("/about", about)
  http.Handle("/", rtr)
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
  http.ListenAndServe(":8080", nil)
}

func main(){
  handleFunc()
} */
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
  "github.com/RamuchiRam/fasole-alpha/logic/repository"
  "github.com/RamuchiRam/fasole-alpha/logic/application"
)

func main() {
	ctx := context.Background()
	dbpool, err := repository.InitDBConn(ctx)
	if err != nil {
		log.Fatalf("%w failed to init DB connection", err)
	}
	defer dbpool.Close()
	a := application.NewApp(ctx, dbpool)
	r := httprouter.New()
	a.Routes(r)
	srv := &http.Server{Addr: "0.0.0.0:8080", Handler: r}
	fmt.Println("It is alive! Try http://localhost:8080")
	srv.ListenAndServe()
}
