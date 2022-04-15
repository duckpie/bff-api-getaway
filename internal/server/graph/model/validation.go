package model

import (
	"regexp"

	"github.com/duckpie/cherry"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (m *Login) Validation() error {
	return validation.ValidateStruct(
		m,

		validation.Field(&m.Login,
			validation.Required,
			validation.Match(regexp.MustCompile(cherry.RegexName)),
		),

		validation.Field(&m.Password,
			validation.Length(8, 20),
			validation.Required,
		),
	)
}

func (m *NewUser) Validation() error {
	return validation.ValidateStruct(
		m,

		validation.Field(&m.Email,
			is.Email,
			validation.Required,
		),

		validation.Field(&m.Login,
			validation.Required,
			validation.Match(regexp.MustCompile(cherry.RegexName)),
		),

		validation.Field(&m.Password,
			validation.Length(8, 20),
			validation.Required,
		),
	)
}

func (m *UpdateUser) Validation() error {
	return validation.ValidateStruct(
		m,

		validation.Field(&m.UUID,
			is.UUIDv4,
			validation.Required,
		),

		validation.Field(&m.Email,
			is.Email,
			validation.Required,
		),

		validation.Field(&m.Login,
			validation.Required,
			validation.Match(regexp.MustCompile(cherry.RegexName)),
		),

		validation.Field(&m.Role,
			validation.Required,
			validation.In(
				0,
				1,
				2,
				3,
				4,
			),
		),
	)
}
