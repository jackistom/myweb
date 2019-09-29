package models

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"github.com/myweb/config"
	"github.com/myweb/utils"
	"path/filepath"
	"strconv"
	"strings"
)

type HomeBlockParam struct {
	Id         int
	Title      string
	Tags       []TagLink
	Short      string
	Content    string
	Author     string
	CreateTime string
	//查看文章的地址
	Link string

	//修改文章的地址
	UpdateLink string
	DeleteLink string

	//记录是否登录
	IsLogin bool
}

type TagLink struct {
	TagName string
	TagUrl  string
}

type HomeFooterPageCode struct {
	HasPre   bool
	HasNext  bool
	ShowPage string
	PreLink  string
	NextLink string
}

//----------首页显示内容---------
func MakeHomeBlocks(articles []Article, isLogin bool) template.HTML {
	htmlHome := ""
	for _, art := range articles {
		//将数据库model转换为首页模板所需要的model
		homeParam := HomeBlockParam{}
		homeParam.Id = art.Id
		homeParam.Title = art.Title
		homeParam.Tags = createTagsLinks(art.Tags)
		fmt.Println("tag-->", art.Tags)
		homeParam.Short = art.Short
		homeParam.Content = art.Content
		homeParam.Author = art.Author
		homeParam.CreateTime = utils.SwitchTimeStampToData(art.Createtime)
		//homeParam.Link = "/article/show/" + strconv.Itoa(art.Id)
		homeParam.Link = "/show/" + strconv.Itoa(art.Id)
		homeParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id)
		homeParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id)
		homeParam.IsLogin = isLogin
		fmt.Println("kkkkk")
		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, err := template.ParseFiles("views/home_block.html")
		fmt.Println(filepath.Abs("views/home_block.html"))
		if err != nil {
			log.Fatal(err)
		}
		//buffer := bytes.Buffer{}
		var buffer bytes.Buffer
		err = t.Execute(&buffer, homeParam)
		if err != nil {
			log.Fatal(err)
		}
		htmlHome += buffer.String()
	}
	fmt.Println("htmlHome-->", htmlHome)
	return template.HTML(htmlHome)
}

//将tags字符串转化成首页模板所需要的数据结构
func createTagsLinks(tags string) []TagLink {
	var tagLink []TagLink
	tagsPamar := strings.Split(tags, "&")
	for _, tag := range tagsPamar {
		tagLink = append(tagLink, TagLink{tag, "/?tag=" + tag})
	}
	return tagLink
}

//-----------翻页-----------
//page是当前的页数
func ConfigHomeFooterPageCode(page int) HomeFooterPageCode {
	pageCode := HomeFooterPageCode{}
	//查询出总的条数
	num := GetArticleRowsNum()
	//从配置文件中读取每页显示的条数
	//pageRow := config.NUM
	//计算出总页数
	//fmt.Println("kkkk")
	allPageNum := (num-1)/config.NUM + 1
	//fmt.Println("kkk")
	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)

	//当前页数小于等于1，那么上一页的按钮不能点击
	if page <= 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}

	//当前页数大于等于总页数，那么下一页的按钮不能点击
	if page >= allPageNum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}
	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
	return pageCode

}