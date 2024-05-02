package prototestutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func RequireEquals(t *testing.T, expected, actual proto.Message, msg ...interface{}) {
	var extraMessage string
	if len(msg) > 1 {
		extraMessage = fmt.Sprintf(msg[0].(string)+"\n", msg[1:]...)
	}
	require.True(t, proto.Equal(expected, actual), "%sProto mismatch:\nexpected:\n%+v\nactual:\n%+v", extraMessage, expected, actual)
}

func RequireListEquals[AT proto.Message, BT proto.Message](t *testing.T, expected []AT, actual []BT, msg ...interface{}) {
	var extraMessage string
	if len(msg) > 1 {
		extraMessage = fmt.Sprintf(msg[0].(string)+"\n", msg[1:]...)
	}
	require.Equal(t, len(expected), len(actual), msg...)
	for idx := range expected {
		RequireEquals(t, expected[idx], actual[idx], "%sIndex %d", extraMessage, idx)
	}
}

func RequireEqualsAnyOrder[AT proto.Message, BT proto.Message](t *testing.T, expected []AT, actual []BT, msg ...interface{}) {
	var extraMessage string
	if len(msg) > 1 {
		extraMessage = fmt.Sprintf(msg[0].(string)+"\n", msg[1:]...)
	}
	require.Equal(t, len(expected), len(actual), msg...)
	used := make([]bool, len(actual))
	for eI := range expected {
		found := false
		for aI := range actual {
			if used[aI] {
				continue
			}
			if proto.Equal(expected[eI], actual[aI]) {
				found = true
				break
			}
		}
		require.True(t, found, "%sMissing element %d from expected list:\n%+v\n%+v", extraMessage, eI, expected, actual)
	}
}

func RequireMapEquals[KT comparable, AT proto.Message, BT proto.Message](t *testing.T, expected map[KT]AT, actual map[KT]BT, msg ...interface{}) {
	var extraMessage string
	if len(msg) > 1 {
		extraMessage = fmt.Sprintf(msg[0].(string)+"\n", msg[1:]...)
	}
	require.Equal(t, len(expected), len(actual), msg...)
	for key := range expected {
		RequireEquals(t, expected[key], actual[key], "%sKey %d", extraMessage, key)
	}
}
