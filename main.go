package main

import (
	"context"
	"fmt"
	"game/dao"
	"game/model"
	"github.com/golang/glog"
	"os"
)

func main() {
	d := dao.NewDao()
	ctx := context.TODO()
	r,err := d.QueryRoleById(ctx,1)
	if err != nil {
		glog.Errorf("query error %s",err)
		os.Exit(0)
	}
	fmt.Print(model.Strengthen(r.Clothing,true))
}
