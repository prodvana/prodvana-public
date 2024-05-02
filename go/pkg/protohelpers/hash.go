package protohelpers

import "google.golang.org/protobuf/proto"

type ProtoHash string

func MustHashProto(msg proto.Message) ProtoHash {
	b, err := proto.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return ProtoHash(b)
}
