package apis

import (
	// "fmt"
	"encoding/json"
	system "golddigger/app/system/models"
	systemReq "golddigger/app/system/models/request"
	request "golddigger/app/system/models/request/common"
	systemRes "golddigger/app/system/models/response"
	response "golddigger/common/models/response"
	"golddigger/global"
	"golddigger/utils"
	"golddigger/utils/jwtutil"
	"golddigger/utils/nonce"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"

	// "github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type UserApi struct{}

// var store = base64Captcha.DefaultMemStore
var store = nonce.NewDefaultRedisStore()

func (w *UserApi) GetRandomSignMsg(c *gin.Context) {
	store = store.UseWithCtx(c)
	data, _ := c.GetRawData()
	var body map[string]string
	_ = json.Unmarshal(data, &body)
	msg := system.NewNonceMsg()
	s := body["address"]
	store.Set(s, msg.GetNonce())
	response.OkWithData(msg, "", c)
}
func (w *UserApi) SignAndLogin(c *gin.Context) {
	store = store.UseWithCtx(c)
	data, _ := c.GetRawData()
	var body map[string]string
	_ = json.Unmarshal(data, &body)
	addr := body["address"]
	if !utils.IsValidAddress(addr) {
		response.FailWithMessage("不是有效的钱包地址！", c)
	}
	sig := body["signature"]
	nonce := store.Get(addr, true)
	if nonce == "" {
		response.FailWithMessage("nonce已过期", c)
	}
	message := system.DefaultNonceMsg().GetMsg() + nonce
	// Hash the unsigned message using EIP-191
	hashedMessage := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)) + message)
	hash := crypto.Keccak256Hash(hashedMessage)

	// Get the bytes of the signed message
	decodedMessage := hexutil.MustDecode(sig)

	// Handles cases where EIP-115 is not implemented (most wallets don't implement it)
	if decodedMessage[64] == 27 || decodedMessage[64] == 28 {
		decodedMessage[64] -= 27
	}

	// Recover a public key from the signed message
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), decodedMessage)
	if sigPublicKeyECDSA == nil {
		response.FailWithMessage("Could not get a public get from the message signature", c)
	}
	if err != nil {
		response.FailWithMessage("wrong messag", c)
	}
	ad := crypto.PubkeyToAddress(*sigPublicKeyECDSA).String()

	if strings.EqualFold(addr, ad) {
		u := &system.SysUser{Address: addr}
		user, ty := userService.Login(u)
		switch ty {
		case 0:
			response.FailWithMessage("登录失败", c)
			return
		case 1:
			response.FailWithMessage("登录失败", c)
			return
		case 2:
			// response.OkWithMessage("登录成功", c)
		}
		// if err != nil {
		// 	// global.GD_LOG.Error("用户："+addr+" 登陆失败!!", zap.Error(err))
		// 	response.FailWithMessage("登录失败", c)
		// 	return
		// }
		if user.Enable != 1 {
			global.GD_LOG.Error("登陆失败! 用户:" + addr + " 被禁止登录!")
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		w.TokenNext(c, *user)
		return
		// response.OkWithMessage("登录成功", c)
	} else {
		response.FailWithMessage("验证失败", c)
	}

}

// TokenNext 登录以后签发jwt
func (b *UserApi) TokenNext(c *gin.Context, user system.SysUser) {
	j := &jwtutil.JWT{SigningKey: []byte(global.GD_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		Address:     user.Address,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GD_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GD_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "成功", c)
		return
	}
	//redis中获取token
	if jwtStr, err := jwtService.GetRedisJWT(user.Address); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Address); err != nil {
			global.GD_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "成功", c)
	} else if err != nil {
		global.GD_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Address); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "成功", c)
	}
}

// Register
// @Tags     SysUser
// @Summary  用户注册账号
// @Produce   application/json
// @Param    data  body      systemReq.Register                                            true  "用户名, 昵称, 密码, 角色ID"
// @Success  200   {object}  response.Response{data=systemRes.SysUserResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router   /user/admin_register [post]
// func (b *UserApi) Register(c *gin.Context) {
// 	var r systemReq.Register
// 	err := c.ShouldBindJSON(&r)
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	err = utils.Verify(r, utils.RegisterVerify)
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	var authorities []system.SysAuthority
// 	for _, v := range r.AuthorityIds {
// 		authorities = append(authorities, system.SysAuthority{
// 			AuthorityId: v,
// 		})
// 	}
// 	user := &system.SysUser{Address: r.Address,AuthorityId: r.AuthorityId, Authorities: authorities, Enable: r.Enable, Phone: r.Phone, Email: r.Email}
// 	userReturn, err := userService.Register(*user)
// 	if err != nil {
// 		global.GD_LOG.Error("注册失败!", zap.Error(err))
// 		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
// 		return
// 	}
// 	response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
// }

// ChangePassword
// @Tags      SysUser
// @Summary   用户修改密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      systemReq.ChangePasswordReq    true  "用户名, 原密码, 新密码"
// @Success   200   {object}  response.Response{msg=string}  "用户修改密码"
// @Router    /user/changePassword [post]
// func (b *UserApi) ChangePassword(c *gin.Context) {
// 	var req systemReq.ChangePasswordReq
// 	err := c.ShouldBindJSON(&req)
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	err = utils.Verify(req, utils.ChangePasswordVerify)
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	uid := jwtutil.GetUserID(c)
// 	u := &system.SysUser{COMMON_MODEL: global.COMMON_MODEL{ID: uid}, Password: req.Password}
// 	_, err = userService.ChangePassword(u, req.NewPassword)
// 	if err != nil {
// 		global.GD_LOG.Error("修改失败!", zap.Error(err))
// 		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
// 		return
// 	}
// 	response.OkWithMessage("修改成功", c)
// }

// GetUserList
// @Tags      SysUser
// @Summary   分页获取用户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取用户列表,返回包括列表,总数,页码,每页数量"
// @Router    /user/getUserList [post]
func (b *UserApi) GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userService.GetUserInfoList(pageInfo)
	if err != nil {
		global.GD_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// SetUserAuthority
// @Tags      SysUser
// @Summary   更改用户权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SetUserAuth          true  "用户UUID, 角色ID"
// @Success   200   {object}  response.Response{msg=string}  "设置用户权限"
// @Router    /user/setUserAuthority [post]
func (b *UserApi) SetUserAuthority(c *gin.Context) {
	var sua systemReq.SetUserAuth
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if UserVerifyErr := utils.Verify(sua, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	userID := jwtutil.GetUserID(c)
	err = userService.SetUserAuthority(userID, sua.AuthorityId)
	if err != nil {
		global.GD_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	claims := jwtutil.GetUserInfo(c)
	j := &jwtutil.JWT{SigningKey: []byte(global.GD_CONFIG.JWT.SigningKey)} // 唯一签名
	claims.AuthorityId = sua.AuthorityId
	if token, err := j.CreateToken(*claims); err != nil {
		global.GD_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		c.Header("new-token", token)
		c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt, 10))
		response.OkWithMessage("修改成功", c)
	}
}

// SetUserAuthorities
// @Tags      SysUser
// @Summary   设置用户权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SetUserAuthorities   true  "用户UUID, 角色ID"
// @Success   200   {object}  response.Response{msg=string}  "设置用户权限"
// @Router    /user/setUserAuthorities [post]
func (b *UserApi) SetUserAuthorities(c *gin.Context) {
	var sua systemReq.SetUserAuthorities
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.SetUserAuthorities(sua.ID, sua.AuthorityIds)
	if err != nil {
		global.GD_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// DeleteUser
// @Tags      SysUser
// @Summary   删除用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "用户ID"
// @Success   200   {object}  response.Response{msg=string}  "删除用户"
// @Router    /user/deleteUser [delete]
func (b *UserApi) DeleteUser(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqId, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jwtId := jwtutil.GetUserID(c)
	if jwtId == uint(reqId.ID) {
		response.FailWithMessage("删除失败", c)
		return
	}
	err = userService.DeleteUser(reqId.ID)
	if err != nil {
		global.GD_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// SetUserInfo
// @Tags      SysUser
// @Summary   设置用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysUser                                             true  "ID, 用户名, 昵称, 头像链接"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "设置用户信息"
// @Router    /user/setUserInfo [put]
func (b *UserApi) SetUserInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(user, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if len(user.AuthorityIds) != 0 {
		err = userService.SetUserAuthorities(user.ID, user.AuthorityIds)
		if err != nil {
			global.GD_LOG.Error("设置失败!", zap.Error(err))
			response.FailWithMessage("设置失败", c)
			return
		}
	}
	err = userService.SetUserInfo(system.SysUser{
		COMMON_MODEL: global.COMMON_MODEL{
			ID: user.ID,
		},
		Enable: user.Enable,
	})
	if err != nil {
		global.GD_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// SetSelfInfo
// @Tags      SysUser
// @Summary   设置用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysUser                                             true  "ID, 用户名, 昵称, 头像链接"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "设置用户信息"
// @Router    /user/SetSelfInfo [put]
func (b *UserApi) SetSelfInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user.ID = jwtutil.GetUserID(c)
	err = userService.SetUserInfo(system.SysUser{
		COMMON_MODEL: global.COMMON_MODEL{
			ID: user.ID,
		},
		Enable: user.Enable,
	})
	if err != nil {
		global.GD_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取用户信息"
// @Router    /user/getUserInfo [get]
func (b *UserApi) GetUserInfo(c *gin.Context) {
	address := jwtutil.GetUserAddress(c)
	ReqUser, err := userService.GetUserInfo(address)
	if err != nil {
		global.GD_LOG.Error("!", zap.Error(err)) //获取失败
		response.FailWithMessage("", c)
		return
	}
	response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "", c) //获取成功

}
func (b *UserApi) SetProviderDefault(c *gin.Context) {
	address := jwtutil.GetUserAddress(c)
	type data struct {
		UrlType string `json:"urlType"`
	}
	var parm data
	err := c.ShouldBindJSON(&parm)
	var err1 error
	if err == nil {
		switch parm.UrlType {
		case "http":
			err1 = userService.SetHttpDefault(address)
		case "wss":
			err1 = userService.SetWssDefault(address)
		default:
			global.GD_LOG.Error("不正确的url type", zap.Error(err1))
			response.FailWithMessage("不正确的url type", c)
			return
		}
		if err1 != nil {
			global.GD_LOG.Error("设置失败!", zap.Error(err1))
			response.FailWithMessage("设置失败", c)
			return
		}
		response.OkWithMessage("设置成功", c)
	} else {
		global.GD_LOG.Error("设置失败!")
		response.FailWithMessage("设置失败", c)
	}

}
func (b *UserApi) SetTheme(c *gin.Context) {
	address := jwtutil.GetUserAddress(c)
	type data struct {
		Theme string `json:"theme"`
	}
	var parm data
	err := c.ShouldBindJSON(&parm)
	if err == nil {
		if parm.Theme == "dark" || parm.Theme == "light"{
			err = userService.SetTheme(address ,parm.Theme)
			if err != nil {
				global.GD_LOG.Error("主题设置失败!", zap.Error(err))
				response.FailWithMessage("主题设置失败", c)
				return
			}
			response.OkWithMessage("主题设置成功", c)
		}
	}else {
		global.GD_LOG.Error(err.Error())
		response.FailWithMessage("主题设置失败", c)
	}
}
func (b *UserApi) SetProviderUrl(c *gin.Context) {
	type data struct {
		UrlType uint   `json:"urlType"`
		Url     string `json:"Url"`
	}
	var parm data
	err := c.ShouldBindJSON(&parm)
	address := jwtutil.GetUserAddress(c)
	var err1 error
	if err == nil {
		switch parm.UrlType {
		case 1:
			err1 = userService.SetHttpProvider(address, parm.Url)
		case 2:
			err1 = userService.SetWssProvider(address, parm.Url)
		default:
			global.GD_LOG.Error("不正确的url type", zap.Error(err1))
			response.FailWithMessage("不正确的url type", c)
			return
		}
		if err1 != nil {
			global.GD_LOG.Error("设置失败!", zap.Error(err1))
			response.FailWithMessage("设置失败", c)
			return
		}
		response.OkWithMessage("设置成功", c)
	} else {
		global.GD_LOG.Error(err.Error())
		response.FailWithMessage("设置失败", c)
	}
}

// @Summary   重置用户密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      system.SysUser                 true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "重置用户密码"
// @Router    /user/resetPassword [post]
func (b *UserApi) ResetPassword(c *gin.Context) {
	var user system.SysUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.ResetPassword(user.ID)
	if err != nil {
		global.GD_LOG.Error("重置失败!", zap.Error(err))
		response.FailWithMessage("重置失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("重置成功", c)
}
func (b *UserApi) Upgrade(c *gin.Context) {
	type Upgrade struct {
		Type uint16 `json:"type"`
		Hash string `json:"hash"`
	}
	address := jwtutil.GetUserAddress(c)
	var up Upgrade
	err := c.ShouldBindJSON(&up)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user, err := userService.Upgrade(up.Type, up.Hash, address)
	if err != nil {
		global.GD_LOG.Error("购买失败!", zap.Error(err))
		response.FailWithMessage("购买失败"+err.Error(), c)
		return
	}
	claims := jwtutil.GetUserInfo(c)
	j := &jwtutil.JWT{SigningKey: []byte(global.GD_CONFIG.JWT.SigningKey)} // 唯一签名
	claims.AuthorityId = user.AuthorityId
	if token, err := j.CreateToken(*claims); err != nil {
		global.GD_LOG.Error("购买失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		c.Header("new-token", token)
		c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt, 10))
		response.OkWithMessage("购买成功!", c)
	}
}
