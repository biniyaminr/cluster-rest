package listing

type FeedFetcher interface {
	FetchFeed(link string) (Feed, error)
	FetchFeedInformation(link string) (FeedInformation, error)
	CheckSupported(link string) (bool, error)
}
