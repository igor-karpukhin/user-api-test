package user

import "errors"

type MockedUserDao struct {
	users map[uint64]*User
}

func NewMockedUserDao() *MockedUserDao {
	return &MockedUserDao{
		users: map[uint64]*User{},
	}
}

func (m *MockedUserDao) SetUsers(users map[uint64]*User) {
	m.users = users
}

func (m *MockedUserDao) GetUser(ID uint64) (*User, error) {
	if u, ok := m.users[ID]; ok {
		return u, nil
	}
	return nil, errors.New("unable to get user")
}

func (m *MockedUserDao) DeleteUser(ID uint64) error {
	delete(m.users, ID)
	return nil
}

func (m *MockedUserDao) CreateUser(user *User) (uint64, error) {
	if _, ok := m.users[user.ID]; ok {
		return 0, errors.New("user with this ID already exists")
	}
	m.users[user.ID] = user
	return user.ID, nil
}

func (m *MockedUserDao) UpdateUser(user *User) error {
	if _, ok := m.users[user.ID]; !ok {
		return errors.New("user doesn't exists")
	}
	m.users[user.ID] = user
	return nil
}

func (m *MockedUserDao) ListUsers() ([]*User, error) {
	users := []*User{}
	for _, v := range m.users {
		users = append(users, v)
	}
	return users, nil
}
