// package main

// import (
// 	"log"
// 	"net/http"
// 	"example.com/myapi/handlers"
// 	"github.com/gorilla/mux"
// )

// func main(){
// 	r := mux.NewRouter()

// 	// http.HandleFunc("/", helloHandler)
// 	r.HandleFunc("/hello", handlers.HelloHandler)
// 	r.HandleFunc("/article", handlers.PostArticleHandler)
// 	r.HandleFunc("/article/list", handlers.ArticleListHandler)
// 	r.HandleFunc("/article/{id:[0~9]+}", handlers.ArticleDetailHandler, handlers.ArticleDetailHandler).Method(http.MethodGet)
// 	r.HandleFunc("/article/nice", handlers.PostNiceHandler)
// 	r.HandleFunc("/comment", handlers.PostCommentHandler)

	
// 	log.Println("sever start at port 8080")
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }

package main

import (
	"log"
	"net/http"
	"example.com/myapi/handlers"
	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}