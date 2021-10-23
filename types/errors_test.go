package types

import (
	"errors"
	"testing"

	. "github.com/onsi/gomega"
)

func TestErrInvalidParameter(t *testing.T) {
	RegisterTestingT(t)

	t.Run("Error function should work as expected", func(t *testing.T) {
		e := &ErrInvalidParameter{Parameter: "123", Cause: errors.New("inner error")}
		Expect(e.Error()).To(Equal(`invalid parameter: "123"`))

		type user struct {
			Name string
		}
		e = &ErrInvalidParameter{Parameter: user{Name: "ann"}, Cause: errors.New("inner error")}
		Expect(e.Error()).To(Equal(`invalid parameter: types.user{Name:"ann"}`))
	})

	t.Run("Unwarp function should work as expected", func(t *testing.T) {
		e := &ErrInvalidParameter{Parameter: "123", Cause: errors.New("inner error")}
		Expect(errors.Unwrap(e)).To(Equal(errors.New("inner error")))

		e = &ErrInvalidParameter{Parameter: "123"}
		Expect(errors.Unwrap(e)).To(BeNil())
	})
}
