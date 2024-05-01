package types

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func StringPtr(s string) *string {
	return &s
}

func StringPtrIfNotEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func StringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func BoolPtr(b bool) *bool {
	return &b
}

func BoolValue(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func Int64Ptr(i int64) *int64 {
	return &i
}

func Int64Value(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

func IntValue(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func TimePtr(time time.Time) *time.Time {
	return &time
}

func TimeValue(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}

func DurationPtr(d time.Duration) *time.Duration {
	return &d
}

func TimeProto(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func Value[T any](v *T) T {
	if v == nil {
		var empty T
		return empty
	}
	return *v
}

func Ptr[T any](v T) *T {
	return &v
}

func PtrIfNotEmpty[T comparable](v T) *T {
	var empty T
	if empty == v {
		return nil
	}
	return &v
}
