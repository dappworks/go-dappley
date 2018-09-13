// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/pb/config.proto

package configpb

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

type Config struct {
	DynastyConfig        *DynastyConfig   `protobuf:"bytes,1,opt,name=dynastyConfig,proto3" json:"dynastyConfig,omitempty"`
	ConsensusConfig      *ConsensusConfig `protobuf:"bytes,2,opt,name=consensusConfig,proto3" json:"consensusConfig,omitempty"`
	NodeConfig           *NodeConfig      `protobuf:"bytes,3,opt,name=nodeConfig,proto3" json:"nodeConfig,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_efa1842917505b21, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetDynastyConfig() *DynastyConfig {
	if m != nil {
		return m.DynastyConfig
	}
	return nil
}

func (m *Config) GetConsensusConfig() *ConsensusConfig {
	if m != nil {
		return m.ConsensusConfig
	}
	return nil
}

func (m *Config) GetNodeConfig() *NodeConfig {
	if m != nil {
		return m.NodeConfig
	}
	return nil
}

type DynastyConfig struct {
	Producers            []string `protobuf:"bytes,1,rep,name=producers,proto3" json:"producers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DynastyConfig) Reset()         { *m = DynastyConfig{} }
func (m *DynastyConfig) String() string { return proto.CompactTextString(m) }
func (*DynastyConfig) ProtoMessage()    {}
func (*DynastyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_efa1842917505b21, []int{1}
}
func (m *DynastyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynastyConfig.Unmarshal(m, b)
}
func (m *DynastyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynastyConfig.Marshal(b, m, deterministic)
}
func (dst *DynastyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynastyConfig.Merge(dst, src)
}
func (m *DynastyConfig) XXX_Size() int {
	return xxx_messageInfo_DynastyConfig.Size(m)
}
func (m *DynastyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DynastyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DynastyConfig proto.InternalMessageInfo

func (m *DynastyConfig) GetProducers() []string {
	if m != nil {
		return m.Producers
	}
	return nil
}

type ConsensusConfig struct {
	MinerAddr            string   `protobuf:"bytes,1,opt,name=minerAddr,proto3" json:"minerAddr,omitempty"`
	PrivKey            string   `protobuf:"bytes,1,opt,name=privKey,proto3" json:"privKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsensusConfig) Reset()         { *m = ConsensusConfig{} }
func (m *ConsensusConfig) String() string { return proto.CompactTextString(m) }
func (*ConsensusConfig) ProtoMessage()    {}
func (*ConsensusConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_efa1842917505b21, []int{2}
}
func (m *ConsensusConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusConfig.Unmarshal(m, b)
}
func (m *ConsensusConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusConfig.Marshal(b, m, deterministic)
}
func (dst *ConsensusConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusConfig.Merge(dst, src)
}
func (m *ConsensusConfig) XXX_Size() int {
	return xxx_messageInfo_ConsensusConfig.Size(m)
}
func (m *ConsensusConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusConfig proto.InternalMessageInfo

func (m *ConsensusConfig) GetMinerAddr() string {
	if m != nil {
		return m.MinerAddr
	}
	return ""
}
func (m *ConsensusConfig) GetPrivKey() string {
	if m != nil {
		return m.PrivKey
	}
	return ""
}

type NodeConfig struct {
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Seed                 string   `protobuf:"bytes,2,opt,name=seed,proto3" json:"seed,omitempty"`
	DbPath               string   `protobuf:"bytes,3,opt,name=dbPath,proto3" json:"dbPath,omitempty"`
	RpcPort              uint32   `protobuf:"varint,4,opt,name=rpcPort,proto3" json:"rpcPort,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeConfig) Reset()         { *m = NodeConfig{} }
func (m *NodeConfig) String() string { return proto.CompactTextString(m) }
func (*NodeConfig) ProtoMessage()    {}
func (*NodeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_efa1842917505b21, []int{3}
}
func (m *NodeConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeConfig.Unmarshal(m, b)
}
func (m *NodeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeConfig.Marshal(b, m, deterministic)
}
func (dst *NodeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeConfig.Merge(dst, src)
}
func (m *NodeConfig) XXX_Size() int {
	return xxx_messageInfo_NodeConfig.Size(m)
}
func (m *NodeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NodeConfig proto.InternalMessageInfo

func (m *NodeConfig) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *NodeConfig) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

func (m *NodeConfig) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *NodeConfig) GetRpcPort() uint32 {
	if m != nil {
		return m.RpcPort
	}
	return 0
}

func (m *NodeConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CliConfig struct {
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CliConfig) Reset()         { *m = CliConfig{} }
func (m *CliConfig) String() string { return proto.CompactTextString(m) }
func (*CliConfig) ProtoMessage()    {}
func (*CliConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_efa1842917505b21, []int{4}
}
func (m *CliConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CliConfig.Unmarshal(m, b)
}
func (m *CliConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CliConfig.Marshal(b, m, deterministic)
}
func (dst *CliConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CliConfig.Merge(dst, src)
}
func (m *CliConfig) XXX_Size() int {
	return xxx_messageInfo_CliConfig.Size(m)
}
func (m *CliConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CliConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CliConfig proto.InternalMessageInfo

func (m *CliConfig) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *CliConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "configpb.Config")
	proto.RegisterType((*DynastyConfig)(nil), "configpb.DynastyConfig")
	proto.RegisterType((*ConsensusConfig)(nil), "configpb.ConsensusConfig")
	proto.RegisterType((*NodeConfig)(nil), "configpb.NodeConfig")
	proto.RegisterType((*CliConfig)(nil), "configpb.CliConfig")
}

func init() { proto.RegisterFile("config/pb/config.proto", fileDescriptor_config_efa1842917505b21) }

var fileDescriptor_config_efa1842917505b21 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0x5b, 0x63, 0x77, 0x24, 0x14, 0x06, 0xa9, 0x51, 0x3c, 0x94, 0x9c, 0x7a, 0x31,
	0x01, 0xf5, 0x26, 0x1e, 0x24, 0x9e, 0xa5, 0xec, 0x3f, 0x48, 0xb2, 0xab, 0x06, 0x74, 0x77, 0xd9,
	0x4d, 0x91, 0x5e, 0xfd, 0x69, 0xfe, 0x32, 0xc9, 0x24, 0xed, 0x26, 0x3d, 0xf4, 0x36, 0xef, 0xed,
	0xfb, 0x5e, 0x66, 0x08, 0x2c, 0x2b, 0xad, 0xde, 0xeb, 0x8f, 0xcc, 0x94, 0x59, 0x37, 0xa5, 0xc6,
	0xea, 0x46, 0xe3, 0xbc, 0x53, 0xa6, 0x4c, 0xfe, 0x02, 0x08, 0x73, 0x12, 0xf8, 0x0c, 0x91, 0xd8,
	0xa9, 0xc2, 0x35, 0xbb, 0xce, 0x88, 0x83, 0x55, 0xb0, 0xbe, 0xb8, 0xbf, 0x4a, 0xf7, 0xe1, 0xf4,
	0x75, 0xf8, 0xcc, 0xc7, 0x69, 0xcc, 0x61, 0x51, 0x69, 0xe5, 0xa4, 0x72, 0x5b, 0xd7, 0x17, 0x4c,
	0xa8, 0xe0, 0xda, 0x17, 0xe4, 0xe3, 0x00, 0x3f, 0x26, 0xf0, 0x11, 0x40, 0x69, 0x21, 0x7b, 0x7e,
	0x4a, 0xfc, 0xa5, 0xe7, 0xdf, 0x0e, 0x6f, 0x7c, 0x90, 0x4b, 0xee, 0x20, 0x1a, 0xad, 0x86, 0xb7,
	0xc0, 0x8c, 0xd5, 0x62, 0x5b, 0x49, 0xeb, 0xe2, 0x60, 0x35, 0x5d, 0x33, 0xee, 0x8d, 0x24, 0x83,
	0xc5, 0xd1, 0x22, 0x2d, 0xf0, 0x5d, 0x2b, 0x69, 0x5f, 0x84, 0xb0, 0x74, 0x37, 0xe3, 0xde, 0x48,
	0x7e, 0x03, 0x00, 0xff, 0x69, 0x44, 0x98, 0x19, 0x6d, 0x1b, 0xca, 0x45, 0x9c, 0xe6, 0xd6, 0x73,
	0x52, 0x0a, 0x3a, 0x99, 0x71, 0x9a, 0x71, 0x09, 0xa1, 0x28, 0x37, 0x45, 0xf3, 0x49, 0x87, 0x30,
	0xde, 0x2b, 0x8c, 0xe1, 0xdc, 0x9a, 0x6a, 0xd3, 0x56, 0xcc, 0xa8, 0x62, 0x2f, 0xf1, 0x06, 0xe6,
	0xa6, 0x70, 0xee, 0x47, 0x5b, 0x11, 0x9f, 0x11, 0x73, 0xd0, 0xc9, 0x13, 0xb0, 0xfc, 0xab, 0x3e,
	0xb1, 0xc2, 0x10, 0x9e, 0x8c, 0xe1, 0x32, 0xa4, 0xff, 0xfe, 0xf0, 0x1f, 0x00, 0x00, 0xff, 0xff,
	0xda, 0xe5, 0x61, 0xf6, 0x11, 0x02, 0x00, 0x00,
}
