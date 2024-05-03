// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/delivery_extension/object.proto

package delivery_extension

import (
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *DeliveryExtension) CloneVT() *DeliveryExtension {
	if m == nil {
		return (*DeliveryExtension)(nil)
	}
	r := new(DeliveryExtension)
	r.Meta = m.Meta.CloneVT()
	r.Config = m.Config.CloneVT()
	r.State = m.State.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *DeliveryExtension) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *DeliveryExtensionState) CloneVT() *DeliveryExtensionState {
	if m == nil {
		return (*DeliveryExtensionState)(nil)
	}
	r := new(DeliveryExtensionState)
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *DeliveryExtensionState) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *DeliveryExtension) StableEqualVT(that *DeliveryExtension) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Meta.StableEqualVT(that.Meta) {
		return false
	}
	if !this.Config.StableEqualVT(that.Config) {
		return false
	}
	if !this.State.StableEqualVT(that.State) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DeliveryExtension) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DeliveryExtension)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *DeliveryExtension) EqualVT(that *DeliveryExtension) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Meta.EqualVT(that.Meta) {
		return false
	}
	if !this.Config.EqualVT(that.Config) {
		return false
	}
	if !this.State.EqualVT(that.State) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DeliveryExtension) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DeliveryExtension)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *DeliveryExtensionState) StableEqualVT(that *DeliveryExtensionState) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DeliveryExtensionState) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DeliveryExtensionState)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *DeliveryExtensionState) EqualVT(that *DeliveryExtensionState) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DeliveryExtensionState) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DeliveryExtensionState)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}