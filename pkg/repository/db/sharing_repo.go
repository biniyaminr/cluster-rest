package db

import (
	"database/sql"
	"github.com/karamazovian/cluster-rest/pkg/domain/sharing"
	"time"
)

type SharingRepoPSQL struct {
	conn *sql.DB
}

func NewSharingRepoPSQL(conn *sql.DB) *SharingRepoPSQL {
	return &SharingRepoPSQL{conn: conn}
}

//FetchPublicEntries returns an array of the specified number fo public entries
func (s SharingRepoPSQL) FetchPublicEntries(offset, limit int, sortBy sharing.SortBy, order sharing.SortOrder) ([]sharing.PublicEntry, error) {
	var entries []sharing.PublicEntry
	results, err := s.conn.Query("SELECT * FROM public_entries ORDER BY $1 $2 LIMIT $3 OFFSET $4", sortBy, order, limit, offset)
	if err != nil {
		return entries, err
	}
	for results.Next() {
		entry := sharing.PublicEntry{}
		//TODO remember that you have to set parseTime=true in the connection string of the db
		err = results.Scan(&entry.EntryID, &entry.ArticleID, &entry.SharedBy, &entry.SharedDate, &entry.LikesCount)
		entries = append(entries, entry)
	}
	return entries, err
}

func (s SharingRepoPSQL) FetchPublicEntry(entryID int) (sharing.PublicEntry, error) {
	entry := sharing.PublicEntry{}
	query := "SELECT * FROM public_entries WHERE entry_id=$1"
	err := s.conn.QueryRow(query, entryID).Scan(&entry.EntryID, &entry.ArticleID, &entry.SharedBy, &entry.SharedDate, &entry.LikesCount)
	return entry, err
}

//FetchComments returns an array of a number of comments on a public entry
func (s SharingRepoPSQL) FetchComments(entryID int, offset, limit int) ([]sharing.Comment, error) {
	var comments []sharing.Comment
	//The limit will be 50 by default
	if limit == 0 {
		limit = 50
	}
	query := "SELECT * FROM comments WHERE entry_id=$1 ORDER BY commented_date LIMIT $2 OFFSET $3"
	results, err := s.conn.Query(query, entryID, limit, offset)
	if err != nil {
		return comments, err
	}
	for results.Next() {
		//var ignore string
		comment := sharing.Comment{}
		//TODO remember that you have to set parseTime=true in the connection string of the db
		err = results.Scan(&comment.CommentID, &comment.PostedBy, nil, &comment.Contents, &comment.CommentedOn)
		if err != nil {
			return comments, nil
		}
		comments = append(comments, comment)
	}
	return comments, err
}

//FetchArticle returns an instance of an Article with the specified ID
func (s SharingRepoPSQL) FetchArticle(articleID int) (sharing.Article, error) {
	article := sharing.Article{}
	result, err := s.conn.Query("SELECT * FROM articles WHERE article_id=$1", articleID)
	if err != nil {
		return article, err
	}
	for result.Next() {
		err = result.Scan(&article.ArticleID, &article.SourceRSS, &article.Title, &article.Link)
	}
	return article, err
}

//AddArticle inserts a new article and returns its id
func (s SharingRepoPSQL) AddArticle(article *sharing.Article) (int, error) {
	var newID int
	query := "INSERT INTO TABLE articles (source_rss, title, link) VALUES($1, $2, $3) RETURNING article_id"
	err := s.conn.QueryRow(query, article.SourceRSS, article.Title, article.Link).Scan(&newID)
	article.ArticleID = newID
	return newID, err
}

//AddPublicEntry adds a single public entry into the db and returns its id
func (s SharingRepoPSQL) AddPublicEntry(entry *sharing.PublicEntry) (int, error) {
	var newID int
	query := "INSERT INTO TABLE public_entries (article_id, shared_by, shared_date, likes_count) VALUES($1, $2, $3, $4)"
	err := s.conn.QueryRow(query, entry.ArticleID, entry.SharedBy, entry.SharedDate, 0).Scan(&newID)
	entry.EntryID = newID
	return newID, err
}

//AddRecommendation adds a recommendation into thd db
func (s SharingRepoPSQL) AddRecommendation(recommendation *sharing.Recommendation) error {
	query := "INSERT INTO TABLE recommended_articles (article_id, sender_username, receiver_username, message)"
	_, err := s.conn.Query(query, recommendation.ArticleID, recommendation.SenderUsername, recommendation.ReceiverUsername, recommendation.Message)
	return err
}

//AddLike adds a like from a user to a public entry
func (s SharingRepoPSQL) AddLike(entryID int, username string) error {
	query := "INSERT INTO TABLE likes (username, entry_id, liked_date) VALUES ($1, $2, $3)"
	err := s.conn.QueryRow(query, username, entryID, time.Now()).Scan()
	return err
}

//DeleteLike deletes a like
func (s SharingRepoPSQL) DeleteLike(entryID int, username string) error {
	query := "DELETE FROM likes WHERE entry_id=$1 AND username=$2"
	err := s.conn.QueryRow(query, entryID, username).Scan()
	return err
}

//AddComment adds a comment to a public entry
func (s SharingRepoPSQL) AddComment(comment sharing.Comment, entryID int) error {
	query := "INSERT INTO comments (username, entry_id, comment, commented_date) VALUES($1, $2, $3, $4)"
	err := s.conn.QueryRow(query, comment.PostedBy, entryID, comment.Contents, comment.CommentedOn).Scan()
	return err
}

//DeleteComment deletes a comment from a public entry
func (s SharingRepoPSQL) DeleteComment(entryID, commentID int) error {
	query := "DELETE FROM comments WHERE comment_id=$1 AND entry_id=$2"
	err := s.conn.QueryRow(query, commentID, entryID).Scan()
	return err
}
