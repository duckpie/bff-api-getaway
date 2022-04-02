package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wrs-news/bff-api-getaway/internal/server/graph/model"
)

func Test_UpdateUser_Validation(t *testing.T) {
	testCases := []struct {
		name    string
		payload model.UpdateUser
		ok      bool
	}{
		{
			name: "empty_uuid",
			payload: model.UpdateUser{
				UUID:  "",
				Email: "wrc@gmail.com",
				Login: "I0HuKc",
				Role:  1,
			},
			ok: false,
		},

		{
			name: "invalid_uuid",
			payload: model.UpdateUser{
				UUID:  "invalid",
				Email: "wrc@gmail.com",
				Login: "I0HuKc",
				Role:  1,
			},
			ok: false,
		},

		{
			name: "uuid_v1",
			payload: model.UpdateUser{
				UUID:  "b08d5878-b1c9-11ec-b909-0242ac120002",
				Email: "wrc@gmail.com",
				Login: "I0HuKc",
				Role:  1,
			},
			ok: false,
		},

		{
			name: "empty_email",
			payload: model.UpdateUser{
				UUID:  "aabb0112-ff7d-4f1c-bb31-69980d2f6218",
				Email: "",
				Login: "I0HuKc",
				Role:  1,
			},
			ok: false,
		},

		{
			name: "invalid_email",
			payload: model.UpdateUser{
				UUID:  "aabb0112-ff7d-4f1c-bb31-69980d2f6218",
				Email: "invalid",
				Login: "I0HuKc",
				Role:  1,
			},
			ok: false,
		},

		{
			name: "empty_login",
			payload: model.UpdateUser{
				UUID:  "aabb0112-ff7d-4f1c-bb31-69980d2f6218",
				Email: "wrc@gmail.com",
				Login: "",
				Role:  1,
			},
			ok: false,
		},

		{
			name: "invalid_login",
			payload: model.UpdateUser{
				UUID:  "aabb0112-ff7d-4f1c-bb31-69980d2f6218",
				Email: "wrc@gmail.com",
				Login: "invalid login",
				Role:  1,
			},
			ok: false,
		},

		{
			name: "invalid_role",
			payload: model.UpdateUser{
				UUID:  "aabb0112-ff7d-4f1c-bb31-69980d2f6218",
				Email: "wrc@gmail.com",
				Login: "I0HuKc",
				Role:  10,
			},
			ok: false,
		},

		{
			name: "valid",
			payload: model.UpdateUser{
				UUID:  "aabb0112-ff7d-4f1c-bb31-69980d2f6218",
				Email: "wrc@gmail.com",
				Login: "I0HuKc",
				Role:  1,
			},
			ok: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.ok {
				assert.NoError(t, tc.payload.Validation())
				return
			}

			assert.Error(t, tc.payload.Validation())
		})
	}
}

func Test_NewUser_Validation(t *testing.T) {
	testCases := []struct {
		name    string
		payload model.NewUser
		ok      bool
	}{
		{
			name: "empty_email",
			payload: model.NewUser{
				Email:    "",
				Login:    "I0HuKc",
				Password: "password",
			},
			ok: false,
		},

		{
			name: "invalid_email",
			payload: model.NewUser{
				Email:    "invalid",
				Login:    "I0HuKc",
				Password: "password",
			},
			ok: false,
		},

		{
			name: "empty_login",
			payload: model.NewUser{
				Email:    "wrc@gmail.com",
				Login:    "",
				Password: "password",
			},
			ok: false,
		},

		{
			name: "invalid_login",
			payload: model.NewUser{
				Email:    "wrc@gmail.com",
				Login:    "invalid login",
				Password: "password",
			},
			ok: false,
		},

		{
			name: "empty_password",
			payload: model.NewUser{
				Email:    "wrc@gmail.com",
				Login:    "I0HuKc",
				Password: "",
			},
			ok: false,
		},

		{
			name: "short_password",
			payload: model.NewUser{
				Email:    "wrc@gmail.com",
				Login:    "I0HuKc",
				Password: "12345",
			},
			ok: false,
		},

		{
			name: "long_password",
			payload: model.NewUser{
				Email:    "wrc@gmail.com",
				Login:    "I0HuKc",
				Password: "12345678901234567890123",
			},
			ok: false,
		},

		{
			name: "valid",
			payload: model.NewUser{
				Email:    "wrc@gmail.com",
				Login:    "I0HuKc",
				Password: "123456789",
			},
			ok: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.ok {
				assert.NoError(t, tc.payload.Validation())
				return
			}

			assert.Error(t, tc.payload.Validation())
		})
	}
}
