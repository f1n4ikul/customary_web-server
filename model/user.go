package model

type User struct {
	Id   int
	Name string
	Surname string
	Age  int

}

var Users []User

var NextId int

func GetAllUsers () (users []User, err error) {
	if len(Users) == 0 {
		Users = []User{
			{1, "Андрей", "Князев", 25},
			{2, "Михаил", "Горшенев", 25},
			{3, "Сергей", "Трошин", 27},
			{4, "Виталий", "Грошев", 25},
			{5, "Павел", "Чернов", 28},
			{6, "Максим", "Морозов", 26},
			{7, "Игорь", "Долгих", 24},
			{8, "Артур", "Шамрай", 23},
			{9, "Александр", "Николаев", 27},
			{10, "Евгений", "Кириллов", 25},
		}
	}
	return Users, nil
}

func AddUser (user *User) {
	user.Id = NextId
	NextId++

	Users = append(Users, *user)
}