package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {

	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}

	userID := snowflake.GenID()
	user := models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	mysql.InsertUser(&user)
	return nil
}
func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return nil, err
	}

	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return nil, err
	}
	user.Token = token
	return
}
