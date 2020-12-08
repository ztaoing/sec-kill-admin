/**
* @Author:zhoutao
* @Date:2020/7/7 下午10:26
 */

package main

import (
	"github.com/ztaoing/sec-kill-admin/sk-admin/setup"
	"github.com/ztaoing/sec-kill-pkg/pkg/bootstrap"
	pkgConfig "github.com/ztaoing/sec-kill-pkg/pkg/config"
	"github.com/ztaoing/sec-kill-pkg/pkg/mysql"
)

func main() {
	mysql.InitMysql(pkgConfig.MysqlConfig.Host, pkgConfig.MysqlConfig.Port, pkgConfig.MysqlConfig.User, pkgConfig.MysqlConfig.Pwd, pkgConfig.MysqlConfig.Db)
	setup.InitZK()
	setup.InitHTTP(bootstrap.HttpConfig.Host, bootstrap.HttpConfig.Port)
}
