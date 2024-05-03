// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/capability/capability.proto

package capability

import (
	common_config "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	pipelines "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/pipelines"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *PerReleaseChannelCapabilityConfig) CloneVT() *PerReleaseChannelCapabilityConfig {
	if m == nil {
		return (*PerReleaseChannelCapabilityConfig)(nil)
	}
	r := new(PerReleaseChannelCapabilityConfig)
	r.ReleaseChannel = m.ReleaseChannel
	if m.DefinitionOneof != nil {
		r.DefinitionOneof = m.DefinitionOneof.(interface {
			CloneVT() isPerReleaseChannelCapabilityConfig_DefinitionOneof
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *PerReleaseChannelCapabilityConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *PerReleaseChannelCapabilityConfig_Inlined) CloneVT() isPerReleaseChannelCapabilityConfig_DefinitionOneof {
	if m == nil {
		return (*PerReleaseChannelCapabilityConfig_Inlined)(nil)
	}
	r := new(PerReleaseChannelCapabilityConfig_Inlined)
	r.Inlined = m.Inlined.CloneVT()
	return r
}

func (m *PerReleaseChannelCapabilityConfig_Ref) CloneVT() isPerReleaseChannelCapabilityConfig_DefinitionOneof {
	if m == nil {
		return (*PerReleaseChannelCapabilityConfig_Ref)(nil)
	}
	r := new(PerReleaseChannelCapabilityConfig_Ref)
	r.Ref = m.Ref.CloneVT()
	return r
}

func (m *CapabilityConfig) CloneVT() *CapabilityConfig {
	if m == nil {
		return (*CapabilityConfig)(nil)
	}
	r := new(CapabilityConfig)
	r.Name = m.Name
	if rhs := m.PerReleaseChannel; rhs != nil {
		tmpContainer := make([]*PerReleaseChannelCapabilityConfig, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.PerReleaseChannel = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *CapabilityConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *CapabilityInstanceConfig) CloneVT() *CapabilityInstanceConfig {
	if m == nil {
		return (*CapabilityInstanceConfig)(nil)
	}
	r := new(CapabilityInstanceConfig)
	r.Name = m.Name
	if rhs := m.Env; rhs != nil {
		tmpContainer := make(map[string]*common_config.EnvValue, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Env = tmpContainer
	}
	if rhs := m.PrePushTasks; rhs != nil {
		tmpContainer := make([]*pipelines.Task, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.PrePushTasks = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *CapabilityInstanceConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *CapabilityInstanceRef) CloneVT() *CapabilityInstanceRef {
	if m == nil {
		return (*CapabilityInstanceRef)(nil)
	}
	r := new(CapabilityInstanceRef)
	r.Name = m.Name
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *CapabilityInstanceRef) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *PerReleaseChannelCapabilityConfig) StableEqualVT(that *PerReleaseChannelCapabilityConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.DefinitionOneof == nil && that.DefinitionOneof != nil {
		return false
	} else if this.DefinitionOneof != nil {
		if that.DefinitionOneof == nil {
			return false
		}
		if !this.DefinitionOneof.(interface {
			StableEqualVT(isPerReleaseChannelCapabilityConfig_DefinitionOneof) bool
		}).StableEqualVT(that.DefinitionOneof) {
			return false
		}
	}
	if this.ReleaseChannel != that.ReleaseChannel {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *PerReleaseChannelCapabilityConfig) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*PerReleaseChannelCapabilityConfig)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *PerReleaseChannelCapabilityConfig_Inlined) StableEqualVT(thatIface isPerReleaseChannelCapabilityConfig_DefinitionOneof) bool {
	that, ok := thatIface.(*PerReleaseChannelCapabilityConfig_Inlined)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Inlined, that.Inlined; p != q {
		if p == nil {
			p = &CapabilityInstanceConfig{}
		}
		if q == nil {
			q = &CapabilityInstanceConfig{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *PerReleaseChannelCapabilityConfig_Ref) StableEqualVT(thatIface isPerReleaseChannelCapabilityConfig_DefinitionOneof) bool {
	that, ok := thatIface.(*PerReleaseChannelCapabilityConfig_Ref)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Ref, that.Ref; p != q {
		if p == nil {
			p = &CapabilityInstanceRef{}
		}
		if q == nil {
			q = &CapabilityInstanceRef{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *PerReleaseChannelCapabilityConfig) EqualVT(that *PerReleaseChannelCapabilityConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.DefinitionOneof == nil && that.DefinitionOneof != nil {
		return false
	} else if this.DefinitionOneof != nil {
		if that.DefinitionOneof == nil {
			return false
		}
		if !this.DefinitionOneof.(interface {
			EqualVT(isPerReleaseChannelCapabilityConfig_DefinitionOneof) bool
		}).EqualVT(that.DefinitionOneof) {
			return false
		}
	}
	if this.ReleaseChannel != that.ReleaseChannel {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *PerReleaseChannelCapabilityConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*PerReleaseChannelCapabilityConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *PerReleaseChannelCapabilityConfig_Inlined) EqualVT(thatIface isPerReleaseChannelCapabilityConfig_DefinitionOneof) bool {
	that, ok := thatIface.(*PerReleaseChannelCapabilityConfig_Inlined)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Inlined, that.Inlined; p != q {
		if p == nil {
			p = &CapabilityInstanceConfig{}
		}
		if q == nil {
			q = &CapabilityInstanceConfig{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *PerReleaseChannelCapabilityConfig_Ref) EqualVT(thatIface isPerReleaseChannelCapabilityConfig_DefinitionOneof) bool {
	that, ok := thatIface.(*PerReleaseChannelCapabilityConfig_Ref)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Ref, that.Ref; p != q {
		if p == nil {
			p = &CapabilityInstanceRef{}
		}
		if q == nil {
			q = &CapabilityInstanceRef{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *CapabilityConfig) StableEqualVT(that *CapabilityConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	if len(this.PerReleaseChannel) != len(that.PerReleaseChannel) {
		return false
	}
	for i, vx := range this.PerReleaseChannel {
		vy := that.PerReleaseChannel[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &PerReleaseChannelCapabilityConfig{}
			}
			if q == nil {
				q = &PerReleaseChannelCapabilityConfig{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *CapabilityConfig) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*CapabilityConfig)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *CapabilityConfig) EqualVT(that *CapabilityConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	if len(this.PerReleaseChannel) != len(that.PerReleaseChannel) {
		return false
	}
	for i, vx := range this.PerReleaseChannel {
		vy := that.PerReleaseChannel[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &PerReleaseChannelCapabilityConfig{}
			}
			if q == nil {
				q = &PerReleaseChannelCapabilityConfig{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *CapabilityConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*CapabilityConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *CapabilityInstanceConfig) StableEqualVT(that *CapabilityInstanceConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Name != that.Name {
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
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if len(this.PrePushTasks) != len(that.PrePushTasks) {
		return false
	}
	for i, vx := range this.PrePushTasks {
		vy := that.PrePushTasks[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &pipelines.Task{}
			}
			if q == nil {
				q = &pipelines.Task{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *CapabilityInstanceConfig) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*CapabilityInstanceConfig)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *CapabilityInstanceConfig) EqualVT(that *CapabilityInstanceConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Name != that.Name {
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
	if len(this.PrePushTasks) != len(that.PrePushTasks) {
		return false
	}
	for i, vx := range this.PrePushTasks {
		vy := that.PrePushTasks[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &pipelines.Task{}
			}
			if q == nil {
				q = &pipelines.Task{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *CapabilityInstanceConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*CapabilityInstanceConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *CapabilityInstanceRef) StableEqualVT(that *CapabilityInstanceRef) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *CapabilityInstanceRef) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*CapabilityInstanceRef)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *CapabilityInstanceRef) EqualVT(that *CapabilityInstanceRef) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *CapabilityInstanceRef) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*CapabilityInstanceRef)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}