package server_test

import (
	"fmt"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/stretchr/testify/assert"
	"github.com/wrs-news/bff-api-getaway/internal/server"
	"github.com/wrs-news/bff-api-getaway/internal/server/graph/generated"
)

var (
	uuid  string
	login string
	email string
	role  int
)

func Test_Create_User(t *testing.T) {
	s := server.InitServer(testConfig)
	assert.NoError(t, s.Config())

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(s.Resolver().Config())))

	t.Run("create_user", func(t *testing.T) {
		var resp struct {
			CreateUser struct {
				UUID  string
				Login string
				Email string
				Role  int
			}
		}

		c.MustPost(`
		mutation {
			createUser(input: {login: "I0HuKc", email: "test@gmail.com", password: "12344321"}){
				uuid
				login
				email
				role
			}
		}
		`, &resp)

		assert.NoError(t, validation.Validate(resp.CreateUser.UUID, is.UUIDv4))
		assert.Equal(t, "I0HuKc", resp.CreateUser.Login)

		uuid = resp.CreateUser.UUID
		login = resp.CreateUser.Login
		email = resp.CreateUser.Email
		role = resp.CreateUser.Role
	})

	t.Run("get_user_by_uuid", func(t *testing.T) {
		var resp struct {
			GetUserByUuid struct {
				UUID  string
				Login string
			}
		}

		c.MustPost(fmt.Sprintf(`
			query {
				getUserByUuid(uuid: "%s"){
				uuid
				login
				}
			}
		`,
			uuid,
		), &resp)

		assert.Equal(t, uuid, resp.GetUserByUuid.UUID)
		assert.Equal(t, login, resp.GetUserByUuid.Login)
	})

	t.Run("get_user_by_login", func(t *testing.T) {
		var resp struct {
			GetUserByLogin struct {
				UUID  string
				Login string
			}
		}

		c.MustPost(fmt.Sprintf(`
			query {
				getUserByLogin(login: "%s"){
				uuid
				login
				}
			}
		`,
			login,
		), &resp)

		assert.Equal(t, uuid, resp.GetUserByLogin.UUID)
		assert.Equal(t, login, resp.GetUserByLogin.Login)
	})

	t.Run("get_users_slice", func(t *testing.T) {
		var resp struct {
			GetUsersSlice struct {
				Limit  int
				Offset int
				Data   []struct {
					UUID  string
					Login string
				}
				Total    int
				LastPage int
			}
		}

		c.MustPost(`
		query {
			getUsersSlice(limit: 30, offset: 0){
			  limit
			  offset
			  total
			  data{
				uuid
				login
			  }
			  lastPage
			}
		  }
		`, &resp)

		assert.Len(t, resp.GetUsersSlice.Data, 1)
		assert.Equal(t, 1, resp.GetUsersSlice.Total)
	})

	t.Run("update_user", func(t *testing.T) {
		var resp struct {
			UpdateUser struct {
				UUID string
				Role int
			}
		}

		c.MustPost(fmt.Sprintf(`
		mutation {
			updateUser(input: {uuid:"%s", login: "%s", email: "%s", role: %d}){
			  uuid			
			  role
			}
		  }
		`,
			uuid,
			login,
			email,
			1,
		), &resp)

		assert.Equal(t, uuid, resp.UpdateUser.UUID)
		assert.Equal(t, 1, resp.UpdateUser.Role)
	})

	t.Run("delete_user", func(t *testing.T) {
		var resp struct {
			DeleteUser struct {
				UUID string
				Role int
			}
		}

		c.MustPost(fmt.Sprintf(`
		mutation {
			deleteUser(uuid: "%s"){
			  uuid
			}
		  }
		`,
			uuid,
		), &resp)

		assert.Equal(t, uuid, resp.DeleteUser.UUID)
	})
}
