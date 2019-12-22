package article

import (
	"github.com/karamazovian/cluster-rest/pkg/domain/article_handling"
	"github.com/thatguystone/swan"
)

type BasicArticleFetcher struct{}

func NewBasicArticleFetcher() *BasicArticleFetcher {
	return &BasicArticleFetcher{}
}

//FetchArticleContents fetches the basic contents of the article from the given url
func (baf *BasicArticleFetcher) FetchArticleContents(url string) (article_handling.ArticleContents, error) {
	content, err := swan.FromURL(url)
	if err != nil {
		return article_handling.ArticleContents{}, err
	}
	articleContents := article_handling.ArticleContents{
		Title:        content.Meta.Title,
		ArticleImage: content.Img.Src,
		Authors:      content.Meta.Authors,
		Description:  content.Meta.Description,
		TextContent:  content.CleanedText,
	}
	return articleContents, err
}

func (baf *BasicArticleFetcher) FetchMultipleArticlesContents(urls ...string) (map[string]article_handling.ArticleContents, error) {
	contents := map[string]article_handling.ArticleContents{}
	for _, url := range urls {
		currentArticle, err := baf.FetchArticleContents(url)
		if err != nil {
			return contents, err
		}
		contents[url] = currentArticle
	}
	return contents, nil
}

//FetchArticleImage fetches a uri to a teaser image for the article from the given url
func (baf *BasicArticleFetcher) FetchArticleImage(url string) (string, error) {
	content, err := swan.FromURL(url)
	if err != nil {
		return "", err
	}
	return content.Img.Src, err
}

//FetchMultipleArticleImages fetches the images of multiple articles and returns a map[articleURL]imageURL
func (baf *BasicArticleFetcher) FetchMultipleArticleImages(urls ...string) (map[string]string, error) {
	imageURLs := map[string]string{}
	for _, url := range urls {
		currentImage, err := baf.FetchArticleImage(url)
		if err != nil {
			return imageURLs, err
		}
		imageURLs[url] = currentImage
	}
	return imageURLs, nil
}
