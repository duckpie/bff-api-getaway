package model

import (
	"regexp"

	"github.com/duckpie/cherry"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (s *NewUser) Validation() error {
	return validation.ValidateStruct(
		s,

		validation.Field(&s.Email,
			is.Email,
			validation.Required,
		),

		validation.Field(&s.Login,
			validation.Required,
			validation.Match(regexp.MustCompile(cherry.RegexName)),
		),

		validation.Field(&s.Password,
			validation.Length(8, 20),
			validation.Required,
		),
	)
}

func (s *UpdateUser) Validation() error {
	return validation.ValidateStruct(
		s,

		validation.Field(&s.UUID,
			is.UUIDv4,
			validation.Required,
		),

		validation.Field(&s.Email,
			is.Email,
			validation.Required,
		),

		validation.Field(&s.Login,
			validation.Required,
			validation.Match(regexp.MustCompile(cherry.RegexName)),
		),

		validation.Field(&s.Role,
			validation.Required,
			validation.In(0, 1, 2, 3),
		),
	)
}
