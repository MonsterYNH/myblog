package user

import (
	"testing"
	"time"
)

func TestLogicUser_Login(t *testing.T) {
	logicUser := LogicUser{}
	token, _ := logicUser.Login("change user", "subscribe user")
	time.Sleep(time.Second * 10)
	logicUser.Logout(token)
}