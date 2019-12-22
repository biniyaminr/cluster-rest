package articlehandling

import "fmt"

type ArticleListingService struct {
	fetcher ArticleFetcherRepo
}

func NewArticleListingService(fetcher ArticleFetcherRepo) *ArticleListingService {
	return &ArticleListingService{fetcher: fetcher}
}

func (als *ArticleListingService) FetchArticleContents(url string) (ArticleContents, error) {
	contents := ArticleContents{}
	contents, err := als.fetcher.FetchArticleContents(url)
	if err != nil {
		fmt.Println("Printing: " + err.Error())
	}
	return contents, err
}

func (als *ArticleListingService) FetchMultipleArticleContents(urls []string) (map[string]ArticleContents, error) {
	articleContents := map[string]ArticleContents{}
	articleContents, err := als.fetcher.FetchMultipleArticlesContents(urls...)
	if err != nil {
		return articleContents, err
	}
	return articleContents, nil
}

func (als *ArticleListingService) FetchMultipleArticleImages(articleURLs []string) (map[string]string, error) {
	var imagesMap map[string]string
	imagesMap, err := als.fetcher.FetchMultipleArticleImages(articleURLs...)
	return imagesMap, err
}

func (als *ArticleListingService) FetchArticleImage(url string) (string, error) {
	imageURL, err := als.fetcher.FetchArticleImage(url)
	return imageURL, err
}
