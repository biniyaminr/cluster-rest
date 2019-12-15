package rest

import (
	"encoding/json"
	//"encoding/json"
	"fmt"
	"github.com/karamazovian/cluster-rest/pkg/listing"
	"log"
	"net/http"
)

type ListingHandler struct {
	lister listing.ArticleLister
}

func NewListingHandler(lister listing.ArticleLister) *ListingHandler {
	return &ListingHandler{lister: lister}
}

func (lh *ListingHandler) FetchArticleContents(w http.ResponseWriter, r *http.Request) {
	articleURL := r.URL.Query().Get("url")
	fmt.Println("Requested URL: " + articleURL)
	response := struct {
		Status  string            `json:"status"`
		Data    []ArticleContents `json:"data"`
		Message string            `json:"message"`
	}{}
	if articleURL == "" {
		response.Status = "Error"
		response.Data = []ArticleContents{}
		response.Message = "You need to specify a URL"
		jsonBytes, _ := json.Marshal(response)
		w.Write(jsonBytes)
		return
	}
	articleContents, err := lh.lister.FetchArticleContents(articleURL)
	contents := (ArticleContents)(articleContents)
	data := []ArticleContents{contents}
	if err != nil {
		log.Fatal("Erred when fetching article contents:\n" + err.Error())
	}

	response.Status = "200 OK"
	response.Data = data
	response.Message = " "
	//log.Fatal("Error at encoder: " + json.NewEncoder(w).Encode(response).Error())
	jsonBytes, err := json.Marshal(response)
	//jsonString := string(jsonBytes)
	w.Write(jsonBytes)
	return
}
