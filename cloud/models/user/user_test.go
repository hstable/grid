package user

import "testing"

func TestUser_Count(t *testing.T) {
	var test User
	t.Log(test.Count())
}

func TestUser_Delete(t *testing.T) {
	var test = User{
		Sub:      "test_username",
		Password: "test_password",
		Name:     "测试用户",
		Admin:    true,
	}
	test.Insert()
	//已存在删除
	err := test.Delete()
	if err != nil {
		t.Fatal(err)
	}
	//不存在删除
	err = test.Delete()
}

func TestUser_Insert(t *testing.T) {
	var test = User{
		Sub:      "test10",
		Password: "test10",
		Name:     "专家10",
		Admin:    true,
	}
	test.Delete()
	err := test.Insert()
	if err != nil {
		t.Fatal(err)
	}
	err = test.Insert()
	t.Log(err)
}

func TestUser_Get(t *testing.T) {
	var test = User{
		Sub: "test_username",
	}
	u, err := test.Find()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}
