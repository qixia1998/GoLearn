package dao

import (
	"testing"
)

func TestUserDAOImpl_Save(t *testing.T) {
	userDAO := &UserDAOImpl{}

	err := InitMysql("127.0.0.1", "3306", "root", "12345678", "user")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	user := &UserEntity{
		Username: "qixia1998",
		Password: "qixia1998",
		Email:    "qixia717138552@mail.com",
	}

	err = userDAO.Save(user)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("new User ID is %d", user.ID)

}

func TestUserDAOImpl_SelectByEmail(t *testing.T) {

	userDAO := &UserDAOImpl{}

	err := InitMysql("127.0.0.1", "3306", "root", "12345678", "user")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	user, err := userDAO.SelectByEmail("qixia717138552@mail.com")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("result uesrname is %s", user.Username)

}
