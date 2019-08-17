package db

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbAlias       = "default"
	mysqlUser     = "admin"
	mysqlPassword = "holamundo123"
	mysqlHost     = "localhost"
	mysqlPort     = 3306
	mysqlDatabase = "sgta"
	mysqlCharset  = "utf8"
)

var (
	mysqlCon = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s",
		mysqlUser,
		mysqlPassword,
		mysqlHost,
		mysqlPort,
		mysqlDatabase,
		mysqlCharset,
	)
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase(dbAlias, "mysql", mysqlCon)
}

//GetSession - conexion de Base de Datos
func GetSession() orm.Ormer {

	/*force := true   // Drop table and re-create.
	verbose := true // Print log
	if err := orm.RunSyncdb(dbAlias, force, verbose); err != nil {
		log.Println(err)
	}*/

	session := orm.NewOrm()
	session.Using("sgta")

	return session
}
