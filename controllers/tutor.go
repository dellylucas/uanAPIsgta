package controllers

import (
	"encoding/json"

	"sgtaUANServices/models"

	"github.com/astaxie/beego"
)

//TutorController Operations about object
type TutorController struct {
	beego.Controller
}

//Post - insert
// @Title Post
// @Description create Tutor
// @Param	body		body 	models.Tutor	true		"The object content"
// @Success 200 models.Tutor
// @Failure 403 body is empty
// @router / [post]
func (o *TutorController) Post() {
	var ob models.Tutor
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	models.InsertTutor(&ob)
	o.Data["json"] = "insert success!"
	o.ServeJSON()
}

//PostLogin - insert
// @Title Post
// @Description create Tutor
// @Param	body		body 	models.Tutor	true		"The object content"
// @Success 200 models.Tutor
// @Failure 403 body is empty
// @router / [post]
func (o *TutorController) PostLogin() {
	var ob models.Usuarios
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	models.ConsultaUsuario(&ob)
	o.Data["json"] = "true"
	o.ServeJSON()
}

/*
// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *ObjectController) GetAll() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (o *ObjectController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Update(objectId, ob.Score)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (o *ObjectController) Delete() {
	objectId := o.Ctx.Input.Param(":objectId")
	models.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}
*/
