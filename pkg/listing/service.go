package listing

type ArticleListingService struct {
	fetcher ArticleFetcher
}

func NewArticleListingService(fetcher ArticleFetcher) *ArticleListingService {
	return &ArticleListingService{fetcher: fetcher}
}

func (als *ArticleListingService) FetchArticleContents(url string) (ArticleContents, error) {
	contents := ArticleContents{}
	contents, err := als.fetcher.FetchArticleContents(url)
	if err != nil {
		return ArticleContents{}, err
	}
	return contents, nil
}

func (als *ArticleListingService) FetchMultipleArticleContents(urls ...string) (map[string]ArticleContents, error) {
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
	if err != nil {
		return imagesMap, err
	}
	return imagesMap, nil
}

func (als *ArticleListingService) FetchArticleImage(url string) (string, error) {
	imageURL, err := als.fetcher.FetchArticleImage(url)
	if err != nil {
		return "", err
	}
	return imageURL, nil
}
