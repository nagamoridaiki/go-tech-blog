package repository

import (
	"go-tech-blog/model"
)

// WriterGetByID ...
func WriterGetByID(id int) (*model.Writer, error) {
	// writers テーブルから筆者データを一件取得します。
	query := `SELECT * FROM writers WHERE id = ?;`
	var writer model.Writer
	if err := db.Get(&writer, query, id); err != nil {
		return nil, err
	}

	// 筆者データの取得に成功したら、筆者 ID を基に複数の記事データを取得します。
	articles, err := ArticleListByWriterID(id)
	if err != nil {
		return nil, err
	}

	// 記事データの取得に成功したら、記事データを筆者の構造体のフィールドに格納します。
	writer.Articles = articles

	return &writer, nil
}
