// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/workflow/integration_config.proto

package workflow

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

func (m *AlertingConfig_PagerDuty) CloneVT() *AlertingConfig_PagerDuty {
	if m == nil {
		return (*AlertingConfig_PagerDuty)(nil)
	}
	r := new(AlertingConfig_PagerDuty)
	r.Service = m.Service
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *AlertingConfig_PagerDuty) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *AlertingConfig) CloneVT() *AlertingConfig {
	if m == nil {
		return (*AlertingConfig)(nil)
	}
	r := new(AlertingConfig)
	r.Pagerduty = m.Pagerduty.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *AlertingConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *AnnotationsConfig_Honeycomb) CloneVT() *AnnotationsConfig_Honeycomb {
	if m == nil {
		return (*AnnotationsConfig_Honeycomb)(nil)
	}
	r := new(AnnotationsConfig_Honeycomb)
	r.Environment = m.Environment
	r.Dataset = m.Dataset
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *AnnotationsConfig_Honeycomb) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *AnnotationsConfig) CloneVT() *AnnotationsConfig {
	if m == nil {
		return (*AnnotationsConfig)(nil)
	}
	r := new(AnnotationsConfig)
	r.Honeycomb = m.Honeycomb.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *AnnotationsConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *TokenConfig) CloneVT() *TokenConfig {
	if m == nil {
		return (*TokenConfig)(nil)
	}
	r := new(TokenConfig)
	r.TokenSecretKey = m.TokenSecretKey
	r.TokenSecretVersion = m.TokenSecretVersion
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *TokenConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *IntegrationConfig) CloneVT() *IntegrationConfig {
	if m == nil {
		return (*IntegrationConfig)(nil)
	}
	r := new(IntegrationConfig)
	if m.ConfigOneof != nil {
		r.ConfigOneof = m.ConfigOneof.(interface {
			CloneVT() isIntegrationConfig_ConfigOneof
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *IntegrationConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *IntegrationConfig_SlackConfig) CloneVT() isIntegrationConfig_ConfigOneof {
	if m == nil {
		return (*IntegrationConfig_SlackConfig)(nil)
	}
	r := new(IntegrationConfig_SlackConfig)
	r.SlackConfig = m.SlackConfig.CloneVT()
	return r
}

func (m *IntegrationConfig_PagerdutyConfig) CloneVT() isIntegrationConfig_ConfigOneof {
	if m == nil {
		return (*IntegrationConfig_PagerdutyConfig)(nil)
	}
	r := new(IntegrationConfig_PagerdutyConfig)
	r.PagerdutyConfig = m.PagerdutyConfig.CloneVT()
	return r
}

func (this *AlertingConfig_PagerDuty) StableEqualVT(that *AlertingConfig_PagerDuty) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Service != that.Service {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *AlertingConfig_PagerDuty) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*AlertingConfig_PagerDuty)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *AlertingConfig) StableEqualVT(that *AlertingConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Pagerduty.StableEqualVT(that.Pagerduty) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *AlertingConfig) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*AlertingConfig)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *AlertingConfig_PagerDuty) EqualVT(that *AlertingConfig_PagerDuty) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Service != that.Service {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *AlertingConfig_PagerDuty) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*AlertingConfig_PagerDuty)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *AlertingConfig) EqualVT(that *AlertingConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Pagerduty.EqualVT(that.Pagerduty) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *AlertingConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*AlertingConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *AnnotationsConfig_Honeycomb) StableEqualVT(that *AnnotationsConfig_Honeycomb) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Environment != that.Environment {
		return false
	}
	if this.Dataset != that.Dataset {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *AnnotationsConfig_Honeycomb) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*AnnotationsConfig_Honeycomb)
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
	if !this.Honeycomb.StableEqualVT(that.Honeycomb) {
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
func (this *AnnotationsConfig_Honeycomb) EqualVT(that *AnnotationsConfig_Honeycomb) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Environment != that.Environment {
		return false
	}
	if this.Dataset != that.Dataset {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *AnnotationsConfig_Honeycomb) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*AnnotationsConfig_Honeycomb)
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
	if !this.Honeycomb.EqualVT(that.Honeycomb) {
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
func (this *TokenConfig) StableEqualVT(that *TokenConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.TokenSecretKey != that.TokenSecretKey {
		return false
	}
	if this.TokenSecretVersion != that.TokenSecretVersion {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *TokenConfig) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*TokenConfig)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *TokenConfig) EqualVT(that *TokenConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.TokenSecretKey != that.TokenSecretKey {
		return false
	}
	if this.TokenSecretVersion != that.TokenSecretVersion {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *TokenConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*TokenConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *IntegrationConfig) StableEqualVT(that *IntegrationConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ConfigOneof == nil && that.ConfigOneof != nil {
		return false
	} else if this.ConfigOneof != nil {
		if that.ConfigOneof == nil {
			return false
		}
		if !this.ConfigOneof.(interface {
			StableEqualVT(isIntegrationConfig_ConfigOneof) bool
		}).StableEqualVT(that.ConfigOneof) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *IntegrationConfig) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*IntegrationConfig)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *IntegrationConfig_SlackConfig) StableEqualVT(thatIface isIntegrationConfig_ConfigOneof) bool {
	that, ok := thatIface.(*IntegrationConfig_SlackConfig)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.SlackConfig, that.SlackConfig; p != q {
		if p == nil {
			p = &TokenConfig{}
		}
		if q == nil {
			q = &TokenConfig{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *IntegrationConfig_PagerdutyConfig) StableEqualVT(thatIface isIntegrationConfig_ConfigOneof) bool {
	that, ok := thatIface.(*IntegrationConfig_PagerdutyConfig)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.PagerdutyConfig, that.PagerdutyConfig; p != q {
		if p == nil {
			p = &TokenConfig{}
		}
		if q == nil {
			q = &TokenConfig{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *IntegrationConfig) EqualVT(that *IntegrationConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ConfigOneof == nil && that.ConfigOneof != nil {
		return false
	} else if this.ConfigOneof != nil {
		if that.ConfigOneof == nil {
			return false
		}
		if !this.ConfigOneof.(interface {
			EqualVT(isIntegrationConfig_ConfigOneof) bool
		}).EqualVT(that.ConfigOneof) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *IntegrationConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*IntegrationConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *IntegrationConfig_SlackConfig) EqualVT(thatIface isIntegrationConfig_ConfigOneof) bool {
	that, ok := thatIface.(*IntegrationConfig_SlackConfig)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.SlackConfig, that.SlackConfig; p != q {
		if p == nil {
			p = &TokenConfig{}
		}
		if q == nil {
			q = &TokenConfig{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *IntegrationConfig_PagerdutyConfig) EqualVT(thatIface isIntegrationConfig_ConfigOneof) bool {
	that, ok := thatIface.(*IntegrationConfig_PagerdutyConfig)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.PagerdutyConfig, that.PagerdutyConfig; p != q {
		if p == nil {
			p = &TokenConfig{}
		}
		if q == nil {
			q = &TokenConfig{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}
