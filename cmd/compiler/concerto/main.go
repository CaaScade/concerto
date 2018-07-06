package main

import (
	"github.com/golang/glog"
	"github.com/koki/concerto/cmd/compiler"
)

func main() {
	if err := compiler.RootCmd.Execute(); err != nil {
		glog.Fatal(err)
	}
}
