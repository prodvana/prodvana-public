// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/application/application_config.proto

package application

import (
	capability "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/capability"
	release_channel "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/release_channel"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *AnnotationsConfig_Last9) CloneVT() *AnnotationsConfig_Last9 {
	if m == nil {
		return (*AnnotationsConfig_Last9)(nil)
	}
	r := new(AnnotationsConfig_Last9)
	r.DataSource = m.DataSource
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *AnnotationsConfig_Last9) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *AnnotationsConfig) CloneVT() *AnnotationsConfig {
	if m == nil {
		return (*AnnotationsConfig)(nil)
	}
	r := new(AnnotationsConfig)
	r.Last9 = m.Last9.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *AnnotationsConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ApplicationConfig) CloneVT() *ApplicationConfig {
	if m == nil {
		return (*ApplicationConfig)(nil)
	}
	r := new(ApplicationConfig)
	r.Name = m.Name
	r.Notifications = m.Notifications.CloneVT()
	r.Annotations = m.Annotations.CloneVT()
	r.ReleaseSettings = m.ReleaseSettings.CloneVT()
	r.Alerts = m.Alerts.CloneVT()
	if rhs := m.ReleaseChannels; rhs != nil {
		tmpContainer := make([]*release_channel.ReleaseChannelConfig, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.ReleaseChannels = tmpContainer
	}
	if rhs := m.ReleaseChannelGroups; rhs != nil {
		tmpContainer := make([]*release_channel.ReleaseChannelGroupGeneratorConfig, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.ReleaseChannelGroups = tmpContainer
	}
	if rhs := m.Capabilities; rhs != nil {
		tmpContainer := make([]*capability.CapabilityConfig, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Capabilities = tmpContainer
	}
	if rhs := m.CapabilityInstances; rhs != nil {
		tmpContainer := make([]*capability.CapabilityInstanceConfig, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.CapabilityInstances = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ApplicationConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *AnnotationsConfig_Last9) StableEqualVT(that *AnnotationsConfig_Last9) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.DataSource != that.DataSource {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *AnnotationsConfig_Last9) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*AnnotationsConfig_Last9)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *AnnotationsConfig) StableEqualVT(that *AnnotationsConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Last9.StableEqualVT(that.Last9) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *AnnotationsConfig) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*AnnotationsConfig)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *AnnotationsConfig_Last9) EqualVT(that *AnnotationsConfig_Last9) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.DataSource != that.DataSource {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *AnnotationsConfig_Last9) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*AnnotationsConfig_Last9)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *AnnotationsConfig) EqualVT(that *AnnotationsConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Last9.EqualVT(that.Last9) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *AnnotationsConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*AnnotationsConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ApplicationConfig) StableEqualVT(that *ApplicationConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	if len(this.ReleaseChannels) != len(that.ReleaseChannels) {
		return false
	}
	for i, vx := range this.ReleaseChannels {
		vy := that.ReleaseChannels[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &release_channel.ReleaseChannelConfig{}
			}
			if q == nil {
				q = &release_channel.ReleaseChannelConfig{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if !this.Notifications.StableEqualVT(that.Notifications) {
		return false
	}
	if !this.Alerts.StableEqualVT(that.Alerts) {
		return false
	}
	if len(this.Capabilities) != len(that.Capabilities) {
		return false
	}
	for i, vx := range this.Capabilities {
		vy := that.Capabilities[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &capability.CapabilityConfig{}
			}
			if q == nil {
				q = &capability.CapabilityConfig{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if len(this.CapabilityInstances) != len(that.CapabilityInstances) {
		return false
	}
	for i, vx := range this.CapabilityInstances {
		vy := that.CapabilityInstances[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &capability.CapabilityInstanceConfig{}
			}
			if q == nil {
				q = &capability.CapabilityInstanceConfig{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if len(this.ReleaseChannelGroups) != len(that.ReleaseChannelGroups) {
		return false
	}
	for i, vx := range this.ReleaseChannelGroups {
		vy := that.ReleaseChannelGroups[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &release_channel.ReleaseChannelGroupGeneratorConfig{}
			}
			if q == nil {
				q = &release_channel.ReleaseChannelGroupGeneratorConfig{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if !this.Annotations.StableEqualVT(that.Annotations) {
		return false
	}
	if !this.ReleaseSettings.StableEqualVT(that.ReleaseSettings) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ApplicationConfig) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ApplicationConfig)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ApplicationConfig) EqualVT(that *ApplicationConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	if len(this.ReleaseChannels) != len(that.ReleaseChannels) {
		return false
	}
	for i, vx := range this.ReleaseChannels {
		vy := that.ReleaseChannels[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &release_channel.ReleaseChannelConfig{}
			}
			if q == nil {
				q = &release_channel.ReleaseChannelConfig{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if !this.Notifications.EqualVT(that.Notifications) {
		return false
	}
	if !this.Alerts.EqualVT(that.Alerts) {
		return false
	}
	if len(this.Capabilities) != len(that.Capabilities) {
		return false
	}
	for i, vx := range this.Capabilities {
		vy := that.Capabilities[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &capability.CapabilityConfig{}
			}
			if q == nil {
				q = &capability.CapabilityConfig{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.CapabilityInstances) != len(that.CapabilityInstances) {
		return false
	}
	for i, vx := range this.CapabilityInstances {
		vy := that.CapabilityInstances[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &capability.CapabilityInstanceConfig{}
			}
			if q == nil {
				q = &capability.CapabilityInstanceConfig{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.ReleaseChannelGroups) != len(that.ReleaseChannelGroups) {
		return false
	}
	for i, vx := range this.ReleaseChannelGroups {
		vy := that.ReleaseChannelGroups[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &release_channel.ReleaseChannelGroupGeneratorConfig{}
			}
			if q == nil {
				q = &release_channel.ReleaseChannelGroupGeneratorConfig{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if !this.Annotations.EqualVT(that.Annotations) {
		return false
	}
	if !this.ReleaseSettings.EqualVT(that.ReleaseSettings) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ApplicationConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ApplicationConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
