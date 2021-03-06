// Code generated by protoc-gen-go.
// source: ipv4_rib_edm_proto.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_isis_as_information is a generated protocol buffer package.

It is generated from these files:
	ipv4_rib_edm_proto.proto

It has these top-level messages:
	Ipv4RibEdmProto_KEYS
	Ipv4RibEdmProto
*/
package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_isis_as_information

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Information of a rib protocol
type Ipv4RibEdmProto_KEYS struct {
	VrfName        string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName         string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName        string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	RouteTableName string `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName" json:"route_table_name,omitempty"`
	As             string `protobuf:"bytes,5,opt,name=as" json:"as,omitempty"`
}

func (m *Ipv4RibEdmProto_KEYS) Reset()                    { *m = Ipv4RibEdmProto_KEYS{} }
func (m *Ipv4RibEdmProto_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmProto_KEYS) ProtoMessage()               {}
func (*Ipv4RibEdmProto_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv4RibEdmProto_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetAs() string {
	if m != nil {
		return m.As
	}
	return ""
}

type Ipv4RibEdmProto struct {
	// Name
	ProtocolNames string `protobuf:"bytes,50,opt,name=protocol_names,json=protocolNames" json:"protocol_names,omitempty"`
	// Instance
	Instance string `protobuf:"bytes,51,opt,name=instance" json:"instance,omitempty"`
	// Proto version
	Version uint32 `protobuf:"varint,52,opt,name=version" json:"version,omitempty"`
	// Number of redist clients
	RedistributionClientCount uint32 `protobuf:"varint,53,opt,name=redistribution_client_count,json=redistributionClientCount" json:"redistribution_client_count,omitempty"`
	// Number of proto clients
	ProtocolClientsCount uint32 `protobuf:"varint,54,opt,name=protocol_clients_count,json=protocolClientsCount" json:"protocol_clients_count,omitempty"`
	// Number of routes (including active, backup and deleted), where, number of backup routes = RoutesCounts - ActiveRoutesCount - DeletedRoutesCount
	RoutesCounts uint32 `protobuf:"varint,55,opt,name=routes_counts,json=routesCounts" json:"routes_counts,omitempty"`
	// Number of active routes (not deleted)
	ActiveRoutesCount uint32 `protobuf:"varint,56,opt,name=active_routes_count,json=activeRoutesCount" json:"active_routes_count,omitempty"`
	// Number of deleted routes
	DeletedRoutesCount uint32 `protobuf:"varint,57,opt,name=deleted_routes_count,json=deletedRoutesCount" json:"deleted_routes_count,omitempty"`
	// Number of paths for all routes
	PathsCount uint32 `protobuf:"varint,58,opt,name=paths_count,json=pathsCount" json:"paths_count,omitempty"`
	// Memory for proto's routes and paths in bytes
	ProtocolRouteMemory uint32 `protobuf:"varint,59,opt,name=protocol_route_memory,json=protocolRouteMemory" json:"protocol_route_memory,omitempty"`
	// Number of backup routes
	BackupRoutesCount uint32 `protobuf:"varint,60,opt,name=backup_routes_count,json=backupRoutesCount" json:"backup_routes_count,omitempty"`
}

func (m *Ipv4RibEdmProto) Reset()                    { *m = Ipv4RibEdmProto{} }
func (m *Ipv4RibEdmProto) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmProto) ProtoMessage()               {}
func (*Ipv4RibEdmProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv4RibEdmProto) GetProtocolNames() string {
	if m != nil {
		return m.ProtocolNames
	}
	return ""
}

func (m *Ipv4RibEdmProto) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv4RibEdmProto) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetRedistributionClientCount() uint32 {
	if m != nil {
		return m.RedistributionClientCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetProtocolClientsCount() uint32 {
	if m != nil {
		return m.ProtocolClientsCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetRoutesCounts() uint32 {
	if m != nil {
		return m.RoutesCounts
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetActiveRoutesCount() uint32 {
	if m != nil {
		return m.ActiveRoutesCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetDeletedRoutesCount() uint32 {
	if m != nil {
		return m.DeletedRoutesCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetProtocolRouteMemory() uint32 {
	if m != nil {
		return m.ProtocolRouteMemory
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetBackupRoutesCount() uint32 {
	if m != nil {
		return m.BackupRoutesCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv4RibEdmProto_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.isis.as.information.ipv4_rib_edm_proto_KEYS")
	proto.RegisterType((*Ipv4RibEdmProto)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.isis.as.information.ipv4_rib_edm_proto")
}

func init() { proto.RegisterFile("ipv4_rib_edm_proto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe5, 0x14, 0x9a, 0x30, 0x90, 0x08, 0xb6, 0x85, 0x6e, 0xe0, 0x40, 0x55, 0x84, 0x94,
	0xd3, 0x0a, 0xb5, 0xe1, 0x3f, 0xe2, 0x52, 0x71, 0x42, 0x70, 0x08, 0x5c, 0x38, 0xad, 0xc6, 0xf6,
	0x5a, 0xac, 0x88, 0xbd, 0xd6, 0xce, 0xda, 0xa2, 0x2f, 0xc1, 0x43, 0xf0, 0x8a, 0xbc, 0x00, 0xf2,
	0xac, 0x1d, 0x12, 0xa2, 0x5e, 0x1c, 0xed, 0xfc, 0x7e, 0xdf, 0xe4, 0xf3, 0x1a, 0xa4, 0xad, 0xdb,
	0xa5, 0xf6, 0x36, 0xd5, 0x26, 0x2f, 0x75, 0xed, 0x5d, 0x70, 0x8a, 0x9f, 0xe2, 0x57, 0x92, 0x59,
	0xca, 0x9c, 0xb6, 0x8e, 0xf4, 0x4f, 0xaf, 0x6d, 0xcd, 0x16, 0xeb, 0xae, 0x36, 0x5e, 0x75, 0x27,
	0x0a, 0x79, 0x7a, 0xa5, 0x5a, 0x5f, 0x50, 0xf7, 0x50, 0x58, 0x90, 0xc2, 0x42, 0x51, 0xf7, 0x4b,
	0x58, 0xa8, 0x3e, 0xe3, 0x5d, 0x13, 0x8c, 0x0e, 0x98, 0xae, 0x8d, 0xae, 0xb0, 0x34, 0x74, 0x1d,
	0x88, 0xff, 0x9c, 0xb9, 0xb5, 0xb2, 0x64, 0x49, 0x21, 0x29, 0x5b, 0x15, 0xce, 0x97, 0x18, 0xac,
	0xab, 0xce, 0x7e, 0x27, 0x70, 0xb2, 0xdf, 0x56, 0x7f, 0xfc, 0xf0, 0xed, 0x8b, 0x98, 0xc3, 0xa4,
	0xf5, 0x05, 0x2f, 0x91, 0xc9, 0x69, 0xb2, 0xb8, 0xb5, 0x1a, 0xb7, 0xbe, 0xf8, 0x8c, 0xa5, 0x11,
	0x27, 0x30, 0xc6, 0x9e, 0x8c, 0x98, 0x1c, 0x62, 0x04, 0x73, 0x98, 0xd0, 0x40, 0x0e, 0x62, 0x86,
	0x7a, 0xb4, 0x80, 0xbb, 0xff, 0x77, 0x93, 0x37, 0x58, 0x99, 0xf1, 0xfc, 0x6b, 0x37, 0x66, 0x73,
	0x06, 0x23, 0x24, 0x79, 0x93, 0xd9, 0x08, 0xe9, 0xec, 0xcf, 0x01, 0x88, 0xfd, 0x92, 0xe2, 0x29,
	0xcc, 0x86, 0x77, 0x8b, 0x57, 0x20, 0xcf, 0x39, 0x32, 0x1d, 0xa6, 0xdd, 0x32, 0x12, 0x0f, 0x61,
	0x62, 0x2b, 0x0a, 0x58, 0x65, 0x46, 0x5e, 0xb0, 0xb0, 0x39, 0x0b, 0x09, 0xe3, 0xd6, 0x78, 0xb2,
	0xae, 0x92, 0xcb, 0xd3, 0x64, 0x31, 0x5d, 0x0d, 0x47, 0xf1, 0x1e, 0x1e, 0x79, 0x93, 0x5b, 0x0a,
	0xde, 0xa6, 0x4d, 0x77, 0x55, 0x3a, 0x5b, 0x5b, 0x53, 0x05, 0x9d, 0xb9, 0xa6, 0x0a, 0xf2, 0x39,
	0xdb, 0xf3, 0x5d, 0xe5, 0x92, 0x8d, 0xcb, 0x4e, 0x10, 0x4b, 0x78, 0xb0, 0x29, 0x17, 0x93, 0xd4,
	0x47, 0x5f, 0x70, 0xf4, 0x78, 0xa0, 0x31, 0x44, 0x31, 0xf5, 0x04, 0xa6, 0x7c, 0x17, 0xbd, 0x4b,
	0xf2, 0x25, 0xcb, 0x77, 0xe2, 0x90, 0x1d, 0x12, 0x0a, 0x8e, 0x30, 0x0b, 0xb6, 0x35, 0x7a, 0xdb,
	0x95, 0xaf, 0x58, 0xbd, 0x17, 0xd1, 0xea, 0x5f, 0x40, 0x3c, 0x83, 0xe3, 0xdc, 0xac, 0x4d, 0x30,
	0xf9, 0x6e, 0xe0, 0x35, 0x07, 0x44, 0xcf, 0xb6, 0x13, 0x8f, 0xe1, 0x76, 0x8d, 0xe1, 0xfb, 0x20,
	0xbe, 0x61, 0x11, 0x78, 0x14, 0x85, 0x73, 0xb8, 0xbf, 0x79, 0xbb, 0xf8, 0x51, 0x4b, 0x53, 0x3a,
	0x7f, 0x25, 0xdf, 0xb2, 0x7a, 0x34, 0x40, 0x5e, 0xfa, 0x89, 0x51, 0x57, 0x3b, 0xc5, 0xec, 0x47,
	0x53, 0xef, 0xb6, 0x78, 0x17, 0x6b, 0x47, 0xb4, 0x55, 0x22, 0x3d, 0xe4, 0x25, 0x17, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x42, 0x1a, 0xf9, 0x82, 0x4e, 0x03, 0x00, 0x00,
}
