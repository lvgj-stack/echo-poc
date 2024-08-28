package main

import (
	"context"
	"flag"

	"github.com/Mr-LvGJ/cloud-native-poc/pkg"
	"github.com/Mr-LvGJ/cloud-native-poc/pkg/common/setting"
)

var (
	confFilePath = flag.String("c", "./etc/echo-poc.yml", "poc conf file path")
)

func main() {
	flag.Parse()
	ctx := context.Background()
	setting.InitGlobalConfig(*confFilePath)
	pkg.InitNecessity(ctx)
	pkg.RunServer(ctx)
}
