package auth

import (
	"buerAdmin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-sql/civil"
)

type Buer_admin struct {
	Id            int            `json:"id"`
	Username      string         `json:"username"`
	Password      string         `json:"password"`
	Phone         string         `json:"phone"`
	LastLoginTime civil.DateTime `json:"lastLoginTime"`
	Power         string         `json:"power"`
	IsDel         int            `json:"isDel"`
	CreateTime    civil.DateTime `json:"createTime"`
	UpdateTime    civil.DateTime `json:"updateTime"`
	DutyUserName  string         `json:"dutyUserName"`
}

func (user Buer_admin) Login(c *gin.Context) {
	model.Conn.First(&user)
	fmt.Println(user.Id)
	fmt.Println(user.Username)
	fmt.Println(user.Password)
}
