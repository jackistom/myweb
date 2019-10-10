package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myproject/models"
	"net/http"
)

func TagsGet(c *gin.Context) {
	//获取session
	islogin := GetSession(c)

	tags := models.QueryArticleWithParam("tags")
	fmt.Println(models.HandleTagsListData(tags))

	//返回html
	if (islogin == true) {
		c.HTML(http.StatusOK, "index.html", gin.H{"Tags": models.HandleTagsListData(tags), "IsLogin": islogin})
	}else{
	   	c.Redirect(http.StatusMovedPermanently,"/login")
	}
}
