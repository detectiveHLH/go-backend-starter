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

// @Summary 获取单个文章
// @Produce  json
// @Param id param int true "ID"
// @Success 200 {string} json "{"code":200,"data":{"id":3,"created_on":1516937037,"modified_on":0,"tag_id":11,"tag":{"id":11,"created_on":1516851591,"modified_on":0,"name":"312321","created_by":"4555","modified_by":"","state":1},"content":"5555","created_by":"2412","modified_by":"","state":1},"msg":"ok"}"
// @Router /api/v1/articles/{id} [get]
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

// @Summary 获取单个文章
// @Produce  json
// @Param id param int true "ID"
// @Success 200 {string} json "{"code":200,"data":{"id":3,"created_on":1516937037,"modified_on":0,"tag_id":11,"tag":{"id":11,"created_on":1516851591,"modified_on":0,"name":"312321","created_by":"4555","modified_by":"","state":1},"content":"5555","created_by":"2412","modified_by":"","state":1},"msg":"ok"}"
// @Router /api/v1/articles/{id} [get]
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