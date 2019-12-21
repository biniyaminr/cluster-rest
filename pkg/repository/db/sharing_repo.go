package db

import (
	"database/sql"
	"github.com/karamazovian/cluster-rest/pkg/sharing"
)

type SharingRepoPSQL struct {
	conn *sql.DB
}

func NewSharingRepoPSQL(conn *sql.DB) *SharingRepoPSQL {
	return &SharingRepoPSQL{conn: conn}
}

//FetchPublicEntries returns an array of the specified number fo public entries
func (s SharingRepoPSQL) FetchPublicEntries(offset, limit int, sortBy string, order sharing.SortOrder) ([]sharing.PublicEntry, error) {
	var orderBy string
	var entries []sharing.PublicEntry
	switch order {
	case sharing.ASC:
		orderBy = "ASC"
	case sharing.DEC:
		orderBy = "DESC"
	}
	results, err := s.conn.Query("SELECT * FROM public_entries ORDER BY shared_date $1 LIMIT $2 OFFSET $3", orderBy, limit, offset)
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

//FetchComments returns an array of a number of comments on a public entry
func (s SharingRepoPSQL) FetchComments(entryID int, offset, limit int) ([]sharing.Comment, error) {
	var comments []sharing.Comment
	//The limit will be 50 by default
	if limit == 0 {
		limit = 50
	}
	results, err := s.conn.Query("SELECT * FROM comments WHERE entry_id=$1 ORDER BY commented_date LIMIT $2 OFFSET $3", entryID, limit, offset)
	if err != nil {
		return comments, err
	}
	for results.Next() {
		//var ignore string
		comment := sharing.Comment{}
		//TODO remember that you have to set parseTime=true in the connection string of the db
		err = results.Scan(&comment.PostedBy, nil, &comment.Contents, &comment.CommentedOn)
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
