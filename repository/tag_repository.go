package repository

import (
	"go-tech-blog/model"

	"github.com/jmoiron/sqlx"
)

// TagListByArticleID ...
func TagListByArticleID(articleID int) ([]*model.Tag, error) {
	// articles_tags テーブルから tag_id を取得します。
	q1 := `SELECT tag_id FROM articles_tags WHERE article_id = ?;`
	var tagIDs []int
	if err := db.Select(&tagIDs, q1, articleID); err != nil {
		return nil, err
	}

	// タグ情報を格納する変数を宣言します。
	var tags []*model.Tag

	// 記事に紐づくタグが一つもない場合は先にリターンします。
	if len(tagIDs) == 0 {
		return tags, nil
	}

	// tags テーブルからタグ情報を取得するクエリ文字列を生成します。
	q2 := `SELECT * FROM tags WHERE id IN(?);`

	// IN 句を利用するクエリを作成するには sqlx パッケージの In() 関数を利用します。
	query, args, err := sqlx.In(q2, tagIDs)
	if err != nil {
		return nil, err
	}

	// sqlx.In() 関数で生成されたクエリ文字列をパラメータを利用して SQL を実行します。
	// sqlx.Select() 関数の第三引数は可変長のパラメータを取ります。
	// args 変数はスライス型なので、...で展開して渡します。
	// 参考：https://golang.org/ref/spec#Passing_arguments_to_..._parameters
	if err := db.Select(&tags, query, args...); err != nil {
		return nil, err
	}

	return tags, nil
}

// TagListMapByArticleIDs ...
func TagListMapByArticleIDs(articleIDs []int) (map[int][]*model.Tag, error) {
	// タグ情報を格納するマップを生成します。
	// マップのキーに記事ID、バリューにタグのスライスを格納します。
	m := make(map[int][]*model.Tag)

	// 引数で渡ってきたスライスのサイズが 0 の場合は即時リターンします。
	if len(articleIDs) == 0 {
		return m, nil
	}

	// articles_tags テーブルからデータを取得します。
	q1 := `SELECT
		at.article_id AS article_id,
		at.tag_id AS tag_id,
		tags.id AS 'tag.id',
		tags.name AS 'tag.name'
	FROM articles_tags AS at
	INNER JOIN tags ON tags.id = at.tag_id
	WHERE article_id IN(?);`

	q2, args, err := sqlx.In(q1, articleIDs)
	if err != nil {
		return nil, err
	}

	var articleTagList []*model.ArticleTag
	if err := db.Select(&articleTagList, q2, args...); err != nil {
		return nil, err
	}

	// 取得したデータを map に格納し直します。
	for _, articleTag := range articleTagList {
		m[articleTag.ArticleID] = append(m[articleTag.ArticleID], articleTag.Tag)
	}

	return m, nil
}
