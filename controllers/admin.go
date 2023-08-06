package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"strings"
	"time"
	"weblog/models"
	"weblog/utils"
)

type AdminController struct {
	baseController
}

// Config 配置信息
func (c *AdminController) Config() {
	var result []*models.Config
	c.o.QueryTable(new(models.Config)).All(&result)
	options := make(map[string]string)
	mp := make(map[string]*models.Config)
	for _, v := range result {
		options[v.Name] = v.Value
		mp[v.Name] = v
	}
	if c.Ctx.Request.Method == "POST" {
		keys := []string{"url", "title", "keywords", "description", "email", "start", "qq"}
		for _, key := range keys {
			val := c.GetString(key)
			if _, ok := mp[key]; !ok {
				options[key] = val
				c.o.Insert(&models.Config{Name: key, Value: val})
			} else {
				opt := mp[key]
				if _, err := c.o.Update(&models.Config{Id: opt.Id, Name: opt.Name, Value: val}); err != nil {
					continue
				}
			}
		}
		c.History("设置数据成功", "")
	}
	c.Data["config"] = options
	c.TplName = c.controllerName + "/config.html"
}

// Login 后台用户登录
func (c *AdminController) Login() {
	if c.Ctx.Request.Method == "POST" {
		username := c.GetString("username")
		password := c.GetString("password")
		user := models.User{Username: username}
		c.o.Read(&user, "username")

		if user.Password == "" {
			c.History("账号不存在", "")
		}

		if utils.Md5(password) != strings.Trim(user.Password, " ") {
			c.History("密码错误", "")
		}
		user.LastIp = c.getClientIp()
		user.LoginCount = user.LoginCount + 1
		if _, err := c.o.Update(&user); err != nil {
			c.History("登录异常", "")
		} else {
			c.History("登录成功", "/admin/main")
		}
		c.SetSession("user", user)

	}
	c.TplName = c.controllerName + "/login.html"
}

func (c *AdminController) Logout() {
	//c.DestroySession()
	c.DelSession("user")
	c.History("退出登录", "/admin/login")
}

// About 单页
func (c *AdminController) About() {
	c.Ctx.WriteString("About")
}

// Index 后台首页
func (c *AdminController) Index() {
	var categories []*models.Category
	c.o.QueryTable(new(models.Category)).All(&categories)
	c.Data["categories"] = categories
	var (
		page     int
		pagesize = 20
		offset   int
		list     []*models.Post
		keyword  string
		cateId   int
	)
	keyword = c.GetString("title")
	cateId, _ = c.GetInt("cate_id")
	if page, _ = c.GetInt("page"); page < 1 {
		page = 1
	}
	offset = (page - 1) * pagesize
	//c.Ctx.WriteString(new(models.Post))
	query := c.o.QueryTable(new(models.Post))
	if keyword != "" {
		cond1 := orm.NewCondition().Or("title__contains", keyword).Or("content__contains", keyword)
		query = query.SetCond(cond1)
		//query = query.Filter("title__contains", keyword)
	}
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-is_top", "-created").Limit(pagesize, offset).All(&list)
	}
	c.Data["keyword"] = keyword
	c.Data["count"] = count
	c.Data["list"] = list
	c.Data["cate_id"] = cateId
	c.Data["pagebar"] = utils.NewPager(page, int(count), pagesize,
		fmt.Sprintf("/admin/index.html?keyword=%s", keyword), true).ToString()
	c.TplName = c.controllerName + "/list.html"
}

// Main 主页
func (c *AdminController) Main() {
	c.TplName = c.controllerName + "/main.tpl"
}

// Article 文章
func (c *AdminController) Article() {
	var categories []*models.Category
	c.o.QueryTable(new(models.Category)).All(&categories)
	id, _ := c.GetInt("id")
	if id != 0 {
		post := models.Post{Id: id}
		c.o.Read(&post)
		c.Data["post"] = post
	}
	c.Data["categories"] = categories
	c.TplName = c.controllerName + "/_form.html"
}

// Upload 上传接口
func (c *AdminController) Upload() {
	f, h, err := c.GetFile("upFilename")
	result := make(map[string]interface{})
	img := ""
	if err == nil {
		exStrArr := strings.Split(h.Filename, ".")
		exStr := strings.ToLower(exStrArr[len(exStrArr)-1])
		if exStr != "jpg" && exStr != "png" && exStr != "gif" {
			result["code"] = 1
			result["message"] = "只能上传jpg、png、gif格式"
		}
		img = "static/upload/" + utils.UniqueId() + "." + exStr
		err = c.SaveToFile("upFilename", img) // 保存位置在 static/upload, 没有文件夹要先创建
		if err != nil {
			result["code"] = 1
			result["message"] = "上传异常" + err.Error()
		} else {
			result["code"] = 0
			result["message"] = img
		}
	} else {
		result["code"] = 2
		result["message"] = "上传异常" + err.Error()
	}
	defer f.Close()
	c.Data["json"] = result
	c.ServeJSON()
}

// Save 保存
func (c *AdminController) Save() {
	post := models.Post{}
	user := c.GetSession("user").(models.User)
	post.User = &user
	post.Title = c.GetString("title")
	post.Content = c.GetString("content")
	post.IsTop, _ = c.GetInt8("is_top")
	post.Types, _ = c.GetInt8("types")
	post.Tags = c.GetString("tags")
	post.Url = c.GetString("url")
	post.CategoryId, _ = c.GetInt("category_id")
	post.Info = c.GetString("info")
	post.Image = c.GetString("Image")

	id, _ := c.GetInt("id")
	if id == 0 {
		if _, err := c.o.Insert(&post); err != nil {
			c.History("博文插入数据错误"+err.Error(), "")
		} else {
			c.History("博文插入数据成功", "/admin/index")
		}
	} else {
		post.Id = id
		if _, err := c.o.Update(&post, "title", "content", "is_top", "types", "tags", "url", "category_id", "info", "image"); err != nil {
			c.History("博文更新数据出错"+err.Error(), "")
		} else {
			c.History("博文更新成功", "/admin/index")
		}
	}
}

func (c *AdminController) Delete() {
	id, err := strconv.Atoi(c.GetString("id"))
	if err != nil {
		c.History("参数错误", "")
	} else {
		if _, err := c.o.Delete(&models.Post{Id: id}); err != nil {
			c.History("未能成功删除", "")
		} else {
			c.History("删除成功", "/admin/index.html")
		}
	}
}

// Category 类目
func (c *AdminController) Category() {
	var categories []*models.Category
	c.o.QueryTable(new(models.Category)).OrderBy("-id").All(&categories)
	c.Data["categories"] = categories
	c.TplName = c.controllerName + "/category.tpl"
}

// Categoryadd 添加修改类目
func (c *AdminController) Categoryadd() {
	id := c.GetString("id")
	if id != "" {
		intId, _ := strconv.Atoi(id)
		cate := models.Category{Id: intId}
		c.o.Read(&cate)
		c.Data["cate"] = cate
	}
	c.TplName = c.controllerName + "/category_add.tpl"
}

// CategorySave 处理插入数据的字段
func (c *AdminController) CategorySave() {
	name := c.GetString("name")
	id := c.GetString("id")
	category := models.Category{}
	category.Name = name
	category.Updated = time.Now()
	if id == "" {
		category.Created = time.Now()
		if _, err := c.o.Insert(&category); err != nil {
			c.History("插入分类数据错误", "")
		} else {
			c.History("插入分类数据成功", "/admin/category")
		}
	} else {
		intId, err := strconv.Atoi(id)
		if err != nil {
			c.History("参数错误", "")
		}
		category.Id = intId
		if _, err := c.o.Update(&category); err != nil {
			c.History("更新分类数据出错", "")
		} else {
			c.History("更新分类成功", "/admin/category")
		}
	}
}

func (c *AdminController) CategoryDel() {
	id, err := strconv.Atoi(c.GetString("id"))
	if err != nil {
		c.History("参数错误", "")
	} else {
		if _, err := c.o.Delete(&models.Category{Id: id}); err != nil {
			c.History("未能成功删除", "")
		} else {
			c.History("删除成功", "/admin/category")
		}
	}
}
