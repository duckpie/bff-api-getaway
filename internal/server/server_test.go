package server_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/BurntSushi/toml"
	"github.com/wrs-news/bff-api-getaway/internal/config"
	pbu "github.com/wrs-news/golang-proto/pkg/proto/user"
)

var (
	testConfig = config.NewConfig()
)

func TestMain(m *testing.M) {
	if _, err := toml.DecodeFile("../../config/config.test.toml", testConfig); err != nil {
		log.Fatalf(err.Error())
	}

	os.Exit(m.Run())
}

func CreateTestUser(t *testing.T, c *client.Client, u *pbu.User) (func(), error) {
	t.Helper()

	// Создание тестового пользователя
	var resp struct {
		CreateUser struct {
			UUID  string
			Login string
			Email string
			Role  int
		}
	}

	c.MustPost(`
	mutation ($login: String!, $email: String!, $password: String!) {
		createUser(input: {login: $login, email: $email, password: $password}){
			uuid
			login
			email
			role
		}
	}`, &resp,
		client.Var("login", u.Login),
		client.Var("email", u.Email),
		client.Var("password", u.Hash),
	)

	// Сохранение данных о созданном пользователе
	u.Uuid = resp.CreateUser.UUID

	// Удаление тестового пользователя.
	// Использовать как отложенную функция в методе
	// где вызывается создание тестового пользователя.
	return func() {
		fmt.Printf("================ Delete user %s ================\n", u.Login)

		var resp struct {
			DeleteUser struct {
				UUID string
				Role int
			}
		}

		c.MustPost(`
		mutation ($uuid: String!){
			deleteUser(uuid: $uuid){
			  uuid
			}
		}`, &resp,
			client.Var("uuid", u.Uuid),
		)
	}, nil
}
