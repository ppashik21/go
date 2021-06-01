package main
import ("fmt"
        "net/http")

func home_home(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello")
  }

func main() {
  http.HandleFunc("/", home_page)
  http.ListenAndServe(":80", nil)
  }
