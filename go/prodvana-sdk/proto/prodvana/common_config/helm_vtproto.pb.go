// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/common_config/helm.proto

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

func (m *RemoteHelmChart) CloneVT() *RemoteHelmChart {
	if m == nil {
		return (*RemoteHelmChart)(nil)
	}
	r := new(RemoteHelmChart)
	r.Chart = m.Chart
	r.ChartVersion = m.ChartVersion
	if m.RepoOneof != nil {
		r.RepoOneof = m.RepoOneof.(interface {
			CloneVT() isRemoteHelmChart_RepoOneof
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *RemoteHelmChart) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *RemoteHelmChart_Repo) CloneVT() isRemoteHelmChart_RepoOneof {
	if m == nil {
		return (*RemoteHelmChart_Repo)(nil)
	}
	r := new(RemoteHelmChart_Repo)
	r.Repo = m.Repo
	return r
}

func (m *HelmValuesOverrides) CloneVT() *HelmValuesOverrides {
	if m == nil {
		return (*HelmValuesOverrides)(nil)
	}
	r := new(HelmValuesOverrides)
	if m.OverrideOneof != nil {
		r.OverrideOneof = m.OverrideOneof.(interface {
			CloneVT() isHelmValuesOverrides_OverrideOneof
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *HelmValuesOverrides) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *HelmValuesOverrides_Inlined) CloneVT() isHelmValuesOverrides_OverrideOneof {
	if m == nil {
		return (*HelmValuesOverrides_Inlined)(nil)
	}
	r := new(HelmValuesOverrides_Inlined)
	r.Inlined = m.Inlined
	return r
}

func (m *HelmValuesOverrides_Local) CloneVT() isHelmValuesOverrides_OverrideOneof {
	if m == nil {
		return (*HelmValuesOverrides_Local)(nil)
	}
	r := new(HelmValuesOverrides_Local)
	r.Local = m.Local.CloneVT()
	return r
}

func (m *HelmValuesOverrides_Remote) CloneVT() isHelmValuesOverrides_OverrideOneof {
	if m == nil {
		return (*HelmValuesOverrides_Remote)(nil)
	}
	r := new(HelmValuesOverrides_Remote)
	r.Remote = m.Remote.CloneVT()
	return r
}

func (m *HelmConfig) CloneVT() *HelmConfig {
	if m == nil {
		return (*HelmConfig)(nil)
	}
	r := new(HelmConfig)
	r.ReleaseName = m.ReleaseName
	r.Namespace = m.Namespace
	r.FixupOwnership = m.FixupOwnership
	if m.ChartOneof != nil {
		r.ChartOneof = m.ChartOneof.(interface {
			CloneVT() isHelmConfig_ChartOneof
		}).CloneVT()
	}
	if rhs := m.ValuesOverrides; rhs != nil {
		tmpContainer := make([]*HelmValuesOverrides, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.ValuesOverrides = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *HelmConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *HelmConfig_Remote) CloneVT() isHelmConfig_ChartOneof {
	if m == nil {
		return (*HelmConfig_Remote)(nil)
	}
	r := new(HelmConfig_Remote)
	r.Remote = m.Remote.CloneVT()
	return r
}

func (m *HelmConfig_Local) CloneVT() isHelmConfig_ChartOneof {
	if m == nil {
		return (*HelmConfig_Local)(nil)
	}
	r := new(HelmConfig_Local)
	r.Local = m.Local.CloneVT()
	return r
}

func (m *HelmConfig_HelmTarballBlobId) CloneVT() isHelmConfig_ChartOneof {
	if m == nil {
		return (*HelmConfig_HelmTarballBlobId)(nil)
	}
	r := new(HelmConfig_HelmTarballBlobId)
	r.HelmTarballBlobId = m.HelmTarballBlobId
	return r
}

func (this *RemoteHelmChart) StableEqualVT(that *RemoteHelmChart) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.RepoOneof == nil && that.RepoOneof != nil {
		return false
	} else if this.RepoOneof != nil {
		if that.RepoOneof == nil {
			return false
		}
		if !this.RepoOneof.(interface {
			StableEqualVT(isRemoteHelmChart_RepoOneof) bool
		}).StableEqualVT(that.RepoOneof) {
			return false
		}
	}
	if this.Chart != that.Chart {
		return false
	}
	if this.ChartVersion != that.ChartVersion {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *RemoteHelmChart) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*RemoteHelmChart)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *RemoteHelmChart_Repo) StableEqualVT(thatIface isRemoteHelmChart_RepoOneof) bool {
	that, ok := thatIface.(*RemoteHelmChart_Repo)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.Repo != that.Repo {
		return false
	}
	return true
}

func (this *RemoteHelmChart) EqualVT(that *RemoteHelmChart) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.RepoOneof == nil && that.RepoOneof != nil {
		return false
	} else if this.RepoOneof != nil {
		if that.RepoOneof == nil {
			return false
		}
		if !this.RepoOneof.(interface {
			EqualVT(isRemoteHelmChart_RepoOneof) bool
		}).EqualVT(that.RepoOneof) {
			return false
		}
	}
	if this.Chart != that.Chart {
		return false
	}
	if this.ChartVersion != that.ChartVersion {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *RemoteHelmChart) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*RemoteHelmChart)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *RemoteHelmChart_Repo) EqualVT(thatIface isRemoteHelmChart_RepoOneof) bool {
	that, ok := thatIface.(*RemoteHelmChart_Repo)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.Repo != that.Repo {
		return false
	}
	return true
}

func (this *HelmValuesOverrides) StableEqualVT(that *HelmValuesOverrides) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.OverrideOneof == nil && that.OverrideOneof != nil {
		return false
	} else if this.OverrideOneof != nil {
		if that.OverrideOneof == nil {
			return false
		}
		if !this.OverrideOneof.(interface {
			StableEqualVT(isHelmValuesOverrides_OverrideOneof) bool
		}).StableEqualVT(that.OverrideOneof) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *HelmValuesOverrides) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*HelmValuesOverrides)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *HelmValuesOverrides_Inlined) StableEqualVT(thatIface isHelmValuesOverrides_OverrideOneof) bool {
	that, ok := thatIface.(*HelmValuesOverrides_Inlined)
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

func (this *HelmValuesOverrides_Local) StableEqualVT(thatIface isHelmValuesOverrides_OverrideOneof) bool {
	that, ok := thatIface.(*HelmValuesOverrides_Local)
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

func (this *HelmValuesOverrides_Remote) StableEqualVT(thatIface isHelmValuesOverrides_OverrideOneof) bool {
	that, ok := thatIface.(*HelmValuesOverrides_Remote)
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

func (this *HelmValuesOverrides) EqualVT(that *HelmValuesOverrides) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.OverrideOneof == nil && that.OverrideOneof != nil {
		return false
	} else if this.OverrideOneof != nil {
		if that.OverrideOneof == nil {
			return false
		}
		if !this.OverrideOneof.(interface {
			EqualVT(isHelmValuesOverrides_OverrideOneof) bool
		}).EqualVT(that.OverrideOneof) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *HelmValuesOverrides) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*HelmValuesOverrides)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *HelmValuesOverrides_Inlined) EqualVT(thatIface isHelmValuesOverrides_OverrideOneof) bool {
	that, ok := thatIface.(*HelmValuesOverrides_Inlined)
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

func (this *HelmValuesOverrides_Local) EqualVT(thatIface isHelmValuesOverrides_OverrideOneof) bool {
	that, ok := thatIface.(*HelmValuesOverrides_Local)
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

func (this *HelmValuesOverrides_Remote) EqualVT(thatIface isHelmValuesOverrides_OverrideOneof) bool {
	that, ok := thatIface.(*HelmValuesOverrides_Remote)
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

func (this *HelmConfig) StableEqualVT(that *HelmConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ChartOneof == nil && that.ChartOneof != nil {
		return false
	} else if this.ChartOneof != nil {
		if that.ChartOneof == nil {
			return false
		}
		if !this.ChartOneof.(interface {
			StableEqualVT(isHelmConfig_ChartOneof) bool
		}).StableEqualVT(that.ChartOneof) {
			return false
		}
	}
	if len(this.ValuesOverrides) != len(that.ValuesOverrides) {
		return false
	}
	for i, vx := range this.ValuesOverrides {
		vy := that.ValuesOverrides[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &HelmValuesOverrides{}
			}
			if q == nil {
				q = &HelmValuesOverrides{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if this.ReleaseName != that.ReleaseName {
		return false
	}
	if this.Namespace != that.Namespace {
		return false
	}
	if this.FixupOwnership != that.FixupOwnership {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *HelmConfig) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*HelmConfig)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *HelmConfig_Remote) StableEqualVT(thatIface isHelmConfig_ChartOneof) bool {
	that, ok := thatIface.(*HelmConfig_Remote)
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
			p = &RemoteHelmChart{}
		}
		if q == nil {
			q = &RemoteHelmChart{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *HelmConfig_Local) StableEqualVT(thatIface isHelmConfig_ChartOneof) bool {
	that, ok := thatIface.(*HelmConfig_Local)
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

func (this *HelmConfig_HelmTarballBlobId) StableEqualVT(thatIface isHelmConfig_ChartOneof) bool {
	that, ok := thatIface.(*HelmConfig_HelmTarballBlobId)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.HelmTarballBlobId != that.HelmTarballBlobId {
		return false
	}
	return true
}

func (this *HelmConfig) EqualVT(that *HelmConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ChartOneof == nil && that.ChartOneof != nil {
		return false
	} else if this.ChartOneof != nil {
		if that.ChartOneof == nil {
			return false
		}
		if !this.ChartOneof.(interface {
			EqualVT(isHelmConfig_ChartOneof) bool
		}).EqualVT(that.ChartOneof) {
			return false
		}
	}
	if len(this.ValuesOverrides) != len(that.ValuesOverrides) {
		return false
	}
	for i, vx := range this.ValuesOverrides {
		vy := that.ValuesOverrides[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &HelmValuesOverrides{}
			}
			if q == nil {
				q = &HelmValuesOverrides{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if this.ReleaseName != that.ReleaseName {
		return false
	}
	if this.Namespace != that.Namespace {
		return false
	}
	if this.FixupOwnership != that.FixupOwnership {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *HelmConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*HelmConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *HelmConfig_Remote) EqualVT(thatIface isHelmConfig_ChartOneof) bool {
	that, ok := thatIface.(*HelmConfig_Remote)
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
			p = &RemoteHelmChart{}
		}
		if q == nil {
			q = &RemoteHelmChart{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *HelmConfig_Local) EqualVT(thatIface isHelmConfig_ChartOneof) bool {
	that, ok := thatIface.(*HelmConfig_Local)
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

func (this *HelmConfig_HelmTarballBlobId) EqualVT(thatIface isHelmConfig_ChartOneof) bool {
	that, ok := thatIface.(*HelmConfig_HelmTarballBlobId)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.HelmTarballBlobId != that.HelmTarballBlobId {
		return false
	}
	return true
}
