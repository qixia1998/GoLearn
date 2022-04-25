package main

// 提供使用模板生成的内容
import (
	"html/template"
	"net/http"
)

func main() {
	tpl, err := template.ParseFiles("template.tpl")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.Execute(w, "John Doe")
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
