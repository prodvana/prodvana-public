// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/common_config/env.proto

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

func (m *EnvValue) CloneVT() *EnvValue {
	if m == nil {
		return (*EnvValue)(nil)
	}
	r := new(EnvValue)
	if m.ValueOneof != nil {
		r.ValueOneof = m.ValueOneof.(interface{ CloneVT() isEnvValue_ValueOneof }).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *EnvValue) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *EnvValue_Value) CloneVT() isEnvValue_ValueOneof {
	if m == nil {
		return (*EnvValue_Value)(nil)
	}
	r := new(EnvValue_Value)
	r.Value = m.Value
	return r
}

func (m *EnvValue_RawSecret) CloneVT() isEnvValue_ValueOneof {
	if m == nil {
		return (*EnvValue_RawSecret)(nil)
	}
	r := new(EnvValue_RawSecret)
	r.RawSecret = m.RawSecret
	return r
}

func (m *EnvValue_Secret) CloneVT() isEnvValue_ValueOneof {
	if m == nil {
		return (*EnvValue_Secret)(nil)
	}
	r := new(EnvValue_Secret)
	r.Secret = m.Secret.CloneVT()
	return r
}

func (m *EnvValue_KubernetesSecret) CloneVT() isEnvValue_ValueOneof {
	if m == nil {
		return (*EnvValue_KubernetesSecret)(nil)
	}
	r := new(EnvValue_KubernetesSecret)
	r.KubernetesSecret = m.KubernetesSecret.CloneVT()
	return r
}

func (m *SecretReferenceValue) CloneVT() *SecretReferenceValue {
	if m == nil {
		return (*SecretReferenceValue)(nil)
	}
	r := new(SecretReferenceValue)
	if m.ValueOneof != nil {
		r.ValueOneof = m.ValueOneof.(interface {
			CloneVT() isSecretReferenceValue_ValueOneof
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *SecretReferenceValue) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *SecretReferenceValue_Secret) CloneVT() isSecretReferenceValue_ValueOneof {
	if m == nil {
		return (*SecretReferenceValue_Secret)(nil)
	}
	r := new(SecretReferenceValue_Secret)
	r.Secret = m.Secret.CloneVT()
	return r
}

func (m *SecretReferenceValue_KubernetesSecret) CloneVT() isSecretReferenceValue_ValueOneof {
	if m == nil {
		return (*SecretReferenceValue_KubernetesSecret)(nil)
	}
	r := new(SecretReferenceValue_KubernetesSecret)
	r.KubernetesSecret = m.KubernetesSecret.CloneVT()
	return r
}

func (m *NativeSecretReferenceValue) CloneVT() *NativeSecretReferenceValue {
	if m == nil {
		return (*NativeSecretReferenceValue)(nil)
	}
	r := new(NativeSecretReferenceValue)
	if m.ValueOneof != nil {
		r.ValueOneof = m.ValueOneof.(interface {
			CloneVT() isNativeSecretReferenceValue_ValueOneof
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *NativeSecretReferenceValue) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *NativeSecretReferenceValue_Secret) CloneVT() isNativeSecretReferenceValue_ValueOneof {
	if m == nil {
		return (*NativeSecretReferenceValue_Secret)(nil)
	}
	r := new(NativeSecretReferenceValue_Secret)
	r.Secret = m.Secret.CloneVT()
	return r
}

func (m *Secret) CloneVT() *Secret {
	if m == nil {
		return (*Secret)(nil)
	}
	r := new(Secret)
	r.Key = m.Key
	r.Version = m.Version
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *Secret) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *KubernetesSecret) CloneVT() *KubernetesSecret {
	if m == nil {
		return (*KubernetesSecret)(nil)
	}
	r := new(KubernetesSecret)
	r.SecretName = m.SecretName
	r.Key = m.Key
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *KubernetesSecret) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *EnvValue) StableEqualVT(that *EnvValue) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ValueOneof == nil && that.ValueOneof != nil {
		return false
	} else if this.ValueOneof != nil {
		if that.ValueOneof == nil {
			return false
		}
		if !this.ValueOneof.(interface {
			StableEqualVT(isEnvValue_ValueOneof) bool
		}).StableEqualVT(that.ValueOneof) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *EnvValue) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*EnvValue)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *EnvValue_Value) StableEqualVT(thatIface isEnvValue_ValueOneof) bool {
	that, ok := thatIface.(*EnvValue_Value)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.Value != that.Value {
		return false
	}
	return true
}

func (this *EnvValue_RawSecret) StableEqualVT(thatIface isEnvValue_ValueOneof) bool {
	that, ok := thatIface.(*EnvValue_RawSecret)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.RawSecret != that.RawSecret {
		return false
	}
	return true
}

func (this *EnvValue_Secret) StableEqualVT(thatIface isEnvValue_ValueOneof) bool {
	that, ok := thatIface.(*EnvValue_Secret)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Secret, that.Secret; p != q {
		if p == nil {
			p = &Secret{}
		}
		if q == nil {
			q = &Secret{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *EnvValue_KubernetesSecret) StableEqualVT(thatIface isEnvValue_ValueOneof) bool {
	that, ok := thatIface.(*EnvValue_KubernetesSecret)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.KubernetesSecret, that.KubernetesSecret; p != q {
		if p == nil {
			p = &KubernetesSecret{}
		}
		if q == nil {
			q = &KubernetesSecret{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *EnvValue) EqualVT(that *EnvValue) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ValueOneof == nil && that.ValueOneof != nil {
		return false
	} else if this.ValueOneof != nil {
		if that.ValueOneof == nil {
			return false
		}
		if !this.ValueOneof.(interface {
			EqualVT(isEnvValue_ValueOneof) bool
		}).EqualVT(that.ValueOneof) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *EnvValue) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*EnvValue)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *EnvValue_Value) EqualVT(thatIface isEnvValue_ValueOneof) bool {
	that, ok := thatIface.(*EnvValue_Value)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.Value != that.Value {
		return false
	}
	return true
}

func (this *EnvValue_RawSecret) EqualVT(thatIface isEnvValue_ValueOneof) bool {
	that, ok := thatIface.(*EnvValue_RawSecret)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.RawSecret != that.RawSecret {
		return false
	}
	return true
}

func (this *EnvValue_Secret) EqualVT(thatIface isEnvValue_ValueOneof) bool {
	that, ok := thatIface.(*EnvValue_Secret)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Secret, that.Secret; p != q {
		if p == nil {
			p = &Secret{}
		}
		if q == nil {
			q = &Secret{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *EnvValue_KubernetesSecret) EqualVT(thatIface isEnvValue_ValueOneof) bool {
	that, ok := thatIface.(*EnvValue_KubernetesSecret)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.KubernetesSecret, that.KubernetesSecret; p != q {
		if p == nil {
			p = &KubernetesSecret{}
		}
		if q == nil {
			q = &KubernetesSecret{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *SecretReferenceValue) StableEqualVT(that *SecretReferenceValue) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ValueOneof == nil && that.ValueOneof != nil {
		return false
	} else if this.ValueOneof != nil {
		if that.ValueOneof == nil {
			return false
		}
		if !this.ValueOneof.(interface {
			StableEqualVT(isSecretReferenceValue_ValueOneof) bool
		}).StableEqualVT(that.ValueOneof) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *SecretReferenceValue) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*SecretReferenceValue)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *SecretReferenceValue_Secret) StableEqualVT(thatIface isSecretReferenceValue_ValueOneof) bool {
	that, ok := thatIface.(*SecretReferenceValue_Secret)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Secret, that.Secret; p != q {
		if p == nil {
			p = &Secret{}
		}
		if q == nil {
			q = &Secret{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *SecretReferenceValue_KubernetesSecret) StableEqualVT(thatIface isSecretReferenceValue_ValueOneof) bool {
	that, ok := thatIface.(*SecretReferenceValue_KubernetesSecret)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.KubernetesSecret, that.KubernetesSecret; p != q {
		if p == nil {
			p = &KubernetesSecret{}
		}
		if q == nil {
			q = &KubernetesSecret{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *SecretReferenceValue) EqualVT(that *SecretReferenceValue) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ValueOneof == nil && that.ValueOneof != nil {
		return false
	} else if this.ValueOneof != nil {
		if that.ValueOneof == nil {
			return false
		}
		if !this.ValueOneof.(interface {
			EqualVT(isSecretReferenceValue_ValueOneof) bool
		}).EqualVT(that.ValueOneof) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *SecretReferenceValue) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*SecretReferenceValue)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *SecretReferenceValue_Secret) EqualVT(thatIface isSecretReferenceValue_ValueOneof) bool {
	that, ok := thatIface.(*SecretReferenceValue_Secret)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Secret, that.Secret; p != q {
		if p == nil {
			p = &Secret{}
		}
		if q == nil {
			q = &Secret{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *SecretReferenceValue_KubernetesSecret) EqualVT(thatIface isSecretReferenceValue_ValueOneof) bool {
	that, ok := thatIface.(*SecretReferenceValue_KubernetesSecret)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.KubernetesSecret, that.KubernetesSecret; p != q {
		if p == nil {
			p = &KubernetesSecret{}
		}
		if q == nil {
			q = &KubernetesSecret{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *NativeSecretReferenceValue) StableEqualVT(that *NativeSecretReferenceValue) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ValueOneof == nil && that.ValueOneof != nil {
		return false
	} else if this.ValueOneof != nil {
		if that.ValueOneof == nil {
			return false
		}
		if !this.ValueOneof.(interface {
			StableEqualVT(isNativeSecretReferenceValue_ValueOneof) bool
		}).StableEqualVT(that.ValueOneof) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *NativeSecretReferenceValue) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*NativeSecretReferenceValue)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *NativeSecretReferenceValue_Secret) StableEqualVT(thatIface isNativeSecretReferenceValue_ValueOneof) bool {
	that, ok := thatIface.(*NativeSecretReferenceValue_Secret)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Secret, that.Secret; p != q {
		if p == nil {
			p = &Secret{}
		}
		if q == nil {
			q = &Secret{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *NativeSecretReferenceValue) EqualVT(that *NativeSecretReferenceValue) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ValueOneof == nil && that.ValueOneof != nil {
		return false
	} else if this.ValueOneof != nil {
		if that.ValueOneof == nil {
			return false
		}
		if !this.ValueOneof.(interface {
			EqualVT(isNativeSecretReferenceValue_ValueOneof) bool
		}).EqualVT(that.ValueOneof) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *NativeSecretReferenceValue) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*NativeSecretReferenceValue)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *NativeSecretReferenceValue_Secret) EqualVT(thatIface isNativeSecretReferenceValue_ValueOneof) bool {
	that, ok := thatIface.(*NativeSecretReferenceValue_Secret)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Secret, that.Secret; p != q {
		if p == nil {
			p = &Secret{}
		}
		if q == nil {
			q = &Secret{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Secret) StableEqualVT(that *Secret) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Key != that.Key {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Secret) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Secret)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *Secret) EqualVT(that *Secret) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Key != that.Key {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Secret) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Secret)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *KubernetesSecret) StableEqualVT(that *KubernetesSecret) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.SecretName != that.SecretName {
		return false
	}
	if this.Key != that.Key {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *KubernetesSecret) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*KubernetesSecret)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *KubernetesSecret) EqualVT(that *KubernetesSecret) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.SecretName != that.SecretName {
		return false
	}
	if this.Key != that.Key {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *KubernetesSecret) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*KubernetesSecret)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}