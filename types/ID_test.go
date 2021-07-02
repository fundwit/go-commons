package types_test

import (
	"encoding/json"
	"testing"

	"github.com/fundwit/go-commons/types"
	. "github.com/onsi/gomega"
)

type testStruct struct {
	Id types.ID
}

func TestIDMarshalJSON(t *testing.T) {
	RegisterTestingT(t)

	t.Run("should be MarshalJSON to string", func(t *testing.T) {
		bytes, err := json.Marshal(&testStruct{Id: types.ID(123)})
		Expect(err).To(BeNil())
		Expect(string(bytes)).To(MatchJSON(`{"Id":"123"}`))
	})
	t.Run("should not be able to MarshalJSON to string when assigned to an interface{} variable", func(t *testing.T) {
		bytes, err := json.Marshal(&testStruct{Id: types.ID(123)})
		Expect(err).To(BeNil())
		Expect(string(bytes)).To(MatchJSON(`{"Id":"123"}`))
	})
}

func TestIDUnmarshalJSON(t *testing.T) {
	RegisterTestingT(t)

	t.Run("should be UnmarshalJSON from string", func(t *testing.T) {
		ts := testStruct{}
		err := json.Unmarshal([]byte(`{"Id":"123"}`), &ts)
		Expect(err).To(BeNil())
		Expect(ts.Id).To(Equal(types.ID(123)))
	})
	t.Run("should be UnmarshalJSON from number", func(t *testing.T) {
		ts := testStruct{}
		err := json.Unmarshal([]byte(`{"Id":123}`), &ts)
		Expect(err).To(BeNil())
		Expect(ts.Id).To(Equal(types.ID(123)))
	})
	t.Run("should be UnmarshalJSON from non number value", func(t *testing.T) {
		ts := testStruct{}
		err := json.Unmarshal([]byte(`{"Id":"abc"}`), &ts)
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal(`strconv.ParseUint: parsing "abc": invalid syntax`))
		Expect(ts.Id).To(BeZero())
	})
}

func TestIDString(t *testing.T) {
	RegisterTestingT(t)

	t.Run("should be able to convert to string", func(t *testing.T) {
		Expect(types.ID(123).String()).To(Equal("123"))
	})

	t.Run("should be able to determine is zero", func(t *testing.T) {
		Expect(types.ID(123).IsZero()).To(BeFalse())
		Expect(types.ID(0).IsZero()).To(BeTrue())
	})
}
