package types_test

import (
	"testing"
	"time"

	"github.com/fundwit/go-commons/types"
	. "github.com/onsi/gomega"
)

func TestTimestampValue(t *testing.T) {
	RegisterTestingT(t)

	t.Run("should be able to calculate value correctly", func(t *testing.T) {
		v, err := types.Timestamp{}.Value()
		Expect(err).To(BeNil())
		Expect(v).To(Equal("0001-01-01 00:00:00.000000000"))

		v, err = types.TimestampOfDate(2021, 5, 6, 12, 30, 40, 666666666, time.Local).Value()
		Expect(err).To(BeNil())
		Expect(v).To(Equal("2021-05-06 12:30:40.666667000"))
	})

	t.Run("should be able to scan value", func(t *testing.T) {
		ts := types.TimestampOfDate(2021, 1, 1, 12, 30, 40, 666666666, time.Local)
		Expect(ts.Scan("0001-01-01 00:00:00.000000000")).To(BeNil())
		Expect(ts.Time().IsZero()).To(BeTrue())
		Expect(ts).To(Equal(types.Timestamp{}))

		Expect(ts.Scan("0001-01-01 01:02:03.004")).To(BeNil())
		Expect(ts.Time().IsZero()).To(BeTrue())
		Expect(ts).To(Equal(types.Timestamp{}))

		Expect(ts.Scan("0000-01-01 00:00:00")).To(BeNil())
		Expect(ts.Time().IsZero()).To(BeTrue())
		Expect(ts).To(Equal(types.Timestamp{}))

		Expect(ts.Scan("2021-05-06 12:30:40.666666666")).To(BeNil())
		Expect(ts).To(Equal(types.TimestampOfDate(2021, 5, 6, 12, 30, 40, 666666666, time.Local)))
	})
}

func TestTimestampCurrentTimestamp(t *testing.T) {
	RegisterTestingT(t)

	t.Run("should be able to calculate value correctly", func(t *testing.T) {
		begin := time.Now().Round(time.Microsecond)
		v := types.CurrentTimestamp()
		end := time.Now().Round(time.Microsecond)

		Expect(v.Time().UnixNano() >= begin.UnixNano()).To(BeTrue())
		Expect(v.Time().UnixNano() <= end.UnixNano()).To(BeTrue())
	})
}

func TestTimestampJSON(t *testing.T) {
	RegisterTestingT(t)

	t.Run("should be able to marshal json", func(t *testing.T) {
		ts := types.TimestampOfDate(2021, 1, 1, 12, 30, 40, 666666666, time.UTC)
		jsonBytes, err := ts.MarshalJSON()
		Expect(err).To(BeNil())
		Expect(string(jsonBytes)).To(Equal(`"2021-01-01T12:30:40.666667Z"`))

		var t1 types.Timestamp
		Expect(t1.UnmarshalJSON(jsonBytes)).To(BeNil())
		Expect(t1).To(Equal(ts))

		jsonBytes, err = types.Timestamp{}.MarshalJSON()
		Expect(err).To(BeNil())
		Expect(string(jsonBytes)).To(Equal(`null`))

		var t2 types.Timestamp
		Expect(t2.UnmarshalJSON(jsonBytes)).To(BeNil())
		Expect(t2.Time().IsZero()).To(BeTrue())
	})
}

func TestTimestampMarshalText(t *testing.T) {
	RegisterTestingT(t)

	t.Run("should be able to marshal text", func(t *testing.T) {
		ts := types.TimestampOfDate(2021, 1, 1, 12, 30, 40, 666666666, time.UTC)
		textBytes, err := ts.MarshalText()
		Expect(err).To(BeNil())
		Expect(string(textBytes)).To(Equal(`2021-01-01T12:30:40.666667Z`))

		var t1 types.Timestamp
		Expect(t1.UnmarshalText(textBytes)).To(BeNil())
		Expect(t1).To(Equal(ts))
	})
}
