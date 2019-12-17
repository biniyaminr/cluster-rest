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

func (s SharingRepoPSQL) FetchComments(entryID string) ([]sharing.Comment, error) {
	var comments []sharing.Comment
	results, err := s.conn.Query("SELECT * FROM comments WHERE entry_id=$1", entryID)
	if err != nil {
		return comments, err
	}
	for results.Next() {
		var ignore string
		comment := sharing.Comment{}
		err = results.Scan(&comment.PostedBy, &ignore, &comment.Contents, &comment.CommentedOn)
		if err != nil {
			return comments, nil
		}
		comments = append(comments, comment)
	}
	return comments, err
}
