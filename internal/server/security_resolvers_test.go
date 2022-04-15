package server_test

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/assert"
	"github.com/wrs-news/bff-api-getaway/internal/server"
	"github.com/wrs-news/bff-api-getaway/internal/server/graph/generated"
	pbu "github.com/wrs-news/golang-proto/pkg/proto/user"
)

func Test_Server_Security(t *testing.T) {
	s := server.InitServer(testConfig)
	assert.NoError(t, s.Config())

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(s.Resolver().Config())))

	u := pbu.User{
		Login: "tester2",
		Email: "tester2@gmail.com",
		Hash:  "12344321", // Password
	}
	teardown, err := CreateTestUser(t, c, &u)
	assert.NoError(t, err)
	defer teardown()

	var authResp struct {
		CreateAuth struct {
			RefreshToken string
			AccessToken  string
		}
	}

	t.Run("create_auth", func(t *testing.T) {
		c.MustPost(`
		mutation ($login: String!, $password: String!) {
			createAuth(input: {login: $login, password: $password}) {
				refreshToken
				accessToken
			}
		}`, &authResp,
			client.Var("login", u.Login),
			client.Var("password", u.Hash),
		)

		assert.NotNil(t, authResp.CreateAuth)
	})

	t.Run("check_auth", func(t *testing.T) {
		var resp struct {
			AuthCheck struct {
				Status string
			}
		}

		c.MustPost(`
		query ($accessToken: String!) {
			authCheck(accessToken: $accessToken) {
				status
			}
		}`, &resp,
			client.Var("accessToken", authResp.CreateAuth.AccessToken),
		)

		assert.Equal(t, "ok", resp.AuthCheck.Status)
	})

	t.Run("refresh_token", func(t *testing.T) {
		var resp struct {
			RefreshToken struct {
				RefreshToken string
				AccessToken  string
			}
		}

		c.MustPost(`
		mutation($refreshToken: String!) {
			refreshToken(token: $refreshToken) {
				accessToken
				refreshToken
			}
		}`, &resp,
			client.Var("refreshToken", authResp.CreateAuth.RefreshToken),
		)

		assert.NotEqual(t, authResp.CreateAuth.AccessToken, resp.RefreshToken.AccessToken)
		assert.NotEqual(t, authResp.CreateAuth.RefreshToken, resp.RefreshToken.RefreshToken)

		authResp.CreateAuth.AccessToken = resp.RefreshToken.AccessToken
		authResp.CreateAuth.RefreshToken = resp.RefreshToken.RefreshToken
	})

	t.Run("logout", func(t *testing.T) {
		var resp struct {
			Logout struct {
				Status string
			}
		}

		c.MustPost(`
		mutation($accessToken: String!) {
			logout(accessToken: $accessToken){
			  status
			}
		  }`, &resp,
			client.Var("accessToken", authResp.CreateAuth.AccessToken),
		)
	})
}
