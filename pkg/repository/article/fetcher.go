package repository

import (
	"github.com/karamazovian/Web-I-RSSManager/pkg/listing"
	"github.com/thatguystone/swan"
	"log"
)

type BasicArticleFetcher struct{}

func NewBasicArticleFetcher() *BasicArticleFetcher {
	return &BasicArticleFetcher{}
}

//FetchArticleContents fetches the basic contents of the article from the given url
func (baf *BasicArticleFetcher) FetchArticleContents(url string) (listing.ArticleContents, error) {
	content, err := swan.FromURL(url)
	if err != nil {
		log.Fatal(err.Error())
		return listing.ArticleContents{}, err
	}
	articleContents := listing.ArticleContents{
		Title:        content.Meta.Title,
		ArticleImage: content.Img.Src,
		Authors:      content.Meta.Authors,
		Description:  content.Meta.Description,
		TextContent:  content.CleanedText,
	}
	return articleContents, nil
}

//FetchArticleImage fetches a uri to a teaser image for the article from the given url
func (baf *BasicArticleFetcher) FetchArticleImage(url string) (string, error) {
	content, err := swan.FromURL(url)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	return content.Img.Src, nil
}
