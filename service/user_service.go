// Package service
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/7/3 20:56
 */
package service

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"r0Website-server/global"
	"r0Website-server/middleware"
	"r0Website-server/models"
	"r0Website-server/models/views"
	"r0Website-server/utils"
	"time"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

const UserColl = "users"

// UserLogin 用户登录
func (u *UserService) UserLogin(params views.LoginVo) (vo *views.LoginResultVo, err error) {
	var ans models.User
	var result views.LoginResultVo
	ans, err = FindObjByEmail(params.Email)
	if err != nil {
		return nil, errors.New("UserLogin 用户email不存在")
	}
	if err = utils.StructCopy(ans, &result); err != nil {
		return nil, errors.New("UserLogin 拷贝结构体失败")
	}
	if !utils.IsPasswordMatch(params.Password, ans.Password) {
		return nil, errors.New("UserLogin 用户密码错误")
	}
	token, err := middleware.GenToken(ans)
	if err != nil {
		return nil, errors.New("UserLogin 构造Token失败，请联系管理员" + global.Config.Author.Email)
	}
	result.Token = token
	return &result, err
}

// UserRegister 用户注册
func (u *UserService) UserRegister(params views.RegisterVo) (vo *views.RegisterResultVo, err error) {
	var input models.User
	var result views.RegisterResultVo
	UpdateUserInputByParams(&input, params)
	if emailCount := EmailCount(input.Email); emailCount > 0 {
		return &result, &models.UniqueError{UniqueField: "email", Msg: input.Email, Count: emailCount}
	}
	_, err = global.DBEngine.Collection(UserColl).InsertOne(context.TODO(), input)
	if err != nil {
		global.Logger.Error(err)
	}
	result.Username = input.Username
	return &result, err
}

// UpdateUserInputByParams 根据参数更新需要输入的用户模型，同时加密密码
func UpdateUserInputByParams(input *models.User, params views.RegisterVo) {
	input.Username = params.Username
	input.Salt, input.Password = utils.Encrypt(params.Password)
	input.Email = params.Email
	input.Phone = params.Phone
	curTime := time.Now()
	input.UpdateTime = curTime
	input.CreateTime = curTime
}

// FindObjByEmail 通过email查找对象
func FindObjByEmail(email string) (models.User, error) {
	filter := bson.M{"email": email}
	var ans models.User
	err := global.DBEngine.Collection(UserColl).FindOne(context.TODO(), &filter).Decode(&ans)
	if err != nil {
		global.Logger.Error(err)
	}
	return ans, err
}

// EmailCount 统计拥有此Email的文档数量
func EmailCount(email string) int64 {
	filter := bson.M{"email": email}
	count, err := global.DBEngine.Collection(UserColl).CountDocuments(context.TODO(), filter)
	if err != nil {
		global.Logger.Error(err)
	}
	return count
}
