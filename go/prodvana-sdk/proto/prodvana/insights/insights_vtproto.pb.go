// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/insights/insights.proto

package insights

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

func (m *Insight_Subject) CloneVT() *Insight_Subject {
	if m == nil {
		return (*Insight_Subject)(nil)
	}
	r := new(Insight_Subject)
	r.Id = m.Id
	r.Name = m.Name
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *Insight_Subject) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *Insight) CloneVT() *Insight {
	if m == nil {
		return (*Insight)(nil)
	}
	r := new(Insight)
	r.Title = m.Title
	r.Description = m.Description
	r.Class = m.Class
	if m.SubjectOneof != nil {
		r.SubjectOneof = m.SubjectOneof.(interface{ CloneVT() isInsight_SubjectOneof }).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *Insight) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *Insight_Service) CloneVT() isInsight_SubjectOneof {
	if m == nil {
		return (*Insight_Service)(nil)
	}
	r := new(Insight_Service)
	r.Service = m.Service.CloneVT()
	return r
}

func (m *Insight_Application) CloneVT() isInsight_SubjectOneof {
	if m == nil {
		return (*Insight_Application)(nil)
	}
	r := new(Insight_Application)
	r.Application = m.Application.CloneVT()
	return r
}

func (m *Insight_Organization) CloneVT() isInsight_SubjectOneof {
	if m == nil {
		return (*Insight_Organization)(nil)
	}
	r := new(Insight_Organization)
	r.Organization = m.Organization.CloneVT()
	return r
}

func (this *Insight_Subject) StableEqualVT(that *Insight_Subject) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Id != that.Id {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Insight_Subject) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Insight_Subject)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *Insight) StableEqualVT(that *Insight) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.SubjectOneof == nil && that.SubjectOneof != nil {
		return false
	} else if this.SubjectOneof != nil {
		if that.SubjectOneof == nil {
			return false
		}
		if !this.SubjectOneof.(interface {
			StableEqualVT(isInsight_SubjectOneof) bool
		}).StableEqualVT(that.SubjectOneof) {
			return false
		}
	}
	if this.Title != that.Title {
		return false
	}
	if this.Description != that.Description {
		return false
	}
	if this.Class != that.Class {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Insight) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Insight)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *Insight_Service) StableEqualVT(thatIface isInsight_SubjectOneof) bool {
	that, ok := thatIface.(*Insight_Service)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Service, that.Service; p != q {
		if p == nil {
			p = &Insight_Subject{}
		}
		if q == nil {
			q = &Insight_Subject{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Insight_Application) StableEqualVT(thatIface isInsight_SubjectOneof) bool {
	that, ok := thatIface.(*Insight_Application)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Application, that.Application; p != q {
		if p == nil {
			p = &Insight_Subject{}
		}
		if q == nil {
			q = &Insight_Subject{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Insight_Organization) StableEqualVT(thatIface isInsight_SubjectOneof) bool {
	that, ok := thatIface.(*Insight_Organization)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Organization, that.Organization; p != q {
		if p == nil {
			p = &Insight_Subject{}
		}
		if q == nil {
			q = &Insight_Subject{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Insight_Subject) EqualVT(that *Insight_Subject) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Id != that.Id {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Insight_Subject) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Insight_Subject)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Insight) EqualVT(that *Insight) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.SubjectOneof == nil && that.SubjectOneof != nil {
		return false
	} else if this.SubjectOneof != nil {
		if that.SubjectOneof == nil {
			return false
		}
		if !this.SubjectOneof.(interface {
			EqualVT(isInsight_SubjectOneof) bool
		}).EqualVT(that.SubjectOneof) {
			return false
		}
	}
	if this.Title != that.Title {
		return false
	}
	if this.Description != that.Description {
		return false
	}
	if this.Class != that.Class {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Insight) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Insight)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Insight_Service) EqualVT(thatIface isInsight_SubjectOneof) bool {
	that, ok := thatIface.(*Insight_Service)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Service, that.Service; p != q {
		if p == nil {
			p = &Insight_Subject{}
		}
		if q == nil {
			q = &Insight_Subject{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Insight_Application) EqualVT(thatIface isInsight_SubjectOneof) bool {
	that, ok := thatIface.(*Insight_Application)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Application, that.Application; p != q {
		if p == nil {
			p = &Insight_Subject{}
		}
		if q == nil {
			q = &Insight_Subject{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Insight_Organization) EqualVT(thatIface isInsight_SubjectOneof) bool {
	that, ok := thatIface.(*Insight_Organization)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Organization, that.Organization; p != q {
		if p == nil {
			p = &Insight_Subject{}
		}
		if q == nil {
			q = &Insight_Subject{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}
