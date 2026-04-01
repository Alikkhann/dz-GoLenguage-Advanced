package jwt

import "github.com/golang-jwt/jwt/v5"

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ // Создаем JWT токен, указывая метод подписи HMAC-SHA256 и включая email в клеймы для его создания // claims(клеймы) - набор данных, которые включены в токен
		"email": email,
	})
	signed, err := token.SignedString([]byte(j.Secret)) // мы подписали созданный jwt и возвращаем строку которая будет содержать подписанный токен в виде строки JWT
	if err != nil {
		return "", err
	}
	return signed, nil
}