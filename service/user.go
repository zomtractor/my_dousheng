package service

import (
	"crypto/md5"
	"encoding/hex"
	"my_dousheng/dao"
	"strconv"
	"time"
)

// Register 用户注册
func Register(username string, password string) (int, string) {
	userdao := dao.NewUserOnceInstance()
	user := userdao.GetUserByUsername(username)
	if user == nil {
		return -1, ""
	}
	salt := strconv.FormatInt(time.Now().Unix(), 10)
	hash := getHashAndSalt(password, salt)
	u := &dao.User{
		Name: username,
		Salt: salt,
		Hash: hash,
	}
	u = userdao.AddUserToSql(u)
	if u == nil {
		return -2, ""
	}
	token := username + salt
	nowAccount := dao.NewNowAccountOnceInstance()
	nowAccount.UpdateAccount(token, u)
	return u.Id, token
}

// LoginByPassword 用户登录
func LoginByPassword(username string, password string) (int, string) {
	userdao := dao.NewUserOnceInstance()
	user := userdao.GetUserByUsername(username)
	if user == nil {
		return -1, ""
	}
	salt := user.Salt
	if getHashAndSalt(password, salt) == user.Hash {
		dao.NewNowAccountOnceInstance().UpdateAccount(username+salt, user)
		return user.Id, username + salt
	} else {
		return -2, ""
	}
}

// GetUserByToken 根据用户token获取用户对象
func GetUserByToken(token string) *dao.User {
	t := CheckToken(token)
	if t == -1 {
		return nil
	}
	accountDao := dao.NewNowAccountOnceInstance()
	return accountDao.NowUser[accountDao.Token[token]]
}

// GetUserByID 根据id获取用户对象
func GetUserByID(id int) *dao.User {
	user := dao.NewUserOnceInstance().GetUserByUserID(id)
	if user == nil {
		return nil
	}
	return filterSensitive(user)
}

// CheckToken 检查token是否有效
func CheckToken(token string) int {
	id, ok := dao.NewNowAccountOnceInstance().Token[token]
	if ok {
		return id
	} else {
		user := dao.NewUserOnceInstance().ValidToken(token)
		if user == nil {
			return -1
		} else {
			dao.NewNowAccountOnceInstance().UpdateAccount(token, user)
			return CheckToken(token)
		}
	}
}

//加密算法
func getHashAndSalt(password string, salt string) string {
	m5 := md5.New()
	m5.Write([]byte(password))
	m5.Write([]byte(salt))
	hash := hex.EncodeToString(m5.Sum(nil))
	return hash
}

//过滤关键信息
func filterSensitive(u *dao.User) *dao.User {
	u.Salt = ""
	u.Hash = ""
	return u
}
