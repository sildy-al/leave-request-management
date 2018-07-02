package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"server/helpers"
	"strconv"
	"strings"

	logic "server/models/logic/user"
	structAPI "server/structs/api"

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

// GetPendingLeave ...
func (c *UserController) GetPendingLeave() {
	var (
		resp structAPI.RespData
	)
	idStr := c.Ctx.Input.Param(":id")
	supervisorID, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert id failed @GetPendingLeave", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	resGet, errGetPend := logic.DBPostUser.GetUserPending(supervisorID)
	if errGetPend != nil {
		resp.Error = errGetPend.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetPendingLeave", err)
	}
}

// GetAcceptLeave ...
func (c *UserController) GetAcceptLeave() {
	var (
		resp structAPI.RespData
	)
	idStr := c.Ctx.Input.Param(":id")
	supervisorID, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert id failed @GetAcceptLeave", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	resGet, errGetAccept := logic.DBPostUser.GetUserAccept(supervisorID)
	if errGetAccept != nil {
		resp.Error = errGetAccept.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetAcceptLeave", err)
	}
}

// GetRejectLeave ...
func (c *UserController) GetRejectLeave() {
	var (
		resp structAPI.RespData
	)
	idStr := c.Ctx.Input.Param(":id")
	supervisorID, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert id failed @GetRejectLeave", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	resGet, errGetReject := logic.DBPostUser.GetUserReject(supervisorID)
	if errGetReject != nil {
		resp.Error = errGetReject.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetRejectLeave", err)
	}
}

// AcceptStatusBySupervisor ...
func (c *UserController) AcceptStatusBySupervisor() {
	var resp structAPI.RespData

	idStr := c.Ctx.Input.Param(":id")
	id, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert id failed @AcceptStatusBySupervisor", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	employeeStr := c.Ctx.Input.Param(":enumber")
	employeeNumber, errCon := strconv.ParseInt(employeeStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert enum failed @AcceptStatusBySupervisor", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	errUpStat := logic.DBPostUser.AcceptBySupervisor(id, employeeNumber)
	if errUpStat != nil {
		resp.Error = errUpStat.Error()
	} else {
		resp.Body = "status leave request has been accepted"
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @AcceptStatusBySupervisor", err)
	}
}

// RejectStatusBySupervisor ...
func (c *UserController) RejectStatusBySupervisor() {
	var (
		resp structAPI.RespData
		// leave structLogic.LeaveReason
	)

	reason := c.Ctx.Input.Param(":reason")
	strReason := reason
	strReason = strings.Replace(strReason, "_", " ", -1)

	beego.Debug("===================", reason)
	idStr := c.Ctx.Input.Param(":id")
	id, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert id failed @AcceptStatusBySupervisor", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	employeeStr := c.Ctx.Input.Param(":enumber")
	employeeNumber, errCon := strconv.ParseInt(employeeStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert enum failed @AcceptStatusBySupervisor", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	errUpStat := logic.DBPostUser.RejectBySupervisor(strReason, id, employeeNumber)
	if errUpStat != nil {
		resp.Error = errUpStat.Error()
	} else {
		resp.Body = "status leave request has been rejected"
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @RejectStatusBySupervisor", err)
	}
}