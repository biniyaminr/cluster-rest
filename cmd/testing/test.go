package main

import (
	"database/sql"
	"fmt"
	"github.com/karamazovian/cluster-rest/pkg/repository/db"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	//Use this to test
	//fetcher := article.NewBasicArticleFetcher()
	//lister := listing.NewArticleListingService(fetcher)
	//listingHandler := rest.NewListingHandler(lister)
	//
	//router := mux.NewRouter()
	//router.HandleFunc("/article", listingHandler.FetchArticleContents).Methods("GET")
	//log.Fatal(http.ListenAndServe(":8181", router).Error())

	dbConn, err := sql.Open("postgres", "dbname=rss_manager_db user=postgres password=root sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}
	thing := db.NewSharingRepoPSQL(dbConn)
	abc, err := thing.FetchArticle(1)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(abc.ArticleID)
	fmt.Println(abc.Link)
	fmt.Println(abc.Title)
	fmt.Println(abc.SourceRSS)

}
