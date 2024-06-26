// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/config_file/config.proto

package config_file

import (
	application "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/application"
	delivery_extension "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/delivery_extension"
	environment "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/environment"
	protection "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/protection"
	recipe "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/recipe"
	release_channel "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/release_channel"
	service "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *ProdvanaConfig) CloneVT() *ProdvanaConfig {
	if m == nil {
		return (*ProdvanaConfig)(nil)
	}
	r := new(ProdvanaConfig)
	if m.ConfigOneof != nil {
		r.ConfigOneof = m.ConfigOneof.(interface {
			CloneVT() isProdvanaConfig_ConfigOneof
		}).CloneVT()
	}
	if m.MetadataOneof != nil {
		r.MetadataOneof = m.MetadataOneof.(interface {
			CloneVT() isProdvanaConfig_MetadataOneof
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ProdvanaConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ProdvanaConfig_Application) CloneVT() isProdvanaConfig_ConfigOneof {
	if m == nil {
		return (*ProdvanaConfig_Application)(nil)
	}
	r := new(ProdvanaConfig_Application)
	r.Application = m.Application.CloneVT()
	return r
}

func (m *ProdvanaConfig_Service) CloneVT() isProdvanaConfig_ConfigOneof {
	if m == nil {
		return (*ProdvanaConfig_Service)(nil)
	}
	r := new(ProdvanaConfig_Service)
	r.Service = m.Service.CloneVT()
	return r
}

func (m *ProdvanaConfig_Protection) CloneVT() isProdvanaConfig_ConfigOneof {
	if m == nil {
		return (*ProdvanaConfig_Protection)(nil)
	}
	r := new(ProdvanaConfig_Protection)
	r.Protection = m.Protection.CloneVT()
	return r
}

func (m *ProdvanaConfig_Runtime) CloneVT() isProdvanaConfig_ConfigOneof {
	if m == nil {
		return (*ProdvanaConfig_Runtime)(nil)
	}
	r := new(ProdvanaConfig_Runtime)
	r.Runtime = m.Runtime.CloneVT()
	return r
}

func (m *ProdvanaConfig_DeliveryExtension) CloneVT() isProdvanaConfig_ConfigOneof {
	if m == nil {
		return (*ProdvanaConfig_DeliveryExtension)(nil)
	}
	r := new(ProdvanaConfig_DeliveryExtension)
	r.DeliveryExtension = m.DeliveryExtension.CloneVT()
	return r
}

func (m *ProdvanaConfig_ReleaseChannel) CloneVT() isProdvanaConfig_ConfigOneof {
	if m == nil {
		return (*ProdvanaConfig_ReleaseChannel)(nil)
	}
	r := new(ProdvanaConfig_ReleaseChannel)
	r.ReleaseChannel = m.ReleaseChannel.CloneVT()
	return r
}

func (m *ProdvanaConfig_Recipe) CloneVT() isProdvanaConfig_ConfigOneof {
	if m == nil {
		return (*ProdvanaConfig_Recipe)(nil)
	}
	r := new(ProdvanaConfig_Recipe)
	r.Recipe = m.Recipe.CloneVT()
	return r
}

func (m *ProdvanaConfig_ApplicationMetadata) CloneVT() isProdvanaConfig_MetadataOneof {
	if m == nil {
		return (*ProdvanaConfig_ApplicationMetadata)(nil)
	}
	r := new(ProdvanaConfig_ApplicationMetadata)
	r.ApplicationMetadata = m.ApplicationMetadata.CloneVT()
	return r
}

func (m *ProdvanaConfig_ServiceMetadata) CloneVT() isProdvanaConfig_MetadataOneof {
	if m == nil {
		return (*ProdvanaConfig_ServiceMetadata)(nil)
	}
	r := new(ProdvanaConfig_ServiceMetadata)
	r.ServiceMetadata = m.ServiceMetadata.CloneVT()
	return r
}

func (this *ProdvanaConfig) StableEqualVT(that *ProdvanaConfig) bool {
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
			StableEqualVT(isProdvanaConfig_ConfigOneof) bool
		}).StableEqualVT(that.ConfigOneof) {
			return false
		}
	}
	if this.MetadataOneof == nil && that.MetadataOneof != nil {
		return false
	} else if this.MetadataOneof != nil {
		if that.MetadataOneof == nil {
			return false
		}
		if !this.MetadataOneof.(interface {
			StableEqualVT(isProdvanaConfig_MetadataOneof) bool
		}).StableEqualVT(that.MetadataOneof) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ProdvanaConfig) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ProdvanaConfig)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ProdvanaConfig_Application) StableEqualVT(thatIface isProdvanaConfig_ConfigOneof) bool {
	that, ok := thatIface.(*ProdvanaConfig_Application)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Application, that.Application; p != q {
		if p == nil {
			p = &application.ApplicationConfig{}
		}
		if q == nil {
			q = &application.ApplicationConfig{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProdvanaConfig_Service) StableEqualVT(thatIface isProdvanaConfig_ConfigOneof) bool {
	that, ok := thatIface.(*ProdvanaConfig_Service)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Service, that.Service; p != q {
		if p == nil {
			p = &service.ServiceConfig{}
		}
		if q == nil {
			q = &service.ServiceConfig{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProdvanaConfig_Protection) StableEqualVT(thatIface isProdvanaConfig_ConfigOneof) bool {
	that, ok := thatIface.(*ProdvanaConfig_Protection)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Protection, that.Protection; p != q {
		if p == nil {
			p = &protection.ProtectionConfig{}
		}
		if q == nil {
			q = &protection.ProtectionConfig{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProdvanaConfig_Runtime) StableEqualVT(thatIface isProdvanaConfig_ConfigOneof) bool {
	that, ok := thatIface.(*ProdvanaConfig_Runtime)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Runtime, that.Runtime; p != q {
		if p == nil {
			p = &environment.ClusterConfig{}
		}
		if q == nil {
			q = &environment.ClusterConfig{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProdvanaConfig_ApplicationMetadata) StableEqualVT(thatIface isProdvanaConfig_MetadataOneof) bool {
	that, ok := thatIface.(*ProdvanaConfig_ApplicationMetadata)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.ApplicationMetadata, that.ApplicationMetadata; p != q {
		if p == nil {
			p = &application.ApplicationUserMetadata{}
		}
		if q == nil {
			q = &application.ApplicationUserMetadata{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProdvanaConfig_ServiceMetadata) StableEqualVT(thatIface isProdvanaConfig_MetadataOneof) bool {
	that, ok := thatIface.(*ProdvanaConfig_ServiceMetadata)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.ServiceMetadata, that.ServiceMetadata; p != q {
		if p == nil {
			p = &service.ServiceUserMetadata{}
		}
		if q == nil {
			q = &service.ServiceUserMetadata{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProdvanaConfig_DeliveryExtension) StableEqualVT(thatIface isProdvanaConfig_ConfigOneof) bool {
	that, ok := thatIface.(*ProdvanaConfig_DeliveryExtension)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.DeliveryExtension, that.DeliveryExtension; p != q {
		if p == nil {
			p = &delivery_extension.DeliveryExtensionConfig{}
		}
		if q == nil {
			q = &delivery_extension.DeliveryExtensionConfig{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProdvanaConfig_ReleaseChannel) StableEqualVT(thatIface isProdvanaConfig_ConfigOneof) bool {
	that, ok := thatIface.(*ProdvanaConfig_ReleaseChannel)
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
			p = &release_channel.ReleaseChannelConfig{}
		}
		if q == nil {
			q = &release_channel.ReleaseChannelConfig{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProdvanaConfig_Recipe) StableEqualVT(thatIface isProdvanaConfig_ConfigOneof) bool {
	that, ok := thatIface.(*ProdvanaConfig_Recipe)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Recipe, that.Recipe; p != q {
		if p == nil {
			p = &recipe.RecipeConfig{}
		}
		if q == nil {
			q = &recipe.RecipeConfig{}
		}
		if !p.StableEqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProdvanaConfig) EqualVT(that *ProdvanaConfig) bool {
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
			EqualVT(isProdvanaConfig_ConfigOneof) bool
		}).EqualVT(that.ConfigOneof) {
			return false
		}
	}
	if this.MetadataOneof == nil && that.MetadataOneof != nil {
		return false
	} else if this.MetadataOneof != nil {
		if that.MetadataOneof == nil {
			return false
		}
		if !this.MetadataOneof.(interface {
			EqualVT(isProdvanaConfig_MetadataOneof) bool
		}).EqualVT(that.MetadataOneof) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ProdvanaConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ProdvanaConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ProdvanaConfig_Application) EqualVT(thatIface isProdvanaConfig_ConfigOneof) bool {
	that, ok := thatIface.(*ProdvanaConfig_Application)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Application, that.Application; p != q {
		if p == nil {
			p = &application.ApplicationConfig{}
		}
		if q == nil {
			q = &application.ApplicationConfig{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProdvanaConfig_Service) EqualVT(thatIface isProdvanaConfig_ConfigOneof) bool {
	that, ok := thatIface.(*ProdvanaConfig_Service)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Service, that.Service; p != q {
		if p == nil {
			p = &service.ServiceConfig{}
		}
		if q == nil {
			q = &service.ServiceConfig{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProdvanaConfig_Protection) EqualVT(thatIface isProdvanaConfig_ConfigOneof) bool {
	that, ok := thatIface.(*ProdvanaConfig_Protection)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Protection, that.Protection; p != q {
		if p == nil {
			p = &protection.ProtectionConfig{}
		}
		if q == nil {
			q = &protection.ProtectionConfig{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProdvanaConfig_Runtime) EqualVT(thatIface isProdvanaConfig_ConfigOneof) bool {
	that, ok := thatIface.(*ProdvanaConfig_Runtime)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Runtime, that.Runtime; p != q {
		if p == nil {
			p = &environment.ClusterConfig{}
		}
		if q == nil {
			q = &environment.ClusterConfig{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProdvanaConfig_ApplicationMetadata) EqualVT(thatIface isProdvanaConfig_MetadataOneof) bool {
	that, ok := thatIface.(*ProdvanaConfig_ApplicationMetadata)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.ApplicationMetadata, that.ApplicationMetadata; p != q {
		if p == nil {
			p = &application.ApplicationUserMetadata{}
		}
		if q == nil {
			q = &application.ApplicationUserMetadata{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProdvanaConfig_ServiceMetadata) EqualVT(thatIface isProdvanaConfig_MetadataOneof) bool {
	that, ok := thatIface.(*ProdvanaConfig_ServiceMetadata)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.ServiceMetadata, that.ServiceMetadata; p != q {
		if p == nil {
			p = &service.ServiceUserMetadata{}
		}
		if q == nil {
			q = &service.ServiceUserMetadata{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProdvanaConfig_DeliveryExtension) EqualVT(thatIface isProdvanaConfig_ConfigOneof) bool {
	that, ok := thatIface.(*ProdvanaConfig_DeliveryExtension)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.DeliveryExtension, that.DeliveryExtension; p != q {
		if p == nil {
			p = &delivery_extension.DeliveryExtensionConfig{}
		}
		if q == nil {
			q = &delivery_extension.DeliveryExtensionConfig{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProdvanaConfig_ReleaseChannel) EqualVT(thatIface isProdvanaConfig_ConfigOneof) bool {
	that, ok := thatIface.(*ProdvanaConfig_ReleaseChannel)
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
			p = &release_channel.ReleaseChannelConfig{}
		}
		if q == nil {
			q = &release_channel.ReleaseChannelConfig{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *ProdvanaConfig_Recipe) EqualVT(thatIface isProdvanaConfig_ConfigOneof) bool {
	that, ok := thatIface.(*ProdvanaConfig_Recipe)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Recipe, that.Recipe; p != q {
		if p == nil {
			p = &recipe.RecipeConfig{}
		}
		if q == nil {
			q = &recipe.RecipeConfig{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}
