package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	"weblog/models"
	"weblog/utils"
)

type BlogController struct {
	baseController
}

func (c *BlogController) list() {
	var (
		page       int
		pagesize   = 6
		offset     int
		list       []*models.Post
		hosts      []*models.Post
		categoryId int
		keyword    string
		categories []*models.Category
	)
	//分类
	c.o.QueryTable(new(models.Category)).All(&categories)
	c.Data["categories"] = categories

	if page, _ := c.GetInt("page"); page < 1 {
		page = 1
	}
	offset = (page - 1) * pagesize

	query := c.o.QueryTable(new(models.Post)).RelatedSel()

	if c.actionName == "resource" {
		query = query.Filter("types", 0)
	} else {
		query = query.Filter("types", 1)
	}

	if categoryId, _ = c.GetInt("category_id"); categoryId != 0 {
		query = query.Filter("category_id", categoryId)
	}
	keyword = c.GetString("keyword")
	if keyword != "" {
		cond1 := orm.NewCondition().Or("title__contains", keyword).Or("content__contains", keyword)
		query = query.SetCond(cond1)
		//query = query.Filter("content__contains", keyword)
	}
	query.OrderBy("-views").Limit(10, 0).All(&hosts)

	if c.actionName == "home" {
		query = query.Filter("is_top", 1)
	}
	count, _ := query.Count()
	c.Data["count"] = count
	query.OrderBy("-created").Limit(pagesize, offset).All(&list)

	c.Data["list"] = list
	c.Data["pagebar"] = utils.NewPager(page, int(count), pagesize, "/"+c.actionName, true).ToString()
	c.Data["hosts"] = hosts
}

// Home 首页
func (c *BlogController) Home() {
	c.IsStart()
	var notices []*models.Post
	c.o.QueryTable(new(models.Post)).Filter("category_id", 2).All(&notices)
	c.Data["notices"] = notices
	c.list()
	c.TplName = c.controllerName + "/home.html"
}

// Article 列表
func (c *BlogController) Article() {
	c.IsStart()
	c.list()
	c.TplName = c.controllerName + "/article.html"
}

// Detail 详情
func (c *BlogController) Detail() {
	c.IsStart()
	if id, _ := c.GetInt("id"); id != 0 {
		post := models.Post{Id: id}
		q := c.o.QueryTable(new(models.Post)).Filter("id", id)
		q.Update(orm.Params{"views": orm.ColValue(orm.ColAdd, 1)})
		q.RelatedSel().One(&post)
		c.Data["post"] = post
		var comments []*models.Comment
		c.o.QueryTable(new(models.Comment)).Filter("post_id", id).OrderBy("-created").All(&comments)
		c.Data["comments"] = comments

		var categories []*models.Category
		c.o.QueryTable(new(models.Category)).All(&categories)
		c.Data["categories"] = categories
		var hosts []*models.Post
		query := c.o.QueryTable(new(models.Post)).Filter("types", 1)
		query.OrderBy("-views").Limit(10, 0).All(&hosts)
		c.Data["hosts"] = hosts

	}
	c.TplName = c.controllerName + "/detail.html"
}

func (c *BlogController) About() {
	c.IsStart()
	post := models.Post{Id: 1}
	c.o.Read(&post)
	c.Data["post"] = post
	c.TplName = c.controllerName + "/about.html"
}

// Timeline 时间线
func (c *BlogController) Timeline() {
	c.IsStart()
	c.TplName = c.controllerName + "/timeline.html"
}

// Resource 资源
func (c *BlogController) Resource() {
	c.IsStart()
	c.list()
	c.TplName = c.controllerName + "/resource.html"
}

// Comment 插入评价
func (c *BlogController) Comment() {
	c.IsStart()
	//获取当前请求的url
	//
	Comment := models.Comment{}
	Comment.Username = c.GetString("username")
	Comment.Content = c.GetString("content")
	Comment.Ip = c.getClientIp()
	Comment.PostId, _ = c.GetInt("post_id")
	q := c.o.QueryTable(new(models.Post)).Filter("id", Comment.PostId)
	q.Update(orm.Params{"comment_count": orm.ColValue(orm.ColAdd, 1)})
	if _, err := c.o.Insert(&Comment); err != nil {
		c.History("发布评价失败"+err.Error(), "")
	} else {
		c.History("发布评价成功", c.Ctx.Request.Referer())
	}
}
