// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: internal/impl/kube.proto

package impl

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BabysitterConfig contains configuration passed to a babysitter.
type BabysitterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace    string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`                                                                                          // Kubernetes namespace
	DeploymentId string            `protobuf:"bytes,2,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`                                                                // globally unique deployment id
	Listeners    map[string]int32  `protobuf:"bytes,3,rep,name=listeners,proto3" json:"listeners,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // a map from listener name to port
	Groups       map[string]string `protobuf:"bytes,4,rep,name=groups,proto3" json:"groups,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`        // a map from component name to group name
	Telemetry    *Telemetry        `protobuf:"bytes,5,opt,name=telemetry,proto3" json:"telemetry,omitempty"`                                                                                          // options to control how the telemetry is being manipulated.
}

func (x *BabysitterConfig) Reset() {
	*x = BabysitterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_impl_kube_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BabysitterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BabysitterConfig) ProtoMessage() {}

func (x *BabysitterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_internal_impl_kube_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BabysitterConfig.ProtoReflect.Descriptor instead.
func (*BabysitterConfig) Descriptor() ([]byte, []int) {
	return file_internal_impl_kube_proto_rawDescGZIP(), []int{0}
}

func (x *BabysitterConfig) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *BabysitterConfig) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *BabysitterConfig) GetListeners() map[string]int32 {
	if x != nil {
		return x.Listeners
	}
	return nil
}

func (x *BabysitterConfig) GetGroups() map[string]string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *BabysitterConfig) GetTelemetry() *Telemetry {
	if x != nil {
		return x.Telemetry
	}
	return nil
}

// Options to control how the telemetry is being manipulated.
type Telemetry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics *MetricOptions `protobuf:"bytes,1,opt,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *Telemetry) Reset() {
	*x = Telemetry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_impl_kube_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Telemetry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Telemetry) ProtoMessage() {}

func (x *Telemetry) ProtoReflect() protoreflect.Message {
	mi := &file_internal_impl_kube_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Telemetry.ProtoReflect.Descriptor instead.
func (*Telemetry) Descriptor() ([]byte, []int) {
	return file_internal_impl_kube_proto_rawDescGZIP(), []int{1}
}

func (x *Telemetry) GetMetrics() *MetricOptions {
	if x != nil {
		return x.Metrics
	}
	return nil
}

// Options to configure how the metrics are handled by the deployer.
type MetricOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, the deployer also exports metrics created implicitly by the
	// framework; e.g., autogenerated metrics, http metrics.
	AutoGenerateMetrics bool `protobuf:"varint,1,opt,name=auto_generate_metrics,json=autoGenerateMetrics,proto3" json:"auto_generate_metrics,omitempty"`
	// How often to export the metrics to Google Cloud.
	ExportInterval *durationpb.Duration `protobuf:"bytes,2,opt,name=export_interval,json=exportInterval,proto3" json:"export_interval,omitempty"`
}

func (x *MetricOptions) Reset() {
	*x = MetricOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_impl_kube_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricOptions) ProtoMessage() {}

func (x *MetricOptions) ProtoReflect() protoreflect.Message {
	mi := &file_internal_impl_kube_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricOptions.ProtoReflect.Descriptor instead.
func (*MetricOptions) Descriptor() ([]byte, []int) {
	return file_internal_impl_kube_proto_rawDescGZIP(), []int{2}
}

func (x *MetricOptions) GetAutoGenerateMetrics() bool {
	if x != nil {
		return x.AutoGenerateMetrics
	}
	return false
}

func (x *MetricOptions) GetExportInterval() *durationpb.Duration {
	if x != nil {
		return x.ExportInterval
	}
	return nil
}

var File_internal_impl_kube_proto protoreflect.FileDescriptor

var file_internal_impl_kube_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6d, 0x70, 0x6c, 0x2f,
	0x6b, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x69, 0x6d, 0x70, 0x6c,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xfe, 0x02, 0x0a, 0x10, 0x42, 0x61, 0x62, 0x79, 0x73, 0x69, 0x74, 0x74, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6d,
	0x70, 0x6c, 0x2e, 0x42, 0x61, 0x62, 0x79, 0x73, 0x69, 0x74, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x3a, 0x0a,
	0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x69, 0x6d, 0x70, 0x6c, 0x2e, 0x42, 0x61, 0x62, 0x79, 0x73, 0x69, 0x74, 0x74, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x2d, 0x0a, 0x09, 0x74, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69,
	0x6d, 0x70, 0x6c, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x09, 0x74,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x1a, 0x3c, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x3a, 0x0a, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x2d,
	0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x69, 0x6d, 0x70, 0x6c, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x87, 0x01,
	0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x32, 0x0a, 0x15, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13,
	0x61, 0x75, 0x74, 0x6f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x12, 0x42, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x57, 0x65, 0x61,
	0x76, 0x65, 0x72, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65, 0x72, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6d, 0x70, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_impl_kube_proto_rawDescOnce sync.Once
	file_internal_impl_kube_proto_rawDescData = file_internal_impl_kube_proto_rawDesc
)

func file_internal_impl_kube_proto_rawDescGZIP() []byte {
	file_internal_impl_kube_proto_rawDescOnce.Do(func() {
		file_internal_impl_kube_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_impl_kube_proto_rawDescData)
	})
	return file_internal_impl_kube_proto_rawDescData
}

var file_internal_impl_kube_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_impl_kube_proto_goTypes = []interface{}{
	(*BabysitterConfig)(nil),    // 0: impl.BabysitterConfig
	(*Telemetry)(nil),           // 1: impl.Telemetry
	(*MetricOptions)(nil),       // 2: impl.MetricOptions
	nil,                         // 3: impl.BabysitterConfig.ListenersEntry
	nil,                         // 4: impl.BabysitterConfig.GroupsEntry
	(*durationpb.Duration)(nil), // 5: google.protobuf.Duration
}
var file_internal_impl_kube_proto_depIdxs = []int32{
	3, // 0: impl.BabysitterConfig.listeners:type_name -> impl.BabysitterConfig.ListenersEntry
	4, // 1: impl.BabysitterConfig.groups:type_name -> impl.BabysitterConfig.GroupsEntry
	1, // 2: impl.BabysitterConfig.telemetry:type_name -> impl.Telemetry
	2, // 3: impl.Telemetry.metrics:type_name -> impl.MetricOptions
	5, // 4: impl.MetricOptions.export_interval:type_name -> google.protobuf.Duration
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_internal_impl_kube_proto_init() }
func file_internal_impl_kube_proto_init() {
	if File_internal_impl_kube_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_impl_kube_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BabysitterConfig); i {
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
		file_internal_impl_kube_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Telemetry); i {
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
		file_internal_impl_kube_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricOptions); i {
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
			RawDescriptor: file_internal_impl_kube_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_impl_kube_proto_goTypes,
		DependencyIndexes: file_internal_impl_kube_proto_depIdxs,
		MessageInfos:      file_internal_impl_kube_proto_msgTypes,
	}.Build()
	File_internal_impl_kube_proto = out.File
	file_internal_impl_kube_proto_rawDesc = nil
	file_internal_impl_kube_proto_goTypes = nil
	file_internal_impl_kube_proto_depIdxs = nil
}
