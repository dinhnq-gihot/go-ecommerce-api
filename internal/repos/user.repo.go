package repos

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (u *UserRepo) GetUserById(id int) string {
	return "dinhnq"
}
