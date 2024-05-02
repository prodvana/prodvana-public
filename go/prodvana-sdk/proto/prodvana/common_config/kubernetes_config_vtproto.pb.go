// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/common_config/kubernetes_config.proto

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

func (m *LocalConfig) CloneVT() *LocalConfig {
	if m == nil {
		return (*LocalConfig)(nil)
	}
	r := new(LocalConfig)
	r.SubPath = m.SubPath
	if m.PathOneof != nil {
		r.PathOneof = m.PathOneof.(interface {
			CloneVT() isLocalConfig_PathOneof
		}).CloneVT()
	}
	if rhs := m.Paths; rhs != nil {
		tmpContainer := make([]string, len(rhs))
		copy(tmpContainer, rhs)
		r.Paths = tmpContainer
	}
	if rhs := m.ExcludePatterns; rhs != nil {
		tmpContainer := make([]string, len(rhs))
		copy(tmpContainer, rhs)
		r.ExcludePatterns = tmpContainer
	}
	if rhs := m.IncludePatterns; rhs != nil {
		tmpContainer := make([]string, len(rhs))
		copy(tmpContainer, rhs)
		r.IncludePatterns = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *LocalConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *LocalConfig_Path) CloneVT() isLocalConfig_PathOneof {
	if m == nil {
		return (*LocalConfig_Path)(nil)
	}
	r := new(LocalConfig_Path)
	r.Path = m.Path
	return r
}

func (m *RemoteConfig) CloneVT() *RemoteConfig {
	if m == nil {
		return (*RemoteConfig)(nil)
	}
	r := new(RemoteConfig)
	r.SubPath = m.SubPath
	if m.RemoteOneof != nil {
		r.RemoteOneof = m.RemoteOneof.(interface {
			CloneVT() isRemoteConfig_RemoteOneof
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *RemoteConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *RemoteConfig_TarballBlobId) CloneVT() isRemoteConfig_RemoteOneof {
	if m == nil {
		return (*RemoteConfig_TarballBlobId)(nil)
	}
	r := new(RemoteConfig_TarballBlobId)
	r.TarballBlobId = m.TarballBlobId
	return r
}

func (m *KubernetesConfig) CloneVT() *KubernetesConfig {
	if m == nil {
		return (*KubernetesConfig)(nil)
	}
	r := new(KubernetesConfig)
	r.Type = m.Type
	r.EnvInjectionMode = m.EnvInjectionMode
	if m.SourceOneof != nil {
		r.SourceOneof = m.SourceOneof.(interface {
			CloneVT() isKubernetesConfig_SourceOneof
		}).CloneVT()
	}
	if rhs := m.Patches; rhs != nil {
		tmpContainer := make([]*KubernetesPatch, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Patches = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *KubernetesConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *KubernetesConfig_Inlined) CloneVT() isKubernetesConfig_SourceOneof {
	if m == nil {
		return (*KubernetesConfig_Inlined)(nil)
	}
	r := new(KubernetesConfig_Inlined)
	r.Inlined = m.Inlined
	return r
}

func (m *KubernetesConfig_Local) CloneVT() isKubernetesConfig_SourceOneof {
	if m == nil {
		return (*KubernetesConfig_Local)(nil)
	}
	r := new(KubernetesConfig_Local)
	r.Local = m.Local.CloneVT()
	return r
}

func (m *KubernetesConfig_Remote) CloneVT() isKubernetesConfig_SourceOneof {
	if m == nil {
		return (*KubernetesConfig_Remote)(nil)
	}
	r := new(KubernetesConfig_Remote)
	r.Remote = m.Remote.CloneVT()
	return r
}

func (m *KubernetesPatch_Target) CloneVT() *KubernetesPatch_Target {
	if m == nil {
		return (*KubernetesPatch_Target)(nil)
	}
	r := new(KubernetesPatch_Target)
	r.Group = m.Group
	r.Version = m.Version
	r.Kind = m.Kind
	r.Name = m.Name
	r.Namespace = m.Namespace
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *KubernetesPatch_Target) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *KubernetesPatch_Replace) CloneVT() *KubernetesPatch_Replace {
	if m == nil {
		return (*KubernetesPatch_Replace)(nil)
	}
	r := new(KubernetesPatch_Replace)
	r.Path = m.Path
	if m.ValueOneof != nil {
		r.ValueOneof = m.ValueOneof.(interface {
			CloneVT() isKubernetesPatch_Replace_ValueOneof
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *KubernetesPatch_Replace) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *KubernetesPatch_Replace_String_) CloneVT() isKubernetesPatch_Replace_ValueOneof {
	if m == nil {
		return (*KubernetesPatch_Replace_String_)(nil)
	}
	r := new(KubernetesPatch_Replace_String_)
	r.String_ = m.String_
	return r
}

func (m *KubernetesPatch_Replace_IntAsString) CloneVT() isKubernetesPatch_Replace_ValueOneof {
	if m == nil {
		return (*KubernetesPatch_Replace_IntAsString)(nil)
	}
	r := new(KubernetesPatch_Replace_IntAsString)
	r.IntAsString = m.IntAsString
	return r
}

func (m *KubernetesPatch) CloneVT() *KubernetesPatch {
	if m == nil {
		return (*KubernetesPatch)(nil)
	}
	r := new(KubernetesPatch)
	r.Target = m.Target.CloneVT()
	if m.PatchOneof != nil {
		r.PatchOneof = m.PatchOneof.(interface {
			CloneVT() isKubernetesPatch_PatchOneof
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *KubernetesPatch) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *KubernetesPatch_Replace_) CloneVT() isKubernetesPatch_PatchOneof {
	if m == nil {
		return (*KubernetesPatch_Replace_)(nil)
	}
	r := new(KubernetesPatch_Replace_)
	r.Replace = m.Replace.CloneVT()
	return r
}

func (this *LocalConfig) StableEqualVT(that *LocalConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.PathOneof == nil && that.PathOneof != nil {
		return false
	} else if this.PathOneof != nil {
		if that.PathOneof == nil {
			return false
		}
		if !this.PathOneof.(interface {
			StableEqualVT(isLocalConfig_PathOneof) bool
		}).StableEqualVT(that.PathOneof) {
			return false
		}
	}
	if len(this.Paths) != len(that.Paths) {
		return false
	}
	for i, vx := range this.Paths {
		vy := that.Paths[i]
		if vx != vy {
			return false
		}
	}
	if this.SubPath != that.SubPath {
		return false
	}
	if len(this.ExcludePatterns) != len(that.ExcludePatterns) {
		return false
	}
	for i, vx := range this.ExcludePatterns {
		vy := that.ExcludePatterns[i]
		if vx != vy {
			return false
		}
	}
	if len(this.IncludePatterns) != len(that.IncludePatterns) {
		return false
	}
	for i, vx := range this.IncludePatterns {
		vy := that.IncludePatterns[i]
		if vx != vy {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *LocalConfig) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*LocalConfig)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *LocalConfig_Path) StableEqualVT(thatIface isLocalConfig_PathOneof) bool {
	that, ok := thatIface.(*LocalConfig_Path)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.Path != that.Path {
		return false
	}
	return true
}

func (this *LocalConfig) EqualVT(that *LocalConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.PathOneof == nil && that.PathOneof != nil {
		return false
	} else if this.PathOneof != nil {
		if that.PathOneof == nil {
			return false
		}
		if !this.PathOneof.(interface {
			EqualVT(isLocalConfig_PathOneof) bool
		}).EqualVT(that.PathOneof) {
			return false
		}
	}
	if len(this.Paths) != len(that.Paths) {
		return false
	}
	for i, vx := range this.Paths {
		vy := that.Paths[i]
		if vx != vy {
			return false
		}
	}
	if this.SubPath != that.SubPath {
		return false
	}
	if len(this.ExcludePatterns) != len(that.ExcludePatterns) {
		return false
	}
	for i, vx := range this.ExcludePatterns {
		vy := that.ExcludePatterns[i]
		if vx != vy {
			return false
		}
	}
	if len(this.IncludePatterns) != len(that.IncludePatterns) {
		return false
	}
	for i, vx := range this.IncludePatterns {
		vy := that.IncludePatterns[i]
		if vx != vy {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *LocalConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*LocalConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *LocalConfig_Path) EqualVT(thatIface isLocalConfig_PathOneof) bool {
	that, ok := thatIface.(*LocalConfig_Path)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.Path != that.Path {
		return false
	}
	return true
}

func (this *RemoteConfig) StableEqualVT(that *RemoteConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.RemoteOneof == nil && that.RemoteOneof != nil {
		return false
	} else if this.RemoteOneof != nil {
		if that.RemoteOneof == nil {
			return false
		}
		if !this.RemoteOneof.(interface {
			StableEqualVT(isRemoteConfig_RemoteOneof) bool
		}).StableEqualVT(that.RemoteOneof) {
			return false
		}
	}
	if this.SubPath != that.SubPath {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *RemoteConfig) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*RemoteConfig)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *RemoteConfig_TarballBlobId) StableEqualVT(thatIface isRemoteConfig_RemoteOneof) bool {
	that, ok := thatIface.(*RemoteConfig_TarballBlobId)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.TarballBlobId != that.TarballBlobId {
		return false
	}
	return true
}

func (this *RemoteConfig) EqualVT(that *RemoteConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.RemoteOneof == nil && that.RemoteOneof != nil {
		return false
	} else if this.RemoteOneof != nil {
		if that.RemoteOneof == nil {
			return false
		}
		if !this.RemoteOneof.(interface {
			EqualVT(isRemoteConfig_RemoteOneof) bool
		}).EqualVT(that.RemoteOneof) {
			return false
		}
	}
	if this.SubPath != that.SubPath {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *RemoteConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*RemoteConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *RemoteConfig_TarballBlobId) EqualVT(thatIface isRemoteConfig_RemoteOneof) bool {
	that, ok := thatIface.(*RemoteConfig_TarballBlobId)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.TarballBlobId != that.TarballBlobId {
		return false
	}
	return true
}

func (this *KubernetesConfig) StableEqualVT(that *KubernetesConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.SourceOneof == nil && that.SourceOneof != nil {
		return false
	} else if this.SourceOneof != nil {
		if that.SourceOneof == nil {
			return false
		}
		if !this.SourceOneof.(interface {
			StableEqualVT(isKubernetesConfig_SourceOneof) bool
		}).StableEqualVT(that.SourceOneof) {
			return false
		}
	}
	if this.Type != that.Type {
		return false
	}
	if this.EnvInjectionMode != that.EnvInjectionMode {
		return false
	}
	if len(this.Patches) != len(that.Patches) {
		return false
	}
	for i, vx := range this.Patches {
		vy := that.Patches[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &KubernetesPatch{}
			}
			if q == nil {
				q = &KubernetesPatch{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *KubernetesConfig) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*KubernetesConfig)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *KubernetesConfig_Inlined) StableEqualVT(thatIface isKubernetesConfig_SourceOneof) bool {
	that, ok := thatIface.(*KubernetesConfig_Inlined)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.Inlined != that.Inlined {
		return false
	}
	return true
}

func (this *KubernetesConfig_Local) StableEqualVT(thatIface isKubernetesConfig_SourceOneof) bool {
	that, ok := thatIface.(*KubernetesConfig_Local)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Local, that.Local; p != q {
		if p == nil {
			p = &LocalConfig{}
		}
		if q == nil {
			q = &LocalConfig{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *KubernetesConfig_Remote) StableEqualVT(thatIface isKubernetesConfig_SourceOneof) bool {
	that, ok := thatIface.(*KubernetesConfig_Remote)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Remote, that.Remote; p != q {
		if p == nil {
			p = &RemoteConfig{}
		}
		if q == nil {
			q = &RemoteConfig{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *KubernetesConfig) EqualVT(that *KubernetesConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.SourceOneof == nil && that.SourceOneof != nil {
		return false
	} else if this.SourceOneof != nil {
		if that.SourceOneof == nil {
			return false
		}
		if !this.SourceOneof.(interface {
			EqualVT(isKubernetesConfig_SourceOneof) bool
		}).EqualVT(that.SourceOneof) {
			return false
		}
	}
	if this.Type != that.Type {
		return false
	}
	if this.EnvInjectionMode != that.EnvInjectionMode {
		return false
	}
	if len(this.Patches) != len(that.Patches) {
		return false
	}
	for i, vx := range this.Patches {
		vy := that.Patches[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &KubernetesPatch{}
			}
			if q == nil {
				q = &KubernetesPatch{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *KubernetesConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*KubernetesConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *KubernetesConfig_Inlined) EqualVT(thatIface isKubernetesConfig_SourceOneof) bool {
	that, ok := thatIface.(*KubernetesConfig_Inlined)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.Inlined != that.Inlined {
		return false
	}
	return true
}

func (this *KubernetesConfig_Local) EqualVT(thatIface isKubernetesConfig_SourceOneof) bool {
	that, ok := thatIface.(*KubernetesConfig_Local)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Local, that.Local; p != q {
		if p == nil {
			p = &LocalConfig{}
		}
		if q == nil {
			q = &LocalConfig{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *KubernetesConfig_Remote) EqualVT(thatIface isKubernetesConfig_SourceOneof) bool {
	that, ok := thatIface.(*KubernetesConfig_Remote)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Remote, that.Remote; p != q {
		if p == nil {
			p = &RemoteConfig{}
		}
		if q == nil {
			q = &RemoteConfig{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *KubernetesPatch_Target) StableEqualVT(that *KubernetesPatch_Target) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Group != that.Group {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	if this.Kind != that.Kind {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	if this.Namespace != that.Namespace {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *KubernetesPatch_Target) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*KubernetesPatch_Target)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *KubernetesPatch_Replace) StableEqualVT(that *KubernetesPatch_Replace) bool {
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
			StableEqualVT(isKubernetesPatch_Replace_ValueOneof) bool
		}).StableEqualVT(that.ValueOneof) {
			return false
		}
	}
	if this.Path != that.Path {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *KubernetesPatch_Replace) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*KubernetesPatch_Replace)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *KubernetesPatch_Replace_String_) StableEqualVT(thatIface isKubernetesPatch_Replace_ValueOneof) bool {
	that, ok := thatIface.(*KubernetesPatch_Replace_String_)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.String_ != that.String_ {
		return false
	}
	return true
}

func (this *KubernetesPatch_Replace_IntAsString) StableEqualVT(thatIface isKubernetesPatch_Replace_ValueOneof) bool {
	that, ok := thatIface.(*KubernetesPatch_Replace_IntAsString)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.IntAsString != that.IntAsString {
		return false
	}
	return true
}

func (this *KubernetesPatch) StableEqualVT(that *KubernetesPatch) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.PatchOneof == nil && that.PatchOneof != nil {
		return false
	} else if this.PatchOneof != nil {
		if that.PatchOneof == nil {
			return false
		}
		if !this.PatchOneof.(interface {
			StableEqualVT(isKubernetesPatch_PatchOneof) bool
		}).StableEqualVT(that.PatchOneof) {
			return false
		}
	}
	if !this.Target.StableEqualVT(that.Target) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *KubernetesPatch) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*KubernetesPatch)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *KubernetesPatch_Replace_) StableEqualVT(thatIface isKubernetesPatch_PatchOneof) bool {
	that, ok := thatIface.(*KubernetesPatch_Replace_)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Replace, that.Replace; p != q {
		if p == nil {
			p = &KubernetesPatch_Replace{}
		}
		if q == nil {
			q = &KubernetesPatch_Replace{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *KubernetesPatch_Target) EqualVT(that *KubernetesPatch_Target) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Group != that.Group {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	if this.Kind != that.Kind {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	if this.Namespace != that.Namespace {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *KubernetesPatch_Target) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*KubernetesPatch_Target)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *KubernetesPatch_Replace) EqualVT(that *KubernetesPatch_Replace) bool {
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
			EqualVT(isKubernetesPatch_Replace_ValueOneof) bool
		}).EqualVT(that.ValueOneof) {
			return false
		}
	}
	if this.Path != that.Path {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *KubernetesPatch_Replace) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*KubernetesPatch_Replace)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *KubernetesPatch_Replace_String_) EqualVT(thatIface isKubernetesPatch_Replace_ValueOneof) bool {
	that, ok := thatIface.(*KubernetesPatch_Replace_String_)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.String_ != that.String_ {
		return false
	}
	return true
}

func (this *KubernetesPatch_Replace_IntAsString) EqualVT(thatIface isKubernetesPatch_Replace_ValueOneof) bool {
	that, ok := thatIface.(*KubernetesPatch_Replace_IntAsString)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.IntAsString != that.IntAsString {
		return false
	}
	return true
}

func (this *KubernetesPatch) EqualVT(that *KubernetesPatch) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.PatchOneof == nil && that.PatchOneof != nil {
		return false
	} else if this.PatchOneof != nil {
		if that.PatchOneof == nil {
			return false
		}
		if !this.PatchOneof.(interface {
			EqualVT(isKubernetesPatch_PatchOneof) bool
		}).EqualVT(that.PatchOneof) {
			return false
		}
	}
	if !this.Target.EqualVT(that.Target) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *KubernetesPatch) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*KubernetesPatch)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *KubernetesPatch_Replace_) EqualVT(thatIface isKubernetesPatch_PatchOneof) bool {
	that, ok := thatIface.(*KubernetesPatch_Replace_)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Replace, that.Replace; p != q {
		if p == nil {
			p = &KubernetesPatch_Replace{}
		}
		if q == nil {
			q = &KubernetesPatch_Replace{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}
