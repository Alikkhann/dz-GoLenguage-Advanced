package resp

import (
	"math/rand"
	"time"
)

func GenerateVerificationCode() int {
	src := rand.NewSource(time.Now().UnixNano()) // Создаем новый источник случайности
	r := rand.New(src)                           // Создаем новый генератор случайных чисел
	return r.Intn(900000) + 100000               //код будет шестизначным
}
