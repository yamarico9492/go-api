package handlers

import (
	"fmt"
	"io"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "helloHandler")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])

		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	}else {
			page = 1
		}

		resString := fmt.Sprintf("Article List (page%d)\n", page)
		io.WriteString(w, resString)
	}
	// io.WriteString(w, "Article List\n")

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}


// func HelloHandler(w http.ResponseWriter, req *http.Request){
// 	if req.Method == http.MethodGet{
// 		io.WriteString(w, "Hello World! \n")
// 	} else {
// 		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
// 	}
// }