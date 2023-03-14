package service

import (
	"GoLearn/Micro-Go/Section10/user/dao"
	"GoLearn/Micro-Go/Section10/user/redis"
	"context"
	"testing"
)

func TestUserServiceImpl_Login(t *testing.T) {

	err := dao.InitMysql("127.0.0.1", "3306", "root", "12345678", "user")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	err = redis.InitRedis("127.0.0.1", "6379", "")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	userService := &UserServiceImpl{
		userDAO: &dao.UserDAOImpl{},
	}

	user, err := userService.Login(context.Background(), "qixia717138552@mail.com", "qixia1998")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("user id is %d", user.ID)

}

func TestUserServiceImpl_Register(t *testing.T) {

	err := dao.InitMysql("127.0.0.1", "3306", "root", "12345678", "user")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	err = redis.InitRedis("127.0.0.1", "6379", "")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	userService := &UserServiceImpl{
		userDAO: &dao.UserDAOImpl{},
	}

	user, err := userService.Register(context.Background(),
		&RegisterUserVO{
			Username: "qixia1998",
			Password: "qixia1998",
			Email:    "qixia717138552@mail.com",
		})

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("user id is %d", user.ID)

}
