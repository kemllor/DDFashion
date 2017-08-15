package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["DDFashion/controllers:UserController"] = append(beego.GlobalControllerRouter["DDFashion/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUserByUid",
			Router: `/getUser/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["DDFashion/controllers:UserController"] = append(beego.GlobalControllerRouter["DDFashion/controllers:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["DDFashion/controllers:UserController"] = append(beego.GlobalControllerRouter["DDFashion/controllers:UserController"],
		beego.ControllerComments{
			Method: "SetUserInfo",
			Router: `/setUserInfo/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
