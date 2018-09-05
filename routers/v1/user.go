package v1

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go-backend-starter/pkg/app"
	"go-backend-starter/pkg/e"
	"go-backend-starter/service/user"
	"net/http"
)

func AddUser (c *gin.Context) {
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
}

func GetUser (c *gin.Context) {
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
}