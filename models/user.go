package models

import (
	"github.com/astaxie/beego/orm"
	"DDFashion/Uitls"
	"time"
)

type User struct {
	Id        int            `json:"-" orm:"pk;auto"`
	Name      string        `json:"name,omitempty" orm:"null"`
	Mobile    string        `json:"mobile,omitempty" orm:"unique;null"` //手机号
	IdCard    string        `json:"idCard,omitempty" orm:"unique;null"` //身份证
	Email     string        `json:"email,omitempty"  orm:"null;unique"`
	Gender    int           `json:"gender,omitempty"  orm:"default(0)"`            //性别 0 默认，1，男 2，女
	QQToken   string        `json:"qqToken,omitempty" orm:"null;column(qq_token)"` //qq授权码
	WXToken   string        `json:"wxToken,omitempty" orm:"null;column(wx_token)"` //微信授权码
	UserId    int            `json:"-" orm:"unique"`
}

type Account struct {
	UserId       int        `json:"-" orm:"pk;auto"`
	UserName     string    `json:"userName,omitempty" orm:"unique"`       //登录帐号名
	Password     string    `json:"-"`                                     //登录密码
	UserToken    string    `json:"userToken,omitempty" orm:"unique"`      //个人识别码，MD5加密串，
	VertualMoney int       `json:"vertMoney,omitempty"`                   //虚拟币
	NickName     string    `json:"nickName,omitempty"  orm:"null;unique"` //用户昵称
	Icon         string    `json:"icon,omitempty" orm:"null;unique"`      //用户头像
	LoginTime    string     `json:"-" orm:"auto_now;type(datetime)"`      //登录时间
	RegisterTime string    `json:"-" orm:"auto_now_add;type(datetime)"`   //注册时间
}

func GetUserByUid(uid string) (account *Account, err error) {
	o := orm.NewOrm()
	account = &Account{}
	err = o.QueryTable(account).Filter("user_id", uid).One(account)
	return
}

//添加帐户
func AddAccount(username string, psd string) (err error) {
	acc := new(Account)
	acc.UserName = username
	acc.Password = Uitls.EnCode(psd)
	acc.UserToken = Uitls.EnCode(username + time.Now().String())
	o := orm.NewOrm()
	_, err = o.Insert(acc)
	return
}

//添加用户信息
func AddUserInfo(user *User) (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(user)
	return
}

func init() {
	orm.RegisterModel(new(User), new(Account))
}
