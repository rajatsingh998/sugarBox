package DataHub

import (
	"sugarBox/Models"
	"sync"
)

var (
	userDataSetInstance UserDataSet
	once                sync.Once
)

type UserDataSet struct {
	users   []*Models.User
	rwMutex sync.RWMutex
}

func (ud *UserDataSet) GetUsers() []*Models.User {
	defer ud.rwMutex.RUnlock()
	ud.rwMutex.RLock()
	return ud.users
}

func (ud *UserDataSet) SetUsers(user *Models.User) {
	defer ud.rwMutex.RUnlock()
	ud.rwMutex.RLock()
	ud.users = append(ud.users, user)
}

func initUserDataSet() {
	userDataSetInstance = UserDataSet{
		users:   make([]*Models.User, 0),
		rwMutex: sync.RWMutex{},
	}
}

func GetUserDataSetInstance() *UserDataSet {
	once.Do(initUserDataSet)
	return &userDataSetInstance
}
