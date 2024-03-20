// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/protection/protection_config.proto

package protection

import (
	durationpb1 "github.com/planetscale/vtprotobuf/types/known/durationpb"
	common_config "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
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

func (m *BuiltinProtectionConfig) CloneVT() *BuiltinProtectionConfig {
	if m == nil {
		return (*BuiltinProtectionConfig)(nil)
	}
	r := new(BuiltinProtectionConfig)
	if m.BuiltinOneof != nil {
		r.BuiltinOneof = m.BuiltinOneof.(interface {
			CloneVT() isBuiltinProtectionConfig_BuiltinOneof
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *BuiltinProtectionConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *BuiltinProtectionConfig_CommitDenylist) CloneVT() isBuiltinProtectionConfig_BuiltinOneof {
	if m == nil {
		return (*BuiltinProtectionConfig_CommitDenylist)(nil)
	}
	r := new(BuiltinProtectionConfig_CommitDenylist)
	r.CommitDenylist = m.CommitDenylist.CloneVT()
	return r
}

func (m *BuiltinProtectionConfig_AllowedTimes) CloneVT() isBuiltinProtectionConfig_BuiltinOneof {
	if m == nil {
		return (*BuiltinProtectionConfig_AllowedTimes)(nil)
	}
	r := new(BuiltinProtectionConfig_AllowedTimes)
	r.AllowedTimes = m.AllowedTimes.CloneVT()
	return r
}

func (m *ProtectionConfig) CloneVT() *ProtectionConfig {
	if m == nil {
		return (*ProtectionConfig)(nil)
	}
	r := new(ProtectionConfig)
	r.Name = m.Name
	r.PollInterval = (*durationpb.Duration)((*durationpb1.Duration)(m.PollInterval).CloneVT())
	r.Timeout = (*durationpb.Duration)((*durationpb1.Duration)(m.Timeout).CloneVT())
	if m.ExecConfig != nil {
		r.ExecConfig = m.ExecConfig.(interface {
			CloneVT() isProtectionConfig_ExecConfig
		}).CloneVT()
	}
	if rhs := m.Parameters; rhs != nil {
		tmpContainer := make([]*common_config.ParameterDefinition, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Parameters = tmpContainer
	}
	if rhs := m.Env; rhs != nil {
		tmpContainer := make(map[string]*common_config.EnvValue, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Env = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ProtectionConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ProtectionConfig_TaskConfig) CloneVT() isProtectionConfig_ExecConfig {
	if m == nil {
		return (*ProtectionConfig_TaskConfig)(nil)
	}
	r := new(ProtectionConfig_TaskConfig)
	r.TaskConfig = m.TaskConfig.CloneVT()
	return r
}

func (m *ProtectionConfig_KubernetesConfig) CloneVT() isProtectionConfig_ExecConfig {
	if m == nil {
		return (*ProtectionConfig_KubernetesConfig)(nil)
	}
	r := new(ProtectionConfig_KubernetesConfig)
	r.KubernetesConfig = m.KubernetesConfig.CloneVT()
	return r
}

func (m *ProtectionConfig_Builtin) CloneVT() isProtectionConfig_ExecConfig {
	if m == nil {
		return (*ProtectionConfig_Builtin)(nil)
	}
	r := new(ProtectionConfig_Builtin)
	r.Builtin = m.Builtin.CloneVT()
	return r
}

func (m *CompiledProtectionAttachmentConfig) CloneVT() *CompiledProtectionAttachmentConfig {
	if m == nil {
		return (*CompiledProtectionAttachmentConfig)(nil)
	}
	r := new(CompiledProtectionAttachmentConfig)
	r.Config = m.Config.CloneVT()
	r.Attachment = m.Attachment.CloneVT()
	r.RuntimeExecution = m.RuntimeExecution.CloneVT()
	if rhs := m.Env; rhs != nil {
		tmpContainer := make(map[string]*common_config.EnvValue, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Env = tmpContainer
	}
	if rhs := m.ParameterValues; rhs != nil {
		tmpContainer := make([]*common_config.ParameterValue, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.ParameterValues = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *CompiledProtectionAttachmentConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ServiceInstanceAttachment) CloneVT() *ServiceInstanceAttachment {
	if m == nil {
		return (*ServiceInstanceAttachment)(nil)
	}
	r := new(ServiceInstanceAttachment)
	r.Service = m.Service
	r.ReleaseChannel = m.ReleaseChannel
	r.Application = m.Application
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ServiceInstanceAttachment) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ReleaseChannelAttachment) CloneVT() *ReleaseChannelAttachment {
	if m == nil {
		return (*ReleaseChannelAttachment)(nil)
	}
	r := new(ReleaseChannelAttachment)
	r.Application = m.Application
	r.ReleaseChannel = m.ReleaseChannel
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ReleaseChannelAttachment) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ConvergenceAttachment) CloneVT() *ConvergenceAttachment {
	if m == nil {
		return (*ConvergenceAttachment)(nil)
	}
	r := new(ConvergenceAttachment)
	r.DesiredStateId = m.DesiredStateId
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ConvergenceAttachment) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ProtectionAttachment) CloneVT() *ProtectionAttachment {
	if m == nil {
		return (*ProtectionAttachment)(nil)
	}
	r := new(ProtectionAttachment)
	if m.Attachment != nil {
		r.Attachment = m.Attachment.(interface {
			CloneVT() isProtectionAttachment_Attachment
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ProtectionAttachment) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ProtectionAttachment_ServiceInstance) CloneVT() isProtectionAttachment_Attachment {
	if m == nil {
		return (*ProtectionAttachment_ServiceInstance)(nil)
	}
	r := new(ProtectionAttachment_ServiceInstance)
	r.ServiceInstance = m.ServiceInstance.CloneVT()
	return r
}

func (m *ProtectionAttachment_ReleaseChannel) CloneVT() isProtectionAttachment_Attachment {
	if m == nil {
		return (*ProtectionAttachment_ReleaseChannel)(nil)
	}
	r := new(ProtectionAttachment_ReleaseChannel)
	r.ReleaseChannel = m.ReleaseChannel.CloneVT()
	return r
}

func (m *ProtectionAttachment_Convergence) CloneVT() isProtectionAttachment_Attachment {
	if m == nil {
		return (*ProtectionAttachment_Convergence)(nil)
	}
	r := new(ProtectionAttachment_Convergence)
	r.Convergence = m.Convergence.CloneVT()
	return r
}

func (this *BuiltinProtectionConfig) EqualVT(that *BuiltinProtectionConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.BuiltinOneof == nil && that.BuiltinOneof != nil {
		return false
	} else if this.BuiltinOneof != nil {
		if that.BuiltinOneof == nil {
			return false
		}
		if !this.BuiltinOneof.(interface {
			EqualVT(isBuiltinProtectionConfig_BuiltinOneof) bool
		}).EqualVT(that.BuiltinOneof) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *BuiltinProtectionConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*BuiltinProtectionConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *BuiltinProtectionConfig_CommitDenylist) EqualVT(thatIface isBuiltinProtectionConfig_BuiltinOneof) bool {
	that, ok := thatIface.(*BuiltinProtectionConfig_CommitDenylist)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.CommitDenylist, that.CommitDenylist; p != q {
		if p == nil {
			p = &CommitDenylistProtectionConfig{}
		}
		if q == nil {
			q = &CommitDenylistProtectionConfig{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *BuiltinProtectionConfig_AllowedTimes) EqualVT(thatIface isBuiltinProtectionConfig_BuiltinOneof) bool {
	that, ok := thatIface.(*BuiltinProtectionConfig_AllowedTimes)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.AllowedTimes, that.AllowedTimes; p != q {
		if p == nil {
			p = &AllowedTimesProtectionConfig{}
		}
		if q == nil {
			q = &AllowedTimesProtectionConfig{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProtectionConfig) EqualVT(that *ProtectionConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ExecConfig == nil && that.ExecConfig != nil {
		return false
	} else if this.ExecConfig != nil {
		if that.ExecConfig == nil {
			return false
		}
		if !this.ExecConfig.(interface {
			EqualVT(isProtectionConfig_ExecConfig) bool
		}).EqualVT(that.ExecConfig) {
			return false
		}
	}
	if this.Name != that.Name {
		return false
	}
	if !(*durationpb1.Duration)(this.PollInterval).EqualVT((*durationpb1.Duration)(that.PollInterval)) {
		return false
	}
	if !(*durationpb1.Duration)(this.Timeout).EqualVT((*durationpb1.Duration)(that.Timeout)) {
		return false
	}
	if len(this.Parameters) != len(that.Parameters) {
		return false
	}
	for i, vx := range this.Parameters {
		vy := that.Parameters[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &common_config.ParameterDefinition{}
			}
			if q == nil {
				q = &common_config.ParameterDefinition{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.Env) != len(that.Env) {
		return false
	}
	for i, vx := range this.Env {
		vy, ok := that.Env[i]
		if !ok {
			return false
		}
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &common_config.EnvValue{}
			}
			if q == nil {
				q = &common_config.EnvValue{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ProtectionConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ProtectionConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ProtectionConfig_TaskConfig) EqualVT(thatIface isProtectionConfig_ExecConfig) bool {
	that, ok := thatIface.(*ProtectionConfig_TaskConfig)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.TaskConfig, that.TaskConfig; p != q {
		if p == nil {
			p = &common_config.TaskConfig{}
		}
		if q == nil {
			q = &common_config.TaskConfig{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProtectionConfig_KubernetesConfig) EqualVT(thatIface isProtectionConfig_ExecConfig) bool {
	that, ok := thatIface.(*ProtectionConfig_KubernetesConfig)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.KubernetesConfig, that.KubernetesConfig; p != q {
		if p == nil {
			p = &common_config.KubernetesConfig{}
		}
		if q == nil {
			q = &common_config.KubernetesConfig{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProtectionConfig_Builtin) EqualVT(thatIface isProtectionConfig_ExecConfig) bool {
	that, ok := thatIface.(*ProtectionConfig_Builtin)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Builtin, that.Builtin; p != q {
		if p == nil {
			p = &BuiltinProtectionConfig{}
		}
		if q == nil {
			q = &BuiltinProtectionConfig{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *CompiledProtectionAttachmentConfig) EqualVT(that *CompiledProtectionAttachmentConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Config.EqualVT(that.Config) {
		return false
	}
	if !this.Attachment.EqualVT(that.Attachment) {
		return false
	}
	if !this.RuntimeExecution.EqualVT(that.RuntimeExecution) {
		return false
	}
	if len(this.Env) != len(that.Env) {
		return false
	}
	for i, vx := range this.Env {
		vy, ok := that.Env[i]
		if !ok {
			return false
		}
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &common_config.EnvValue{}
			}
			if q == nil {
				q = &common_config.EnvValue{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.ParameterValues) != len(that.ParameterValues) {
		return false
	}
	for i, vx := range this.ParameterValues {
		vy := that.ParameterValues[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &common_config.ParameterValue{}
			}
			if q == nil {
				q = &common_config.ParameterValue{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *CompiledProtectionAttachmentConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*CompiledProtectionAttachmentConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ServiceInstanceAttachment) EqualVT(that *ServiceInstanceAttachment) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Service != that.Service {
		return false
	}
	if this.ReleaseChannel != that.ReleaseChannel {
		return false
	}
	if this.Application != that.Application {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ServiceInstanceAttachment) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ServiceInstanceAttachment)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ReleaseChannelAttachment) EqualVT(that *ReleaseChannelAttachment) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Application != that.Application {
		return false
	}
	if this.ReleaseChannel != that.ReleaseChannel {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ReleaseChannelAttachment) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ReleaseChannelAttachment)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ConvergenceAttachment) EqualVT(that *ConvergenceAttachment) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.DesiredStateId != that.DesiredStateId {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ConvergenceAttachment) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ConvergenceAttachment)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ProtectionAttachment) EqualVT(that *ProtectionAttachment) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Attachment == nil && that.Attachment != nil {
		return false
	} else if this.Attachment != nil {
		if that.Attachment == nil {
			return false
		}
		if !this.Attachment.(interface {
			EqualVT(isProtectionAttachment_Attachment) bool
		}).EqualVT(that.Attachment) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ProtectionAttachment) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ProtectionAttachment)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ProtectionAttachment_ServiceInstance) EqualVT(thatIface isProtectionAttachment_Attachment) bool {
	that, ok := thatIface.(*ProtectionAttachment_ServiceInstance)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.ServiceInstance, that.ServiceInstance; p != q {
		if p == nil {
			p = &ServiceInstanceAttachment{}
		}
		if q == nil {
			q = &ServiceInstanceAttachment{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProtectionAttachment_ReleaseChannel) EqualVT(thatIface isProtectionAttachment_Attachment) bool {
	that, ok := thatIface.(*ProtectionAttachment_ReleaseChannel)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.ReleaseChannel, that.ReleaseChannel; p != q {
		if p == nil {
			p = &ReleaseChannelAttachment{}
		}
		if q == nil {
			q = &ReleaseChannelAttachment{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProtectionAttachment_Convergence) EqualVT(thatIface isProtectionAttachment_Attachment) bool {
	that, ok := thatIface.(*ProtectionAttachment_Convergence)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Convergence, that.Convergence; p != q {
		if p == nil {
			p = &ConvergenceAttachment{}
		}
		if q == nil {
			q = &ConvergenceAttachment{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}
