package dao

func NewNowAccountOnceInstance() *NowAccount {
	accountOnce.Do(
		func() {
			account = &NowAccount{
				make(map[int]*User, 0),
				make(map[string]int, 0),
			}
			//account.Token["aaaaaa1653464011"] = 12
			account.Token[""] = 0
		})
	return account
}

// UpdateAccount 更新信息
func (acc *NowAccount) UpdateAccount(token string, user *User) {
	mapLock.Lock()
	acc.NowUser[user.Id] = user
	acc.Token[token] = user.Id
	mapLock.Unlock()
}
