package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["sgtaUANServices/controllers:TutorController"] = append(beego.GlobalControllerRouter["sgtaUANServices/controllers:TutorController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["sgtaUANServices/controllers:TutorController"] = append(beego.GlobalControllerRouter["sgtaUANServices/controllers:TutorController"],
		beego.ControllerComments{
			Method:           "PostLogin",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

}
