package main

import (
	"context"
	"game/dao"
	"game/model"
	"github.com/golang/glog"
)

func main() {
	d := dao.NewDao()
	ctx := context.TODO()
	//if err := d.Register(ctx,&model.User{
	//	Uid:        10050680,
	//	Username:   "heidashuai",
	//	Password:   "123222456",
	//	Email:      "1005088680@qq.com",
	//	Phone:      "17356582776",
	//	Age:        18,
	//	Sex:        1,
	//});err != nil {
	//	glog.Error(err)
	//}

	if err := d.Login(ctx,&model.User{
		Uid:        10050680,
		Password:   "123222456",
	});err != nil {
		glog.Error(err)
	}
}
