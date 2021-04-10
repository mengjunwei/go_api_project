module github.com/mengjunwei/go_api_project

go 1.15

require (
	github.com/StackExchange/wmi v0.0.0-20210224194228-fe8f1750fd46 // indirect
	github.com/beego/beego/v2 v2.0.1
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/mengjunwei/go-utils v0.0.0-20210330100332-6ae500945626
	github.com/shirou/gopsutil v3.21.3+incompatible
	github.com/smartystreets/goconvey v1.6.4
)

replace (
	github.com/coreos/bbolt v1.3.5 => go.etcd.io/bbolt v1.3.4
	google.golang.org/grpc v1.36.0 => google.golang.org/grpc v1.26.0
)
