package controllers

import "github.com/astaxie/beego"
import (
	"encoding/json"
	"testApi/models"
)

type StudentController struct {
	beego.Controller
}

// @Title 获得所有学生
// @Description 返回所有的学生数据
// @Success 200 {object} models.Student
// @router / [get]
func (u *StudentController) GetAll() {
	ss := models.GetAllStudents()
	u.Data["json"] = ss
	u.ServeJSON()
}

// @Title 获得一个学生
// @Description 返回某学生数据
// @Param      id            path   int    true          "The key for staticblock"
// @Success 200 {object} models.Student
// @Failure 403 no data
// @router /:id [get]
func (u *StudentController) GetById() {
	id, _ := u.GetInt(":id")
	s := models.GetStudentById(id)
	u.Data["json"] = s
	u.ServeJSON()
}

// @Title 创建用户
// @Description 创建用户的描述
// @Param      body          body   models.Student true          "body for user content"
// @Success 200 {string} models.Student.Name
// @Failure 403 no data
// @router / [post]
func (u *StudentController) Post() {
	var s models.Student
	json.Unmarshal(u.Ctx.Input.RequestBody, &s)
	uname := models.AddStudent(&s)
	u.Data["json"] = uname
	u.ServeJSON()
}

// @Title 修改用户
// @Description 修改用户的内容
// @Param      body          body   models.Student true          "body for user content"
// @Success 200 {int} models.Student.Id
// @Failure 403 body is empty
// @router / [put]
func (u *StudentController) Update() {
	var s models.Student
	json.Unmarshal(u.Ctx.Input.RequestBody, &s)
	models.UpdateStudent(&s)
	u.Data["json"] = s
	u.ServeJSON()
}

// @Title 删除一个学生
// @Description 删除某学生数据
// @Param      id            path   int    true          "The key for staticblock"
// @Success 200 {object} models.Student
// @router /:id [delete]
func (u *StudentController) Delete() {
	id, _ := u.GetInt(":id")
	models.DeleteStudent(id)
	u.Data["json"] = true
	u.ServeJSON()
}
