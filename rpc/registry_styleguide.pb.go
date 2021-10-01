// Copyright 2021 Google LLC
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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: google/cloud/apigee/registry/applications/v1alpha1/registry_styleguide.proto

package rpc

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Possible severities for this guideline.
type Guideline_Status int32

const (
	// The default value, unused.
	Guideline_STATUS_UNSPECIFIED Guideline_Status = 0
	// The guideline is being proposed, and shouldn't yet
	// be enforced.
	Guideline_PROPOSED Guideline_Status = 1
	// The guideline is active and should be enforced.
	Guideline_ACTIVE Guideline_Status = 2
	// The guideline is deprecated as of the recent version,
	// and shouldn't be enforced.
	Guideline_DEPRECATED Guideline_Status = 3
	// The guideline has been disabled and shouldn't be enforced.
	Guideline_DISABLED Guideline_Status = 4
)

// Enum value maps for Guideline_Status.
var (
	Guideline_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "PROPOSED",
		2: "ACTIVE",
		3: "DEPRECATED",
		4: "DISABLED",
	}
	Guideline_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"PROPOSED":           1,
		"ACTIVE":             2,
		"DEPRECATED":         3,
		"DISABLED":           4,
	}
)

func (x Guideline_Status) Enum() *Guideline_Status {
	p := new(Guideline_Status)
	*p = x
	return p
}

func (x Guideline_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Guideline_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_enumTypes[0].Descriptor()
}

func (Guideline_Status) Type() protoreflect.EnumType {
	return &file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_enumTypes[0]
}

func (x Guideline_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Guideline_Status.Descriptor instead.
func (Guideline_Status) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_rawDescGZIP(), []int{1, 0}
}

// Possible severities for the violation of a rule.
type Rule_Severity int32

const (
	// The default value, unused.
	Rule_SEVERITY_UNSPECIFIED Rule_Severity = 0
	// Violation of the rule is an error that must be fixed.
	Rule_ERROR Rule_Severity = 1
	// Violation of the rule is a pattern that is wrong,
	// and should be warned about.
	Rule_WARNING Rule_Severity = 2
	// Violation of the rule is not necessarily a bad pattern
	// or error, but information the user should be aware of.
	Rule_INFO Rule_Severity = 3
	// Violation of the rule is a hint that is provided to
	// the user to fix their spec's design.
	Rule_HINT Rule_Severity = 4
)

// Enum value maps for Rule_Severity.
var (
	Rule_Severity_name = map[int32]string{
		0: "SEVERITY_UNSPECIFIED",
		1: "ERROR",
		2: "WARNING",
		3: "INFO",
		4: "HINT",
	}
	Rule_Severity_value = map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"ERROR":                1,
		"WARNING":              2,
		"INFO":                 3,
		"HINT":                 4,
	}
)

func (x Rule_Severity) Enum() *Rule_Severity {
	p := new(Rule_Severity)
	*p = x
	return p
}

func (x Rule_Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Rule_Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_enumTypes[1].Descriptor()
}

func (Rule_Severity) Type() protoreflect.EnumType {
	return &file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_enumTypes[1]
}

func (x Rule_Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Rule_Severity.Descriptor instead.
func (Rule_Severity) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_rawDescGZIP(), []int{2, 0}
}

// StyleGuide defines a set of guidelines and linters that govern the
// static analysis of API Specs with specified mime types.
type StyleGuide struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier of the style guide.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Human-meaningful name of the style guide.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// This field lists the MIME types of the specs that the style guide applies to.
	// It is a list containing style (format) descriptors that are specified
	// as a Media Type (https://en.wikipedia.org/wiki/Media_type).
	// Possible values include "application/vnd.apigee.proto",
	// "application/vnd.apigee.openapi", and "application/vnd.apigee.graphql",
	// with possible suffixes representing compression types. These hypothetical
	// names are defined in the vendor tree defined in RFC6838
	// (https://tools.ietf.org/html/rfc6838) and are not final.
	MimeTypes []string `protobuf:"bytes,3,rep,name=mime_types,json=mimeTypes,proto3" json:"mime_types,omitempty"`
	// A list of guidelines that are specified by this style guide.
	Guidelines []*Guideline `protobuf:"bytes,4,rep,name=guidelines,proto3" json:"guidelines,omitempty"`
	// A list of linters that this style guide uses.
	Linters []*Linter `protobuf:"bytes,5,rep,name=linters,proto3" json:"linters,omitempty"`
}

func (x *StyleGuide) Reset() {
	*x = StyleGuide{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StyleGuide) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StyleGuide) ProtoMessage() {}

func (x *StyleGuide) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StyleGuide.ProtoReflect.Descriptor instead.
func (*StyleGuide) Descriptor() ([]byte, []int) {
	return file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_rawDescGZIP(), []int{0}
}

func (x *StyleGuide) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StyleGuide) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *StyleGuide) GetMimeTypes() []string {
	if x != nil {
		return x.MimeTypes
	}
	return nil
}

func (x *StyleGuide) GetGuidelines() []*Guideline {
	if x != nil {
		return x.Guidelines
	}
	return nil
}

func (x *StyleGuide) GetLinters() []*Linter {
	if x != nil {
		return x.Linters
	}
	return nil
}

// Guideline defines a set of rules that achieve a common design
// goal for API specs.
type Guideline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier of the guideline.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Human-meaningful name of the guideline.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// A detailed description of the guideline.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// A list of rules that this guideline specifies.
	Rules []*Rule `protobuf:"bytes,4,rep,name=rules,proto3" json:"rules,omitempty"`
	// Indicates the status of the guideline.
	Status Guideline_Status `protobuf:"varint,5,opt,name=status,proto3,enum=google.cloud.apigee.registry.applications.v1alpha1.Guideline_Status" json:"status,omitempty"`
}

func (x *Guideline) Reset() {
	*x = Guideline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Guideline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Guideline) ProtoMessage() {}

func (x *Guideline) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Guideline.ProtoReflect.Descriptor instead.
func (*Guideline) Descriptor() ([]byte, []int) {
	return file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_rawDescGZIP(), []int{1}
}

func (x *Guideline) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Guideline) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Guideline) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Guideline) GetRules() []*Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *Guideline) GetStatus() Guideline_Status {
	if x != nil {
		return x.Status
	}
	return Guideline_STATUS_UNSPECIFIED
}

// Rule is a specific design rule that can be applied to an API spec,
// and is enforced by a specified linter.
type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier of the rule.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Human-meaningful name of the rule.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// A detailed description of the rule.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Name of the linter that enforces this rule.
	Linter string `protobuf:"bytes,4,opt,name=linter,proto3" json:"linter,omitempty"`
	// Name of the rule on the linter that enforces it.
	LinterRulename string `protobuf:"bytes,5,opt,name=linter_rulename,json=linterRulename,proto3" json:"linter_rulename,omitempty"`
	// Severity of violating this rule.
	Severity Rule_Severity `protobuf:"varint,6,opt,name=severity,proto3,enum=google.cloud.apigee.registry.applications.v1alpha1.Rule_Severity" json:"severity,omitempty"`
	// A link to additional documentation relating to this rule.
	DocUri string `protobuf:"bytes,7,opt,name=doc_uri,json=docUri,proto3" json:"doc_uri,omitempty"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_rawDescGZIP(), []int{2}
}

func (x *Rule) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Rule) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Rule) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Rule) GetLinter() string {
	if x != nil {
		return x.Linter
	}
	return ""
}

func (x *Rule) GetLinterRulename() string {
	if x != nil {
		return x.LinterRulename
	}
	return ""
}

func (x *Rule) GetSeverity() Rule_Severity {
	if x != nil {
		return x.Severity
	}
	return Rule_SEVERITY_UNSPECIFIED
}

func (x *Rule) GetDocUri() string {
	if x != nil {
		return x.DocUri
	}
	return ""
}

// Linter contains the name and source code / documentation of specific
// linter that a style guide uses.
type Linter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the linter.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A uri to the linter source code or documentation.
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *Linter) Reset() {
	*x = Linter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Linter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Linter) ProtoMessage() {}

func (x *Linter) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Linter.ProtoReflect.Descriptor instead.
func (*Linter) Descriptor() ([]byte, []int) {
	return file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_rawDescGZIP(), []int{3}
}

func (x *Linter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Linter) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

var File_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto protoreflect.FileDescriptor

var file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x67, 0x65, 0x65, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x73, 0x74,
	0x79, 0x6c, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x32,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x67, 0x65, 0x65, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x02, 0x0a, 0x0a, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x47, 0x75, 0x69,
	0x64, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x0a, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x12, 0x5d, 0x0a, 0x0a, 0x67, 0x75, 0x69, 0x64, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x75, 0x69,
	0x64, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x0a, 0x67, 0x75, 0x69, 0x64, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x73, 0x12, 0x54, 0x0a, 0x07, 0x6c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x52,
	0x07, 0x6c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x22, 0xf2, 0x02, 0x0a, 0x09, 0x47, 0x75, 0x69,
	0x64, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x5c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x75, 0x69, 0x64, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x58, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x4f, 0x50, 0x4f, 0x53, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12,
	0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12,
	0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x22, 0xf5, 0x02,
	0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1b, 0x0a, 0x06, 0x6c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a,
	0x0f, 0x6c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x6c, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x5d, 0x0a, 0x08, 0x73,
	0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x41, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x67, 0x65, 0x65, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x6f,
	0x63, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x63,
	0x55, 0x72, 0x69, 0x22, 0x50, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x18, 0x0a, 0x14, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x48,
	0x49, 0x4e, 0x54, 0x10, 0x04, 0x22, 0x38, 0x0a, 0x06, 0x4c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x03, 0x75, 0x72, 0x69, 0x42,
	0x77, 0x0a, 0x36, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2f, 0x72, 0x70, 0x63, 0x3b, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_rawDescOnce sync.Once
	file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_rawDescData = file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_rawDesc
)

func file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_rawDescGZIP() []byte {
	file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_rawDescOnce.Do(func() {
		file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_rawDescData)
	})
	return file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_rawDescData
}

var file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_goTypes = []interface{}{
	(Guideline_Status)(0), // 0: google.cloud.apigee.registry.applications.v1alpha1.Guideline.Status
	(Rule_Severity)(0),    // 1: google.cloud.apigee.registry.applications.v1alpha1.Rule.Severity
	(*StyleGuide)(nil),    // 2: google.cloud.apigee.registry.applications.v1alpha1.StyleGuide
	(*Guideline)(nil),     // 3: google.cloud.apigee.registry.applications.v1alpha1.Guideline
	(*Rule)(nil),          // 4: google.cloud.apigee.registry.applications.v1alpha1.Rule
	(*Linter)(nil),        // 5: google.cloud.apigee.registry.applications.v1alpha1.Linter
}
var file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_depIdxs = []int32{
	3, // 0: google.cloud.apigee.registry.applications.v1alpha1.StyleGuide.guidelines:type_name -> google.cloud.apigee.registry.applications.v1alpha1.Guideline
	5, // 1: google.cloud.apigee.registry.applications.v1alpha1.StyleGuide.linters:type_name -> google.cloud.apigee.registry.applications.v1alpha1.Linter
	4, // 2: google.cloud.apigee.registry.applications.v1alpha1.Guideline.rules:type_name -> google.cloud.apigee.registry.applications.v1alpha1.Rule
	0, // 3: google.cloud.apigee.registry.applications.v1alpha1.Guideline.status:type_name -> google.cloud.apigee.registry.applications.v1alpha1.Guideline.Status
	1, // 4: google.cloud.apigee.registry.applications.v1alpha1.Rule.severity:type_name -> google.cloud.apigee.registry.applications.v1alpha1.Rule.Severity
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_init() }
func file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_init() {
	if File_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StyleGuide); i {
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
		file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Guideline); i {
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
		file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
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
		file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Linter); i {
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
			RawDescriptor: file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_goTypes,
		DependencyIndexes: file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_depIdxs,
		EnumInfos:         file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_enumTypes,
		MessageInfos:      file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_msgTypes,
	}.Build()
	File_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto = out.File
	file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_rawDesc = nil
	file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_goTypes = nil
	file_google_cloud_apigee_registry_applications_v1alpha1_registry_styleguide_proto_depIdxs = nil
}