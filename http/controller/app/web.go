// Copyright 2018 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// https://studygolang.com
// Author: polaris	polaris@studygolang.com

package app

import (
	"net/url"

	"github.com/studygolang/studygolang/context"
	. "github.com/studygolang/studygolang/http"
	"github.com/studygolang/studygolang/logic"

	echo "github.com/labstack/echo/v4"
	"github.com/polaris1119/logger"
)

type WebController struct{}

// RegisterRoute 注册路由
func (self WebController) RegisterRoute(g *echo.Group) {
	g.GET("/check_session", self.CheckSession)
	g.POST("/register", self.Register)
	g.POST("/login", self.Login)
}

// CheckSession 校验小程序 session
func (WebController) CheckSession(ctx echo.Context) error {
	code := ctx.QueryParam("code")

	wechatUser, err := logic.DefaultWechat.CheckSession(context.EchoContext(ctx), code)
	if err != nil {
		return fail(ctx, err.Error())
	}

	if wechatUser.Uid > 0 {
		data := map[string]interface{}{
			"token":    GenToken(wechatUser.Uid),
			"uid":      wechatUser.Uid,
			"nickname": wechatUser.Nickname,
			"avatar":   wechatUser.Avatar,
		}

		return success(ctx, data)
	}

	data := map[string]interface{}{
		"unbind_token": GenToken(wechatUser.Id),
	}

	return success(ctx, data)
}

// Login 通过系统用户登录
func (WebController) Login(ctx echo.Context) error {
	unbindToken := ctx.FormValue("unbind_token")
	id, ok := ParseToken(unbindToken)
	if !ok {
		return fail(ctx, "无效请求!")
	}

	username := ctx.FormValue("username")
	if username == "" {
		return fail(ctx, "用户名为空")
	}

	// 处理用户登录
	passwd := ctx.FormValue("passwd")
	userLogin, err := logic.DefaultUser.Login(context.EchoContext(ctx), username, passwd)
	if err != nil {
		return fail(ctx, err.Error())
	}

	userInfo := ctx.FormValue("userInfo")

	wechatUser, err := logic.DefaultWechat.Bind(context.EchoContext(ctx), id, userLogin.Uid, userInfo)
	if err != nil {
		return fail(ctx, err.Error())
	}

	data := map[string]interface{}{
		"token":    GenToken(wechatUser.Uid),
		"uid":      wechatUser.Uid,
		"nickname": wechatUser.Nickname,
		"avatar":   wechatUser.Avatar,
	}

	return success(ctx, data)
}

// Register 注册系统账号
func (WebController) Register(ctx echo.Context) error {
	// unbindToken := ctx.FormValue("unbind_token")
	// id, ok := ParseToken(unbindToken)
	// if !ok {
	// 	return fail(ctx, "无效请求!")
	// }

	// passwd := ctx.FormValue("passwd")
	// pass2 := ctx.FormValue("pass2")
	// if passwd != pass2 {
	// 	return fail(ctx, "确认密码不一致", 1)
	// }
	logger.Infoln("注册系统账号")
	logger.Infoln(ctx)

	fields := []string{"username", "email", "passwd"}
	form := url.Values{}
	for _, field := range fields {
		form.Set(field, ctx.FormValue(field))
	}
	// form.Set("id", strconv.Itoa(id))

	errMsg, err := logic.DefaultUser.CreateUser(context.EchoContext(ctx), form)
	if err != nil {
		return fail(ctx, errMsg, 2)
	}

	return success(ctx, nil)
}
