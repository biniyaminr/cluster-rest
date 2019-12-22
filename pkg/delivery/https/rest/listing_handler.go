package rest

import (
	"encoding/json"
	//"encoding/json"
	"fmt"
	"github.com/karamazovian/cluster-rest/pkg/domain/article_handling"
	"net/http"
)

type ListingHandler struct {
	lister article_handling.ArticleLister
}

func NewListingHandler(lister article_handling.ArticleLister) *ListingHandler {
	return &ListingHandler{lister: lister}
}

func (lh *ListingHandler) FetchArticleContents(w http.ResponseWriter, r *http.Request) {
	queryURL := r.URL.Query()["url"]
	if len(queryURL[0]) != 0 {
		fmt.Println("Requested URL: " + queryURL[0])
	}
	response := struct {
		Status  string                             `json:"status"`
		Data    []article_handling.ArticleContents `json:"data"`
		Message string                             `json:"message"`
	}{}
	if len(queryURL[0]) == 0 {
		response.Status = "Error"
		response.Data = []article_handling.ArticleContents{}
		response.Message = "You need to specify a URL"
		_ = json.NewEncoder(w).Encode(response)
		return
	}
	articleContents, err := lh.lister.FetchArticleContents(queryURL[0])
	if err != nil {
		response.Status = "Error"
		response.Data = []article_handling.ArticleContents{}
		response.Message = err.Error()
		_ = json.NewEncoder(w).Encode(response)
		return
	}
	contents := articleContents
	data := []article_handling.ArticleContents{contents}
	response.Data = data
	_ = json.NewEncoder(w).Encode(response)
}
