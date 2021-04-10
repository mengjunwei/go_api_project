package routers

import (
	beego "github.com/beego/beego/v2/server/web"

	"github.com/mengjunwei/go_api_project/controllers"
)

func systemSourceRouters() beego.LinkNamespace {
	ns := beego.NSNamespace("/linux",
		beego.NSRouter("/systemSource", &controllers.LinuxSystemSourceSetController{}),
	)
	return ns
}
