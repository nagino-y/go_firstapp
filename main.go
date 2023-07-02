package main
import(
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	router:= gin.Default()

	// 自動的にファイルを返す設定
	router.StaticFS("/static", http.Dir("static"))

	// ルートなら/static/index.htmlにリダイレクト
	router.GET("/",func(ctx *gin.Context){
		ctx.Redirect(302,"/static/index.html")
	})

	// フォームの内容を受け取って挨拶する
	router.GET("/hello", func(ctx *gin.Context){
		name := ctx.Query("name")
		ctx.Header("Context-Type", "text/html; charset=UTF-8")
		ctx.String(200, "<h1>Hello, "+ name + "</h1>")
	})

	// サーバーを起動
	err := router.Run("127.0.0.1:8888")
	if err != nil {
		log.Fatal("サーバー起動に失敗", err)
	}
}