package main

import (
	"fmt"
	"net/http"
	"html/template"
	"models"
)

var posts map[string] *models.Post				//map хранит ссылки на посты

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "<h1>Hello World</h1>")
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil{
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)	// выполняем наш шаблон
}

func writeHandler(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil{
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "write", nil)	// выполняем наш шаблон
}

func savePostHandler(w http.ResponseWriter, r *http.Request){
	// читаем поля
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	post := models.NewPost(id, title, content)
	posts[post.Id] = post
}


func main() {
	fmt.Println("Listening on port: 3000")

	posts = make(map[string]*models.Post, 0)
	// если файл начинается на assets то искать в папке assets
	// обрезаем assets вначале
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/SavePost", savePostHandler)

	// вызывается на корень сайта - ResponseWriter - для вывода Response-а
	// все что мы запишем будет передано в Response
	// Request - хранит инфу по реквесту

	http.ListenAndServe(":3000", nil)


}
