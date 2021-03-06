// Code generated by protoc-gen-gogo.
// source: internal/meta.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	internal/meta.proto

It has these top-level messages:
	Data
	NodeInfo
	DatabaseInfo
	RetentionPolicyInfo
	ShardGroupInfo
	ShardInfo
	ContinuousQueryInfo
	User
	UserPrivilege
	Command
	CreateNodeCommand
	DeleteNodeCommand
	CreateDatabaseCommand
	DropDatabaseCommand
	CreateRetentionPolicyCommand
	DropRetentionPolicyCommand
	SetDefaultRetentionPolicyCommand
	UpdateRetentionPolicyCommand
	CreateShardGroupCommand
	DeleteShardGroupCommand
	CreateContinuousQueryCommand
	DropContinuousQueryCommand
	CreateUserCommand
	DropUserCommand
	UpdateUserCommand
	SetPrivilegeCommand
*/
package internal

import proto "github.com/gogo/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Command_Type int32

const (
	Command_CreateNodeCommand                Command_Type = 1
	Command_DeleteNodeCommand                Command_Type = 2
	Command_CreateDatabaseCommand            Command_Type = 3
	Command_DropDatabaseCommand              Command_Type = 4
	Command_CreateRetentionPolicyCommand     Command_Type = 5
	Command_DropRetentionPolicyCommand       Command_Type = 6
	Command_SetDefaultRetentionPolicyCommand Command_Type = 7
	Command_UpdateRetentionPolicyCommand     Command_Type = 8
	Command_CreateShardGroupCommand          Command_Type = 9
	Command_DeleteShardGroupCommand          Command_Type = 10
	Command_CreateContinuousQueryCommand     Command_Type = 11
	Command_DropContinuousQueryCommand       Command_Type = 12
	Command_CreateUserCommand                Command_Type = 13
	Command_DropUserCommand                  Command_Type = 14
	Command_UpdateUserCommand                Command_Type = 15
	Command_SetPrivilegeCommand              Command_Type = 16
)

var Command_Type_name = map[int32]string{
	1:  "CreateNodeCommand",
	2:  "DeleteNodeCommand",
	3:  "CreateDatabaseCommand",
	4:  "DropDatabaseCommand",
	5:  "CreateRetentionPolicyCommand",
	6:  "DropRetentionPolicyCommand",
	7:  "SetDefaultRetentionPolicyCommand",
	8:  "UpdateRetentionPolicyCommand",
	9:  "CreateShardGroupCommand",
	10: "DeleteShardGroupCommand",
	11: "CreateContinuousQueryCommand",
	12: "DropContinuousQueryCommand",
	13: "CreateUserCommand",
	14: "DropUserCommand",
	15: "UpdateUserCommand",
	16: "SetPrivilegeCommand",
}
var Command_Type_value = map[string]int32{
	"CreateNodeCommand":                1,
	"DeleteNodeCommand":                2,
	"CreateDatabaseCommand":            3,
	"DropDatabaseCommand":              4,
	"CreateRetentionPolicyCommand":     5,
	"DropRetentionPolicyCommand":       6,
	"SetDefaultRetentionPolicyCommand": 7,
	"UpdateRetentionPolicyCommand":     8,
	"CreateShardGroupCommand":          9,
	"DeleteShardGroupCommand":          10,
	"CreateContinuousQueryCommand":     11,
	"DropContinuousQueryCommand":       12,
	"CreateUserCommand":                13,
	"DropUserCommand":                  14,
	"UpdateUserCommand":                15,
	"SetPrivilegeCommand":              16,
}

func (x Command_Type) Enum() *Command_Type {
	p := new(Command_Type)
	*p = x
	return p
}
func (x Command_Type) String() string {
	return proto.EnumName(Command_Type_name, int32(x))
}
func (x *Command_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Command_Type_value, data, "Command_Type")
	if err != nil {
		return err
	}
	*x = Command_Type(value)
	return nil
}

type Data struct {
	Nodes            []*NodeInfo     `protobuf:"bytes,1,rep" json:"Nodes,omitempty"`
	Databases        []*DatabaseInfo `protobuf:"bytes,2,rep" json:"Databases,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}

func (m *Data) GetNodes() []*NodeInfo {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *Data) GetDatabases() []*DatabaseInfo {
	if m != nil {
		return m.Databases
	}
	return nil
}

type NodeInfo struct {
	ID               *uint64 `protobuf:"varint,1,req" json:"ID,omitempty"`
	Host             *string `protobuf:"bytes,2,req" json:"Host,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NodeInfo) Reset()         { *m = NodeInfo{} }
func (m *NodeInfo) String() string { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()    {}

func (m *NodeInfo) GetID() uint64 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *NodeInfo) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

type DatabaseInfo struct {
	Name                   *string                `protobuf:"bytes,1,req" json:"Name,omitempty"`
	DefaultRetentionPolicy *string                `protobuf:"bytes,2,req" json:"DefaultRetentionPolicy,omitempty"`
	Polices                []*RetentionPolicyInfo `protobuf:"bytes,3,rep" json:"Polices,omitempty"`
	ContinuousQueries      []*ContinuousQueryInfo `protobuf:"bytes,4,rep" json:"ContinuousQueries,omitempty"`
	XXX_unrecognized       []byte                 `json:"-"`
}

func (m *DatabaseInfo) Reset()         { *m = DatabaseInfo{} }
func (m *DatabaseInfo) String() string { return proto.CompactTextString(m) }
func (*DatabaseInfo) ProtoMessage()    {}

func (m *DatabaseInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *DatabaseInfo) GetDefaultRetentionPolicy() string {
	if m != nil && m.DefaultRetentionPolicy != nil {
		return *m.DefaultRetentionPolicy
	}
	return ""
}

func (m *DatabaseInfo) GetPolices() []*RetentionPolicyInfo {
	if m != nil {
		return m.Polices
	}
	return nil
}

func (m *DatabaseInfo) GetContinuousQueries() []*ContinuousQueryInfo {
	if m != nil {
		return m.ContinuousQueries
	}
	return nil
}

type RetentionPolicyInfo struct {
	Name               *string           `protobuf:"bytes,1,req" json:"Name,omitempty"`
	Duration           *int64            `protobuf:"varint,2,req" json:"Duration,omitempty"`
	ShardGroupDuration *int64            `protobuf:"varint,3,req" json:"ShardGroupDuration,omitempty"`
	ReplicaN           *uint32           `protobuf:"varint,4,req" json:"ReplicaN,omitempty"`
	ShardGroups        []*ShardGroupInfo `protobuf:"bytes,5,rep" json:"ShardGroups,omitempty"`
	XXX_unrecognized   []byte            `json:"-"`
}

func (m *RetentionPolicyInfo) Reset()         { *m = RetentionPolicyInfo{} }
func (m *RetentionPolicyInfo) String() string { return proto.CompactTextString(m) }
func (*RetentionPolicyInfo) ProtoMessage()    {}

func (m *RetentionPolicyInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *RetentionPolicyInfo) GetDuration() int64 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *RetentionPolicyInfo) GetShardGroupDuration() int64 {
	if m != nil && m.ShardGroupDuration != nil {
		return *m.ShardGroupDuration
	}
	return 0
}

func (m *RetentionPolicyInfo) GetReplicaN() uint32 {
	if m != nil && m.ReplicaN != nil {
		return *m.ReplicaN
	}
	return 0
}

func (m *RetentionPolicyInfo) GetShardGroups() []*ShardGroupInfo {
	if m != nil {
		return m.ShardGroups
	}
	return nil
}

type ShardGroupInfo struct {
	ID               *uint64      `protobuf:"varint,1,req" json:"ID,omitempty"`
	StartTime        *int64       `protobuf:"varint,2,req" json:"StartTime,omitempty"`
	EndTime          *int64       `protobuf:"varint,3,req" json:"EndTime,omitempty"`
	Shards           []*ShardInfo `protobuf:"bytes,4,rep" json:"Shards,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ShardGroupInfo) Reset()         { *m = ShardGroupInfo{} }
func (m *ShardGroupInfo) String() string { return proto.CompactTextString(m) }
func (*ShardGroupInfo) ProtoMessage()    {}

func (m *ShardGroupInfo) GetID() uint64 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *ShardGroupInfo) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *ShardGroupInfo) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *ShardGroupInfo) GetShards() []*ShardInfo {
	if m != nil {
		return m.Shards
	}
	return nil
}

type ShardInfo struct {
	ID               *uint64 `protobuf:"varint,1,req" json:"ID,omitempty"`
	OwnerIDs         *uint64 `protobuf:"varint,2,req" json:"OwnerIDs,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ShardInfo) Reset()         { *m = ShardInfo{} }
func (m *ShardInfo) String() string { return proto.CompactTextString(m) }
func (*ShardInfo) ProtoMessage()    {}

func (m *ShardInfo) GetID() uint64 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *ShardInfo) GetOwnerIDs() uint64 {
	if m != nil && m.OwnerIDs != nil {
		return *m.OwnerIDs
	}
	return 0
}

type ContinuousQueryInfo struct {
	Query            *string `protobuf:"bytes,1,req" json:"Query,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ContinuousQueryInfo) Reset()         { *m = ContinuousQueryInfo{} }
func (m *ContinuousQueryInfo) String() string { return proto.CompactTextString(m) }
func (*ContinuousQueryInfo) ProtoMessage()    {}

func (m *ContinuousQueryInfo) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

type User struct {
	Name             *string          `protobuf:"bytes,1,req" json:"Name,omitempty"`
	Hash             *string          `protobuf:"bytes,2,req" json:"Hash,omitempty"`
	Admin            *bool            `protobuf:"varint,3,req" json:"Admin,omitempty"`
	Privileges       []*UserPrivilege `protobuf:"bytes,4,rep" json:"Privileges,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}

func (m *User) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *User) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func (m *User) GetAdmin() bool {
	if m != nil && m.Admin != nil {
		return *m.Admin
	}
	return false
}

func (m *User) GetPrivileges() []*UserPrivilege {
	if m != nil {
		return m.Privileges
	}
	return nil
}

type UserPrivilege struct {
	Database         *string `protobuf:"bytes,1,req" json:"Database,omitempty"`
	Privilege        *int32  `protobuf:"varint,2,req" json:"Privilege,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserPrivilege) Reset()         { *m = UserPrivilege{} }
func (m *UserPrivilege) String() string { return proto.CompactTextString(m) }
func (*UserPrivilege) ProtoMessage()    {}

func (m *UserPrivilege) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

func (m *UserPrivilege) GetPrivilege() int32 {
	if m != nil && m.Privilege != nil {
		return *m.Privilege
	}
	return 0
}

type Command struct {
	Type             *Command_Type             `protobuf:"varint,1,req,name=type,enum=internal.Command_Type" json:"type,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}

var extRange_Command = []proto.ExtensionRange{
	{100, 536870911},
}

func (*Command) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Command
}
func (m *Command) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *Command) GetType() Command_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Command_CreateNodeCommand
}

type CreateNodeCommand struct {
	Host             *string `protobuf:"bytes,1,req" json:"Host,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateNodeCommand) Reset()         { *m = CreateNodeCommand{} }
func (m *CreateNodeCommand) String() string { return proto.CompactTextString(m) }
func (*CreateNodeCommand) ProtoMessage()    {}

func (m *CreateNodeCommand) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

var E_CreateNodeCommand_Command = &proto.ExtensionDesc{
	ExtendedType:  (*Command)(nil),
	ExtensionType: (*CreateNodeCommand)(nil),
	Field:         101,
	Name:          "internal.CreateNodeCommand.command",
	Tag:           "bytes,101,opt,name=command",
}

type DeleteNodeCommand struct {
	ID               *uint64 `protobuf:"varint,1,req" json:"ID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DeleteNodeCommand) Reset()         { *m = DeleteNodeCommand{} }
func (m *DeleteNodeCommand) String() string { return proto.CompactTextString(m) }
func (*DeleteNodeCommand) ProtoMessage()    {}

func (m *DeleteNodeCommand) GetID() uint64 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

var E_DeleteNodeCommand_Command = &proto.ExtensionDesc{
	ExtendedType:  (*Command)(nil),
	ExtensionType: (*DeleteNodeCommand)(nil),
	Field:         102,
	Name:          "internal.DeleteNodeCommand.command",
	Tag:           "bytes,102,opt,name=command",
}

type CreateDatabaseCommand struct {
	Name             *string `protobuf:"bytes,1,req" json:"Name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateDatabaseCommand) Reset()         { *m = CreateDatabaseCommand{} }
func (m *CreateDatabaseCommand) String() string { return proto.CompactTextString(m) }
func (*CreateDatabaseCommand) ProtoMessage()    {}

func (m *CreateDatabaseCommand) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

var E_CreateDatabaseCommand_Command = &proto.ExtensionDesc{
	ExtendedType:  (*Command)(nil),
	ExtensionType: (*CreateDatabaseCommand)(nil),
	Field:         103,
	Name:          "internal.CreateDatabaseCommand.command",
	Tag:           "bytes,103,opt,name=command",
}

type DropDatabaseCommand struct {
	Name             *string `protobuf:"bytes,1,req" json:"Name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DropDatabaseCommand) Reset()         { *m = DropDatabaseCommand{} }
func (m *DropDatabaseCommand) String() string { return proto.CompactTextString(m) }
func (*DropDatabaseCommand) ProtoMessage()    {}

func (m *DropDatabaseCommand) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

var E_DropDatabaseCommand_Command = &proto.ExtensionDesc{
	ExtendedType:  (*Command)(nil),
	ExtensionType: (*DropDatabaseCommand)(nil),
	Field:         104,
	Name:          "internal.DropDatabaseCommand.command",
	Tag:           "bytes,104,opt,name=command",
}

type CreateRetentionPolicyCommand struct {
	Database         *string              `protobuf:"bytes,1,req" json:"Database,omitempty"`
	RetentionPolicy  *RetentionPolicyInfo `protobuf:"bytes,2,req" json:"RetentionPolicy,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *CreateRetentionPolicyCommand) Reset()         { *m = CreateRetentionPolicyCommand{} }
func (m *CreateRetentionPolicyCommand) String() string { return proto.CompactTextString(m) }
func (*CreateRetentionPolicyCommand) ProtoMessage()    {}

func (m *CreateRetentionPolicyCommand) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

func (m *CreateRetentionPolicyCommand) GetRetentionPolicy() *RetentionPolicyInfo {
	if m != nil {
		return m.RetentionPolicy
	}
	return nil
}

var E_CreateRetentionPolicyCommand_Command = &proto.ExtensionDesc{
	ExtendedType:  (*Command)(nil),
	ExtensionType: (*CreateRetentionPolicyCommand)(nil),
	Field:         105,
	Name:          "internal.CreateRetentionPolicyCommand.command",
	Tag:           "bytes,105,opt,name=command",
}

type DropRetentionPolicyCommand struct {
	Database         *string `protobuf:"bytes,1,req" json:"Database,omitempty"`
	Name             *string `protobuf:"bytes,2,req" json:"Name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DropRetentionPolicyCommand) Reset()         { *m = DropRetentionPolicyCommand{} }
func (m *DropRetentionPolicyCommand) String() string { return proto.CompactTextString(m) }
func (*DropRetentionPolicyCommand) ProtoMessage()    {}

func (m *DropRetentionPolicyCommand) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

func (m *DropRetentionPolicyCommand) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

var E_DropRetentionPolicyCommand_Command = &proto.ExtensionDesc{
	ExtendedType:  (*Command)(nil),
	ExtensionType: (*DropRetentionPolicyCommand)(nil),
	Field:         106,
	Name:          "internal.DropRetentionPolicyCommand.command",
	Tag:           "bytes,106,opt,name=command",
}

type SetDefaultRetentionPolicyCommand struct {
	Database         *string `protobuf:"bytes,1,req" json:"Database,omitempty"`
	Name             *string `protobuf:"bytes,2,req" json:"Name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SetDefaultRetentionPolicyCommand) Reset()         { *m = SetDefaultRetentionPolicyCommand{} }
func (m *SetDefaultRetentionPolicyCommand) String() string { return proto.CompactTextString(m) }
func (*SetDefaultRetentionPolicyCommand) ProtoMessage()    {}

func (m *SetDefaultRetentionPolicyCommand) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

func (m *SetDefaultRetentionPolicyCommand) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

var E_SetDefaultRetentionPolicyCommand_Command = &proto.ExtensionDesc{
	ExtendedType:  (*Command)(nil),
	ExtensionType: (*SetDefaultRetentionPolicyCommand)(nil),
	Field:         107,
	Name:          "internal.SetDefaultRetentionPolicyCommand.command",
	Tag:           "bytes,107,opt,name=command",
}

type UpdateRetentionPolicyCommand struct {
	Database         *string `protobuf:"bytes,1,req" json:"Database,omitempty"`
	Name             *string `protobuf:"bytes,2,req" json:"Name,omitempty"`
	NewName          *string `protobuf:"bytes,3,opt" json:"NewName,omitempty"`
	Duration         *int64  `protobuf:"varint,4,opt" json:"Duration,omitempty"`
	ReplicaN         *uint32 `protobuf:"varint,5,opt" json:"ReplicaN,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UpdateRetentionPolicyCommand) Reset()         { *m = UpdateRetentionPolicyCommand{} }
func (m *UpdateRetentionPolicyCommand) String() string { return proto.CompactTextString(m) }
func (*UpdateRetentionPolicyCommand) ProtoMessage()    {}

func (m *UpdateRetentionPolicyCommand) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

func (m *UpdateRetentionPolicyCommand) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *UpdateRetentionPolicyCommand) GetNewName() string {
	if m != nil && m.NewName != nil {
		return *m.NewName
	}
	return ""
}

func (m *UpdateRetentionPolicyCommand) GetDuration() int64 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *UpdateRetentionPolicyCommand) GetReplicaN() uint32 {
	if m != nil && m.ReplicaN != nil {
		return *m.ReplicaN
	}
	return 0
}

var E_UpdateRetentionPolicyCommand_Command = &proto.ExtensionDesc{
	ExtendedType:  (*Command)(nil),
	ExtensionType: (*UpdateRetentionPolicyCommand)(nil),
	Field:         108,
	Name:          "internal.UpdateRetentionPolicyCommand.command",
	Tag:           "bytes,108,opt,name=command",
}

type CreateShardGroupCommand struct {
	Database         *string `protobuf:"bytes,1,req" json:"Database,omitempty"`
	Policy           *string `protobuf:"bytes,2,req" json:"Policy,omitempty"`
	Timestamp        *int64  `protobuf:"varint,3,req" json:"Timestamp,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateShardGroupCommand) Reset()         { *m = CreateShardGroupCommand{} }
func (m *CreateShardGroupCommand) String() string { return proto.CompactTextString(m) }
func (*CreateShardGroupCommand) ProtoMessage()    {}

func (m *CreateShardGroupCommand) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

func (m *CreateShardGroupCommand) GetPolicy() string {
	if m != nil && m.Policy != nil {
		return *m.Policy
	}
	return ""
}

func (m *CreateShardGroupCommand) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

var E_CreateShardGroupCommand_Command = &proto.ExtensionDesc{
	ExtendedType:  (*Command)(nil),
	ExtensionType: (*CreateShardGroupCommand)(nil),
	Field:         109,
	Name:          "internal.CreateShardGroupCommand.command",
	Tag:           "bytes,109,opt,name=command",
}

type DeleteShardGroupCommand struct {
	Database         *string `protobuf:"bytes,1,req" json:"Database,omitempty"`
	Policy           *string `protobuf:"bytes,2,req" json:"Policy,omitempty"`
	ShardGroupID     *uint64 `protobuf:"varint,3,req" json:"ShardGroupID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DeleteShardGroupCommand) Reset()         { *m = DeleteShardGroupCommand{} }
func (m *DeleteShardGroupCommand) String() string { return proto.CompactTextString(m) }
func (*DeleteShardGroupCommand) ProtoMessage()    {}

func (m *DeleteShardGroupCommand) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

func (m *DeleteShardGroupCommand) GetPolicy() string {
	if m != nil && m.Policy != nil {
		return *m.Policy
	}
	return ""
}

func (m *DeleteShardGroupCommand) GetShardGroupID() uint64 {
	if m != nil && m.ShardGroupID != nil {
		return *m.ShardGroupID
	}
	return 0
}

var E_DeleteShardGroupCommand_Command = &proto.ExtensionDesc{
	ExtendedType:  (*Command)(nil),
	ExtensionType: (*DeleteShardGroupCommand)(nil),
	Field:         110,
	Name:          "internal.DeleteShardGroupCommand.command",
	Tag:           "bytes,110,opt,name=command",
}

type CreateContinuousQueryCommand struct {
	Database         *string `protobuf:"bytes,1,req" json:"Database,omitempty"`
	Name             *string `protobuf:"bytes,2,req" json:"Name,omitempty"`
	Query            *string `protobuf:"bytes,3,req" json:"Query,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateContinuousQueryCommand) Reset()         { *m = CreateContinuousQueryCommand{} }
func (m *CreateContinuousQueryCommand) String() string { return proto.CompactTextString(m) }
func (*CreateContinuousQueryCommand) ProtoMessage()    {}

func (m *CreateContinuousQueryCommand) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

func (m *CreateContinuousQueryCommand) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CreateContinuousQueryCommand) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

var E_CreateContinuousQueryCommand_Command = &proto.ExtensionDesc{
	ExtendedType:  (*Command)(nil),
	ExtensionType: (*CreateContinuousQueryCommand)(nil),
	Field:         111,
	Name:          "internal.CreateContinuousQueryCommand.command",
	Tag:           "bytes,111,opt,name=command",
}

type DropContinuousQueryCommand struct {
	Database         *string `protobuf:"bytes,1,req" json:"Database,omitempty"`
	Name             *string `protobuf:"bytes,2,req" json:"Name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DropContinuousQueryCommand) Reset()         { *m = DropContinuousQueryCommand{} }
func (m *DropContinuousQueryCommand) String() string { return proto.CompactTextString(m) }
func (*DropContinuousQueryCommand) ProtoMessage()    {}

func (m *DropContinuousQueryCommand) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

func (m *DropContinuousQueryCommand) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

var E_DropContinuousQueryCommand_Command = &proto.ExtensionDesc{
	ExtendedType:  (*Command)(nil),
	ExtensionType: (*DropContinuousQueryCommand)(nil),
	Field:         112,
	Name:          "internal.DropContinuousQueryCommand.command",
	Tag:           "bytes,112,opt,name=command",
}

type CreateUserCommand struct {
	Name             *string `protobuf:"bytes,1,req" json:"Name,omitempty"`
	Hash             *string `protobuf:"bytes,2,req" json:"Hash,omitempty"`
	Admin            *bool   `protobuf:"varint,3,req" json:"Admin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateUserCommand) Reset()         { *m = CreateUserCommand{} }
func (m *CreateUserCommand) String() string { return proto.CompactTextString(m) }
func (*CreateUserCommand) ProtoMessage()    {}

func (m *CreateUserCommand) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CreateUserCommand) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func (m *CreateUserCommand) GetAdmin() bool {
	if m != nil && m.Admin != nil {
		return *m.Admin
	}
	return false
}

var E_CreateUserCommand_Command = &proto.ExtensionDesc{
	ExtendedType:  (*Command)(nil),
	ExtensionType: (*CreateUserCommand)(nil),
	Field:         113,
	Name:          "internal.CreateUserCommand.command",
	Tag:           "bytes,113,opt,name=command",
}

type DropUserCommand struct {
	Name             *string `protobuf:"bytes,1,req" json:"Name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DropUserCommand) Reset()         { *m = DropUserCommand{} }
func (m *DropUserCommand) String() string { return proto.CompactTextString(m) }
func (*DropUserCommand) ProtoMessage()    {}

func (m *DropUserCommand) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

var E_DropUserCommand_Command = &proto.ExtensionDesc{
	ExtendedType:  (*Command)(nil),
	ExtensionType: (*DropUserCommand)(nil),
	Field:         114,
	Name:          "internal.DropUserCommand.command",
	Tag:           "bytes,114,opt,name=command",
}

type UpdateUserCommand struct {
	Name             *string `protobuf:"bytes,1,req" json:"Name,omitempty"`
	Hash             *string `protobuf:"bytes,2,req" json:"Hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UpdateUserCommand) Reset()         { *m = UpdateUserCommand{} }
func (m *UpdateUserCommand) String() string { return proto.CompactTextString(m) }
func (*UpdateUserCommand) ProtoMessage()    {}

func (m *UpdateUserCommand) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *UpdateUserCommand) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

var E_UpdateUserCommand_Command = &proto.ExtensionDesc{
	ExtendedType:  (*Command)(nil),
	ExtensionType: (*UpdateUserCommand)(nil),
	Field:         115,
	Name:          "internal.UpdateUserCommand.command",
	Tag:           "bytes,115,opt,name=command",
}

type SetPrivilegeCommand struct {
	Username         *string `protobuf:"bytes,1,req" json:"Username,omitempty"`
	Database         *string `protobuf:"bytes,2,req" json:"Database,omitempty"`
	Privilege        *int32  `protobuf:"varint,3,req" json:"Privilege,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SetPrivilegeCommand) Reset()         { *m = SetPrivilegeCommand{} }
func (m *SetPrivilegeCommand) String() string { return proto.CompactTextString(m) }
func (*SetPrivilegeCommand) ProtoMessage()    {}

func (m *SetPrivilegeCommand) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *SetPrivilegeCommand) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

func (m *SetPrivilegeCommand) GetPrivilege() int32 {
	if m != nil && m.Privilege != nil {
		return *m.Privilege
	}
	return 0
}

var E_SetPrivilegeCommand_Command = &proto.ExtensionDesc{
	ExtendedType:  (*Command)(nil),
	ExtensionType: (*SetPrivilegeCommand)(nil),
	Field:         116,
	Name:          "internal.SetPrivilegeCommand.command",
	Tag:           "bytes,116,opt,name=command",
}

func init() {
	proto.RegisterEnum("internal.Command_Type", Command_Type_name, Command_Type_value)
	proto.RegisterExtension(E_CreateNodeCommand_Command)
	proto.RegisterExtension(E_DeleteNodeCommand_Command)
	proto.RegisterExtension(E_CreateDatabaseCommand_Command)
	proto.RegisterExtension(E_DropDatabaseCommand_Command)
	proto.RegisterExtension(E_CreateRetentionPolicyCommand_Command)
	proto.RegisterExtension(E_DropRetentionPolicyCommand_Command)
	proto.RegisterExtension(E_SetDefaultRetentionPolicyCommand_Command)
	proto.RegisterExtension(E_UpdateRetentionPolicyCommand_Command)
	proto.RegisterExtension(E_CreateShardGroupCommand_Command)
	proto.RegisterExtension(E_DeleteShardGroupCommand_Command)
	proto.RegisterExtension(E_CreateContinuousQueryCommand_Command)
	proto.RegisterExtension(E_DropContinuousQueryCommand_Command)
	proto.RegisterExtension(E_CreateUserCommand_Command)
	proto.RegisterExtension(E_DropUserCommand_Command)
	proto.RegisterExtension(E_UpdateUserCommand_Command)
	proto.RegisterExtension(E_SetPrivilegeCommand_Command)
}
