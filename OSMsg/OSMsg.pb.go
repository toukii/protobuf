// Code generated by protoc-gen-go.
// source: OSMsg.proto
// DO NOT EDIT!

/*
Package OSMsg is a generated protocol buffer package.

It is generated from these files:
	OSMsg.proto

It has these top-level messages:
	OSMsg
*/
package OSMsg

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type OSMsg_MsgType int32

const (
	OSMsg_Logup      OSMsg_MsgType = 0
	OSMsg_Login      OSMsg_MsgType = 1
	OSMsg_Logout     OSMsg_MsgType = 2
	OSMsg_Individual OSMsg_MsgType = 3
	OSMsg_Group      OSMsg_MsgType = 4
)

var OSMsg_MsgType_name = map[int32]string{
	0: "Logup",
	1: "Login",
	2: "Logout",
	3: "Individual",
	4: "Group",
}
var OSMsg_MsgType_value = map[string]int32{
	"Logup":      0,
	"Login":      1,
	"Logout":     2,
	"Individual": 3,
	"Group":      4,
}

func (x OSMsg_MsgType) Enum() *OSMsg_MsgType {
	p := new(OSMsg_MsgType)
	*p = x
	return p
}
func (x OSMsg_MsgType) String() string {
	return proto.EnumName(OSMsg_MsgType_name, int32(x))
}
func (x *OSMsg_MsgType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(OSMsg_MsgType_value, data, "OSMsg_MsgType")
	if err != nil {
		return err
	}
	*x = OSMsg_MsgType(value)
	return nil
}

type OSMsg struct {
	Pattern          *OSMsg_MsgType   `protobuf:"varint,1,req,name=pattern,enum=OSMsg_MsgType,def=3" json:"pattern,omitempty"`
	Logup            *OSMsg_LogupMsg  `protobuf:"bytes,2,opt,name=logup" json:"logup,omitempty"`
	Login            *OSMsg_LoginMsg  `protobuf:"bytes,3,opt,name=login" json:"login,omitempty"`
	Logout           *OSMsg_LogoutMsg `protobuf:"bytes,4,opt,name=logout" json:"logout,omitempty"`
	Ind              *OSMsg_IndMsg    `protobuf:"bytes,5,opt,name=ind" json:"ind,omitempty"`
	Group            *OSMsg_GroupMsg  `protobuf:"bytes,6,opt,name=group" json:"group,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *OSMsg) Reset()         { *m = OSMsg{} }
func (m *OSMsg) String() string { return proto.CompactTextString(m) }
func (*OSMsg) ProtoMessage()    {}

const Default_OSMsg_Pattern OSMsg_MsgType = OSMsg_Individual

func (m *OSMsg) GetPattern() OSMsg_MsgType {
	if m != nil && m.Pattern != nil {
		return *m.Pattern
	}
	return Default_OSMsg_Pattern
}

func (m *OSMsg) GetLogup() *OSMsg_LogupMsg {
	if m != nil {
		return m.Logup
	}
	return nil
}

func (m *OSMsg) GetLogin() *OSMsg_LoginMsg {
	if m != nil {
		return m.Login
	}
	return nil
}

func (m *OSMsg) GetLogout() *OSMsg_LogoutMsg {
	if m != nil {
		return m.Logout
	}
	return nil
}

func (m *OSMsg) GetInd() *OSMsg_IndMsg {
	if m != nil {
		return m.Ind
	}
	return nil
}

func (m *OSMsg) GetGroup() *OSMsg_GroupMsg {
	if m != nil {
		return m.Group
	}
	return nil
}

type OSMsg_IndMsg struct {
	From             *int32  `protobuf:"varint,1,req,name=from" json:"from,omitempty"`
	To               *int32  `protobuf:"varint,2,req,name=to" json:"to,omitempty"`
	Msg              *string `protobuf:"bytes,3,req,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *OSMsg_IndMsg) Reset()         { *m = OSMsg_IndMsg{} }
func (m *OSMsg_IndMsg) String() string { return proto.CompactTextString(m) }
func (*OSMsg_IndMsg) ProtoMessage()    {}

func (m *OSMsg_IndMsg) GetFrom() int32 {
	if m != nil && m.From != nil {
		return *m.From
	}
	return 0
}

func (m *OSMsg_IndMsg) GetTo() int32 {
	if m != nil && m.To != nil {
		return *m.To
	}
	return 0
}

func (m *OSMsg_IndMsg) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

type OSMsg_LogupMsg struct {
	From             *int32  `protobuf:"varint,1,req,name=from" json:"from,omitempty"`
	To               *int32  `protobuf:"varint,2,req,name=to" json:"to,omitempty"`
	Msg              *string `protobuf:"bytes,3,req,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *OSMsg_LogupMsg) Reset()         { *m = OSMsg_LogupMsg{} }
func (m *OSMsg_LogupMsg) String() string { return proto.CompactTextString(m) }
func (*OSMsg_LogupMsg) ProtoMessage()    {}

func (m *OSMsg_LogupMsg) GetFrom() int32 {
	if m != nil && m.From != nil {
		return *m.From
	}
	return 0
}

func (m *OSMsg_LogupMsg) GetTo() int32 {
	if m != nil && m.To != nil {
		return *m.To
	}
	return 0
}

func (m *OSMsg_LogupMsg) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

type OSMsg_LoginMsg struct {
	From             *int32  `protobuf:"varint,1,req,name=from" json:"from,omitempty"`
	To               *int32  `protobuf:"varint,2,req,name=to" json:"to,omitempty"`
	Msg              *string `protobuf:"bytes,3,req,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *OSMsg_LoginMsg) Reset()         { *m = OSMsg_LoginMsg{} }
func (m *OSMsg_LoginMsg) String() string { return proto.CompactTextString(m) }
func (*OSMsg_LoginMsg) ProtoMessage()    {}

func (m *OSMsg_LoginMsg) GetFrom() int32 {
	if m != nil && m.From != nil {
		return *m.From
	}
	return 0
}

func (m *OSMsg_LoginMsg) GetTo() int32 {
	if m != nil && m.To != nil {
		return *m.To
	}
	return 0
}

func (m *OSMsg_LoginMsg) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

type OSMsg_LogoutMsg struct {
	From             *int32  `protobuf:"varint,1,req,name=from" json:"from,omitempty"`
	To               *int32  `protobuf:"varint,2,req,name=to" json:"to,omitempty"`
	Msg              *string `protobuf:"bytes,3,req,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *OSMsg_LogoutMsg) Reset()         { *m = OSMsg_LogoutMsg{} }
func (m *OSMsg_LogoutMsg) String() string { return proto.CompactTextString(m) }
func (*OSMsg_LogoutMsg) ProtoMessage()    {}

func (m *OSMsg_LogoutMsg) GetFrom() int32 {
	if m != nil && m.From != nil {
		return *m.From
	}
	return 0
}

func (m *OSMsg_LogoutMsg) GetTo() int32 {
	if m != nil && m.To != nil {
		return *m.To
	}
	return 0
}

func (m *OSMsg_LogoutMsg) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

type OSMsg_GroupMsg struct {
	From             *int32  `protobuf:"varint,1,req,name=from" json:"from,omitempty"`
	To               *int32  `protobuf:"varint,2,req,name=to" json:"to,omitempty"`
	Msg              *string `protobuf:"bytes,3,req,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *OSMsg_GroupMsg) Reset()         { *m = OSMsg_GroupMsg{} }
func (m *OSMsg_GroupMsg) String() string { return proto.CompactTextString(m) }
func (*OSMsg_GroupMsg) ProtoMessage()    {}

func (m *OSMsg_GroupMsg) GetFrom() int32 {
	if m != nil && m.From != nil {
		return *m.From
	}
	return 0
}

func (m *OSMsg_GroupMsg) GetTo() int32 {
	if m != nil && m.To != nil {
		return *m.To
	}
	return 0
}

func (m *OSMsg_GroupMsg) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

func init() {
	proto.RegisterEnum("OSMsg_MsgType", OSMsg_MsgType_name, OSMsg_MsgType_value)
}