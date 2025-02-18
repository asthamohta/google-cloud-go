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
// source: google/cloud/run/v2/execution_template.proto

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

// ExecutionTemplate describes the data an execution should have when created
// from a template.
type ExecutionTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// KRM-style labels for the resource.
	Labels map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// KRM-style annotations for the resource.
	Annotations map[string]string `protobuf:"bytes,2,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Specifies the maximum desired number of tasks the execution should run at
	// given time. Must be <= task_count.
	// When the job is run, if this field is 0 or unset, the maximum possible
	// value will be used for that execution.
	// The actual number of tasks running in steady state will be less than this
	// number when there are fewer tasks waiting to be completed remaining,
	// i.e. when the work left to do is less than max parallelism.
	Parallelism int32 `protobuf:"varint,3,opt,name=parallelism,proto3" json:"parallelism,omitempty"`
	// Specifies the desired number of tasks the execution should run.
	// Setting to 1 means that parallelism is limited to 1 and the success of
	// that task signals the success of the execution.
	// More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	TaskCount int32 `protobuf:"varint,4,opt,name=task_count,json=taskCount,proto3" json:"task_count,omitempty"`
	// Required. Describes the task(s) that will be created when executing an execution.
	Template *TaskTemplate `protobuf:"bytes,5,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *ExecutionTemplate) Reset() {
	*x = ExecutionTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v2_execution_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionTemplate) ProtoMessage() {}

func (x *ExecutionTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v2_execution_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionTemplate.ProtoReflect.Descriptor instead.
func (*ExecutionTemplate) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v2_execution_template_proto_rawDescGZIP(), []int{0}
}

func (x *ExecutionTemplate) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ExecutionTemplate) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *ExecutionTemplate) GetParallelism() int32 {
	if x != nil {
		return x.Parallelism
	}
	return 0
}

func (x *ExecutionTemplate) GetTaskCount() int32 {
	if x != nil {
		return x.TaskCount
	}
	return 0
}

func (x *ExecutionTemplate) GetTemplate() *TaskTemplate {
	if x != nil {
		return x.Template
	}
	return nil
}

var File_google_cloud_run_v2_execution_template_proto protoreflect.FileDescriptor

var file_google_cloud_run_v2_execution_template_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x75, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e,
	0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x03,
	0x0a, 0x11, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x59, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61,
	0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x69, 0x73, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x69, 0x73, 0x6d, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x08, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e,
	0x2e, 0x76, 0x32, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a,
	0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x6b, 0x0a, 0x17, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72,
	0x75, 0x6e, 0x2e, 0x76, 0x32, 0x42, 0x16, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x75, 0x6e,
	0x2f, 0x76, 0x32, 0x3b, 0x72, 0x75, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_run_v2_execution_template_proto_rawDescOnce sync.Once
	file_google_cloud_run_v2_execution_template_proto_rawDescData = file_google_cloud_run_v2_execution_template_proto_rawDesc
)

func file_google_cloud_run_v2_execution_template_proto_rawDescGZIP() []byte {
	file_google_cloud_run_v2_execution_template_proto_rawDescOnce.Do(func() {
		file_google_cloud_run_v2_execution_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_run_v2_execution_template_proto_rawDescData)
	})
	return file_google_cloud_run_v2_execution_template_proto_rawDescData
}

var file_google_cloud_run_v2_execution_template_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_run_v2_execution_template_proto_goTypes = []interface{}{
	(*ExecutionTemplate)(nil), // 0: google.cloud.run.v2.ExecutionTemplate
	nil,                       // 1: google.cloud.run.v2.ExecutionTemplate.LabelsEntry
	nil,                       // 2: google.cloud.run.v2.ExecutionTemplate.AnnotationsEntry
	(*TaskTemplate)(nil),      // 3: google.cloud.run.v2.TaskTemplate
}
var file_google_cloud_run_v2_execution_template_proto_depIdxs = []int32{
	1, // 0: google.cloud.run.v2.ExecutionTemplate.labels:type_name -> google.cloud.run.v2.ExecutionTemplate.LabelsEntry
	2, // 1: google.cloud.run.v2.ExecutionTemplate.annotations:type_name -> google.cloud.run.v2.ExecutionTemplate.AnnotationsEntry
	3, // 2: google.cloud.run.v2.ExecutionTemplate.template:type_name -> google.cloud.run.v2.TaskTemplate
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_run_v2_execution_template_proto_init() }
func file_google_cloud_run_v2_execution_template_proto_init() {
	if File_google_cloud_run_v2_execution_template_proto != nil {
		return
	}
	file_google_cloud_run_v2_task_template_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_run_v2_execution_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionTemplate); i {
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
			RawDescriptor: file_google_cloud_run_v2_execution_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_run_v2_execution_template_proto_goTypes,
		DependencyIndexes: file_google_cloud_run_v2_execution_template_proto_depIdxs,
		MessageInfos:      file_google_cloud_run_v2_execution_template_proto_msgTypes,
	}.Build()
	File_google_cloud_run_v2_execution_template_proto = out.File
	file_google_cloud_run_v2_execution_template_proto_rawDesc = nil
	file_google_cloud_run_v2_execution_template_proto_goTypes = nil
	file_google_cloud_run_v2_execution_template_proto_depIdxs = nil
}
