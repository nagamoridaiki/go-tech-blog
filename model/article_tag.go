package model

// ArticleTag ...
type ArticleTag struct {
	ArticleID int  `db:"article_id"`
	TagID     int  `db:"tag_id"`
	Tag       *Tag `db:"tag"`
}
