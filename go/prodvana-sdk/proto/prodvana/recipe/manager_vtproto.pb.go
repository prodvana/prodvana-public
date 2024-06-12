// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/recipe/manager.proto

package recipe

import (
	common_config "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *ConfigureRecipeReq) CloneVT() *ConfigureRecipeReq {
	if m == nil {
		return (*ConfigureRecipeReq)(nil)
	}
	r := new(ConfigureRecipeReq)
	r.Source = m.Source
	r.SourceMetadata = m.SourceMetadata.CloneVT()
	r.RecipeConfig = m.RecipeConfig.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ConfigureRecipeReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ConfigureRecipeResp) CloneVT() *ConfigureRecipeResp {
	if m == nil {
		return (*ConfigureRecipeResp)(nil)
	}
	r := new(ConfigureRecipeResp)
	r.RecipeId = m.RecipeId
	r.Version = m.Version
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ConfigureRecipeResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ValidateConfigureRecipeResp) CloneVT() *ValidateConfigureRecipeResp {
	if m == nil {
		return (*ValidateConfigureRecipeResp)(nil)
	}
	r := new(ValidateConfigureRecipeResp)
	r.CompiledConfig = m.CompiledConfig.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ValidateConfigureRecipeResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ListRecipesReq_ByService) CloneVT() *ListRecipesReq_ByService {
	if m == nil {
		return (*ListRecipesReq_ByService)(nil)
	}
	r := new(ListRecipesReq_ByService)
	r.Application = m.Application
	r.Service = m.Service
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ListRecipesReq_ByService) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ListRecipesReq) CloneVT() *ListRecipesReq {
	if m == nil {
		return (*ListRecipesReq)(nil)
	}
	r := new(ListRecipesReq)
	r.PageToken = m.PageToken
	r.PageSize = m.PageSize
	if m.Filter != nil {
		r.Filter = m.Filter.(interface {
			CloneVT() isListRecipesReq_Filter
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ListRecipesReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ListRecipesReq_ServiceFilter) CloneVT() isListRecipesReq_Filter {
	if m == nil {
		return (*ListRecipesReq_ServiceFilter)(nil)
	}
	r := new(ListRecipesReq_ServiceFilter)
	r.ServiceFilter = m.ServiceFilter.CloneVT()
	return r
}

func (m *ListRecipesResp) CloneVT() *ListRecipesResp {
	if m == nil {
		return (*ListRecipesResp)(nil)
	}
	r := new(ListRecipesResp)
	r.NextPageToken = m.NextPageToken
	if rhs := m.Recipes; rhs != nil {
		tmpContainer := make([]*Recipe, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Recipes = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ListRecipesResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetRecipeReq) CloneVT() *GetRecipeReq {
	if m == nil {
		return (*GetRecipeReq)(nil)
	}
	r := new(GetRecipeReq)
	r.Recipe = m.Recipe
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetRecipeReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetRecipeResp) CloneVT() *GetRecipeResp {
	if m == nil {
		return (*GetRecipeResp)(nil)
	}
	r := new(GetRecipeResp)
	r.Recipe = m.Recipe.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetRecipeResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetRecipeConfigReq) CloneVT() *GetRecipeConfigReq {
	if m == nil {
		return (*GetRecipeConfigReq)(nil)
	}
	r := new(GetRecipeConfigReq)
	r.Recipe = m.Recipe
	r.Version = m.Version
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetRecipeConfigReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetRecipeConfigResp) CloneVT() *GetRecipeConfigResp {
	if m == nil {
		return (*GetRecipeConfigResp)(nil)
	}
	r := new(GetRecipeConfigResp)
	r.InputConfig = m.InputConfig.CloneVT()
	r.CompiledConfig = m.CompiledConfig.CloneVT()
	r.Version = m.Version
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetRecipeConfigResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ApplyRecipeParametersReq) CloneVT() *ApplyRecipeParametersReq {
	if m == nil {
		return (*ApplyRecipeParametersReq)(nil)
	}
	r := new(ApplyRecipeParametersReq)
	r.Source = m.Source
	r.SourceMetadata = m.SourceMetadata.CloneVT()
	r.Recipe = m.Recipe
	if rhs := m.Parameters; rhs != nil {
		tmpContainer := make([]*common_config.ParameterValue, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Parameters = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ApplyRecipeParametersReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ApplyRecipeParametersResp_ServiceVersion) CloneVT() *ApplyRecipeParametersResp_ServiceVersion {
	if m == nil {
		return (*ApplyRecipeParametersResp_ServiceVersion)(nil)
	}
	r := new(ApplyRecipeParametersResp_ServiceVersion)
	r.Service = m.Service
	r.ServiceId = m.ServiceId
	r.ServiceVersion = m.ServiceVersion
	r.Application = m.Application
	r.ApplicationId = m.ApplicationId
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ApplyRecipeParametersResp_ServiceVersion) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ApplyRecipeParametersResp) CloneVT() *ApplyRecipeParametersResp {
	if m == nil {
		return (*ApplyRecipeParametersResp)(nil)
	}
	r := new(ApplyRecipeParametersResp)
	if rhs := m.Versions; rhs != nil {
		tmpContainer := make([]*ApplyRecipeParametersResp_ServiceVersion, len(rhs))
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

func (m *ApplyRecipeParametersResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *Recipe) CloneVT() *Recipe {
	if m == nil {
		return (*Recipe)(nil)
	}
	r := new(Recipe)
	r.Meta = m.Meta.CloneVT()
	r.Config = m.Config.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *Recipe) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *ConfigureRecipeReq) StableEqualVT(that *ConfigureRecipeReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Source != that.Source {
		return false
	}
	if !this.SourceMetadata.StableEqualVT(that.SourceMetadata) {
		return false
	}
	if !this.RecipeConfig.StableEqualVT(that.RecipeConfig) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ConfigureRecipeReq) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ConfigureRecipeReq)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ConfigureRecipeReq) EqualVT(that *ConfigureRecipeReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Source != that.Source {
		return false
	}
	if !this.SourceMetadata.EqualVT(that.SourceMetadata) {
		return false
	}
	if !this.RecipeConfig.EqualVT(that.RecipeConfig) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ConfigureRecipeReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ConfigureRecipeReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ConfigureRecipeResp) StableEqualVT(that *ConfigureRecipeResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.RecipeId != that.RecipeId {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ConfigureRecipeResp) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ConfigureRecipeResp)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ConfigureRecipeResp) EqualVT(that *ConfigureRecipeResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.RecipeId != that.RecipeId {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ConfigureRecipeResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ConfigureRecipeResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ValidateConfigureRecipeResp) StableEqualVT(that *ValidateConfigureRecipeResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.CompiledConfig.StableEqualVT(that.CompiledConfig) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ValidateConfigureRecipeResp) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ValidateConfigureRecipeResp)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ValidateConfigureRecipeResp) EqualVT(that *ValidateConfigureRecipeResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.CompiledConfig.EqualVT(that.CompiledConfig) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ValidateConfigureRecipeResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ValidateConfigureRecipeResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ListRecipesReq_ByService) StableEqualVT(that *ListRecipesReq_ByService) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Application != that.Application {
		return false
	}
	if this.Service != that.Service {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListRecipesReq_ByService) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListRecipesReq_ByService)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ListRecipesReq) StableEqualVT(that *ListRecipesReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Filter == nil && that.Filter != nil {
		return false
	} else if this.Filter != nil {
		if that.Filter == nil {
			return false
		}
		if !this.Filter.(interface {
			StableEqualVT(isListRecipesReq_Filter) bool
		}).StableEqualVT(that.Filter) {
			return false
		}
	}
	if this.PageToken != that.PageToken {
		return false
	}
	if this.PageSize != that.PageSize {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListRecipesReq) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListRecipesReq)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ListRecipesReq_ServiceFilter) StableEqualVT(thatIface isListRecipesReq_Filter) bool {
	that, ok := thatIface.(*ListRecipesReq_ServiceFilter)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.ServiceFilter, that.ServiceFilter; p != q {
		if p == nil {
			p = &ListRecipesReq_ByService{}
		}
		if q == nil {
			q = &ListRecipesReq_ByService{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ListRecipesReq_ByService) EqualVT(that *ListRecipesReq_ByService) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Application != that.Application {
		return false
	}
	if this.Service != that.Service {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListRecipesReq_ByService) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListRecipesReq_ByService)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ListRecipesReq) EqualVT(that *ListRecipesReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Filter == nil && that.Filter != nil {
		return false
	} else if this.Filter != nil {
		if that.Filter == nil {
			return false
		}
		if !this.Filter.(interface {
			EqualVT(isListRecipesReq_Filter) bool
		}).EqualVT(that.Filter) {
			return false
		}
	}
	if this.PageToken != that.PageToken {
		return false
	}
	if this.PageSize != that.PageSize {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListRecipesReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListRecipesReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ListRecipesReq_ServiceFilter) EqualVT(thatIface isListRecipesReq_Filter) bool {
	that, ok := thatIface.(*ListRecipesReq_ServiceFilter)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.ServiceFilter, that.ServiceFilter; p != q {
		if p == nil {
			p = &ListRecipesReq_ByService{}
		}
		if q == nil {
			q = &ListRecipesReq_ByService{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ListRecipesResp) StableEqualVT(that *ListRecipesResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Recipes) != len(that.Recipes) {
		return false
	}
	for i, vx := range this.Recipes {
		vy := that.Recipes[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Recipe{}
			}
			if q == nil {
				q = &Recipe{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if this.NextPageToken != that.NextPageToken {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListRecipesResp) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListRecipesResp)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ListRecipesResp) EqualVT(that *ListRecipesResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Recipes) != len(that.Recipes) {
		return false
	}
	for i, vx := range this.Recipes {
		vy := that.Recipes[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Recipe{}
			}
			if q == nil {
				q = &Recipe{}
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

func (this *ListRecipesResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListRecipesResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetRecipeReq) StableEqualVT(that *GetRecipeReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Recipe != that.Recipe {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetRecipeReq) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetRecipeReq)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *GetRecipeReq) EqualVT(that *GetRecipeReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Recipe != that.Recipe {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetRecipeReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetRecipeReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetRecipeResp) StableEqualVT(that *GetRecipeResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Recipe.StableEqualVT(that.Recipe) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetRecipeResp) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetRecipeResp)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *GetRecipeResp) EqualVT(that *GetRecipeResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Recipe.EqualVT(that.Recipe) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetRecipeResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetRecipeResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetRecipeConfigReq) StableEqualVT(that *GetRecipeConfigReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Recipe != that.Recipe {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetRecipeConfigReq) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetRecipeConfigReq)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *GetRecipeConfigReq) EqualVT(that *GetRecipeConfigReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Recipe != that.Recipe {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetRecipeConfigReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetRecipeConfigReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetRecipeConfigResp) StableEqualVT(that *GetRecipeConfigResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.InputConfig.StableEqualVT(that.InputConfig) {
		return false
	}
	if !this.CompiledConfig.StableEqualVT(that.CompiledConfig) {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetRecipeConfigResp) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetRecipeConfigResp)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *GetRecipeConfigResp) EqualVT(that *GetRecipeConfigResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.InputConfig.EqualVT(that.InputConfig) {
		return false
	}
	if !this.CompiledConfig.EqualVT(that.CompiledConfig) {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetRecipeConfigResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetRecipeConfigResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ApplyRecipeParametersReq) StableEqualVT(that *ApplyRecipeParametersReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Source != that.Source {
		return false
	}
	if !this.SourceMetadata.StableEqualVT(that.SourceMetadata) {
		return false
	}
	if this.Recipe != that.Recipe {
		return false
	}
	if len(this.Parameters) != len(that.Parameters) {
		return false
	}
	for i, vx := range this.Parameters {
		vy := that.Parameters[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &common_config.ParameterValue{}
			}
			if q == nil {
				q = &common_config.ParameterValue{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ApplyRecipeParametersReq) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ApplyRecipeParametersReq)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ApplyRecipeParametersReq) EqualVT(that *ApplyRecipeParametersReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Source != that.Source {
		return false
	}
	if !this.SourceMetadata.EqualVT(that.SourceMetadata) {
		return false
	}
	if this.Recipe != that.Recipe {
		return false
	}
	if len(this.Parameters) != len(that.Parameters) {
		return false
	}
	for i, vx := range this.Parameters {
		vy := that.Parameters[i]
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

func (this *ApplyRecipeParametersReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ApplyRecipeParametersReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ApplyRecipeParametersResp_ServiceVersion) StableEqualVT(that *ApplyRecipeParametersResp_ServiceVersion) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Service != that.Service {
		return false
	}
	if this.ServiceId != that.ServiceId {
		return false
	}
	if this.ServiceVersion != that.ServiceVersion {
		return false
	}
	if this.Application != that.Application {
		return false
	}
	if this.ApplicationId != that.ApplicationId {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ApplyRecipeParametersResp_ServiceVersion) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ApplyRecipeParametersResp_ServiceVersion)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ApplyRecipeParametersResp) StableEqualVT(that *ApplyRecipeParametersResp) bool {
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
				p = &ApplyRecipeParametersResp_ServiceVersion{}
			}
			if q == nil {
				q = &ApplyRecipeParametersResp_ServiceVersion{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ApplyRecipeParametersResp) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ApplyRecipeParametersResp)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ApplyRecipeParametersResp_ServiceVersion) EqualVT(that *ApplyRecipeParametersResp_ServiceVersion) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Service != that.Service {
		return false
	}
	if this.ServiceId != that.ServiceId {
		return false
	}
	if this.ServiceVersion != that.ServiceVersion {
		return false
	}
	if this.Application != that.Application {
		return false
	}
	if this.ApplicationId != that.ApplicationId {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ApplyRecipeParametersResp_ServiceVersion) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ApplyRecipeParametersResp_ServiceVersion)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ApplyRecipeParametersResp) EqualVT(that *ApplyRecipeParametersResp) bool {
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
				p = &ApplyRecipeParametersResp_ServiceVersion{}
			}
			if q == nil {
				q = &ApplyRecipeParametersResp_ServiceVersion{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ApplyRecipeParametersResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ApplyRecipeParametersResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Recipe) StableEqualVT(that *Recipe) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Meta.StableEqualVT(that.Meta) {
		return false
	}
	if !this.Config.StableEqualVT(that.Config) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Recipe) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Recipe)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *Recipe) EqualVT(that *Recipe) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Meta.EqualVT(that.Meta) {
		return false
	}
	if !this.Config.EqualVT(that.Config) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Recipe) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Recipe)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
