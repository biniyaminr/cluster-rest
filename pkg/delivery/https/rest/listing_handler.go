package rest

import (
	"encoding/json"
	//"encoding/json"
	"fmt"
	"github.com/karamazovian/cluster-rest/pkg/domain/articlehandling"
	"net/http"
)

type ListingHandler struct {
	lister articlehandling.ArticleLister
}

func NewListingHandler(lister articlehandling.ArticleLister) *ListingHandler {
	return &ListingHandler{lister: lister}
}

func (lh *ListingHandler) FetchArticleContents(w http.ResponseWriter, r *http.Request) {
	queryURL := r.URL.Query()["url"]
	if len(queryURL[0]) != 0 {
		fmt.Println("Requested URL: " + queryURL[0])
	}
	response := struct {
		Status  string                            `json:"status"`
		Data    []articlehandling.ArticleContents `json:"data"`
		Message string                            `json:"message"`
	}{}
	if len(queryURL[0]) == 0 {
		response.Status = "Error"
		response.Data = []articlehandling.ArticleContents{}
		response.Message = "You need to specify a URL"
		_ = json.NewEncoder(w).Encode(response)
		return
	}
	articleContents, err := lh.lister.FetchArticleContents(queryURL[0])
	if err != nil {
		response.Status = "Error"
		response.Data = []articlehandling.ArticleContents{}
		response.Message = err.Error()
		_ = json.NewEncoder(w).Encode(response)
		return
	}
	contents := articleContents
	data := []articlehandling.ArticleContents{contents}
	response.Data = data
	_ = json.NewEncoder(w).Encode(response)
}
