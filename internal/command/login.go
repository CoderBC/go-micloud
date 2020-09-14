package command

import (
	"errors"
	"github.com/urfave/cli/v2"
	"go-micloud/internal/user"
	"go-micloud/pkg/zlog"
)

func (r *Command) Login() *cli.Command {
	return &cli.Command{
		Name:  "login",
		Usage: "登录小米云服务账号",
		Action: func(ctx *cli.Context) error {
			if r.HttpApi.User.IsLogin {
				return errors.New("您已登录，账号为：" + r.HttpApi.User.UserName)
			}
			if r.HttpApi.User.AutoLogin() != nil {
				err := r.HttpApi.User.Login(false)
				if err != nil {
					if err == user.ErrorPwd {
						zlog.Error("账号或密码错误,请重新输入账号密码")
						err := r.HttpApi.User.Login(true)
						if err != nil {
							return err
						}
					} else {
						return err
					}
				}
				_ = r.List().Run(ctx)
			}
			return nil
		},
	}
}