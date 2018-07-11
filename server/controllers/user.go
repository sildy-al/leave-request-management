package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"server/helpers"
	"strconv"

	logic "server/models/logic/user"
	structAPI "server/structs/api"
	structLogic "server/structs/logic"

	"github.com/astaxie/beego"
)

//UserController ...
type UserController struct {
	beego.Controller
}

// Login ...
func (c *UserController) Login() {
	var resp structAPI.RespData
	var reqLogin structAPI.ReqLogin

	body := c.Ctx.Input.RequestBody
	fmt.Println("LOGIN=======>", string(body))
	err := json.Unmarshal(body, &reqLogin)
	if err != nil {
		helpers.CheckErr("unmarshall req body failed @Login", err)
		resp.Error = errors.New("type request malform").Error()
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	result, errLogin := logic.DBPostUser.GetJWT(reqLogin)
	if errLogin != nil {
		resp.Error = errLogin.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = result
	}

	err = c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @Login", err)
	}
}

// GetRequestPending ...
func (c *UserController) GetRequestPending() {
	var (
		resp structAPI.RespData
	)
	idStr := c.Ctx.Input.Param(":id")
	employeeNumber, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert id failed @GetRequestPending", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	resGet, errGetPend := logic.DBPostUser.GetPendingRequest(employeeNumber)
	if errGetPend != nil {
		resp.Error = errGetPend.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetRequestPending", err)
	}
}

// GetRequestAccept ...
func (c *UserController) GetRequestAccept() {
	var (
		resp structAPI.RespData
	)
	idStr := c.Ctx.Input.Param(":id")
	employeeNumber, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert id failed @GetRequestAccept", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	resGet, errGetAccept := logic.DBPostUser.GetAcceptRequest(employeeNumber)
	if errGetAccept != nil {
		resp.Error = errGetAccept.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetRequestAccept", err)
	}
}

// GetRequestReject ...
func (c *UserController) GetRequestReject() {
	var (
		resp structAPI.RespData
	)
	idStr := c.Ctx.Input.Param(":id")
	employeeNumber, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert id failed @GetRequestReject", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	resGet, errGetReject := logic.DBPostUser.GetRejectRequest(employeeNumber)
	if errGetReject != nil {
		resp.Error = errGetReject.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetRequestReject", err)
	}
}

// UpdateNewPassword ...
func (c *UserController) UpdateNewPassword() {
	var (
		resp   structAPI.RespData
		newPwd structLogic.NewPassword
	)

	body := c.Ctx.Input.RequestBody
	fmt.Println("NEW-PASSWORD=======>", string(body))

	errMarshal := json.Unmarshal(body, &newPwd)
	if errMarshal != nil {
		helpers.CheckErr("unmarshall req body failed @UpdateNewPassword", errMarshal)
		resp.Error = errors.New("type request malform").Error()
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	employeeStr := c.Ctx.Input.Param(":id")
	employeeNumber, errCon := strconv.ParseInt(employeeStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert enum failed @UpdateNewPassword", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	errUpPassword := logic.DBPostUser.UpdatePassword(&newPwd, employeeNumber)
	if errUpPassword != nil {
		resp.Error = errUpPassword.Error()
	} else {
		resp.Body = "update password success"
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @UpdateNewPassword", err)
	}
}

// GetTypeLeave ...
func (c *UserController) GetTypeLeave() {
	var resp structAPI.RespData
	res, errGet := logic.DBPostUser.GetTypeLeave()
	if errGet != nil {
		resp.Error = errGet.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = res
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetTypeLeave", err)
	}
}

// GetSupervisors ...
func (c *UserController) GetSupervisors() {
	var resp structAPI.RespData

	res, errGet := logic.DBPostUser.GetSupervisors()
	if errGet != nil {
		resp.Error = errGet.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = res
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetSupervisors", err)
	}
}
