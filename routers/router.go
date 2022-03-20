// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"resbox_go/controllers"
)

func init() {
	//匹配控制器

	beego.Router("/v1/healthy",&controllers.UserInfoController{},"Get:Healthy")
	beego.Router("/v1/userInfo",&controllers.UserInfoController{})
	beego.Router("/v1/allUserInfo",&controllers.UserInfoController{},"Post:GetAll")

}


