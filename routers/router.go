package routers

import (
	"../pkg/app"
	"../service/user"
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	type Person struct {
		Name  string
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
			"status": gin.H{
				"status_code": http.StatusOK,
				"status":      "ok",
			},
			"message": message,
			"name":    name,
		})
	})

	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"id":   1,
			"name": "Tom",
		})
	})

	router.POST("/addTag", func(c *gin.Context) {
		appGin := app.Gin{c}
		name := c.PostForm("name")
		age := com.StrTo(c.PostForm("age")).MustInt()
		state := com.StrTo(c.PostForm("state")).MustInt()

		userService := user.User{
			Name: name,
			Age: age,
			State: state,
		}
		if err := userService.Add(); err != nil {
			appGin.Response(http.StatusOK, e.ERROR_ADD_ARTICLE_FAIL, nil)
			return
		}
		appGin.Response(http.StatusOK, e.SUCCESS, userService)
	})

	router.GET("/getUser", func(c *gin.Context) {
		appGin := app.Gin{c}
		id := com.StrTo(c.Query("id")).MustInt()
		valid := validation.Validation{}
		valid.Min(id, 0, "id").Message("Id必须大于0")


		if valid.HasErrors() {
			fmt.Println(valid.Errors)
			appGin.Response(http.StatusOK, e.INVALID_PARAMS, nil)
			return
		}

		userService := user.User{Id: id}
		user, err := userService.Get()
		if err != nil {
			appGin.Response(http.StatusOK, e.ERROR_GET_ARTICLE_FAIL, nil)
			return
		}

		appGin.Response(http.StatusOK, e.SUCCESS, user)
	})

	return router
}