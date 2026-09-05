package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTData struct {
	Phone string
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(data *JWTData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ // Создаем JWT токен, указывая метод подписи HMAC-SHA256 и включая email в клеймы для его создания // claims(клеймы) - набор данных, которые включены в токен
		"phone": data.Phone,
	})
	signed, err := token.SignedString([]byte(j.Secret)) // мы подписали созданный jwt и возвращаем строку которая будет содержать подписанный токен в виде строки JWT
	if err != nil {
		return "", err
	}
	return signed, nil
}

func (j *JWT) Parse(token string) (bool, *JWTData) {
	tkn, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {	
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}
	phone := tkn.Claims.(jwt.MapClaims)["phone"]
	return tkn.Valid, &JWTData{
		Phone: phone.(string),
	}
}