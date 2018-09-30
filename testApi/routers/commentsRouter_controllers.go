package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["testApi/controllers:StudentController"] = append(beego.GlobalControllerRouter["testApi/controllers:StudentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["testApi/controllers:StudentController"] = append(beego.GlobalControllerRouter["testApi/controllers:StudentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["testApi/controllers:StudentController"] = append(beego.GlobalControllerRouter["testApi/controllers:StudentController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["testApi/controllers:StudentController"] = append(beego.GlobalControllerRouter["testApi/controllers:StudentController"],
		beego.ControllerComments{
			Method: "GetById",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["testApi/controllers:StudentController"] = append(beego.GlobalControllerRouter["testApi/controllers:StudentController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
