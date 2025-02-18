// Copyright 2019 Google LLC.
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
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: google/cloud/asset/v1p2beta1/assets.proto

package assetpb

import (
	reflect "reflect"
	sync "sync"

	v1 "google.golang.org/genproto/googleapis/iam/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Temporal asset. In addition to the asset, the temporal asset includes the
// status of the asset and valid from and to time of it.
type TemporalAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The time window when the asset data and state was observed.
	Window *TimeWindow `protobuf:"bytes,1,opt,name=window,proto3" json:"window,omitempty"`
	// If the asset is deleted or not.
	Deleted bool `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// Asset.
	Asset *Asset `protobuf:"bytes,3,opt,name=asset,proto3" json:"asset,omitempty"`
}

func (x *TemporalAsset) Reset() {
	*x = TemporalAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_asset_v1p2beta1_assets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemporalAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemporalAsset) ProtoMessage() {}

func (x *TemporalAsset) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_asset_v1p2beta1_assets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemporalAsset.ProtoReflect.Descriptor instead.
func (*TemporalAsset) Descriptor() ([]byte, []int) {
	return file_google_cloud_asset_v1p2beta1_assets_proto_rawDescGZIP(), []int{0}
}

func (x *TemporalAsset) GetWindow() *TimeWindow {
	if x != nil {
		return x.Window
	}
	return nil
}

func (x *TemporalAsset) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *TemporalAsset) GetAsset() *Asset {
	if x != nil {
		return x.Asset
	}
	return nil
}

// A time window of (start_time, end_time].
type TimeWindow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Start time of the time window (exclusive).
	StartTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// End time of the time window (inclusive).
	// Current timestamp if not specified.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *TimeWindow) Reset() {
	*x = TimeWindow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_asset_v1p2beta1_assets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeWindow) ProtoMessage() {}

func (x *TimeWindow) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_asset_v1p2beta1_assets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeWindow.ProtoReflect.Descriptor instead.
func (*TimeWindow) Descriptor() ([]byte, []int) {
	return file_google_cloud_asset_v1p2beta1_assets_proto_rawDescGZIP(), []int{1}
}

func (x *TimeWindow) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *TimeWindow) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

// Cloud asset. This includes all Google Cloud Platform resources,
// Cloud IAM policies, and other non-GCP assets.
type Asset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The full name of the asset. For example:
	// `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`.
	// See [Resource
	// Names](https://cloud.google.com/apis/design/resource_names#full_resource_name)
	// for more information.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Type of the asset. Example: "compute.googleapis.com/Disk".
	AssetType string `protobuf:"bytes,2,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	// Representation of the resource.
	Resource *Resource `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	// Representation of the actual Cloud IAM policy set on a cloud resource. For
	// each resource, there must be at most one Cloud IAM policy set on it.
	IamPolicy *v1.Policy `protobuf:"bytes,4,opt,name=iam_policy,json=iamPolicy,proto3" json:"iam_policy,omitempty"`
	// Asset's ancestry path in Cloud Resource Manager (CRM) hierarchy,
	// represented as a list of relative resource names. Ancestry path starts with
	// the closest CRM ancestor and ends at root. If the asset is a CRM
	// project/folder/organization, this starts from the asset itself.
	//
	// Example: ["projects/123456789", "folders/5432", "organizations/1234"]
	Ancestors []string `protobuf:"bytes,6,rep,name=ancestors,proto3" json:"ancestors,omitempty"`
}

func (x *Asset) Reset() {
	*x = Asset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_asset_v1p2beta1_assets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Asset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Asset) ProtoMessage() {}

func (x *Asset) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_asset_v1p2beta1_assets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Asset.ProtoReflect.Descriptor instead.
func (*Asset) Descriptor() ([]byte, []int) {
	return file_google_cloud_asset_v1p2beta1_assets_proto_rawDescGZIP(), []int{2}
}

func (x *Asset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Asset) GetAssetType() string {
	if x != nil {
		return x.AssetType
	}
	return ""
}

func (x *Asset) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *Asset) GetIamPolicy() *v1.Policy {
	if x != nil {
		return x.IamPolicy
	}
	return nil
}

func (x *Asset) GetAncestors() []string {
	if x != nil {
		return x.Ancestors
	}
	return nil
}

// Representation of a cloud resource.
type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The API version. Example: "v1".
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The URL of the discovery document containing the resource's JSON schema.
	// For example:
	// `"https://www.googleapis.com/discovery/v1/apis/compute/v1/rest"`.
	// It will be left unspecified for resources without a discovery-based API,
	// such as Cloud Bigtable.
	DiscoveryDocumentUri string `protobuf:"bytes,2,opt,name=discovery_document_uri,json=discoveryDocumentUri,proto3" json:"discovery_document_uri,omitempty"`
	// The JSON schema name listed in the discovery document.
	// Example: "Project". It will be left unspecified for resources (such as
	// Cloud Bigtable) without a discovery-based API.
	DiscoveryName string `protobuf:"bytes,3,opt,name=discovery_name,json=discoveryName,proto3" json:"discovery_name,omitempty"`
	// The REST URL for accessing the resource. An HTTP GET operation using this
	// URL returns the resource itself.
	// Example:
	// `https://cloudresourcemanager.googleapis.com/v1/projects/my-project-123`.
	// It will be left unspecified for resources without a REST API.
	ResourceUrl string `protobuf:"bytes,4,opt,name=resource_url,json=resourceUrl,proto3" json:"resource_url,omitempty"`
	// The full name of the immediate parent of this resource. See
	// [Resource
	// Names](https://cloud.google.com/apis/design/resource_names#full_resource_name)
	// for more information.
	//
	// For GCP assets, it is the parent resource defined in the [Cloud IAM policy
	// hierarchy](https://cloud.google.com/iam/docs/overview#policy_hierarchy).
	// For example:
	// `"//cloudresourcemanager.googleapis.com/projects/my_project_123"`.
	//
	// For third-party assets, it is up to the users to define.
	Parent string `protobuf:"bytes,5,opt,name=parent,proto3" json:"parent,omitempty"`
	// The content of the resource, in which some sensitive fields are scrubbed
	// away and may not be present.
	Data *structpb.Struct `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_asset_v1p2beta1_assets_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_asset_v1p2beta1_assets_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_google_cloud_asset_v1p2beta1_assets_proto_rawDescGZIP(), []int{3}
}

func (x *Resource) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Resource) GetDiscoveryDocumentUri() string {
	if x != nil {
		return x.DiscoveryDocumentUri
	}
	return ""
}

func (x *Resource) GetDiscoveryName() string {
	if x != nil {
		return x.DiscoveryName
	}
	return ""
}

func (x *Resource) GetResourceUrl() string {
	if x != nil {
		return x.ResourceUrl
	}
	return ""
}

func (x *Resource) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *Resource) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_google_cloud_asset_v1p2beta1_assets_proto protoreflect.FileDescriptor

var file_google_cloud_asset_v1p2beta1_assets_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x70, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x70, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x0d, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x40, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x70, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x52, 0x06, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x12, 0x39, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x70, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x22, 0x7e, 0x0a,
	0x0a, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x39, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xd2, 0x01,
	0x0a, 0x05, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x42, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x70, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x34,
	0x0a, 0x0a, 0x69, 0x61, 0x6d, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x09, 0x69, 0x61, 0x6d, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x73, 0x22, 0xe9, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x16, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x72, 0x69, 0x12,
	0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0xb4,
	0x01, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x70, 0x32, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x42, 0x0a, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x70, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x61,
	0x73, 0x73, 0x65, 0x74, 0xf8, 0x01, 0x01, 0xaa, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x70,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x5c, 0x56, 0x31, 0x70, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_asset_v1p2beta1_assets_proto_rawDescOnce sync.Once
	file_google_cloud_asset_v1p2beta1_assets_proto_rawDescData = file_google_cloud_asset_v1p2beta1_assets_proto_rawDesc
)

func file_google_cloud_asset_v1p2beta1_assets_proto_rawDescGZIP() []byte {
	file_google_cloud_asset_v1p2beta1_assets_proto_rawDescOnce.Do(func() {
		file_google_cloud_asset_v1p2beta1_assets_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_asset_v1p2beta1_assets_proto_rawDescData)
	})
	return file_google_cloud_asset_v1p2beta1_assets_proto_rawDescData
}

var file_google_cloud_asset_v1p2beta1_assets_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_asset_v1p2beta1_assets_proto_goTypes = []interface{}{
	(*TemporalAsset)(nil),         // 0: google.cloud.asset.v1p2beta1.TemporalAsset
	(*TimeWindow)(nil),            // 1: google.cloud.asset.v1p2beta1.TimeWindow
	(*Asset)(nil),                 // 2: google.cloud.asset.v1p2beta1.Asset
	(*Resource)(nil),              // 3: google.cloud.asset.v1p2beta1.Resource
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*v1.Policy)(nil),             // 5: google.iam.v1.Policy
	(*structpb.Struct)(nil),       // 6: google.protobuf.Struct
}
var file_google_cloud_asset_v1p2beta1_assets_proto_depIdxs = []int32{
	1, // 0: google.cloud.asset.v1p2beta1.TemporalAsset.window:type_name -> google.cloud.asset.v1p2beta1.TimeWindow
	2, // 1: google.cloud.asset.v1p2beta1.TemporalAsset.asset:type_name -> google.cloud.asset.v1p2beta1.Asset
	4, // 2: google.cloud.asset.v1p2beta1.TimeWindow.start_time:type_name -> google.protobuf.Timestamp
	4, // 3: google.cloud.asset.v1p2beta1.TimeWindow.end_time:type_name -> google.protobuf.Timestamp
	3, // 4: google.cloud.asset.v1p2beta1.Asset.resource:type_name -> google.cloud.asset.v1p2beta1.Resource
	5, // 5: google.cloud.asset.v1p2beta1.Asset.iam_policy:type_name -> google.iam.v1.Policy
	6, // 6: google.cloud.asset.v1p2beta1.Resource.data:type_name -> google.protobuf.Struct
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_cloud_asset_v1p2beta1_assets_proto_init() }
func file_google_cloud_asset_v1p2beta1_assets_proto_init() {
	if File_google_cloud_asset_v1p2beta1_assets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_asset_v1p2beta1_assets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemporalAsset); i {
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
		file_google_cloud_asset_v1p2beta1_assets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeWindow); i {
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
		file_google_cloud_asset_v1p2beta1_assets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Asset); i {
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
		file_google_cloud_asset_v1p2beta1_assets_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
			RawDescriptor: file_google_cloud_asset_v1p2beta1_assets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_asset_v1p2beta1_assets_proto_goTypes,
		DependencyIndexes: file_google_cloud_asset_v1p2beta1_assets_proto_depIdxs,
		MessageInfos:      file_google_cloud_asset_v1p2beta1_assets_proto_msgTypes,
	}.Build()
	File_google_cloud_asset_v1p2beta1_assets_proto = out.File
	file_google_cloud_asset_v1p2beta1_assets_proto_rawDesc = nil
	file_google_cloud_asset_v1p2beta1_assets_proto_goTypes = nil
	file_google_cloud_asset_v1p2beta1_assets_proto_depIdxs = nil
}
