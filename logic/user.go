package logic

import (
	"bluebell/dao/mysql"
	"bluebell/pkg/snowflake"
)

func SignUp() {
	mysql.QueryUserByUsername()
	snowflake.GenID()
	mysql.InsertUser()
}
