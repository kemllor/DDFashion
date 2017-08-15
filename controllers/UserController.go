package controllers

import (
	"github.com/astaxie/beego"
	"DDFashion/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

func init() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"user.log"}`)
}

type UserController struct {
	beego.Controller
}

// @Title 根据获取指定用户
// @Description 获取指定用户
// @Param	uid  query 	string	true		"用户ID"
// @Success 200 {object} models.User       "操作成功"
// @router /getUser/ [get]
func (c *UserController) GetUserByUid() {
	uid := c.GetString("uid")
	fmt.Print("uid=====" + uid)
	if uid == "" {
		c.Data["json"] = "uid  can't  null"
	} else {
		if user, err := models.GetUserByUid(uid); err == nil {
			c.Data["json"] = user
		} else {
			c.Data["json"] = "can't find  user by this uid"
		}
	}
	c.ServeJSON()
}

// @Title  注册
// @Description 注册接口
// @Param	username		query 	string	true		"帐号"
// @Param	password		query 	string	true		"密码"
// @Success 200 {object} models.Account
// @router /register/ [get]
func (c *UserController) Register() {
	name := c.GetString("username")
	psd := c.GetString("password")
	if name == "" || psd == "" {
		c.Data["json"] = "username or password  is null"
	} else {
		if err := models.AddAccount(name, psd); err == nil {
			c.Data["json"] = "success"
		} else {
			c.Data["json"] = err
		}
	}
	c.ServeJSON()
}

// @Title  用户信息
// @Description 用户信息
// @Param	utoken		header 	string	false		"userToken"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.
// @router /setUserInfo/ [post]
func (c *UserController) SetUserInfo() {
	utoken := c.Ctx.Input.Header("utoken")
	if utoken == "" {
		c.Data["json"] = "user token is null"
	} else {
		var err error
		o := orm.NewOrm()
		acc := &models.Account{}
		if err = o.QueryTable(acc).Filter("UserToken", utoken).One(acc); err == nil {
			var user models.User
			if err = json.Unmarshal(c.Ctx.Input.RequestBody, &user); err == nil {
				user.UserId = acc.UserId
				if _, err = o.Insert(&user); err == nil {
					c.Data["json"] = "success"
				} else {
					c.Data["json"] = "insert user info fail,because the user info can only be edited  once"
				}
			} else {
				c.Data["json"] = "can't parse  user info from requestbody"
			}
		} else {
			c.Data["json"] = "can't find this user by token"
		}
		if err != nil {
			logs.Debug(err.Error())
		}
	}
	c.ServeJSON()
}
