// Code generated by protoc-gen-gogo.
// source: core.proto
// DO NOT EDIT!

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Metadata struct {
	DutyRX1     uint32  `protobuf:"varint,1,opt,name=DutyRX1,json=dutyRX1,proto3" json:"DutyRX1,omitempty"`
	DutyRX2     uint32  `protobuf:"varint,2,opt,name=DutyRX2,json=dutyRX2,proto3" json:"DutyRX2,omitempty"`
	Frequency   float32 `protobuf:"fixed32,3,opt,name=Frequency,json=frequency,proto3" json:"Frequency,omitempty"`
	DataRate    string  `protobuf:"bytes,4,opt,name=DataRate,json=dataRate,proto3" json:"DataRate,omitempty"`
	CodingRate  string  `protobuf:"bytes,5,opt,name=CodingRate,json=codingRate,proto3" json:"CodingRate,omitempty"`
	Timestamp   uint32  `protobuf:"varint,6,opt,name=Timestamp,json=timestamp,proto3" json:"Timestamp,omitempty"`
	Rssi        int32   `protobuf:"varint,7,opt,name=Rssi,json=rssi,proto3" json:"Rssi,omitempty"`
	Lsnr        float32 `protobuf:"fixed32,8,opt,name=Lsnr,json=lsnr,proto3" json:"Lsnr,omitempty"`
	PayloadSize uint32  `protobuf:"varint,9,opt,name=PayloadSize,json=payloadSize,proto3" json:"PayloadSize,omitempty"`
	Time        string  `protobuf:"bytes,10,opt,name=Time,json=time,proto3" json:"Time,omitempty"`
	RFChain     uint32  `protobuf:"varint,11,opt,name=RFChain,json=rFChain,proto3" json:"RFChain,omitempty"`
	CRCStatus   int32   `protobuf:"varint,12,opt,name=CRCStatus,json=cRCStatus,proto3" json:"CRCStatus,omitempty"`
	Modulation  string  `protobuf:"bytes,13,opt,name=Modulation,json=modulation,proto3" json:"Modulation,omitempty"`
	InvPolarity bool    `protobuf:"varint,14,opt,name=InvPolarity,json=invPolarity,proto3" json:"InvPolarity,omitempty"`
	Power       uint32  `protobuf:"varint,15,opt,name=Power,json=power,proto3" json:"Power,omitempty"`
	Channel     uint32  `protobuf:"varint,16,opt,name=Channel,json=channel,proto3" json:"Channel,omitempty"`
	GatewayEUI  string  `protobuf:"bytes,20,opt,name=GatewayEUI,json=gatewayEUI,proto3" json:"GatewayEUI,omitempty"`
	Altitude    int32   `protobuf:"varint,21,opt,name=Altitude,json=altitude,proto3" json:"Altitude,omitempty"`
	Longitude   float32 `protobuf:"fixed32,22,opt,name=Longitude,json=longitude,proto3" json:"Longitude,omitempty"`
	Latitude    float32 `protobuf:"fixed32,23,opt,name=Latitude,json=latitude,proto3" json:"Latitude,omitempty"`
	ServerTime  string  `protobuf:"bytes,31,opt,name=ServerTime,json=serverTime,proto3" json:"ServerTime,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptorCore, []int{0} }

type StatsMetadata struct {
	Altitude  int32   `protobuf:"varint,1,opt,name=Altitude,json=altitude,proto3" json:"Altitude,omitempty"`
	Longitude float32 `protobuf:"fixed32,2,opt,name=Longitude,json=longitude,proto3" json:"Longitude,omitempty"`
	Latitude  float32 `protobuf:"fixed32,3,opt,name=Latitude,json=latitude,proto3" json:"Latitude,omitempty"`
}

func (m *StatsMetadata) Reset()                    { *m = StatsMetadata{} }
func (m *StatsMetadata) String() string            { return proto.CompactTextString(m) }
func (*StatsMetadata) ProtoMessage()               {}
func (*StatsMetadata) Descriptor() ([]byte, []int) { return fileDescriptorCore, []int{1} }

func init() {
	proto.RegisterType((*Metadata)(nil), "core.Metadata")
	proto.RegisterType((*StatsMetadata)(nil), "core.StatsMetadata")
}
func (m *Metadata) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Metadata) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DutyRX1 != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintCore(data, i, uint64(m.DutyRX1))
	}
	if m.DutyRX2 != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintCore(data, i, uint64(m.DutyRX2))
	}
	if m.Frequency != 0 {
		data[i] = 0x1d
		i++
		i = encodeFixed32Core(data, i, uint32(math.Float32bits(float32(m.Frequency))))
	}
	if len(m.DataRate) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintCore(data, i, uint64(len(m.DataRate)))
		i += copy(data[i:], m.DataRate)
	}
	if len(m.CodingRate) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintCore(data, i, uint64(len(m.CodingRate)))
		i += copy(data[i:], m.CodingRate)
	}
	if m.Timestamp != 0 {
		data[i] = 0x30
		i++
		i = encodeVarintCore(data, i, uint64(m.Timestamp))
	}
	if m.Rssi != 0 {
		data[i] = 0x38
		i++
		i = encodeVarintCore(data, i, uint64(m.Rssi))
	}
	if m.Lsnr != 0 {
		data[i] = 0x45
		i++
		i = encodeFixed32Core(data, i, uint32(math.Float32bits(float32(m.Lsnr))))
	}
	if m.PayloadSize != 0 {
		data[i] = 0x48
		i++
		i = encodeVarintCore(data, i, uint64(m.PayloadSize))
	}
	if len(m.Time) > 0 {
		data[i] = 0x52
		i++
		i = encodeVarintCore(data, i, uint64(len(m.Time)))
		i += copy(data[i:], m.Time)
	}
	if m.RFChain != 0 {
		data[i] = 0x58
		i++
		i = encodeVarintCore(data, i, uint64(m.RFChain))
	}
	if m.CRCStatus != 0 {
		data[i] = 0x60
		i++
		i = encodeVarintCore(data, i, uint64(m.CRCStatus))
	}
	if len(m.Modulation) > 0 {
		data[i] = 0x6a
		i++
		i = encodeVarintCore(data, i, uint64(len(m.Modulation)))
		i += copy(data[i:], m.Modulation)
	}
	if m.InvPolarity {
		data[i] = 0x70
		i++
		if m.InvPolarity {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.Power != 0 {
		data[i] = 0x78
		i++
		i = encodeVarintCore(data, i, uint64(m.Power))
	}
	if m.Channel != 0 {
		data[i] = 0x80
		i++
		data[i] = 0x1
		i++
		i = encodeVarintCore(data, i, uint64(m.Channel))
	}
	if len(m.GatewayEUI) > 0 {
		data[i] = 0xa2
		i++
		data[i] = 0x1
		i++
		i = encodeVarintCore(data, i, uint64(len(m.GatewayEUI)))
		i += copy(data[i:], m.GatewayEUI)
	}
	if m.Altitude != 0 {
		data[i] = 0xa8
		i++
		data[i] = 0x1
		i++
		i = encodeVarintCore(data, i, uint64(m.Altitude))
	}
	if m.Longitude != 0 {
		data[i] = 0xb5
		i++
		data[i] = 0x1
		i++
		i = encodeFixed32Core(data, i, uint32(math.Float32bits(float32(m.Longitude))))
	}
	if m.Latitude != 0 {
		data[i] = 0xbd
		i++
		data[i] = 0x1
		i++
		i = encodeFixed32Core(data, i, uint32(math.Float32bits(float32(m.Latitude))))
	}
	if len(m.ServerTime) > 0 {
		data[i] = 0xfa
		i++
		data[i] = 0x1
		i++
		i = encodeVarintCore(data, i, uint64(len(m.ServerTime)))
		i += copy(data[i:], m.ServerTime)
	}
	return i, nil
}

func (m *StatsMetadata) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StatsMetadata) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Altitude != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintCore(data, i, uint64(m.Altitude))
	}
	if m.Longitude != 0 {
		data[i] = 0x15
		i++
		i = encodeFixed32Core(data, i, uint32(math.Float32bits(float32(m.Longitude))))
	}
	if m.Latitude != 0 {
		data[i] = 0x1d
		i++
		i = encodeFixed32Core(data, i, uint32(math.Float32bits(float32(m.Latitude))))
	}
	return i, nil
}

func encodeFixed64Core(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Core(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintCore(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Metadata) Size() (n int) {
	var l int
	_ = l
	if m.DutyRX1 != 0 {
		n += 1 + sovCore(uint64(m.DutyRX1))
	}
	if m.DutyRX2 != 0 {
		n += 1 + sovCore(uint64(m.DutyRX2))
	}
	if m.Frequency != 0 {
		n += 5
	}
	l = len(m.DataRate)
	if l > 0 {
		n += 1 + l + sovCore(uint64(l))
	}
	l = len(m.CodingRate)
	if l > 0 {
		n += 1 + l + sovCore(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovCore(uint64(m.Timestamp))
	}
	if m.Rssi != 0 {
		n += 1 + sovCore(uint64(m.Rssi))
	}
	if m.Lsnr != 0 {
		n += 5
	}
	if m.PayloadSize != 0 {
		n += 1 + sovCore(uint64(m.PayloadSize))
	}
	l = len(m.Time)
	if l > 0 {
		n += 1 + l + sovCore(uint64(l))
	}
	if m.RFChain != 0 {
		n += 1 + sovCore(uint64(m.RFChain))
	}
	if m.CRCStatus != 0 {
		n += 1 + sovCore(uint64(m.CRCStatus))
	}
	l = len(m.Modulation)
	if l > 0 {
		n += 1 + l + sovCore(uint64(l))
	}
	if m.InvPolarity {
		n += 2
	}
	if m.Power != 0 {
		n += 1 + sovCore(uint64(m.Power))
	}
	if m.Channel != 0 {
		n += 2 + sovCore(uint64(m.Channel))
	}
	l = len(m.GatewayEUI)
	if l > 0 {
		n += 2 + l + sovCore(uint64(l))
	}
	if m.Altitude != 0 {
		n += 2 + sovCore(uint64(m.Altitude))
	}
	if m.Longitude != 0 {
		n += 6
	}
	if m.Latitude != 0 {
		n += 6
	}
	l = len(m.ServerTime)
	if l > 0 {
		n += 2 + l + sovCore(uint64(l))
	}
	return n
}

func (m *StatsMetadata) Size() (n int) {
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovCore(uint64(m.Altitude))
	}
	if m.Longitude != 0 {
		n += 5
	}
	if m.Latitude != 0 {
		n += 5
	}
	return n
}

func sovCore(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCore(x uint64) (n int) {
	return sovCore(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Metadata) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCore
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DutyRX1", wireType)
			}
			m.DutyRX1 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.DutyRX1 |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DutyRX2", wireType)
			}
			m.DutyRX2 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.DutyRX2 |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frequency", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(data[iNdEx-4])
			v |= uint32(data[iNdEx-3]) << 8
			v |= uint32(data[iNdEx-2]) << 16
			v |= uint32(data[iNdEx-1]) << 24
			m.Frequency = float32(math.Float32frombits(v))
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCore
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataRate = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodingRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCore
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodingRate = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Timestamp |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rssi", wireType)
			}
			m.Rssi = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Rssi |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lsnr", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(data[iNdEx-4])
			v |= uint32(data[iNdEx-3]) << 8
			v |= uint32(data[iNdEx-2]) << 16
			v |= uint32(data[iNdEx-1]) << 24
			m.Lsnr = float32(math.Float32frombits(v))
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayloadSize", wireType)
			}
			m.PayloadSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.PayloadSize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCore
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Time = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RFChain", wireType)
			}
			m.RFChain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RFChain |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CRCStatus", wireType)
			}
			m.CRCStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.CRCStatus |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Modulation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCore
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Modulation = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InvPolarity", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InvPolarity = bool(v != 0)
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Power |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			m.Channel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Channel |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GatewayEUI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCore
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GatewayEUI = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 21:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Altitude", wireType)
			}
			m.Altitude = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Altitude |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 22:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Longitude", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(data[iNdEx-4])
			v |= uint32(data[iNdEx-3]) << 8
			v |= uint32(data[iNdEx-2]) << 16
			v |= uint32(data[iNdEx-1]) << 24
			m.Longitude = float32(math.Float32frombits(v))
		case 23:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Latitude", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(data[iNdEx-4])
			v |= uint32(data[iNdEx-3]) << 8
			v |= uint32(data[iNdEx-2]) << 16
			v |= uint32(data[iNdEx-1]) << 24
			m.Latitude = float32(math.Float32frombits(v))
		case 31:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCore
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerTime = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCore(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCore
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatsMetadata) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCore
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatsMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatsMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Altitude", wireType)
			}
			m.Altitude = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Altitude |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Longitude", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(data[iNdEx-4])
			v |= uint32(data[iNdEx-3]) << 8
			v |= uint32(data[iNdEx-2]) << 16
			v |= uint32(data[iNdEx-1]) << 24
			m.Longitude = float32(math.Float32frombits(v))
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Latitude", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(data[iNdEx-4])
			v |= uint32(data[iNdEx-3]) << 8
			v |= uint32(data[iNdEx-2]) << 16
			v |= uint32(data[iNdEx-1]) << 24
			m.Latitude = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipCore(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCore
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCore(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCore
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCore
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCore
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCore
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCore
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCore(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCore = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCore   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorCore = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x52, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0xa5, 0x5d, 0xba, 0x26, 0x2e, 0x85, 0xc9, 0x1a, 0x70, 0x85, 0x50, 0x99, 0xf6, 0xc4, 0x13,
	0x12, 0xe3, 0x0b, 0x20, 0x63, 0x68, 0x52, 0x27, 0x55, 0x2e, 0x48, 0xbc, 0x9a, 0xc4, 0x14, 0x4b,
	0xae, 0x1d, 0x1c, 0x67, 0x53, 0xf8, 0x12, 0x3e, 0x89, 0x47, 0x1e, 0xf8, 0x00, 0x04, 0x3f, 0x82,
	0xaf, 0x9d, 0xa4, 0xdb, 0x4b, 0x1f, 0x22, 0xf9, 0x9c, 0x93, 0x7b, 0xcf, 0xb9, 0xd7, 0x26, 0xa4,
	0x30, 0x56, 0xbc, 0xac, 0xac, 0x71, 0x86, 0x26, 0x78, 0x3e, 0xfd, 0x9d, 0x90, 0xf4, 0x4a, 0x38,
	0x5e, 0x72, 0xc7, 0x29, 0x90, 0xe9, 0x79, 0xe3, 0x5a, 0xf6, 0xe9, 0x15, 0x8c, 0x4e, 0x46, 0x2f,
	0xe6, 0x6c, 0x5a, 0x46, 0xb8, 0x53, 0xce, 0x60, 0x7c, 0x5b, 0x39, 0xa3, 0xcf, 0x48, 0x76, 0x61,
	0xc5, 0xb7, 0x46, 0xe8, 0xa2, 0x85, 0x03, 0xaf, 0x8d, 0x59, 0xf6, 0xa5, 0x27, 0xe8, 0x53, 0x92,
	0x9e, 0xfb, 0xce, 0x8c, 0x3b, 0x01, 0x89, 0x17, 0x33, 0x96, 0x96, 0x1d, 0xa6, 0x0b, 0x42, 0x72,
	0x53, 0x4a, 0xbd, 0x09, 0xea, 0x24, 0xa8, 0x3e, 0x60, 0xcf, 0x60, 0xe7, 0x0f, 0x72, 0x2b, 0x6a,
	0xc7, 0xb7, 0x15, 0x1c, 0x06, 0xd7, 0xcc, 0xf5, 0x04, 0xa5, 0x24, 0x61, 0x75, 0x2d, 0x61, 0xea,
	0x85, 0x09, 0x4b, 0xac, 0x3f, 0x23, 0xb7, 0xac, 0xb5, 0x85, 0x34, 0xc4, 0x48, 0x94, 0x3f, 0xd3,
	0x13, 0x32, 0x5b, 0xf1, 0x56, 0x19, 0x5e, 0xae, 0xe5, 0x77, 0x01, 0x59, 0xe8, 0x33, 0xab, 0x76,
	0x14, 0x56, 0xa1, 0x0f, 0x90, 0x90, 0x20, 0x41, 0x0b, 0x9c, 0x97, 0x5d, 0xe4, 0x5f, 0xb9, 0xd4,
	0x30, 0x8b, 0xf3, 0xda, 0x08, 0x31, 0x55, 0xce, 0xf2, 0xb5, 0xe3, 0xae, 0xa9, 0xe1, 0x7e, 0x30,
	0xcf, 0x8a, 0x9e, 0xc0, 0x99, 0xae, 0x4c, 0xd9, 0x28, 0xee, 0xa4, 0xd1, 0x30, 0x8f, 0x33, 0x6d,
	0x07, 0x06, 0xd3, 0x5c, 0xea, 0xeb, 0x95, 0x51, 0xdc, 0x4a, 0xd7, 0xc2, 0x03, 0xff, 0x43, 0xca,
	0x66, 0x72, 0x47, 0xd1, 0x63, 0x32, 0x59, 0x99, 0x1b, 0x61, 0xe1, 0x61, 0xf0, 0x9d, 0x54, 0x08,
	0x30, 0x8f, 0xb7, 0xd7, 0x5a, 0x28, 0x38, 0x8a, 0x79, 0x8a, 0x08, 0xd1, 0xf1, 0xbd, 0xdf, 0xd6,
	0x0d, 0x6f, 0xdf, 0x7d, 0xbc, 0x84, 0xe3, 0xe8, 0xb8, 0x19, 0x18, 0xbc, 0x81, 0x37, 0xca, 0x49,
	0xd7, 0x94, 0x02, 0x1e, 0x85, 0xb8, 0x29, 0xef, 0x30, 0xce, 0xb2, 0x34, 0x7a, 0x13, 0xc5, 0xc7,
	0xf1, 0xee, 0x54, 0x4f, 0x60, 0xe5, 0x92, 0x77, 0x95, 0x4f, 0x82, 0x98, 0xaa, 0x0e, 0xa3, 0xeb,
	0x5a, 0xd8, 0x6b, 0x61, 0xc3, 0xe6, 0x9e, 0x47, 0xd7, 0x7a, 0x60, 0x4e, 0x05, 0x99, 0xe3, 0x46,
	0xea, 0xe1, 0x69, 0xdd, 0x8e, 0x31, 0xda, 0x17, 0x63, 0xbc, 0x2f, 0xc6, 0xc1, 0xdd, 0x18, 0x6f,
	0x8f, 0x7e, 0xfe, 0x5d, 0x8c, 0x7e, 0xf9, 0xef, 0x8f, 0xff, 0x7e, 0xfc, 0x5b, 0xdc, 0xfb, 0x7c,
	0x18, 0x1e, 0xf7, 0xeb, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x23, 0x2c, 0xd8, 0xcc, 0xea, 0x02,
	0x00, 0x00,
}
