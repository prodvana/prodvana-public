// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/application/application_manager.proto

package application

import (
	durationpb1 "github.com/planetscale/vtprotobuf/types/known/durationpb"
	timestamppb1 "github.com/planetscale/vtprotobuf/types/known/timestamppb"
	common_config "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	insights "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/insights"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *ListApplicationsReq) CloneVT() *ListApplicationsReq {
	if m == nil {
		return (*ListApplicationsReq)(nil)
	}
	r := new(ListApplicationsReq)
	r.Detailed = m.Detailed
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ListApplicationsReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ListApplicationsResp) CloneVT() *ListApplicationsResp {
	if m == nil {
		return (*ListApplicationsResp)(nil)
	}
	r := new(ListApplicationsResp)
	if rhs := m.Applications; rhs != nil {
		tmpContainer := make([]*Application, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Applications = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ListApplicationsResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetApplicationReq) CloneVT() *GetApplicationReq {
	if m == nil {
		return (*GetApplicationReq)(nil)
	}
	r := new(GetApplicationReq)
	r.Application = m.Application
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetApplicationReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetApplicationResp) CloneVT() *GetApplicationResp {
	if m == nil {
		return (*GetApplicationResp)(nil)
	}
	r := new(GetApplicationResp)
	r.Application = m.Application.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetApplicationResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ConfigureApplicationReq) CloneVT() *ConfigureApplicationReq {
	if m == nil {
		return (*ConfigureApplicationReq)(nil)
	}
	r := new(ConfigureApplicationReq)
	r.ApplicationConfig = m.ApplicationConfig.CloneVT()
	r.Source = m.Source
	r.SourceMetadata = m.SourceMetadata.CloneVT()
	r.BaseVersion = m.BaseVersion
	r.ForceCreateNewVersion = m.ForceCreateNewVersion
	if rhs := m.ApprovedDangerousActionIds; rhs != nil {
		tmpContainer := make([]string, len(rhs))
		copy(tmpContainer, rhs)
		r.ApprovedDangerousActionIds = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ConfigureApplicationReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ConfigureApplicationResp) CloneVT() *ConfigureApplicationResp {
	if m == nil {
		return (*ConfigureApplicationResp)(nil)
	}
	r := new(ConfigureApplicationResp)
	r.Meta = m.Meta.CloneVT()
	r.CreatedNewVersion = m.CreatedNewVersion
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ConfigureApplicationResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ValidateConfigureApplicationResp) CloneVT() *ValidateConfigureApplicationResp {
	if m == nil {
		return (*ValidateConfigureApplicationResp)(nil)
	}
	r := new(ValidateConfigureApplicationResp)
	r.Config = m.Config.CloneVT()
	r.CompiledConfig = m.CompiledConfig.CloneVT()
	if rhs := m.DangerousActions; rhs != nil {
		tmpContainer := make([]*common_config.DangerousAction, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.DangerousActions = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ValidateConfigureApplicationResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetApplicationConfigReq) CloneVT() *GetApplicationConfigReq {
	if m == nil {
		return (*GetApplicationConfigReq)(nil)
	}
	r := new(GetApplicationConfigReq)
	r.Application = m.Application
	r.Version = m.Version
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetApplicationConfigReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetApplicationConfigResp) CloneVT() *GetApplicationConfigResp {
	if m == nil {
		return (*GetApplicationConfigResp)(nil)
	}
	r := new(GetApplicationConfigResp)
	r.Config = m.Config.CloneVT()
	r.Version = m.Version
	r.CompiledConfig = m.CompiledConfig.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetApplicationConfigResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ListApplicationConfigVersionsReq) CloneVT() *ListApplicationConfigVersionsReq {
	if m == nil {
		return (*ListApplicationConfigVersionsReq)(nil)
	}
	r := new(ListApplicationConfigVersionsReq)
	r.Application = m.Application
	r.PageToken = m.PageToken
	r.PageSize = m.PageSize
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ListApplicationConfigVersionsReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ListApplicationConfigVersionsResp_VersionMetadata) CloneVT() *ListApplicationConfigVersionsResp_VersionMetadata {
	if m == nil {
		return (*ListApplicationConfigVersionsResp_VersionMetadata)(nil)
	}
	r := new(ListApplicationConfigVersionsResp_VersionMetadata)
	r.Version = m.Version
	r.CreationTimestamp = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.CreationTimestamp).CloneVT())
	r.Source = m.Source
	r.SourceMetadata = m.SourceMetadata.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ListApplicationConfigVersionsResp_VersionMetadata) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ListApplicationConfigVersionsResp) CloneVT() *ListApplicationConfigVersionsResp {
	if m == nil {
		return (*ListApplicationConfigVersionsResp)(nil)
	}
	r := new(ListApplicationConfigVersionsResp)
	r.NextPageToken = m.NextPageToken
	if rhs := m.Versions; rhs != nil {
		tmpContainer := make([]*ListApplicationConfigVersionsResp_VersionMetadata, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Versions = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ListApplicationConfigVersionsResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *DeleteApplicationReq) CloneVT() *DeleteApplicationReq {
	if m == nil {
		return (*DeleteApplicationReq)(nil)
	}
	r := new(DeleteApplicationReq)
	r.Application = m.Application
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *DeleteApplicationReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *DeleteApplicationResp) CloneVT() *DeleteApplicationResp {
	if m == nil {
		return (*DeleteApplicationResp)(nil)
	}
	r := new(DeleteApplicationResp)
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *DeleteApplicationResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetApplicationMetricsReq) CloneVT() *GetApplicationMetricsReq {
	if m == nil {
		return (*GetApplicationMetricsReq)(nil)
	}
	r := new(GetApplicationMetricsReq)
	r.Application = m.Application
	r.StartTimestamp = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.StartTimestamp).CloneVT())
	r.EndTimestamp = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.EndTimestamp).CloneVT())
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetApplicationMetricsReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetApplicationMetricsResp) CloneVT() *GetApplicationMetricsResp {
	if m == nil {
		return (*GetApplicationMetricsResp)(nil)
	}
	r := new(GetApplicationMetricsResp)
	r.DeploymentMetrics = m.DeploymentMetrics.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetApplicationMetricsResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetApplicationInsightsReq) CloneVT() *GetApplicationInsightsReq {
	if m == nil {
		return (*GetApplicationInsightsReq)(nil)
	}
	r := new(GetApplicationInsightsReq)
	r.Application = m.Application
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetApplicationInsightsReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetApplicationInsightsResp) CloneVT() *GetApplicationInsightsResp {
	if m == nil {
		return (*GetApplicationInsightsResp)(nil)
	}
	r := new(GetApplicationInsightsResp)
	if rhs := m.Insights; rhs != nil {
		tmpContainer := make([]*insights.Insight, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Insights = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetApplicationInsightsResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *SnoozeApplicationInsightReq) CloneVT() *SnoozeApplicationInsightReq {
	if m == nil {
		return (*SnoozeApplicationInsightReq)(nil)
	}
	r := new(SnoozeApplicationInsightReq)
	r.Application = m.Application
	r.Class = m.Class
	r.Duration = (*durationpb.Duration)((*durationpb1.Duration)(m.Duration).CloneVT())
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *SnoozeApplicationInsightReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *SnoozeApplicationInsightResp) CloneVT() *SnoozeApplicationInsightResp {
	if m == nil {
		return (*SnoozeApplicationInsightResp)(nil)
	}
	r := new(SnoozeApplicationInsightResp)
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *SnoozeApplicationInsightResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetApplicationMetadataReq) CloneVT() *GetApplicationMetadataReq {
	if m == nil {
		return (*GetApplicationMetadataReq)(nil)
	}
	r := new(GetApplicationMetadataReq)
	r.Application = m.Application
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetApplicationMetadataReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetApplicationMetadataResp) CloneVT() *GetApplicationMetadataResp {
	if m == nil {
		return (*GetApplicationMetadataResp)(nil)
	}
	r := new(GetApplicationMetadataResp)
	r.Metadata = m.Metadata.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetApplicationMetadataResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *SetApplicationMetadataReq) CloneVT() *SetApplicationMetadataReq {
	if m == nil {
		return (*SetApplicationMetadataReq)(nil)
	}
	r := new(SetApplicationMetadataReq)
	r.Application = m.Application
	r.Metadata = m.Metadata.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *SetApplicationMetadataReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *SetApplicationMetadataResp) CloneVT() *SetApplicationMetadataResp {
	if m == nil {
		return (*SetApplicationMetadataResp)(nil)
	}
	r := new(SetApplicationMetadataResp)
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *SetApplicationMetadataResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *ListApplicationsReq) EqualVT(that *ListApplicationsReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Detailed != that.Detailed {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListApplicationsReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListApplicationsReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ListApplicationsResp) EqualVT(that *ListApplicationsResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Applications) != len(that.Applications) {
		return false
	}
	for i, vx := range this.Applications {
		vy := that.Applications[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Application{}
			}
			if q == nil {
				q = &Application{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListApplicationsResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListApplicationsResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetApplicationReq) EqualVT(that *GetApplicationReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Application != that.Application {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetApplicationReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetApplicationReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetApplicationResp) EqualVT(that *GetApplicationResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Application.EqualVT(that.Application) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetApplicationResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetApplicationResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ConfigureApplicationReq) EqualVT(that *ConfigureApplicationReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.ApplicationConfig.EqualVT(that.ApplicationConfig) {
		return false
	}
	if len(this.ApprovedDangerousActionIds) != len(that.ApprovedDangerousActionIds) {
		return false
	}
	for i, vx := range this.ApprovedDangerousActionIds {
		vy := that.ApprovedDangerousActionIds[i]
		if vx != vy {
			return false
		}
	}
	if this.Source != that.Source {
		return false
	}
	if !this.SourceMetadata.EqualVT(that.SourceMetadata) {
		return false
	}
	if this.BaseVersion != that.BaseVersion {
		return false
	}
	if this.ForceCreateNewVersion != that.ForceCreateNewVersion {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ConfigureApplicationReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ConfigureApplicationReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ConfigureApplicationResp) EqualVT(that *ConfigureApplicationResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Meta.EqualVT(that.Meta) {
		return false
	}
	if this.CreatedNewVersion != that.CreatedNewVersion {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ConfigureApplicationResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ConfigureApplicationResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ValidateConfigureApplicationResp) EqualVT(that *ValidateConfigureApplicationResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Config.EqualVT(that.Config) {
		return false
	}
	if !this.CompiledConfig.EqualVT(that.CompiledConfig) {
		return false
	}
	if len(this.DangerousActions) != len(that.DangerousActions) {
		return false
	}
	for i, vx := range this.DangerousActions {
		vy := that.DangerousActions[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &common_config.DangerousAction{}
			}
			if q == nil {
				q = &common_config.DangerousAction{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ValidateConfigureApplicationResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ValidateConfigureApplicationResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetApplicationConfigReq) EqualVT(that *GetApplicationConfigReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Application != that.Application {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetApplicationConfigReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetApplicationConfigReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetApplicationConfigResp) EqualVT(that *GetApplicationConfigResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Config.EqualVT(that.Config) {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	if !this.CompiledConfig.EqualVT(that.CompiledConfig) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetApplicationConfigResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetApplicationConfigResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ListApplicationConfigVersionsReq) EqualVT(that *ListApplicationConfigVersionsReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Application != that.Application {
		return false
	}
	if this.PageToken != that.PageToken {
		return false
	}
	if this.PageSize != that.PageSize {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListApplicationConfigVersionsReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListApplicationConfigVersionsReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ListApplicationConfigVersionsResp_VersionMetadata) EqualVT(that *ListApplicationConfigVersionsResp_VersionMetadata) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.CreationTimestamp).EqualVT((*timestamppb1.Timestamp)(that.CreationTimestamp)) {
		return false
	}
	if this.Source != that.Source {
		return false
	}
	if !this.SourceMetadata.EqualVT(that.SourceMetadata) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListApplicationConfigVersionsResp_VersionMetadata) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListApplicationConfigVersionsResp_VersionMetadata)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ListApplicationConfigVersionsResp) EqualVT(that *ListApplicationConfigVersionsResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Versions) != len(that.Versions) {
		return false
	}
	for i, vx := range this.Versions {
		vy := that.Versions[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &ListApplicationConfigVersionsResp_VersionMetadata{}
			}
			if q == nil {
				q = &ListApplicationConfigVersionsResp_VersionMetadata{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if this.NextPageToken != that.NextPageToken {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListApplicationConfigVersionsResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListApplicationConfigVersionsResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *DeleteApplicationReq) EqualVT(that *DeleteApplicationReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Application != that.Application {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DeleteApplicationReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DeleteApplicationReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *DeleteApplicationResp) EqualVT(that *DeleteApplicationResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DeleteApplicationResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DeleteApplicationResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetApplicationMetricsReq) EqualVT(that *GetApplicationMetricsReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Application != that.Application {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.StartTimestamp).EqualVT((*timestamppb1.Timestamp)(that.StartTimestamp)) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.EndTimestamp).EqualVT((*timestamppb1.Timestamp)(that.EndTimestamp)) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetApplicationMetricsReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetApplicationMetricsReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetApplicationMetricsResp) EqualVT(that *GetApplicationMetricsResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.DeploymentMetrics.EqualVT(that.DeploymentMetrics) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetApplicationMetricsResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetApplicationMetricsResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetApplicationInsightsReq) EqualVT(that *GetApplicationInsightsReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Application != that.Application {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetApplicationInsightsReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetApplicationInsightsReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetApplicationInsightsResp) EqualVT(that *GetApplicationInsightsResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Insights) != len(that.Insights) {
		return false
	}
	for i, vx := range this.Insights {
		vy := that.Insights[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &insights.Insight{}
			}
			if q == nil {
				q = &insights.Insight{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetApplicationInsightsResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetApplicationInsightsResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *SnoozeApplicationInsightReq) EqualVT(that *SnoozeApplicationInsightReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Application != that.Application {
		return false
	}
	if this.Class != that.Class {
		return false
	}
	if !(*durationpb1.Duration)(this.Duration).EqualVT((*durationpb1.Duration)(that.Duration)) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *SnoozeApplicationInsightReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*SnoozeApplicationInsightReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *SnoozeApplicationInsightResp) EqualVT(that *SnoozeApplicationInsightResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *SnoozeApplicationInsightResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*SnoozeApplicationInsightResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetApplicationMetadataReq) EqualVT(that *GetApplicationMetadataReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Application != that.Application {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetApplicationMetadataReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetApplicationMetadataReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetApplicationMetadataResp) EqualVT(that *GetApplicationMetadataResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Metadata.EqualVT(that.Metadata) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetApplicationMetadataResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetApplicationMetadataResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *SetApplicationMetadataReq) EqualVT(that *SetApplicationMetadataReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Application != that.Application {
		return false
	}
	if !this.Metadata.EqualVT(that.Metadata) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *SetApplicationMetadataReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*SetApplicationMetadataReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *SetApplicationMetadataResp) EqualVT(that *SetApplicationMetadataResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *SetApplicationMetadataResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*SetApplicationMetadataResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
