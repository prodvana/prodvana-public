// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/desired_state/model/entity.proto

package model

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

func (m *Notifications_ProgramFailures) CloneVT() *Notifications_ProgramFailures {
	if m == nil {
		return (*Notifications_ProgramFailures)(nil)
	}
	r := new(Notifications_ProgramFailures)
	r.FailureCount = m.FailureCount
	r.MostRecentFailure = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.MostRecentFailure).CloneVT())
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *Notifications_ProgramFailures) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *Notifications_RuntimeFailure) CloneVT() *Notifications_RuntimeFailure {
	if m == nil {
		return (*Notifications_RuntimeFailure)(nil)
	}
	r := new(Notifications_RuntimeFailure)
	r.FailureType = m.FailureType
	r.Message = m.Message
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *Notifications_RuntimeFailure) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *Notifications_ProtectionFailure) CloneVT() *Notifications_ProtectionFailure {
	if m == nil {
		return (*Notifications_ProtectionFailure)(nil)
	}
	r := new(Notifications_ProtectionFailure)
	r.ProtectionId = m.ProtectionId.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *Notifications_ProtectionFailure) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *Notifications_ConvergenceExtensionFailure) CloneVT() *Notifications_ConvergenceExtensionFailure {
	if m == nil {
		return (*Notifications_ConvergenceExtensionFailure)(nil)
	}
	r := new(Notifications_ConvergenceExtensionFailure)
	r.ExtensionId = m.ExtensionId.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *Notifications_ConvergenceExtensionFailure) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *Notifications_DelayedConvergence) CloneVT() *Notifications_DelayedConvergence {
	if m == nil {
		return (*Notifications_DelayedConvergence)(nil)
	}
	r := new(Notifications_DelayedConvergence)
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *Notifications_DelayedConvergence) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *Notifications) CloneVT() *Notifications {
	if m == nil {
		return (*Notifications)(nil)
	}
	r := new(Notifications)
	r.ProgramFailures = m.ProgramFailures.CloneVT()
	r.DelayedConvergence = m.DelayedConvergence.CloneVT()
	if rhs := m.RuntimeFailures; rhs != nil {
		tmpContainer := make([]*Notifications_RuntimeFailure, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.RuntimeFailures = tmpContainer
	}
	if rhs := m.ProtectionFailure; rhs != nil {
		tmpContainer := make([]*Notifications_ProtectionFailure, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.ProtectionFailure = tmpContainer
	}
	if rhs := m.ConvergenceExtensionFailure; rhs != nil {
		tmpContainer := make([]*Notifications_ConvergenceExtensionFailure, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.ConvergenceExtensionFailure = tmpContainer
	}
	if rhs := m.ConcurrencyLimitExceededErrors; rhs != nil {
		tmpContainer := make([]*ConcurrencyLimitExceeded, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.ConcurrencyLimitExceededErrors = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *Notifications) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *Dependency) CloneVT() *Dependency {
	if m == nil {
		return (*Dependency)(nil)
	}
	r := new(Dependency)
	r.Type = m.Type
	r.Id = m.Id.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *Dependency) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *Entity) CloneVT() *Entity {
	if m == nil {
		return (*Entity)(nil)
	}
	r := new(Entity)
	r.Id = m.Id.CloneVT()
	r.DesiredStateId = m.DesiredStateId
	r.RootDesiredStateId = m.RootDesiredStateId
	r.ParentDesiredStateId = m.ParentDesiredStateId
	r.ReleaseId = m.ReleaseId
	r.Annotations = m.Annotations.CloneVT()
	r.Status = m.Status
	r.SimpleStatus = m.SimpleStatus
	r.StartingState = m.StartingState.CloneVT()
	r.LastSeenState = m.LastSeenState.CloneVT()
	r.DesiredState = m.DesiredState.CloneVT()
	r.TargetState = m.TargetState.CloneVT()
	r.ActionExplanation = m.ActionExplanation.CloneVT()
	r.LastUpdateTimestamp = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.LastUpdateTimestamp).CloneVT())
	r.LastFetchedTimestamp = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.LastFetchedTimestamp).CloneVT())
	r.LastAppliedTimestamp = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.LastAppliedTimestamp).CloneVT())
	r.ExpectedNextApplyTimestamp = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.ExpectedNextApplyTimestamp).CloneVT())
	r.HasWork = m.HasWork
	r.HasWorkReason = m.HasWorkReason
	r.Lifecycle = m.Lifecycle
	r.MissingApproval = m.MissingApproval.CloneVT()
	r.ApplyError = m.ApplyError.CloneVT()
	r.Notifications = m.Notifications.CloneVT()
	r.LastKeyDeliveryDecision = m.LastKeyDeliveryDecision.CloneVT()
	r.LastRollbackKeyDeliveryDecision = m.LastRollbackKeyDeliveryDecision.CloneVT()
	if rhs := m.PreconditionStatuses; rhs != nil {
		tmpContainer := make([]*ConditionState, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.PreconditionStatuses = tmpContainer
	}
	if rhs := m.StatusExplanations; rhs != nil {
		tmpContainer := make([]*StatusExplanation, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.StatusExplanations = tmpContainer
	}
	if rhs := m.Logs; rhs != nil {
		tmpContainer := make([]*DebugLog, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Logs = tmpContainer
	}
	if rhs := m.Dependencies; rhs != nil {
		tmpContainer := make([]*Identifier, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Dependencies = tmpContainer
	}
	if rhs := m.DirectDependencies; rhs != nil {
		tmpContainer := make([]*Identifier, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.DirectDependencies = tmpContainer
	}
	if rhs := m.DependenciesWithType; rhs != nil {
		tmpContainer := make([]*Dependency, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.DependenciesWithType = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *Entity) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *EntityGraph) CloneVT() *EntityGraph {
	if m == nil {
		return (*EntityGraph)(nil)
	}
	r := new(EntityGraph)
	r.Root = m.Root.CloneVT()
	r.ReplacedTimestamp = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.ReplacedTimestamp).CloneVT())
	if rhs := m.Entities; rhs != nil {
		tmpContainer := make([]*Entity, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Entities = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *EntityGraph) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *Notifications_ProgramFailures) StableEqualVT(that *Notifications_ProgramFailures) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.FailureCount != that.FailureCount {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.MostRecentFailure).StableEqualVT((*timestamppb1.Timestamp)(that.MostRecentFailure)) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Notifications_ProgramFailures) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Notifications_ProgramFailures)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *Notifications_RuntimeFailure) StableEqualVT(that *Notifications_RuntimeFailure) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.FailureType != that.FailureType {
		return false
	}
	if this.Message != that.Message {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Notifications_RuntimeFailure) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Notifications_RuntimeFailure)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *Notifications_ProtectionFailure) StableEqualVT(that *Notifications_ProtectionFailure) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.ProtectionId.StableEqualVT(that.ProtectionId) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Notifications_ProtectionFailure) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Notifications_ProtectionFailure)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *Notifications_ConvergenceExtensionFailure) StableEqualVT(that *Notifications_ConvergenceExtensionFailure) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.ExtensionId.StableEqualVT(that.ExtensionId) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Notifications_ConvergenceExtensionFailure) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Notifications_ConvergenceExtensionFailure)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *Notifications_DelayedConvergence) StableEqualVT(that *Notifications_DelayedConvergence) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Notifications_DelayedConvergence) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Notifications_DelayedConvergence)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *Notifications) StableEqualVT(that *Notifications) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.ProgramFailures.StableEqualVT(that.ProgramFailures) {
		return false
	}
	if len(this.RuntimeFailures) != len(that.RuntimeFailures) {
		return false
	}
	for i, vx := range this.RuntimeFailures {
		vy := that.RuntimeFailures[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Notifications_RuntimeFailure{}
			}
			if q == nil {
				q = &Notifications_RuntimeFailure{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if len(this.ProtectionFailure) != len(that.ProtectionFailure) {
		return false
	}
	for i, vx := range this.ProtectionFailure {
		vy := that.ProtectionFailure[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Notifications_ProtectionFailure{}
			}
			if q == nil {
				q = &Notifications_ProtectionFailure{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if len(this.ConvergenceExtensionFailure) != len(that.ConvergenceExtensionFailure) {
		return false
	}
	for i, vx := range this.ConvergenceExtensionFailure {
		vy := that.ConvergenceExtensionFailure[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Notifications_ConvergenceExtensionFailure{}
			}
			if q == nil {
				q = &Notifications_ConvergenceExtensionFailure{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if !this.DelayedConvergence.StableEqualVT(that.DelayedConvergence) {
		return false
	}
	if len(this.ConcurrencyLimitExceededErrors) != len(that.ConcurrencyLimitExceededErrors) {
		return false
	}
	for i, vx := range this.ConcurrencyLimitExceededErrors {
		vy := that.ConcurrencyLimitExceededErrors[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &ConcurrencyLimitExceeded{}
			}
			if q == nil {
				q = &ConcurrencyLimitExceeded{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Notifications) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Notifications)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *Notifications_ProgramFailures) EqualVT(that *Notifications_ProgramFailures) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.FailureCount != that.FailureCount {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.MostRecentFailure).EqualVT((*timestamppb1.Timestamp)(that.MostRecentFailure)) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Notifications_ProgramFailures) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Notifications_ProgramFailures)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Notifications_RuntimeFailure) EqualVT(that *Notifications_RuntimeFailure) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.FailureType != that.FailureType {
		return false
	}
	if this.Message != that.Message {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Notifications_RuntimeFailure) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Notifications_RuntimeFailure)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Notifications_ProtectionFailure) EqualVT(that *Notifications_ProtectionFailure) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.ProtectionId.EqualVT(that.ProtectionId) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Notifications_ProtectionFailure) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Notifications_ProtectionFailure)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Notifications_ConvergenceExtensionFailure) EqualVT(that *Notifications_ConvergenceExtensionFailure) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.ExtensionId.EqualVT(that.ExtensionId) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Notifications_ConvergenceExtensionFailure) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Notifications_ConvergenceExtensionFailure)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Notifications_DelayedConvergence) EqualVT(that *Notifications_DelayedConvergence) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Notifications_DelayedConvergence) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Notifications_DelayedConvergence)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Notifications) EqualVT(that *Notifications) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.ProgramFailures.EqualVT(that.ProgramFailures) {
		return false
	}
	if len(this.RuntimeFailures) != len(that.RuntimeFailures) {
		return false
	}
	for i, vx := range this.RuntimeFailures {
		vy := that.RuntimeFailures[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Notifications_RuntimeFailure{}
			}
			if q == nil {
				q = &Notifications_RuntimeFailure{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.ProtectionFailure) != len(that.ProtectionFailure) {
		return false
	}
	for i, vx := range this.ProtectionFailure {
		vy := that.ProtectionFailure[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Notifications_ProtectionFailure{}
			}
			if q == nil {
				q = &Notifications_ProtectionFailure{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.ConvergenceExtensionFailure) != len(that.ConvergenceExtensionFailure) {
		return false
	}
	for i, vx := range this.ConvergenceExtensionFailure {
		vy := that.ConvergenceExtensionFailure[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Notifications_ConvergenceExtensionFailure{}
			}
			if q == nil {
				q = &Notifications_ConvergenceExtensionFailure{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if !this.DelayedConvergence.EqualVT(that.DelayedConvergence) {
		return false
	}
	if len(this.ConcurrencyLimitExceededErrors) != len(that.ConcurrencyLimitExceededErrors) {
		return false
	}
	for i, vx := range this.ConcurrencyLimitExceededErrors {
		vy := that.ConcurrencyLimitExceededErrors[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &ConcurrencyLimitExceeded{}
			}
			if q == nil {
				q = &ConcurrencyLimitExceeded{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Notifications) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Notifications)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Dependency) StableEqualVT(that *Dependency) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Type != that.Type {
		return false
	}
	if !this.Id.StableEqualVT(that.Id) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Dependency) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Dependency)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *Dependency) EqualVT(that *Dependency) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Type != that.Type {
		return false
	}
	if !this.Id.EqualVT(that.Id) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Dependency) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Dependency)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Entity) StableEqualVT(that *Entity) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Id.StableEqualVT(that.Id) {
		return false
	}
	if this.DesiredStateId != that.DesiredStateId {
		return false
	}
	if this.RootDesiredStateId != that.RootDesiredStateId {
		return false
	}
	if !this.Annotations.StableEqualVT(that.Annotations) {
		return false
	}
	if this.Status != that.Status {
		return false
	}
	if !this.StartingState.StableEqualVT(that.StartingState) {
		return false
	}
	if !this.LastSeenState.StableEqualVT(that.LastSeenState) {
		return false
	}
	if !this.DesiredState.StableEqualVT(that.DesiredState) {
		return false
	}
	if len(this.PreconditionStatuses) != len(that.PreconditionStatuses) {
		return false
	}
	for i, vx := range this.PreconditionStatuses {
		vy := that.PreconditionStatuses[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &ConditionState{}
			}
			if q == nil {
				q = &ConditionState{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if len(this.StatusExplanations) != len(that.StatusExplanations) {
		return false
	}
	for i, vx := range this.StatusExplanations {
		vy := that.StatusExplanations[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &StatusExplanation{}
			}
			if q == nil {
				q = &StatusExplanation{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if len(this.Logs) != len(that.Logs) {
		return false
	}
	for i, vx := range this.Logs {
		vy := that.Logs[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &DebugLog{}
			}
			if q == nil {
				q = &DebugLog{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if !this.ActionExplanation.StableEqualVT(that.ActionExplanation) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.LastUpdateTimestamp).StableEqualVT((*timestamppb1.Timestamp)(that.LastUpdateTimestamp)) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.LastFetchedTimestamp).StableEqualVT((*timestamppb1.Timestamp)(that.LastFetchedTimestamp)) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.LastAppliedTimestamp).StableEqualVT((*timestamppb1.Timestamp)(that.LastAppliedTimestamp)) {
		return false
	}
	if len(this.Dependencies) != len(that.Dependencies) {
		return false
	}
	for i, vx := range this.Dependencies {
		vy := that.Dependencies[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Identifier{}
			}
			if q == nil {
				q = &Identifier{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if this.SimpleStatus != that.SimpleStatus {
		return false
	}
	if this.Lifecycle != that.Lifecycle {
		return false
	}
	if !this.TargetState.StableEqualVT(that.TargetState) {
		return false
	}
	if !this.MissingApproval.StableEqualVT(that.MissingApproval) {
		return false
	}
	if !this.ApplyError.StableEqualVT(that.ApplyError) {
		return false
	}
	if !this.Notifications.StableEqualVT(that.Notifications) {
		return false
	}
	if len(this.DirectDependencies) != len(that.DirectDependencies) {
		return false
	}
	for i, vx := range this.DirectDependencies {
		vy := that.DirectDependencies[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Identifier{}
			}
			if q == nil {
				q = &Identifier{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if this.ReleaseId != that.ReleaseId {
		return false
	}
	if this.ParentDesiredStateId != that.ParentDesiredStateId {
		return false
	}
	if len(this.DependenciesWithType) != len(that.DependenciesWithType) {
		return false
	}
	for i, vx := range this.DependenciesWithType {
		vy := that.DependenciesWithType[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Dependency{}
			}
			if q == nil {
				q = &Dependency{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if !(*timestamppb1.Timestamp)(this.ExpectedNextApplyTimestamp).StableEqualVT((*timestamppb1.Timestamp)(that.ExpectedNextApplyTimestamp)) {
		return false
	}
	if this.HasWork != that.HasWork {
		return false
	}
	if this.HasWorkReason != that.HasWorkReason {
		return false
	}
	if !this.LastKeyDeliveryDecision.StableEqualVT(that.LastKeyDeliveryDecision) {
		return false
	}
	if !this.LastRollbackKeyDeliveryDecision.StableEqualVT(that.LastRollbackKeyDeliveryDecision) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Entity) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Entity)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *Entity) EqualVT(that *Entity) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Id.EqualVT(that.Id) {
		return false
	}
	if this.DesiredStateId != that.DesiredStateId {
		return false
	}
	if this.RootDesiredStateId != that.RootDesiredStateId {
		return false
	}
	if !this.Annotations.EqualVT(that.Annotations) {
		return false
	}
	if this.Status != that.Status {
		return false
	}
	if !this.StartingState.EqualVT(that.StartingState) {
		return false
	}
	if !this.LastSeenState.EqualVT(that.LastSeenState) {
		return false
	}
	if !this.DesiredState.EqualVT(that.DesiredState) {
		return false
	}
	if len(this.PreconditionStatuses) != len(that.PreconditionStatuses) {
		return false
	}
	for i, vx := range this.PreconditionStatuses {
		vy := that.PreconditionStatuses[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &ConditionState{}
			}
			if q == nil {
				q = &ConditionState{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.StatusExplanations) != len(that.StatusExplanations) {
		return false
	}
	for i, vx := range this.StatusExplanations {
		vy := that.StatusExplanations[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &StatusExplanation{}
			}
			if q == nil {
				q = &StatusExplanation{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.Logs) != len(that.Logs) {
		return false
	}
	for i, vx := range this.Logs {
		vy := that.Logs[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &DebugLog{}
			}
			if q == nil {
				q = &DebugLog{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if !this.ActionExplanation.EqualVT(that.ActionExplanation) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.LastUpdateTimestamp).EqualVT((*timestamppb1.Timestamp)(that.LastUpdateTimestamp)) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.LastFetchedTimestamp).EqualVT((*timestamppb1.Timestamp)(that.LastFetchedTimestamp)) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.LastAppliedTimestamp).EqualVT((*timestamppb1.Timestamp)(that.LastAppliedTimestamp)) {
		return false
	}
	if len(this.Dependencies) != len(that.Dependencies) {
		return false
	}
	for i, vx := range this.Dependencies {
		vy := that.Dependencies[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Identifier{}
			}
			if q == nil {
				q = &Identifier{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if this.SimpleStatus != that.SimpleStatus {
		return false
	}
	if this.Lifecycle != that.Lifecycle {
		return false
	}
	if !this.TargetState.EqualVT(that.TargetState) {
		return false
	}
	if !this.MissingApproval.EqualVT(that.MissingApproval) {
		return false
	}
	if !this.ApplyError.EqualVT(that.ApplyError) {
		return false
	}
	if !this.Notifications.EqualVT(that.Notifications) {
		return false
	}
	if len(this.DirectDependencies) != len(that.DirectDependencies) {
		return false
	}
	for i, vx := range this.DirectDependencies {
		vy := that.DirectDependencies[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Identifier{}
			}
			if q == nil {
				q = &Identifier{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if this.ReleaseId != that.ReleaseId {
		return false
	}
	if this.ParentDesiredStateId != that.ParentDesiredStateId {
		return false
	}
	if len(this.DependenciesWithType) != len(that.DependenciesWithType) {
		return false
	}
	for i, vx := range this.DependenciesWithType {
		vy := that.DependenciesWithType[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Dependency{}
			}
			if q == nil {
				q = &Dependency{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if !(*timestamppb1.Timestamp)(this.ExpectedNextApplyTimestamp).EqualVT((*timestamppb1.Timestamp)(that.ExpectedNextApplyTimestamp)) {
		return false
	}
	if this.HasWork != that.HasWork {
		return false
	}
	if this.HasWorkReason != that.HasWorkReason {
		return false
	}
	if !this.LastKeyDeliveryDecision.EqualVT(that.LastKeyDeliveryDecision) {
		return false
	}
	if !this.LastRollbackKeyDeliveryDecision.EqualVT(that.LastRollbackKeyDeliveryDecision) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Entity) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Entity)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *EntityGraph) StableEqualVT(that *EntityGraph) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Root.StableEqualVT(that.Root) {
		return false
	}
	if len(this.Entities) != len(that.Entities) {
		return false
	}
	for i, vx := range this.Entities {
		vy := that.Entities[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Entity{}
			}
			if q == nil {
				q = &Entity{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if !(*timestamppb1.Timestamp)(this.ReplacedTimestamp).StableEqualVT((*timestamppb1.Timestamp)(that.ReplacedTimestamp)) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *EntityGraph) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*EntityGraph)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *EntityGraph) EqualVT(that *EntityGraph) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Root.EqualVT(that.Root) {
		return false
	}
	if len(this.Entities) != len(that.Entities) {
		return false
	}
	for i, vx := range this.Entities {
		vy := that.Entities[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Entity{}
			}
			if q == nil {
				q = &Entity{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if !(*timestamppb1.Timestamp)(this.ReplacedTimestamp).EqualVT((*timestamppb1.Timestamp)(that.ReplacedTimestamp)) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *EntityGraph) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*EntityGraph)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}