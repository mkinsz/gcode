package server

import (
	"fmt"
	"gcode/graph"
	"gcode/graph/generated"
	"log"
	"net/http"
	"os"
	"strings"

	"gcode/orm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func graphqlHandler(orm *orm.ORM) gin.HandlerFunc {
	c := generated.Config{Resolvers: &graph.Resolver{ORM: orm}}
	h := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*") // 这是允许访问所有域
			//服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			//  header的类型
			c.Header("Access-Control-Allow-Headers", `Authorization, Content-Length, X-CSRF-Token, 
					Token, session, X_Requested_With, Accept, Origin, Connection, Accept-Encoding, 
					Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, 
					If-Modified-Since, Cache-Control, Content-Type, Host, Pragma`)
			// 允许跨域设置  可以返回其他子段
			c.Header("Access-Control-Expose-Headers", `Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, 
					Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar`) // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")          // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false") //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")             // 设置返回格式是json
		}

		if method == "OPTIONS" { //放行所有OPTIONS方法
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}

func init() {
	// host := utils.MustGet("SERVER_HOST")
	// port := utils.MustGet("SERVER_PORT")
	// fmt.Println("Init: ", host, port)
}

// Run spins up the server
func Run(orm *orm.ORM) {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := gin.Default()
	r.Use(cors())
	r.Use(static.Serve("/", static.LocalFile("static", false)))
	r.POST("/graphql", graphqlHandler(orm))
	r.GET("/playground", playgroundHandler())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", port)
	r.Run(":" + port)
}
