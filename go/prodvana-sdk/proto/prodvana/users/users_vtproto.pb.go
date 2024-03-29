// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/users/users.proto

package users

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

func (m *User) CloneVT() *User {
	if m == nil {
		return (*User)(nil)
	}
	r := new(User)
	r.UserId = m.UserId
	r.ApiUser = m.ApiUser
	r.Name = m.Name
	r.FirstName = m.FirstName
	r.LastName = m.LastName
	if rhs := m.Emails; rhs != nil {
		tmpContainer := make([]string, len(rhs))
		copy(tmpContainer, rhs)
		r.Emails = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *User) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *User) StableEqualVT(that *User) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.UserId != that.UserId {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	if this.FirstName != that.FirstName {
		return false
	}
	if this.LastName != that.LastName {
		return false
	}
	if len(this.Emails) != len(that.Emails) {
		return false
	}
	for i, vx := range this.Emails {
		vy := that.Emails[i]
		if vx != vy {
			return false
		}
	}
	if this.ApiUser != that.ApiUser {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *User) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*User)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *User) EqualVT(that *User) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.UserId != that.UserId {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	if this.FirstName != that.FirstName {
		return false
	}
	if this.LastName != that.LastName {
		return false
	}
	if len(this.Emails) != len(that.Emails) {
		return false
	}
	for i, vx := range this.Emails {
		vy := that.Emails[i]
		if vx != vy {
			return false
		}
	}
	if this.ApiUser != that.ApiUser {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *User) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*User)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
