package models

import (
	"log"
	"sgtaUANServices/db"

	"github.com/astaxie/beego/orm"
)

//Tutor
type Tutor struct {
	ID      int    `orm:"column(ID)"`
	Nombre  string `orm:"column(Nombre)" json:"nombre"`
	Materia string `orm:"column(Materia)" json:"materia"`
	Horario string `orm:"column(Horario)" json:"horario"`
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Tutor))
}
func InsertTutor(tutor *Tutor) {
	session := db.GetSession()

	if _, err := session.Insert(tutor); err != nil {
		log.Println(err)
	}
}
