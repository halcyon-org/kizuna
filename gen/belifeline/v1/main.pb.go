// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: belifeline/v1/main.proto

package mainv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_belifeline_v1_main_proto protoreflect.FileDescriptor

var file_belifeline_v1_main_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x62, 0x65, 0x6c, 0x69,
	0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x62, 0x65, 0x6c, 0x69, 0x66,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xb3, 0x08, 0x0a, 0x11, 0x42, 0x65, 0x4c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62,
	0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x20, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x59, 0x0a, 0x0c, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x62, 0x65, 0x6c,
	0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x76, 0x6f, 0x6b, 0x65, 0x12, 0x22, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x76, 0x6f,
	0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62, 0x65, 0x6c, 0x69,
	0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x5c, 0x0a, 0x0d, 0x45, 0x78, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x23, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x78, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58,
	0x0a, 0x0b, 0x45, 0x78, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e,
	0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x5c, 0x0a, 0x0d, 0x45, 0x78, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x62, 0x65, 0x6c, 0x69,
	0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0a, 0x4b, 0x6f, 0x79, 0x6f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x08, 0x4b,
	0x6f, 0x79, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x53, 0x0a, 0x0a,
	0x4b, 0x6f, 0x79, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x62, 0x65, 0x6c,
	0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79, 0x6f, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62,
	0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79,
	0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5c, 0x0a, 0x0d, 0x4b, 0x6f, 0x79, 0x6f, 0x41, 0x70, 0x69, 0x52, 0x65, 0x76, 0x6f,
	0x6b, 0x65, 0x12, 0x23, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79, 0x6f, 0x41, 0x70, 0x69, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79, 0x6f, 0x41, 0x70, 0x69, 0x52,
	0x65, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x47, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x2e, 0x62, 0x65, 0x6c, 0x69,
	0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x6c, 0x63, 0x79, 0x6f, 0x6e, 0x2d, 0x6f,
	0x72, 0x67, 0x2f, 0x6b, 0x69, 0x7a, 0x75, 0x6e, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x65,
	0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x61, 0x69, 0x6e,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_belifeline_v1_main_proto_goTypes = []any{
	(*ClientCreateRequest)(nil),   // 0: belifeline.v1.ClientCreateRequest
	(*ClientListRequest)(nil),     // 1: belifeline.v1.ClientListRequest
	(*ClientDeleteRequest)(nil),   // 2: belifeline.v1.ClientDeleteRequest
	(*ClientRevokeRequest)(nil),   // 3: belifeline.v1.ClientRevokeRequest
	(*ExtInfoCreateRequest)(nil),  // 4: belifeline.v1.ExtInfoCreateRequest
	(*ExtInfoListRequest)(nil),    // 5: belifeline.v1.ExtInfoListRequest
	(*ExtInfoDeleteRequest)(nil),  // 6: belifeline.v1.ExtInfoDeleteRequest
	(*KoyoCreateRequest)(nil),     // 7: belifeline.v1.KoyoCreateRequest
	(*KoyoListRequest)(nil),       // 8: belifeline.v1.KoyoListRequest
	(*KoyoDeleteRequest)(nil),     // 9: belifeline.v1.KoyoDeleteRequest
	(*KoyoApiRevokeRequest)(nil),  // 10: belifeline.v1.KoyoApiRevokeRequest
	(*StatusRequest)(nil),         // 11: belifeline.v1.StatusRequest
	(*ClientCreateResponse)(nil),  // 12: belifeline.v1.ClientCreateResponse
	(*ClientListResponse)(nil),    // 13: belifeline.v1.ClientListResponse
	(*ClientDeleteResponse)(nil),  // 14: belifeline.v1.ClientDeleteResponse
	(*ClientRevokeResponse)(nil),  // 15: belifeline.v1.ClientRevokeResponse
	(*ExtInfoCreateResponse)(nil), // 16: belifeline.v1.ExtInfoCreateResponse
	(*ExtInfoListResponse)(nil),   // 17: belifeline.v1.ExtInfoListResponse
	(*ExtInfoDeleteResponse)(nil), // 18: belifeline.v1.ExtInfoDeleteResponse
	(*KoyoCreateResponse)(nil),    // 19: belifeline.v1.KoyoCreateResponse
	(*KoyoListResponse)(nil),      // 20: belifeline.v1.KoyoListResponse
	(*KoyoDeleteResponse)(nil),    // 21: belifeline.v1.KoyoDeleteResponse
	(*KoyoApiRevokeResponse)(nil), // 22: belifeline.v1.KoyoApiRevokeResponse
	(*StatusResponse)(nil),        // 23: belifeline.v1.StatusResponse
}
var file_belifeline_v1_main_proto_depIdxs = []int32{
	0,  // 0: belifeline.v1.BeLifelineService.ClientCreate:input_type -> belifeline.v1.ClientCreateRequest
	1,  // 1: belifeline.v1.BeLifelineService.ClientList:input_type -> belifeline.v1.ClientListRequest
	2,  // 2: belifeline.v1.BeLifelineService.ClientDelete:input_type -> belifeline.v1.ClientDeleteRequest
	3,  // 3: belifeline.v1.BeLifelineService.ClientRevoke:input_type -> belifeline.v1.ClientRevokeRequest
	4,  // 4: belifeline.v1.BeLifelineService.ExtInfoCreate:input_type -> belifeline.v1.ExtInfoCreateRequest
	5,  // 5: belifeline.v1.BeLifelineService.ExtInfoList:input_type -> belifeline.v1.ExtInfoListRequest
	6,  // 6: belifeline.v1.BeLifelineService.ExtInfoDelete:input_type -> belifeline.v1.ExtInfoDeleteRequest
	7,  // 7: belifeline.v1.BeLifelineService.KoyoCreate:input_type -> belifeline.v1.KoyoCreateRequest
	8,  // 8: belifeline.v1.BeLifelineService.KoyoList:input_type -> belifeline.v1.KoyoListRequest
	9,  // 9: belifeline.v1.BeLifelineService.KoyoDelete:input_type -> belifeline.v1.KoyoDeleteRequest
	10, // 10: belifeline.v1.BeLifelineService.KoyoApiRevoke:input_type -> belifeline.v1.KoyoApiRevokeRequest
	11, // 11: belifeline.v1.BeLifelineService.Status:input_type -> belifeline.v1.StatusRequest
	12, // 12: belifeline.v1.BeLifelineService.ClientCreate:output_type -> belifeline.v1.ClientCreateResponse
	13, // 13: belifeline.v1.BeLifelineService.ClientList:output_type -> belifeline.v1.ClientListResponse
	14, // 14: belifeline.v1.BeLifelineService.ClientDelete:output_type -> belifeline.v1.ClientDeleteResponse
	15, // 15: belifeline.v1.BeLifelineService.ClientRevoke:output_type -> belifeline.v1.ClientRevokeResponse
	16, // 16: belifeline.v1.BeLifelineService.ExtInfoCreate:output_type -> belifeline.v1.ExtInfoCreateResponse
	17, // 17: belifeline.v1.BeLifelineService.ExtInfoList:output_type -> belifeline.v1.ExtInfoListResponse
	18, // 18: belifeline.v1.BeLifelineService.ExtInfoDelete:output_type -> belifeline.v1.ExtInfoDeleteResponse
	19, // 19: belifeline.v1.BeLifelineService.KoyoCreate:output_type -> belifeline.v1.KoyoCreateResponse
	20, // 20: belifeline.v1.BeLifelineService.KoyoList:output_type -> belifeline.v1.KoyoListResponse
	21, // 21: belifeline.v1.BeLifelineService.KoyoDelete:output_type -> belifeline.v1.KoyoDeleteResponse
	22, // 22: belifeline.v1.BeLifelineService.KoyoApiRevoke:output_type -> belifeline.v1.KoyoApiRevokeResponse
	23, // 23: belifeline.v1.BeLifelineService.Status:output_type -> belifeline.v1.StatusResponse
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_belifeline_v1_main_proto_init() }
func file_belifeline_v1_main_proto_init() {
	if File_belifeline_v1_main_proto != nil {
		return
	}
	file_belifeline_v1_api_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_belifeline_v1_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_belifeline_v1_main_proto_goTypes,
		DependencyIndexes: file_belifeline_v1_main_proto_depIdxs,
	}.Build()
	File_belifeline_v1_main_proto = out.File
	file_belifeline_v1_main_proto_rawDesc = nil
	file_belifeline_v1_main_proto_goTypes = nil
	file_belifeline_v1_main_proto_depIdxs = nil
}
