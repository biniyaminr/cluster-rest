package repository

import (
	"encoding/xml"
	"github.com/karamazovian/Web-I-RSSManager/pkg/listing"
	"io/ioutil"
	"log"
	"net/http"
)

type basicRssFetcher struct{}

func NewFetcher() *basicRssFetcher {
	return &basicRssFetcher{}
}

func castFeed(feed MappedFeed) listing.Feed {
	newFeed := listing.Feed{
		Title: feed.Title,
		Link:  feed.Link,
	}
	for _, article := range feed.Articles {
		newFeed.Articles = append(newFeed.Articles, listing.Article(article))
	}
	return newFeed
}

//FetchFeed fetches feed information and the contents of the feed
func (p basicRssFetcher) FetchFeed(link string) (listing.Feed, error) {
	response, err := http.Get(link)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err.Error())
		return listing.Feed{}, err
	}
	newFeed := MappedFeed{}
	err = xml.Unmarshal(bodyBytes, &newFeed)
	if err != nil {
		log.Fatal(err.Error())
		return listing.Feed{}, err
	}

	return castFeed(newFeed), nil
}

//FetchFeedInformation fetches information on the feed only
func (p basicRssFetcher) FetchFeedInformation(link string) (listing.FeedInformation, error) {
	feedInfo := MappedFeedInformation{}
	response, err := http.Get(link)
	if err != nil {
		log.Fatal(err.Error())
		return listing.FeedInformation{}, err
	}
	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err.Error())
		return listing.FeedInformation{}, err
	}
	err = xml.Unmarshal(bodyBytes, &feedInfo)
	if err != nil {
		log.Fatal(err.Error())
	}
	return (listing.FeedInformation)(feedInfo), nil
}

//CheckSupported checks whether the link provided is a supported rss feed format
func (p basicRssFetcher) CheckSupported(link string) (bool, error) {
	//TODO implement this function
	return true, nil
}
