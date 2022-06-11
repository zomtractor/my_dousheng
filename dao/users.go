package dao

func NewUserOnceInstance() *userDao {
	userOnce.Do(
		func() {
			uDao = &userDao{}
		})
	return uDao
}

// AddUserToSql 添加用户
func (*userDao) AddUserToSql(user *User) *User {
	userlock.Lock()
	if db.Create(user).Error != nil {
		userlock.Unlock()
		return nil
	}
	userlock.Unlock()
	return user
}

// GetUserByUsername 获取用户
func (*userDao) GetUserByUsername(username string) *User {
	user := &User{}
	if db.Where("name=?", username).Find(user).Error != nil {
		return nil
	}
	return user
}

// GetUserByUserID 获取用户
func (*userDao) GetUserByUserID(id int) *User {
	user := &User{}
	if db.Where("id=?", id).Find(user).Error != nil {
		return nil
	}
	return user
}

// ValidToken 检查用户token
func (*userDao) ValidToken(token string) *User {
	user := &User{}
	if db.Where("concat(name,salt)=?", token).Find(user).Error != nil {
		return nil
	}
	return user
}
