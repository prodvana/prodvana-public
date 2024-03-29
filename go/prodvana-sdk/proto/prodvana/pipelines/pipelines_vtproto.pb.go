// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/pipelines/pipelines.proto

package pipelines

import (
	durationpb1 "github.com/planetscale/vtprotobuf/types/known/durationpb"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *PushTask) CloneVT() *PushTask {
	if m == nil {
		return (*PushTask)(nil)
	}
	r := new(PushTask)
	r.ServiceId = m.ServiceId
	r.Service = m.Service
	r.ReleaseChannelId = m.ReleaseChannelId
	r.ReleaseChannel = m.ReleaseChannel
	r.Rollback = m.Rollback
	r.ApplicationId = m.ApplicationId
	r.Application = m.Application
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *PushTask) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *WaitTask) CloneVT() *WaitTask {
	if m == nil {
		return (*WaitTask)(nil)
	}
	r := new(WaitTask)
	r.Duration = (*durationpb.Duration)((*durationpb1.Duration)(m.Duration).CloneVT())
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *WaitTask) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ParallelTask_TaskTrack) CloneVT() *ParallelTask_TaskTrack {
	if m == nil {
		return (*ParallelTask_TaskTrack)(nil)
	}
	r := new(ParallelTask_TaskTrack)
	if rhs := m.Tasks; rhs != nil {
		tmpContainer := make([]*Task, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Tasks = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ParallelTask_TaskTrack) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ParallelTask) CloneVT() *ParallelTask {
	if m == nil {
		return (*ParallelTask)(nil)
	}
	r := new(ParallelTask)
	if rhs := m.Tracks; rhs != nil {
		tmpContainer := make([]*ParallelTask_TaskTrack, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Tracks = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ParallelTask) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ManualApprovalTask) CloneVT() *ManualApprovalTask {
	if m == nil {
		return (*ManualApprovalTask)(nil)
	}
	r := new(ManualApprovalTask)
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ManualApprovalTask) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *CustomTask) CloneVT() *CustomTask {
	if m == nil {
		return (*CustomTask)(nil)
	}
	r := new(CustomTask)
	r.Description = m.Description
	r.Program = m.Program.CloneVT()
	r.RetryConfig = m.RetryConfig.CloneVT()
	r.Application = m.Application
	r.ApplicationId = m.ApplicationId
	r.ReleaseChannel = m.ReleaseChannel
	r.ReleaseChannelId = m.ReleaseChannelId
	r.FinalCompiledProgram = m.FinalCompiledProgram
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *CustomTask) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *Task_Metadata) CloneVT() *Task_Metadata {
	if m == nil {
		return (*Task_Metadata)(nil)
	}
	r := new(Task_Metadata)
	r.Id = m.Id
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *Task_Metadata) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *Task) CloneVT() *Task {
	if m == nil {
		return (*Task)(nil)
	}
	r := new(Task)
	r.Metadata = m.Metadata.CloneVT()
	if m.TaskOneof != nil {
		r.TaskOneof = m.TaskOneof.(interface{ CloneVT() isTask_TaskOneof }).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *Task) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *Task_PushTask) CloneVT() isTask_TaskOneof {
	if m == nil {
		return (*Task_PushTask)(nil)
	}
	r := new(Task_PushTask)
	r.PushTask = m.PushTask.CloneVT()
	return r
}

func (m *Task_WaitTask) CloneVT() isTask_TaskOneof {
	if m == nil {
		return (*Task_WaitTask)(nil)
	}
	r := new(Task_WaitTask)
	r.WaitTask = m.WaitTask.CloneVT()
	return r
}

func (m *Task_ParallelTask) CloneVT() isTask_TaskOneof {
	if m == nil {
		return (*Task_ParallelTask)(nil)
	}
	r := new(Task_ParallelTask)
	r.ParallelTask = m.ParallelTask.CloneVT()
	return r
}

func (m *Task_ManualApprovalTask) CloneVT() isTask_TaskOneof {
	if m == nil {
		return (*Task_ManualApprovalTask)(nil)
	}
	r := new(Task_ManualApprovalTask)
	r.ManualApprovalTask = m.ManualApprovalTask.CloneVT()
	return r
}

func (m *Task_CustomTask) CloneVT() isTask_TaskOneof {
	if m == nil {
		return (*Task_CustomTask)(nil)
	}
	r := new(Task_CustomTask)
	r.CustomTask = m.CustomTask.CloneVT()
	return r
}

func (m *PipelineConfig) CloneVT() *PipelineConfig {
	if m == nil {
		return (*PipelineConfig)(nil)
	}
	r := new(PipelineConfig)
	r.Name = m.Name
	r.Rollback = m.Rollback
	r.Notifications = m.Notifications.CloneVT()
	if rhs := m.Tasks; rhs != nil {
		tmpContainer := make([]*Task, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Tasks = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *PipelineConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *PipelineTemplate) CloneVT() *PipelineTemplate {
	if m == nil {
		return (*PipelineTemplate)(nil)
	}
	r := new(PipelineTemplate)
	r.NameSuffix = m.NameSuffix
	r.Rollback = m.Rollback
	r.Notifications = m.Notifications.CloneVT()
	if rhs := m.Tasks; rhs != nil {
		tmpContainer := make([]*Task, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Tasks = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *PipelineTemplate) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *PushTask) StableEqualVT(that *PushTask) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ServiceId != that.ServiceId {
		return false
	}
	if this.Service != that.Service {
		return false
	}
	if this.ReleaseChannelId != that.ReleaseChannelId {
		return false
	}
	if this.ReleaseChannel != that.ReleaseChannel {
		return false
	}
	if this.Rollback != that.Rollback {
		return false
	}
	if this.ApplicationId != that.ApplicationId {
		return false
	}
	if this.Application != that.Application {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *PushTask) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*PushTask)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *PushTask) EqualVT(that *PushTask) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ServiceId != that.ServiceId {
		return false
	}
	if this.Service != that.Service {
		return false
	}
	if this.ReleaseChannelId != that.ReleaseChannelId {
		return false
	}
	if this.ReleaseChannel != that.ReleaseChannel {
		return false
	}
	if this.Rollback != that.Rollback {
		return false
	}
	if this.ApplicationId != that.ApplicationId {
		return false
	}
	if this.Application != that.Application {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *PushTask) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*PushTask)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *WaitTask) StableEqualVT(that *WaitTask) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !(*durationpb1.Duration)(this.Duration).StableEqualVT((*durationpb1.Duration)(that.Duration)) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *WaitTask) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*WaitTask)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *WaitTask) EqualVT(that *WaitTask) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !(*durationpb1.Duration)(this.Duration).EqualVT((*durationpb1.Duration)(that.Duration)) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *WaitTask) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*WaitTask)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ParallelTask_TaskTrack) StableEqualVT(that *ParallelTask_TaskTrack) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Tasks) != len(that.Tasks) {
		return false
	}
	for i, vx := range this.Tasks {
		vy := that.Tasks[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Task{}
			}
			if q == nil {
				q = &Task{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ParallelTask_TaskTrack) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ParallelTask_TaskTrack)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ParallelTask) StableEqualVT(that *ParallelTask) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Tracks) != len(that.Tracks) {
		return false
	}
	for i, vx := range this.Tracks {
		vy := that.Tracks[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &ParallelTask_TaskTrack{}
			}
			if q == nil {
				q = &ParallelTask_TaskTrack{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ParallelTask) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ParallelTask)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ParallelTask_TaskTrack) EqualVT(that *ParallelTask_TaskTrack) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Tasks) != len(that.Tasks) {
		return false
	}
	for i, vx := range this.Tasks {
		vy := that.Tasks[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Task{}
			}
			if q == nil {
				q = &Task{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ParallelTask_TaskTrack) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ParallelTask_TaskTrack)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ParallelTask) EqualVT(that *ParallelTask) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Tracks) != len(that.Tracks) {
		return false
	}
	for i, vx := range this.Tracks {
		vy := that.Tracks[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &ParallelTask_TaskTrack{}
			}
			if q == nil {
				q = &ParallelTask_TaskTrack{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ParallelTask) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ParallelTask)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ManualApprovalTask) StableEqualVT(that *ManualApprovalTask) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ManualApprovalTask) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ManualApprovalTask)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ManualApprovalTask) EqualVT(that *ManualApprovalTask) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ManualApprovalTask) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ManualApprovalTask)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *CustomTask) StableEqualVT(that *CustomTask) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Program.StableEqualVT(that.Program) {
		return false
	}
	if !this.RetryConfig.StableEqualVT(that.RetryConfig) {
		return false
	}
	if this.Application != that.Application {
		return false
	}
	if this.ApplicationId != that.ApplicationId {
		return false
	}
	if this.ReleaseChannel != that.ReleaseChannel {
		return false
	}
	if this.ReleaseChannelId != that.ReleaseChannelId {
		return false
	}
	if this.Description != that.Description {
		return false
	}
	if this.FinalCompiledProgram != that.FinalCompiledProgram {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *CustomTask) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*CustomTask)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *CustomTask) EqualVT(that *CustomTask) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Program.EqualVT(that.Program) {
		return false
	}
	if !this.RetryConfig.EqualVT(that.RetryConfig) {
		return false
	}
	if this.Application != that.Application {
		return false
	}
	if this.ApplicationId != that.ApplicationId {
		return false
	}
	if this.ReleaseChannel != that.ReleaseChannel {
		return false
	}
	if this.ReleaseChannelId != that.ReleaseChannelId {
		return false
	}
	if this.Description != that.Description {
		return false
	}
	if this.FinalCompiledProgram != that.FinalCompiledProgram {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *CustomTask) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*CustomTask)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Task_Metadata) StableEqualVT(that *Task_Metadata) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Id != that.Id {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Task_Metadata) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Task_Metadata)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *Task) StableEqualVT(that *Task) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.TaskOneof == nil && that.TaskOneof != nil {
		return false
	} else if this.TaskOneof != nil {
		if that.TaskOneof == nil {
			return false
		}
		if !this.TaskOneof.(interface{ StableEqualVT(isTask_TaskOneof) bool }).StableEqualVT(that.TaskOneof) {
			return false
		}
	}
	if !this.Metadata.StableEqualVT(that.Metadata) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Task) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Task)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *Task_PushTask) StableEqualVT(thatIface isTask_TaskOneof) bool {
	that, ok := thatIface.(*Task_PushTask)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.PushTask, that.PushTask; p != q {
		if p == nil {
			p = &PushTask{}
		}
		if q == nil {
			q = &PushTask{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Task_WaitTask) StableEqualVT(thatIface isTask_TaskOneof) bool {
	that, ok := thatIface.(*Task_WaitTask)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.WaitTask, that.WaitTask; p != q {
		if p == nil {
			p = &WaitTask{}
		}
		if q == nil {
			q = &WaitTask{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Task_ParallelTask) StableEqualVT(thatIface isTask_TaskOneof) bool {
	that, ok := thatIface.(*Task_ParallelTask)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.ParallelTask, that.ParallelTask; p != q {
		if p == nil {
			p = &ParallelTask{}
		}
		if q == nil {
			q = &ParallelTask{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Task_ManualApprovalTask) StableEqualVT(thatIface isTask_TaskOneof) bool {
	that, ok := thatIface.(*Task_ManualApprovalTask)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.ManualApprovalTask, that.ManualApprovalTask; p != q {
		if p == nil {
			p = &ManualApprovalTask{}
		}
		if q == nil {
			q = &ManualApprovalTask{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Task_CustomTask) StableEqualVT(thatIface isTask_TaskOneof) bool {
	that, ok := thatIface.(*Task_CustomTask)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.CustomTask, that.CustomTask; p != q {
		if p == nil {
			p = &CustomTask{}
		}
		if q == nil {
			q = &CustomTask{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Task_Metadata) EqualVT(that *Task_Metadata) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Id != that.Id {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Task_Metadata) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Task_Metadata)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Task) EqualVT(that *Task) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.TaskOneof == nil && that.TaskOneof != nil {
		return false
	} else if this.TaskOneof != nil {
		if that.TaskOneof == nil {
			return false
		}
		if !this.TaskOneof.(interface{ EqualVT(isTask_TaskOneof) bool }).EqualVT(that.TaskOneof) {
			return false
		}
	}
	if !this.Metadata.EqualVT(that.Metadata) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Task) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Task)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Task_PushTask) EqualVT(thatIface isTask_TaskOneof) bool {
	that, ok := thatIface.(*Task_PushTask)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.PushTask, that.PushTask; p != q {
		if p == nil {
			p = &PushTask{}
		}
		if q == nil {
			q = &PushTask{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Task_WaitTask) EqualVT(thatIface isTask_TaskOneof) bool {
	that, ok := thatIface.(*Task_WaitTask)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.WaitTask, that.WaitTask; p != q {
		if p == nil {
			p = &WaitTask{}
		}
		if q == nil {
			q = &WaitTask{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Task_ParallelTask) EqualVT(thatIface isTask_TaskOneof) bool {
	that, ok := thatIface.(*Task_ParallelTask)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.ParallelTask, that.ParallelTask; p != q {
		if p == nil {
			p = &ParallelTask{}
		}
		if q == nil {
			q = &ParallelTask{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Task_ManualApprovalTask) EqualVT(thatIface isTask_TaskOneof) bool {
	that, ok := thatIface.(*Task_ManualApprovalTask)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.ManualApprovalTask, that.ManualApprovalTask; p != q {
		if p == nil {
			p = &ManualApprovalTask{}
		}
		if q == nil {
			q = &ManualApprovalTask{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Task_CustomTask) EqualVT(thatIface isTask_TaskOneof) bool {
	that, ok := thatIface.(*Task_CustomTask)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.CustomTask, that.CustomTask; p != q {
		if p == nil {
			p = &CustomTask{}
		}
		if q == nil {
			q = &CustomTask{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *PipelineConfig) StableEqualVT(that *PipelineConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	if len(this.Tasks) != len(that.Tasks) {
		return false
	}
	for i, vx := range this.Tasks {
		vy := that.Tasks[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Task{}
			}
			if q == nil {
				q = &Task{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if this.Rollback != that.Rollback {
		return false
	}
	if !this.Notifications.StableEqualVT(that.Notifications) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *PipelineConfig) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*PipelineConfig)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *PipelineConfig) EqualVT(that *PipelineConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	if len(this.Tasks) != len(that.Tasks) {
		return false
	}
	for i, vx := range this.Tasks {
		vy := that.Tasks[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Task{}
			}
			if q == nil {
				q = &Task{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if this.Rollback != that.Rollback {
		return false
	}
	if !this.Notifications.EqualVT(that.Notifications) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *PipelineConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*PipelineConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *PipelineTemplate) StableEqualVT(that *PipelineTemplate) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.NameSuffix != that.NameSuffix {
		return false
	}
	if len(this.Tasks) != len(that.Tasks) {
		return false
	}
	for i, vx := range this.Tasks {
		vy := that.Tasks[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Task{}
			}
			if q == nil {
				q = &Task{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if this.Rollback != that.Rollback {
		return false
	}
	if !this.Notifications.StableEqualVT(that.Notifications) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *PipelineTemplate) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*PipelineTemplate)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *PipelineTemplate) EqualVT(that *PipelineTemplate) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.NameSuffix != that.NameSuffix {
		return false
	}
	if len(this.Tasks) != len(that.Tasks) {
		return false
	}
	for i, vx := range this.Tasks {
		vy := that.Tasks[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Task{}
			}
			if q == nil {
				q = &Task{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if this.Rollback != that.Rollback {
		return false
	}
	if !this.Notifications.EqualVT(that.Notifications) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *PipelineTemplate) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*PipelineTemplate)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
