package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {
	testCases := []struct {
		email    string
		expected error
	}{
		{
			"muratov@gmail.com",
			nil,
		},
		{
			"muratovgmail.com",
			ErrInvalidEmail,
		},
		{
			"muratov@gmail.c",
			ErrInvalidEmail,
		},
		{
			"",
			ErrInvalidEmail,
		},
		{
			"мuratov@.ru",
			ErrInvalidEmail,
		},
	}
	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, validateEmail(testCase.email))
	}
}
