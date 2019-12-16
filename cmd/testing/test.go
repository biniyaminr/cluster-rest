package main

import (
	"github.com/gorilla/mux"
	"github.com/karamazovian/cluster-rest/pkg/delivery/https/rest"
	"github.com/karamazovian/cluster-rest/pkg/listing"
	"github.com/karamazovian/cluster-rest/pkg/repository/article"
	"log"
	"net/http"
)

func main() {
	//Use this to test
	fetcher := article.NewBasicArticleFetcher()
	lister := listing.NewArticleListingService(fetcher)
	listingHandler := rest.NewListingHandler(lister)

	router := mux.NewRouter()
	router.HandleFunc("/article", listingHandler.FetchArticleContents).Methods("GET")
	log.Fatal(http.ListenAndServe(":8181", router).Error())

}
