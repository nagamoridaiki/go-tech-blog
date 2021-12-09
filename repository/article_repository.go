package repository

import (
	"database/sql"
	"go-tech-blog/model"
	"math"
	"time"
)

// ArticleCreate ...
func ArticleCreate(article *model.Article) (sql.Result, error) {
	// 現在日時を取得します
	now := time.Now()

	// 構造体に現在日時を設定します。
	article.Created = now
	article.Updated = now

	// クエリ文字列を生成します。
	query := `INSERT INTO articles (title, body, created, updated)
	VALUES (:title, :body, :created, :updated);`

	// トランザクションを開始します。
	tx := db.MustBegin()

	// クエリ文字列と構造体を引数に渡して SQL を実行します。
	// クエリ文字列内の「:title」「:body」「:created」「:updated」は構造体の値で置換されます。
	// 構造体タグで指定してあるフィールドが対象となります。（`db:"title"` など）
	res, err := tx.NamedExec(query, article)
	if err != nil {
		// エラーが発生した場合はロールバックします。
		tx.Rollback()

		// エラー内容を返却します。
		return nil, err
	}

	// SQL の実行に成功した場合はコミットします。
	tx.Commit()

	// SQL の実行結果を返却します。
	return res, nil
}

// ArticleListByCursor ...
func ArticleListByCursor(cursor int) ([]*model.Article, error) {
	// 引数で渡されたカーソルの値が 0 以下の場合は、代わりに int 型の最大値で置き換えます。
	if cursor <= 0 {
		cursor = math.MaxInt32
	}

	// ID の降順に記事データを 10 件取得するクエリ文字列を生成します。
	query := `SELECT *
	FROM articles
	WHERE id < ?
	ORDER BY id desc
	LIMIT 10`

	// クエリ結果を格納するスライスを初期化します。
	// 10 件取得すると決まっているため、サイズとキャパシティを指定しています。
	articles := make([]*model.Article, 0, 10)

	// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
	if err := db.Select(&articles, query, cursor); err != nil {
		return nil, err
	}

	return articles, nil
}

// ArticleDelete ...
func ArticleDelete(id int) error {
	// 記事データを削除するクエリ文字列を生成します。
	query := "DELETE FROM articles WHERE id = ?"

	// トランザクションを開始します。
	tx := db.MustBegin()

	// クエリ文字列とパラメータを指定して SQL を実行します。
	if _, err := tx.Exec(query, id); err != nil {
		// エラーが発生した場合はロールバックします。
		tx.Rollback()

		// エラー内容を返却します。
		return err
	}

	// エラーがない場合はコミットします。
	return tx.Commit()
}

// ArticleGetByID ...
func ArticleGetByID(id int) (*model.Article, error) {
	// クエリ文字列を生成します。
	query := `SELECT *
	FROM articles
	WHERE id = ?;`

	// クエリ結果を格納する変数を宣言します。
	// 複数件取得の場合はスライスでしたが、一件取得の場合は構造体になります。
	var article model.Article

	// 結果を格納する構造体、クエリ文字列、パラメータを指定して SQL を実行します。
	// 複数件の取得の場合は db.Select() でしたが、一件取得の場合は db.Get() になります。
	if err := db.Get(&article, query, id); err != nil {
		// エラーが発生した場合はエラーを返却します。
		return nil, err
	}

	// エラーがない場合は記事データを返却します。
	return &article, nil
}

// ArticleUpdate ...
func ArticleUpdate(article *model.Article) (sql.Result, error) {
	// 現在日時を取得します
	now := time.Now()

	// 構造体に現在日時を設定します。
	article.Updated = now

	// クエリ文字列を生成します。
	query := `UPDATE articles
	SET title = :title,
		body = :body,
		updated = :updated
	WHERE id = :id;`

	// トランザクションを開始します。
	tx := db.MustBegin()

	// クエリ文字列と引数で渡ってきた構造体を指定して、SQL を実行します。
	// クエリ文字列内の :title, :body, :id には、
	// 第 2 引数の Article 構造体の Title, Body, ID が bind されます。
	// 構造体に db タグで指定した値が紐付けされます。
	res, err := tx.NamedExec(query, article)

	if err != nil {
		// エラーが発生した場合はロールバックします。
		tx.Rollback()

		// エラーを返却します。
		return nil, err
	}

	// エラーがない場合はコミットします。
	tx.Commit()

	// SQL の実行結果を返却します。
	return res, nil
}

// ArticleGetWithWriterName ...
func ArticleGetWithWriterName(id int) (*model.Article, error) {
	// クエリ文字列を生成します。
	// 取得カラムは AS 句でリネームします。
	// リネーム後の名称は Article 構造体の db タグで指定した名称とします。
	// Null の可能性のあるカラムは COALESCE 関数を使って初期値を指定すると Go でのエラーを回避できます。
	query := `SELECT
		articles.id AS id,
		articles.title AS title,
		COALESCE(writers.name, '') AS writer_name
	FROM articles
	INNER JOIN writers ON writers.id = articles.writer_id
	WHERE articles.id = ? AND articles.writer_id IS NOT NULL;`

	var article model.Article
	if err := db.Get(&article, query, id); err != nil {
		return nil, err
	}
	return &article, nil
}

// ArticleGetWithWriter ...
func ArticleGetWithWriter(id int) (*model.Article, error) {
	// 構造体を階層化した状態でデータを取得する場合は、
	// AS 句でのリネームでドット繋ぎの名称にします。
	// Article 構造体の db タグで指定した `writer` にドットで続けて、
	// Writer 構造体の db タグで指定した `id` と `name` を指定します。
	query := `SELECT
		articles.id AS id,
		articles.title AS title,
		writers.id AS 'writer.id',
		writers.name AS 'writer.name'
	FROM articles
	INNER JOIN writers ON writers.id = articles.writer_id
	WHERE articles.id = ?;`

	var article model.Article
	if err := db.Get(&article, query, id); err != nil {
		return nil, err
	}
	return &article, nil
}

// ArticleListByWriterID ...
func ArticleListByWriterID(writerID int) ([]*model.Article, error) {
	query := `SELECT * FROM articles WHERE writer_id = ?;`
	var articles []*model.Article
	if err := db.Select(&articles, query, writerID); err != nil {
		return nil, err
	}
	return articles, nil
}

// ArticleGetWithTags ...
func ArticleGetWithTags(id int) (*model.Article, error) {
	// 記事データを取得します。
	article, err := ArticleGetByID(id)
	if err != nil {
		return nil, err
	}

	// タグデータを取得します。
	tags, err := TagListByArticleID(id)
	if err != nil {
		return nil, err
	}

	// 記事の構造体にタグ情報を格納します。
	article.Tags = tags

	return article, nil
}

// ArticleListWithTags ...
func ArticleListWithTags() ([]*model.Article, error) {
	// 記事の一覧データを取得します。
	q1 := `SELECT id, title FROM articles;`

	var articles []*model.Article
	if err := db.Select(&articles, q1); err != nil {
		return nil, err
	}

	// 取得できた記事データ一覧から記事 ID を抽出します。
	articleIDs := make([]int, len(articles))
	for i, article := range articles {
		articleIDs[i] = article.ID
	}

	// タグ情報を map で取得します。
	tagListMap, err := TagListMapByArticleIDs(articleIDs)
	if err != nil {
		return nil, err
	}

	// 記事の一覧データにタグ情報を格納します。
	for _, article := range articles {
		article.Tags = tagListMap[article.ID]
	}

	return articles, nil
}
