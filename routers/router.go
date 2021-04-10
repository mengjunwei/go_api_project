package routers

import (
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
		systemSourceRouters(),

	)

	beego.AddNamespace(ns)
}
