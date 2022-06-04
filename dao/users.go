package dao

func NewUserOnceInstance() *userDao {
	userOnce.Do(
		func() {
			uDao = &userDao{}
		})
	return uDao
}

func (*userDao) AddUserToSql(user *User) *User {
	userlock.Lock()
	if db.Create(user).Error != nil {
		userlock.Unlock()
		return nil
	}
	userlock.Unlock()
	return user
}

func (*userDao) GetUserByUsername(username string) *User {
	user := &User{}
	if db.Where("name=?", username).Find(user).Error != nil {
		return nil
	}
	return user
}

func (*userDao) GetUserByUserID(id int) *User {
	user := &User{}
	if db.Where("id=?", id).Find(user).Error != nil {
		return nil
	}
	return user
}
func (*userDao) ValidToken(token string) *User {
	user := &User{}
	if db.Where("concat(name,salt)=?", token).Find(user).Error != nil {
		return nil
	}
	return user
}
