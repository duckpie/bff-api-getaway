package server_test

import (
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

func Test_Server_User(t *testing.T) {
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
		mutation ($login: String!, $email: String!, $password: String!){
			createUser(input: {login: $login, email: $email, password: $password}){
				uuid
				login
				email
				role
			}
		}`, &resp,
			client.Var("login", "tester1"),
			client.Var("email", "tester1@gmail.com"),
			client.Var("password", "12344321"),
		)

		assert.NoError(t, validation.Validate(resp.CreateUser.UUID, is.UUIDv4))
		assert.Equal(t, "tester1", resp.CreateUser.Login)

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

		c.MustPost(`
		query ($uuid: String!){
			getUserByUuid(uuid: $uuid){
				uuid
				login
			}
		}`, &resp,
			client.Var("uuid", uuid),
		)

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

		c.MustPost(`
		query ($login: String!){
			getUserByLogin(login: $login){
				uuid
				login
			}
		}
		`, &resp,
			client.Var("login", login),
		)

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
		query ($limit: Int!, $offset: Int!) {
			getUsersSlice(limit: $limit, offset: $offset){
			  limit
			  offset
			  total
			  data{
				uuid
				login
			  }
			  lastPage
			}
		}`, &resp,
			client.Var("limit", 30),
			client.Var("offset", 0),
		)

		assert.NotNil(t, resp)
	})

	t.Run("update_user", func(t *testing.T) {
		var resp struct {
			UpdateUser struct {
				UUID string
				Role int
			}
		}

		c.MustPost(`
		mutation ($uuid: String!, $login: String!, $email: String!, $role: Int!) {
			updateUser(input: {uuid: $uuid, login: $login, email: $email, role: $role}){
			  uuid			
			  role
			}
		}`, &resp,
			client.Var("uuid", uuid),
			client.Var("login", login),
			client.Var("email", email),
			client.Var("role", 1),
		)

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

		c.MustPost(`
		mutation ($uuid: String!){
			deleteUser(uuid: $uuid){
			  uuid
			}
		}`, &resp,
			client.Var("uuid", uuid),
		)

		assert.Equal(t, uuid, resp.DeleteUser.UUID)
	})
}
