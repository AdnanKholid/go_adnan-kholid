package main

type User struct {
	ID       int
	Username int
	Password int
}

type UserService struct {
	Users []User // sebelumnya t []user harus Kapital dan typenya ditulis jelas yaitu "Users".
}

func (u UserService) GetAllUsers() []User {
	return u.Users // sebelumnya u.t diganti karna typenya harus jelas.
}

func (u UserService) GetUserByID(id int) User { // harus Kapital awalnya karna untuk menentukan.
	for _, user := range u.Users { // r diganti jadi user bedanya dengan type diawal harus kapital.
		if id == user.ID { // r diganti juga dengan user karena untuk menentukan.
			return user // r diganti juga dengan user karena untuk menentukan.
		}
	}
	return User{}
}
