package dao

func NewNowAccountOnceInstance() *nowAccount {
	accountOnce.Do(
		func() {
			account = &nowAccount{
				make(map[int]*User, 0),
				make(map[string]int, 0),
			}
			//account.Token["aaaaaa1653464011"] = 12
			account.Token[""] = 0
		})
	return account
}

func (acc *nowAccount) UpdateAccount(token string, user *User) {
	mapLock.Lock()
	acc.NowUser[user.Id] = user
	acc.Token[token] = user.Id
	mapLock.Unlock()
}
