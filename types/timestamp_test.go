package types_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/fundwit/go-commons/types"
	. "github.com/onsi/gomega"
)

func TestTimestampValue(t *testing.T) {
	RegisterTestingT(t)

	t.Run("should be able to calculate database value correctly", func(t *testing.T) {
		v, err := types.Timestamp{}.Value()
		Expect(err).To(BeNil())
		Expect(v).To(Equal("0001-01-01 00:00:00.000000000"))

		ts := types.TimestampOfDate(2021, 5, 6, 12, 30, 40, 666666666, types.TimeZoneCST8)
		v, err = ts.Value()
		Expect(err).To(BeNil())
		Expect(v).To(Equal("2021-05-06 12:30:40.666667000"))                                             // Local time
		Expect(v).To(Equal("2021-05-06 " + strconv.Itoa(ts.Time().Local().Hour()) + ":30:40.666667000")) // Local time

		ts = types.TimestampOfDate(2021, 5, 6, 12, 30, 40, 666666666, time.UTC)
		v, err = ts.Value()
		Expect(err).To(BeNil())
		Expect(v).To(Equal("2021-05-06 20:30:40.666667000"))                                             // Local time
		Expect(v).To(Equal("2021-05-06 " + strconv.Itoa(ts.Time().Local().Hour()) + ":30:40.666667000")) // Local time
	})

	t.Run("should be able to scan value", func(t *testing.T) {
		// mysql datetime format
		ts := types.TimestampOfDate(2021, 1, 1, 12, 30, 40, 666666666, time.Local)
		Expect(ts.Scan("0001-01-01 00:00:00.000000000")).To(BeNil())
		Expect(ts.IsZero()).To(BeTrue())
		Expect(ts.Time().IsZero()).To(BeTrue())
		Expect(ts).To(Equal(types.Timestamp{}))

		Expect(ts.Scan("0001-01-01 01:02:03.004")).To(BeNil())
		Expect(ts.IsZero()).To(BeTrue())
		Expect(ts.Time().IsZero()).To(BeTrue())
		Expect(ts).To(Equal(types.Timestamp{}))

		Expect(ts.Scan("0000-01-01 00:00:00")).To(BeNil())
		Expect(ts.IsZero()).To(BeTrue())
		Expect(ts.Time().IsZero()).To(BeTrue())
		Expect(ts).To(Equal(types.Timestamp{}))

		Expect(ts.Scan("2021-05-06 12:30:40.666666666")).To(BeNil())
		Expect(ts.IsZero()).To(BeFalse())
		Expect(ts).To(Equal(types.TimestampOfDate(2021, 5, 6, 12, 30, 40, 666666666, time.Local)))
		Expect(ts.String()).To(Equal(ts.Time().String()))

		// RFC3339 time format
		Expect(ts.Scan("2021-05-06T12:30:40.666666666Z")).To(BeNil())
		Expect(ts).To(Equal(types.TimestampOfDate(2021, 5, 6, 12, 30, 40, 666666666, time.UTC)))

		Expect(ts.Scan("2021-05-06T12:30:40Z")).To(BeNil())
		Expect(ts).To(Equal(types.TimestampOfDate(2021, 5, 6, 12, 30, 40, 0, time.UTC)))

		Expect(ts.Scan("2021-05-06T12:30:40.001Z")).To(BeNil())
		Expect(ts).To(Equal(types.TimestampOfDate(2021, 5, 6, 12, 30, 40, 1000000, time.UTC)))

		Expect(ts.Scan("2021-05-06T12:30:40+07:01")).To(BeNil())
		Expect(ts).To(Equal(types.TimestampOfDate(2021, 5, 6, 12, 30, 40, 0, time.FixedZone("", int(7*time.Hour/time.Second+60)))))

		Expect(ts.Scan("2021-05-06T12:30:40-07:01")).To(BeNil())
		Expect(ts).To(Equal(types.TimestampOfDate(2021, 5, 6, 12, 30, 40, 0, time.FixedZone("", -int(7*time.Hour/time.Second+60)))))

		// change location
		ts.Time().In(time.UTC).Equal(time.Date(2021, 5, 6, 19, 31, 40, 0, time.UTC))

		// error
		Expect(ts.Scan(1.24).Error()).To(Equal("invalid parameter: 1.24"))
		Expect(ts.Scan("someT123").Error()).To(Equal(`invalid parameter: "someT123"`))
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
