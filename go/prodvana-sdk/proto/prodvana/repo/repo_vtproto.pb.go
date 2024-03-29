// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/repo/repo.proto

package repo

import (
	timestamppb1 "github.com/planetscale/vtprotobuf/types/known/timestamppb"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
	r.Name = m.Name
	r.Email = m.Email
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *User) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *Commit) CloneVT() *Commit {
	if m == nil {
		return (*Commit)(nil)
	}
	r := new(Commit)
	r.CommitId = m.CommitId
	r.Url = m.Url
	r.Message = m.Message
	r.Author = m.Author.CloneVT()
	r.ImpactAnalysis = m.ImpactAnalysis.CloneVT()
	r.CommitTimestamp = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.CommitTimestamp).CloneVT())
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *Commit) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *FuzzyCommit) CloneVT() *FuzzyCommit {
	if m == nil {
		return (*FuzzyCommit)(nil)
	}
	r := new(FuzzyCommit)
	r.CommitIsh = m.CommitIsh
	r.Source = m.Source
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *FuzzyCommit) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *User) StableEqualVT(that *User) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	if this.Email != that.Email {
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
	if this.Name != that.Name {
		return false
	}
	if this.Email != that.Email {
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
func (this *Commit) StableEqualVT(that *Commit) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.CommitId != that.CommitId {
		return false
	}
	if this.Url != that.Url {
		return false
	}
	if this.Message != that.Message {
		return false
	}
	if !this.Author.StableEqualVT(that.Author) {
		return false
	}
	if !this.ImpactAnalysis.StableEqualVT(that.ImpactAnalysis) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.CommitTimestamp).StableEqualVT((*timestamppb1.Timestamp)(that.CommitTimestamp)) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Commit) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Commit)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *Commit) EqualVT(that *Commit) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.CommitId != that.CommitId {
		return false
	}
	if this.Url != that.Url {
		return false
	}
	if this.Message != that.Message {
		return false
	}
	if !this.Author.EqualVT(that.Author) {
		return false
	}
	if !this.ImpactAnalysis.EqualVT(that.ImpactAnalysis) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.CommitTimestamp).EqualVT((*timestamppb1.Timestamp)(that.CommitTimestamp)) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Commit) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Commit)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *FuzzyCommit) StableEqualVT(that *FuzzyCommit) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.CommitIsh != that.CommitIsh {
		return false
	}
	if this.Source != that.Source {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *FuzzyCommit) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*FuzzyCommit)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *FuzzyCommit) EqualVT(that *FuzzyCommit) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.CommitIsh != that.CommitIsh {
		return false
	}
	if this.Source != that.Source {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *FuzzyCommit) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*FuzzyCommit)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
