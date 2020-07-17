// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: config.proto

package config

import (
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Policies []*Policy `protobuf:"bytes,2,rep,name=policies,proto3" json:"policies,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Config) GetPolicies() []*Policy {
	if x != nil {
		return x.Policies
	}
	return nil
}

type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                             string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	From                             string             `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                               string             `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	AllowedUsers                     []string           `protobuf:"bytes,4,rep,name=allowed_users,json=allowedUsers,proto3" json:"allowed_users,omitempty"`
	AllowedGroups                    []string           `protobuf:"bytes,5,rep,name=allowed_groups,json=allowedGroups,proto3" json:"allowed_groups,omitempty"`
	AllowedDomains                   []string           `protobuf:"bytes,6,rep,name=allowed_domains,json=allowedDomains,proto3" json:"allowed_domains,omitempty"`
	Prefix                           string             `protobuf:"bytes,7,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Path                             string             `protobuf:"bytes,8,opt,name=path,proto3" json:"path,omitempty"`
	Regex                            string             `protobuf:"bytes,9,opt,name=regex,proto3" json:"regex,omitempty"`
	CorsAllowPreflight               bool               `protobuf:"varint,10,opt,name=cors_allow_preflight,json=corsAllowPreflight,proto3" json:"cors_allow_preflight,omitempty"`
	AllowPublicUnauthenticatedAccess bool               `protobuf:"varint,11,opt,name=allow_public_unauthenticated_access,json=allowPublicUnauthenticatedAccess,proto3" json:"allow_public_unauthenticated_access,omitempty"`
	Timeout                          *duration.Duration `protobuf:"bytes,12,opt,name=timeout,proto3" json:"timeout,omitempty"`
	AllowWebsockets                  bool               `protobuf:"varint,13,opt,name=allow_websockets,json=allowWebsockets,proto3" json:"allow_websockets,omitempty"`
	TlsSkipVerify                    bool               `protobuf:"varint,14,opt,name=tls_skip_verify,json=tlsSkipVerify,proto3" json:"tls_skip_verify,omitempty"`
	TlsServerName                    string             `protobuf:"bytes,15,opt,name=tls_server_name,json=tlsServerName,proto3" json:"tls_server_name,omitempty"`
	TlsCustomCa                      string             `protobuf:"bytes,16,opt,name=tls_custom_ca,json=tlsCustomCa,proto3" json:"tls_custom_ca,omitempty"`
	TlsCustomCaFile                  string             `protobuf:"bytes,17,opt,name=tls_custom_ca_file,json=tlsCustomCaFile,proto3" json:"tls_custom_ca_file,omitempty"`
	TlsClientCert                    string             `protobuf:"bytes,18,opt,name=tls_client_cert,json=tlsClientCert,proto3" json:"tls_client_cert,omitempty"`
	TlsClientKey                     string             `protobuf:"bytes,19,opt,name=tls_client_key,json=tlsClientKey,proto3" json:"tls_client_key,omitempty"`
	TlsClientCertFile                string             `protobuf:"bytes,20,opt,name=tls_client_cert_file,json=tlsClientCertFile,proto3" json:"tls_client_cert_file,omitempty"`
	TlsClientKeyFile                 string             `protobuf:"bytes,21,opt,name=tls_client_key_file,json=tlsClientKeyFile,proto3" json:"tls_client_key_file,omitempty"`
	SetRequestHeaders                map[string]string  `protobuf:"bytes,22,rep,name=set_request_headers,json=setRequestHeaders,proto3" json:"set_request_headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RemoveRequestHeaders             []string           `protobuf:"bytes,23,rep,name=remove_request_headers,json=removeRequestHeaders,proto3" json:"remove_request_headers,omitempty"`
	PreserveHostHeader               bool               `protobuf:"varint,24,opt,name=preserve_host_header,json=preserveHostHeader,proto3" json:"preserve_host_header,omitempty"`
	PassIdentityHeaders              bool               `protobuf:"varint,25,opt,name=pass_identity_headers,json=passIdentityHeaders,proto3" json:"pass_identity_headers,omitempty"`
	KubernetesServiceAccountToken    string             `protobuf:"bytes,26,opt,name=kubernetes_service_account_token,json=kubernetesServiceAccountToken,proto3" json:"kubernetes_service_account_token,omitempty"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

func (x *Policy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Policy) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Policy) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Policy) GetAllowedUsers() []string {
	if x != nil {
		return x.AllowedUsers
	}
	return nil
}

func (x *Policy) GetAllowedGroups() []string {
	if x != nil {
		return x.AllowedGroups
	}
	return nil
}

func (x *Policy) GetAllowedDomains() []string {
	if x != nil {
		return x.AllowedDomains
	}
	return nil
}

func (x *Policy) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *Policy) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Policy) GetRegex() string {
	if x != nil {
		return x.Regex
	}
	return ""
}

func (x *Policy) GetCorsAllowPreflight() bool {
	if x != nil {
		return x.CorsAllowPreflight
	}
	return false
}

func (x *Policy) GetAllowPublicUnauthenticatedAccess() bool {
	if x != nil {
		return x.AllowPublicUnauthenticatedAccess
	}
	return false
}

func (x *Policy) GetTimeout() *duration.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Policy) GetAllowWebsockets() bool {
	if x != nil {
		return x.AllowWebsockets
	}
	return false
}

func (x *Policy) GetTlsSkipVerify() bool {
	if x != nil {
		return x.TlsSkipVerify
	}
	return false
}

func (x *Policy) GetTlsServerName() string {
	if x != nil {
		return x.TlsServerName
	}
	return ""
}

func (x *Policy) GetTlsCustomCa() string {
	if x != nil {
		return x.TlsCustomCa
	}
	return ""
}

func (x *Policy) GetTlsCustomCaFile() string {
	if x != nil {
		return x.TlsCustomCaFile
	}
	return ""
}

func (x *Policy) GetTlsClientCert() string {
	if x != nil {
		return x.TlsClientCert
	}
	return ""
}

func (x *Policy) GetTlsClientKey() string {
	if x != nil {
		return x.TlsClientKey
	}
	return ""
}

func (x *Policy) GetTlsClientCertFile() string {
	if x != nil {
		return x.TlsClientCertFile
	}
	return ""
}

func (x *Policy) GetTlsClientKeyFile() string {
	if x != nil {
		return x.TlsClientKeyFile
	}
	return ""
}

func (x *Policy) GetSetRequestHeaders() map[string]string {
	if x != nil {
		return x.SetRequestHeaders
	}
	return nil
}

func (x *Policy) GetRemoveRequestHeaders() []string {
	if x != nil {
		return x.RemoveRequestHeaders
	}
	return nil
}

func (x *Policy) GetPreserveHostHeader() bool {
	if x != nil {
		return x.PreserveHostHeader
	}
	return false
}

func (x *Policy) GetPassIdentityHeaders() bool {
	if x != nil {
		return x.PassIdentityHeaders
	}
	return false
}

func (x *Policy) GetKubernetesServiceAccountToken() string {
	if x != nil {
		return x.KubernetesServiceAccountToken
	}
	return ""
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x51, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a,
	0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x22, 0xb2, 0x09, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x12, 0x30, 0x0a, 0x14,
	0x63, 0x6f, 0x72, 0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x63, 0x6f, 0x72, 0x73,
	0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x65, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x4d,
	0x0a, 0x23, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x75,
	0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x20, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x77, 0x65, 0x62, 0x73,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x57, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x26, 0x0a,
	0x0f, 0x74, 0x6c, 0x73, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x74, 0x6c, 0x73, 0x53, 0x6b, 0x69, 0x70, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6c, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x74, 0x6c, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a,
	0x0d, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x63, 0x61, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6c, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43,
	0x61, 0x12, 0x2b, 0x0a, 0x12, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x63, 0x61, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74,
	0x6c, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x26,
	0x0a, 0x0f, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72,
	0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x6c, 0x73, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x74, 0x6c, 0x73, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x14,
	0x74, 0x6c, 0x73, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x6c, 0x73, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2d, 0x0a,
	0x13, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x6c, 0x73, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x5e, 0x0a, 0x13,
	0x73, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x16, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x6f, 0x6d, 0x65,
	0x72, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x73, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x34, 0x0a, 0x16,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x17, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x5f, 0x68,
	0x6f, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x18, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x12, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x15, 0x70, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x19, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x13, 0x70, 0x61, 0x73, 0x73, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x47, 0x0a, 0x20, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x1a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x1d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x1a, 0x44, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x70,
	0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_config_proto_goTypes = []interface{}{
	(*Config)(nil),            // 0: pomerium.config.Config
	(*Policy)(nil),            // 1: pomerium.config.Policy
	nil,                       // 2: pomerium.config.Policy.SetRequestHeadersEntry
	(*duration.Duration)(nil), // 3: google.protobuf.Duration
}
var file_config_proto_depIdxs = []int32{
	1, // 0: pomerium.config.Config.policies:type_name -> pomerium.config.Policy
	3, // 1: pomerium.config.Policy.timeout:type_name -> google.protobuf.Duration
	2, // 2: pomerium.config.Policy.set_request_headers:type_name -> pomerium.config.Policy.SetRequestHeadersEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}
