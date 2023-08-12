package users

import "sync"

type userRepository struct {
	mu    sync.Mutex
	users []*User
}

func (this *userRepository) findByEmail(email string) *User {
	this.mu.Lock()
	defer this.mu.Unlock()

	for _, user := range this.users {
		if user.Email == email {
			return user
		}
	}
	return nil
}

func (this *userRepository) findById(id int) *User {
	this.mu.Lock()
	defer this.mu.Unlock()

	for _, user := range this.users {
		if user.Id == id {
			return user
		}
	}
	return nil
}

func (this *userRepository) create(user *User) {
	this.mu.Lock()
	defer this.mu.Unlock()

	user.Id = len(this.users) + 1
	this.users = append(this.users, user)
}

func (this *userRepository) update(user *User) {
	this.mu.Lock()
	defer this.mu.Unlock()

	for i, existingUser := range this.users {
		if existingUser.Id == user.Id {
			this.users[i] = user
		}
	}
}

var UserRepository = userRepository{users: []*User{}}
