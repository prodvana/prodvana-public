package types

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TimestampProto(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return nil
	}
	return timestamppb.New(t)
}

func TimestampProtoFromPtr(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return TimestampProto(*t)
}
