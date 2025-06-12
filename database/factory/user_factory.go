package factory

import (
	"math/rand"
	"time"

	"github.com/username/go-app/internal/auth/domain"

	"github.com/bxcodec/faker/v3"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func GenerateFakeUser() *domain.User {
    email := faker.Email()
    password := "password123"

    return &domain.User{
        ID:       uint(rand.Intn(1000) + 1),
        Name:     faker.Name(),
        Email:    email,
        Password: password,
    }
}

func GenerateManyFakeUsers(n int) []*domain.User {
    var users []*domain.User
    for i := 0; i < n; i++ {
        users = append(users, GenerateFakeUser())
    }
    return users
}
