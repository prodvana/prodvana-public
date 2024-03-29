// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/common_config/notification.proto

package common_config

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

func (m *NotificationConfig_Slack) CloneVT() *NotificationConfig_Slack {
	if m == nil {
		return (*NotificationConfig_Slack)(nil)
	}
	r := new(NotificationConfig_Slack)
	r.Channel = m.Channel
	r.ErrorChannel = m.ErrorChannel
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *NotificationConfig_Slack) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *NotificationConfig) CloneVT() *NotificationConfig {
	if m == nil {
		return (*NotificationConfig)(nil)
	}
	r := new(NotificationConfig)
	r.Slack = m.Slack.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *NotificationConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *NotificationConfig_Slack) StableEqualVT(that *NotificationConfig_Slack) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Channel != that.Channel {
		return false
	}
	if this.ErrorChannel != that.ErrorChannel {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *NotificationConfig_Slack) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*NotificationConfig_Slack)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *NotificationConfig) StableEqualVT(that *NotificationConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Slack.StableEqualVT(that.Slack) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *NotificationConfig) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*NotificationConfig)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *NotificationConfig_Slack) EqualVT(that *NotificationConfig_Slack) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Channel != that.Channel {
		return false
	}
	if this.ErrorChannel != that.ErrorChannel {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *NotificationConfig_Slack) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*NotificationConfig_Slack)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *NotificationConfig) EqualVT(that *NotificationConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Slack.EqualVT(that.Slack) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *NotificationConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*NotificationConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
