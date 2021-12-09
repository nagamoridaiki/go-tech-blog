package main

import (
	"log"
	"os"

	"go-tech-blog/handler"
	"go-tech-blog/repository"

	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

var authUser = os.Getenv("AUTH_USER")
var authPassword = os.Getenv("AUTH_PASSWORD")

var db *sqlx.DB
var e = createMux()

func main() {
	db = connectDB()
	repository.SetDB(db)

	// ルーティングのグループを作成します。
	auth := e.Group("")

	// 作成したルーティンググループに対して Basic 認証のミドルウェアを追加します。
	auth.Use(basicAuth())

	// 認証が必要なルーティングに対しては auth グループを利用します。

	// TOP ページに記事の一覧を表示します。
	e.GET("/", handler.ArticleIndex)

	// 記事に関するページは "/articles" で開始するようにします。
	// 記事一覧画面には "/" と "/articles" の両方でアクセスできるようにします。
	// パスパラメータの ":id" も ":articleID" と明確にしています。
	e.GET("/articles", handler.ArticleIndex)                   // 一覧画面
	auth.GET("/articles/new", handler.ArticleNew)              // 新規作成画面
	e.GET("/articles/:articleID", handler.ArticleShow)         // 詳細画面
	auth.GET("/articles/:articleID/edit", handler.ArticleEdit) // 編集画面

	// HTML ではなく JSON を返却する処理は "/api" で開始するようにします。
	// 記事に関する処理なので "/articles" を続けます。
	e.GET("/api/articles", handler.ArticleList)                    // 一覧
	auth.POST("/api/articles", handler.ArticleCreate)              // 作成
	auth.DELETE("/api/articles/:articleID", handler.ArticleDelete) // 削除
	auth.PATCH("/api/articles/:articleID", handler.ArticleUpdate)  // 更新

	e.GET("/test", handler.Test)

	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CSRF())

	e.Static("/css", "src/css")
	e.Static("/js", "src/js")

	e.Validator = &CustomValidator{validator: validator.New()}

	return e
}

func connectDB() *sqlx.DB {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")
	return db
}

// CustomValidator ...
type CustomValidator struct {
	validator *validator.Validate
}

// Validate ...
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func basicAuth() echo.MiddlewareFunc {
	var basicAuthValidator middleware.BasicAuthValidator

	// basicAuthValidator という変数に、認証成功・認証失敗を判定する関数を代入します。
	basicAuthValidator = func(username, password string, c echo.Context) (bool, error) {
		// ユーザー名が "joe"、パスワードが "secret" の場合認証に成功します。
		if username == authUser && password == authPassword {
			return true, nil
		}
		return false, nil
	}

	// middleware パッケージの BasicAuth() 関数は、
	// Basic 認証判定用の関数を引数に取り、MiddlewareFunc 型を返却します。
	middlewareFunc := middleware.BasicAuth(basicAuthValidator)

	// MiddlewareFunc 型を返却します。
	return middlewareFunc
}
