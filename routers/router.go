package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"fmt"
	"strings"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(Cors())

	type Person struct {
		Name string
		Phone string
	}



	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello, welcome to index ")
	})

	// url参数由query获取
	router.GET("/welcome", func(ctx *gin.Context) {
		firstname := ctx.DefaultQuery("firstname", "Guest")
		lastname := ctx.Query("lastname")
		ctx.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.POST("/add", func(ctx *gin.Context) {
		id := ctx.Query("id")
		page := ctx.DefaultQuery("page", "0")
		name := ctx.PostForm("name")
		message := ctx.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		ctx.JSON(200, gin.H{
			"status":  gin.H{
				"status_code": http.StatusOK,
				"status":      "ok",
			},
			"message": message,
			"name":    name,
		})
	})

	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"id": 1,
			"name": "Tom",
		})
	})

	router.GET("/mongo", func(context *gin.Context){
		//可本地可远程，不指定协议时默认为http协议访问，此时需要设置 mongodb 的nohttpinterface=false来打开httpinterface。
		//也可以指定mongodb协议，如 "mongodb://127.0.0.1:27017"
		MOGODB_URI := "127.0.0.1:27017"
		//连接
		session, err := mgo.Dial(MOGODB_URI)
		//连接失败时终止
		if err != nil {
			panic(err)
		}
		//延迟关闭，释放资源
		defer session.Close()
		//设置模式
		session.SetMode(mgo.Monotonic, true)
		//选择数据库与集合
		c := session.DB("adatabase").C("acollection")
		//插入文档
		err = c.Insert(&Person{Name:"Ale", Phone:"+55 53 8116 9639"},
			&Person{Name:"Cla",  Phone:"+55 53 8402 8510"})
		//出错判断
		if err != nil {
			log.Fatal(err)
		}
		//查询文档
		result := Person{}
		//注意mongodb存储后的字段大小写问题
		err = c.Find(bson.M{"name": "Ale"}).One(&result)
		//出错判断
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Phone:", result.Phone)
	})

	return router
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method      //请求方法
		origin := c.Request.Header.Get("Origin")        //请求头部
		var headerKeys []string                             // 声明请求头keys
		for k, _ := range c.Request.Header {
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
			c.Header("Access-Control-Allow-Origin", "*")        // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")      //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")      // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")        // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")       //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")       // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()        //  处理请求
	}
}