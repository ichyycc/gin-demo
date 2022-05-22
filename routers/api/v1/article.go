package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/ichyycc/gin-demo/models"
	"github.com/ichyycc/gin-demo/pkg/e"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

// GetArticle 获取单个文章
func GetArticle(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// GetArticles 获取多个文章
func GetArticles(ctx *gin.Context) {
}

// AddArticle 新增文章
func AddArticle(ctx *gin.Context) {
}

// EditArticle 修改文章
func EditArticle(ctx *gin.Context) {
}

// DeleteArticle 删除文章
func DeleteArticle(ctx *gin.Context) {
}
