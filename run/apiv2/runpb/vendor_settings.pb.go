// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: google/cloud/run/v2/vendor_settings.proto

package runpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Allowed ingress traffic for the Container.
type IngressTraffic int32

const (
	// Unspecified
	IngressTraffic_INGRESS_TRAFFIC_UNSPECIFIED IngressTraffic = 0
	// All inbound traffic is allowed.
	IngressTraffic_INGRESS_TRAFFIC_ALL IngressTraffic = 1
	// Only internal traffic is allowed.
	IngressTraffic_INGRESS_TRAFFIC_INTERNAL_ONLY IngressTraffic = 2
	// Both internal and Google Cloud Load Balancer traffic is allowed.
	IngressTraffic_INGRESS_TRAFFIC_INTERNAL_LOAD_BALANCER IngressTraffic = 3
)

// Enum value maps for IngressTraffic.
var (
	IngressTraffic_name = map[int32]string{
		0: "INGRESS_TRAFFIC_UNSPECIFIED",
		1: "INGRESS_TRAFFIC_ALL",
		2: "INGRESS_TRAFFIC_INTERNAL_ONLY",
		3: "INGRESS_TRAFFIC_INTERNAL_LOAD_BALANCER",
	}
	IngressTraffic_value = map[string]int32{
		"INGRESS_TRAFFIC_UNSPECIFIED":            0,
		"INGRESS_TRAFFIC_ALL":                    1,
		"INGRESS_TRAFFIC_INTERNAL_ONLY":          2,
		"INGRESS_TRAFFIC_INTERNAL_LOAD_BALANCER": 3,
	}
)

func (x IngressTraffic) Enum() *IngressTraffic {
	p := new(IngressTraffic)
	*p = x
	return p
}

func (x IngressTraffic) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IngressTraffic) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_run_v2_vendor_settings_proto_enumTypes[0].Descriptor()
}

func (IngressTraffic) Type() protoreflect.EnumType {
	return &file_google_cloud_run_v2_vendor_settings_proto_enumTypes[0]
}

func (x IngressTraffic) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IngressTraffic.Descriptor instead.
func (IngressTraffic) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_run_v2_vendor_settings_proto_rawDescGZIP(), []int{0}
}

// Alternatives for execution environments.
type ExecutionEnvironment int32

const (
	// Unspecified
	ExecutionEnvironment_EXECUTION_ENVIRONMENT_UNSPECIFIED ExecutionEnvironment = 0
	// Uses the First Generation environment.
	ExecutionEnvironment_EXECUTION_ENVIRONMENT_GEN1 ExecutionEnvironment = 1
	// Uses Second Generation environment.
	ExecutionEnvironment_EXECUTION_ENVIRONMENT_GEN2 ExecutionEnvironment = 2
)

// Enum value maps for ExecutionEnvironment.
var (
	ExecutionEnvironment_name = map[int32]string{
		0: "EXECUTION_ENVIRONMENT_UNSPECIFIED",
		1: "EXECUTION_ENVIRONMENT_GEN1",
		2: "EXECUTION_ENVIRONMENT_GEN2",
	}
	ExecutionEnvironment_value = map[string]int32{
		"EXECUTION_ENVIRONMENT_UNSPECIFIED": 0,
		"EXECUTION_ENVIRONMENT_GEN1":        1,
		"EXECUTION_ENVIRONMENT_GEN2":        2,
	}
)

func (x ExecutionEnvironment) Enum() *ExecutionEnvironment {
	p := new(ExecutionEnvironment)
	*p = x
	return p
}

func (x ExecutionEnvironment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExecutionEnvironment) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_run_v2_vendor_settings_proto_enumTypes[1].Descriptor()
}

func (ExecutionEnvironment) Type() protoreflect.EnumType {
	return &file_google_cloud_run_v2_vendor_settings_proto_enumTypes[1]
}

func (x ExecutionEnvironment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExecutionEnvironment.Descriptor instead.
func (ExecutionEnvironment) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_run_v2_vendor_settings_proto_rawDescGZIP(), []int{1}
}

// Egress options for VPC access.
type VpcAccess_VpcEgress int32

const (
	// Unspecified
	VpcAccess_VPC_EGRESS_UNSPECIFIED VpcAccess_VpcEgress = 0
	// All outbound traffic is routed through the VPC connector.
	VpcAccess_ALL_TRAFFIC VpcAccess_VpcEgress = 1
	// Only private IP ranges are routed through the VPC connector.
	VpcAccess_PRIVATE_RANGES_ONLY VpcAccess_VpcEgress = 2
)

// Enum value maps for VpcAccess_VpcEgress.
var (
	VpcAccess_VpcEgress_name = map[int32]string{
		0: "VPC_EGRESS_UNSPECIFIED",
		1: "ALL_TRAFFIC",
		2: "PRIVATE_RANGES_ONLY",
	}
	VpcAccess_VpcEgress_value = map[string]int32{
		"VPC_EGRESS_UNSPECIFIED": 0,
		"ALL_TRAFFIC":            1,
		"PRIVATE_RANGES_ONLY":    2,
	}
)

func (x VpcAccess_VpcEgress) Enum() *VpcAccess_VpcEgress {
	p := new(VpcAccess_VpcEgress)
	*p = x
	return p
}

func (x VpcAccess_VpcEgress) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VpcAccess_VpcEgress) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_run_v2_vendor_settings_proto_enumTypes[2].Descriptor()
}

func (VpcAccess_VpcEgress) Type() protoreflect.EnumType {
	return &file_google_cloud_run_v2_vendor_settings_proto_enumTypes[2]
}

func (x VpcAccess_VpcEgress) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VpcAccess_VpcEgress.Descriptor instead.
func (VpcAccess_VpcEgress) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_run_v2_vendor_settings_proto_rawDescGZIP(), []int{0, 0}
}

// VPC Access settings. For more information on creating a VPC Connector, visit
// https://cloud.google.com/vpc/docs/configure-serverless-vpc-access For
// information on how to configure Cloud Run with an existing VPC Connector,
// visit https://cloud.google.com/run/docs/configuring/connecting-vpc
type VpcAccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// VPC Access connector name.
	// Format: projects/{project}/locations/{location}/connectors/{connector},
	// where {project} can be project id or number.
	Connector string `protobuf:"bytes,1,opt,name=connector,proto3" json:"connector,omitempty"`
	// Traffic VPC egress settings.
	Egress VpcAccess_VpcEgress `protobuf:"varint,2,opt,name=egress,proto3,enum=google.cloud.run.v2.VpcAccess_VpcEgress" json:"egress,omitempty"`
}

func (x *VpcAccess) Reset() {
	*x = VpcAccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v2_vendor_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VpcAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VpcAccess) ProtoMessage() {}

func (x *VpcAccess) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v2_vendor_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VpcAccess.ProtoReflect.Descriptor instead.
func (*VpcAccess) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v2_vendor_settings_proto_rawDescGZIP(), []int{0}
}

func (x *VpcAccess) GetConnector() string {
	if x != nil {
		return x.Connector
	}
	return ""
}

func (x *VpcAccess) GetEgress() VpcAccess_VpcEgress {
	if x != nil {
		return x.Egress
	}
	return VpcAccess_VPC_EGRESS_UNSPECIFIED
}

// Settings for Binary Authorization feature.
type BinaryAuthorization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to BinauthzMethod:
	//
	//	*BinaryAuthorization_UseDefault
	BinauthzMethod isBinaryAuthorization_BinauthzMethod `protobuf_oneof:"binauthz_method"`
	// If present, indicates to use Breakglass using this justification.
	// If use_default is False, then it must be empty.
	// For more information on breakglass, see
	// https://cloud.google.com/binary-authorization/docs/using-breakglass
	BreakglassJustification string `protobuf:"bytes,2,opt,name=breakglass_justification,json=breakglassJustification,proto3" json:"breakglass_justification,omitempty"`
}

func (x *BinaryAuthorization) Reset() {
	*x = BinaryAuthorization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v2_vendor_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryAuthorization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryAuthorization) ProtoMessage() {}

func (x *BinaryAuthorization) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v2_vendor_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryAuthorization.ProtoReflect.Descriptor instead.
func (*BinaryAuthorization) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v2_vendor_settings_proto_rawDescGZIP(), []int{1}
}

func (m *BinaryAuthorization) GetBinauthzMethod() isBinaryAuthorization_BinauthzMethod {
	if m != nil {
		return m.BinauthzMethod
	}
	return nil
}

func (x *BinaryAuthorization) GetUseDefault() bool {
	if x, ok := x.GetBinauthzMethod().(*BinaryAuthorization_UseDefault); ok {
		return x.UseDefault
	}
	return false
}

func (x *BinaryAuthorization) GetBreakglassJustification() string {
	if x != nil {
		return x.BreakglassJustification
	}
	return ""
}

type isBinaryAuthorization_BinauthzMethod interface {
	isBinaryAuthorization_BinauthzMethod()
}

type BinaryAuthorization_UseDefault struct {
	// If True, indicates to use the default project's binary authorization
	// policy. If False, binary authorization will be disabled.
	UseDefault bool `protobuf:"varint,1,opt,name=use_default,json=useDefault,proto3,oneof"`
}

func (*BinaryAuthorization_UseDefault) isBinaryAuthorization_BinauthzMethod() {}

// Settings for revision-level scaling settings.
type RevisionScaling struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Minimum number of serving instances that this resource should have.
	MinInstanceCount int32 `protobuf:"varint,1,opt,name=min_instance_count,json=minInstanceCount,proto3" json:"min_instance_count,omitempty"`
	// Maximum number of serving instances that this resource should have.
	MaxInstanceCount int32 `protobuf:"varint,2,opt,name=max_instance_count,json=maxInstanceCount,proto3" json:"max_instance_count,omitempty"`
}

func (x *RevisionScaling) Reset() {
	*x = RevisionScaling{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v2_vendor_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevisionScaling) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevisionScaling) ProtoMessage() {}

func (x *RevisionScaling) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v2_vendor_settings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevisionScaling.ProtoReflect.Descriptor instead.
func (*RevisionScaling) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v2_vendor_settings_proto_rawDescGZIP(), []int{2}
}

func (x *RevisionScaling) GetMinInstanceCount() int32 {
	if x != nil {
		return x.MinInstanceCount
	}
	return 0
}

func (x *RevisionScaling) GetMaxInstanceCount() int32 {
	if x != nil {
		return x.MaxInstanceCount
	}
	return 0
}

var File_google_cloud_run_v2_vendor_settings_proto protoreflect.FileDescriptor

var file_google_cloud_run_v2_vendor_settings_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x75, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x09,
	0x56, 0x70, 0x63, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x45, 0x0a, 0x09, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xfa, 0x41,
	0x24, 0x0a, 0x22, 0x76, 0x70, 0x63, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x40, 0x0a, 0x06, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x70, 0x63, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2e, 0x56, 0x70, 0x63, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x06, 0x65, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x51, 0x0a, 0x09, 0x56, 0x70, 0x63, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1a, 0x0a, 0x16, 0x56, 0x50, 0x43, 0x5f, 0x45, 0x47, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x41,
	0x4c, 0x4c, 0x5f, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13,
	0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x53, 0x5f, 0x4f,
	0x4e, 0x4c, 0x59, 0x10, 0x02, 0x22, 0x86, 0x01, 0x0a, 0x13, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a,
	0x0b, 0x75, 0x73, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x00, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x12, 0x39, 0x0a, 0x18, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x67, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6a,
	0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x17, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x67, 0x6c, 0x61, 0x73, 0x73, 0x4a, 0x75,
	0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x11, 0x0a, 0x0f, 0x62,
	0x69, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x6d,
	0x0a, 0x0f, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x61, 0x6c, 0x69, 0x6e,
	0x67, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6d,
	0x69, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2c, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6d, 0x61, 0x78,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x99, 0x01,
	0x0a, 0x0e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x4e, 0x47, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x46,
	0x46, 0x49, 0x43, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x47, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x52, 0x41,
	0x46, 0x46, 0x49, 0x43, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x49, 0x4e,
	0x47, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x02, 0x12, 0x2a, 0x0a,
	0x26, 0x49, 0x4e, 0x47, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43,
	0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x42,
	0x41, 0x4c, 0x41, 0x4e, 0x43, 0x45, 0x52, 0x10, 0x03, 0x2a, 0x7d, 0x0a, 0x14, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x25, 0x0a, 0x21, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45,
	0x4e, 0x56, 0x49, 0x52, 0x4f, 0x4e, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x58, 0x45, 0x43,
	0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x4e, 0x56, 0x49, 0x52, 0x4f, 0x4e, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x47, 0x45, 0x4e, 0x31, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x58, 0x45, 0x43,
	0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x4e, 0x56, 0x49, 0x52, 0x4f, 0x4e, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x47, 0x45, 0x4e, 0x32, 0x10, 0x02, 0x42, 0x68, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e,
	0x2e, 0x76, 0x32, 0x42, 0x13, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x76, 0x32, 0x3b, 0x72,
	0x75, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_run_v2_vendor_settings_proto_rawDescOnce sync.Once
	file_google_cloud_run_v2_vendor_settings_proto_rawDescData = file_google_cloud_run_v2_vendor_settings_proto_rawDesc
)

func file_google_cloud_run_v2_vendor_settings_proto_rawDescGZIP() []byte {
	file_google_cloud_run_v2_vendor_settings_proto_rawDescOnce.Do(func() {
		file_google_cloud_run_v2_vendor_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_run_v2_vendor_settings_proto_rawDescData)
	})
	return file_google_cloud_run_v2_vendor_settings_proto_rawDescData
}

var file_google_cloud_run_v2_vendor_settings_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_google_cloud_run_v2_vendor_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_run_v2_vendor_settings_proto_goTypes = []interface{}{
	(IngressTraffic)(0),         // 0: google.cloud.run.v2.IngressTraffic
	(ExecutionEnvironment)(0),   // 1: google.cloud.run.v2.ExecutionEnvironment
	(VpcAccess_VpcEgress)(0),    // 2: google.cloud.run.v2.VpcAccess.VpcEgress
	(*VpcAccess)(nil),           // 3: google.cloud.run.v2.VpcAccess
	(*BinaryAuthorization)(nil), // 4: google.cloud.run.v2.BinaryAuthorization
	(*RevisionScaling)(nil),     // 5: google.cloud.run.v2.RevisionScaling
}
var file_google_cloud_run_v2_vendor_settings_proto_depIdxs = []int32{
	2, // 0: google.cloud.run.v2.VpcAccess.egress:type_name -> google.cloud.run.v2.VpcAccess.VpcEgress
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_run_v2_vendor_settings_proto_init() }
func file_google_cloud_run_v2_vendor_settings_proto_init() {
	if File_google_cloud_run_v2_vendor_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_run_v2_vendor_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VpcAccess); i {
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
		file_google_cloud_run_v2_vendor_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryAuthorization); i {
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
		file_google_cloud_run_v2_vendor_settings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevisionScaling); i {
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
	file_google_cloud_run_v2_vendor_settings_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*BinaryAuthorization_UseDefault)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_run_v2_vendor_settings_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_run_v2_vendor_settings_proto_goTypes,
		DependencyIndexes: file_google_cloud_run_v2_vendor_settings_proto_depIdxs,
		EnumInfos:         file_google_cloud_run_v2_vendor_settings_proto_enumTypes,
		MessageInfos:      file_google_cloud_run_v2_vendor_settings_proto_msgTypes,
	}.Build()
	File_google_cloud_run_v2_vendor_settings_proto = out.File
	file_google_cloud_run_v2_vendor_settings_proto_rawDesc = nil
	file_google_cloud_run_v2_vendor_settings_proto_goTypes = nil
	file_google_cloud_run_v2_vendor_settings_proto_depIdxs = nil
}
