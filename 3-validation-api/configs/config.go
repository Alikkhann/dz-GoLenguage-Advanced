package configs

type Email struct {
	Email string
	Password string
	Address string
}

func NewConfig() *Email {
		return &Email{
		Email : "workingprogramms508@gmail.com",
		Password: "wbnzlkdjegdotqwp",
		Address: "smtp.gmail.com:587",
	}
}