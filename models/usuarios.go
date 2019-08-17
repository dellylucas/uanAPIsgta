package models

import (
	"github.com/astaxie/beego/orm"
)

//Tutor
type Usuarios struct {
	ID      int    `orm:"column(ID)"`
	Usuario string `orm:"column(Usuario)" json:"usuario"`
	Clave   string `orm:"column(Clave)" json:"clave"`
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Usuarios))
}
func ConsultaUsuario(usuario *Usuarios) {
	//session := db.GetSession()
	/*
		if _, err := session.QueryTable(Usuarios.usuario).Exist; err != nil {
			log.Println(err)
		}*/
}
