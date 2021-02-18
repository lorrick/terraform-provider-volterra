// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/log/console_login_events/types.proto

/*
	Package console_login_events is a generated protocol buffer package.

	It is generated from these files:
		ves.io/schema/log/console_login_events/types.proto

	It has these top-level messages:
		DateAggregation
		AggregationRequest
*/
package console_login_events

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Date Aggregation
//
// x-displayName: Date Aggregation
// Aggregate console login events based on timestamp in the log
type DateAggregation struct {
	// step
	//
	// x-displayName: "Step"
	// x-required
	// x-example: "5m"
	//
	// step is the resolution width, which determines the number of the data points [x-axis (time)] to be returned in the response.
	// The timestamps in the response will be t1=start_time, t2=t1+step, ... tn=tn-1+step, where tn <= end_time.
	// Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days
	Step string `protobuf:"bytes,1,opt,name=step,proto3" json:"step,omitempty"`
}

func (m *DateAggregation) Reset()                    { *m = DateAggregation{} }
func (*DateAggregation) ProtoMessage()               {}
func (*DateAggregation) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

func (m *DateAggregation) GetStep() string {
	if m != nil {
		return m.Step
	}
	return ""
}

// Aggregation Request
//
// x-displayName: "Aggregation Request"
// Aggregation request to provide analytics data over the log response
type AggregationRequest struct {
	// aggregation type
	//
	// x-displayName: "Aggregation Type"
	// Specify one of the aggregation types
	//
	// Types that are valid to be assigned to AggregationType:
	//	*AggregationRequest_DateAggregation
	AggregationType isAggregationRequest_AggregationType `protobuf_oneof:"aggregation_type"`
}

func (m *AggregationRequest) Reset()                    { *m = AggregationRequest{} }
func (*AggregationRequest) ProtoMessage()               {}
func (*AggregationRequest) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

type isAggregationRequest_AggregationType interface {
	isAggregationRequest_AggregationType()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type AggregationRequest_DateAggregation struct {
	DateAggregation *DateAggregation `protobuf:"bytes,1,opt,name=date_aggregation,json=dateAggregation,oneof"`
}

func (*AggregationRequest_DateAggregation) isAggregationRequest_AggregationType() {}

func (m *AggregationRequest) GetAggregationType() isAggregationRequest_AggregationType {
	if m != nil {
		return m.AggregationType
	}
	return nil
}

func (m *AggregationRequest) GetDateAggregation() *DateAggregation {
	if x, ok := m.GetAggregationType().(*AggregationRequest_DateAggregation); ok {
		return x.DateAggregation
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AggregationRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AggregationRequest_OneofMarshaler, _AggregationRequest_OneofUnmarshaler, _AggregationRequest_OneofSizer, []interface{}{
		(*AggregationRequest_DateAggregation)(nil),
	}
}

func _AggregationRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AggregationRequest)
	// aggregation_type
	switch x := m.AggregationType.(type) {
	case *AggregationRequest_DateAggregation:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DateAggregation); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AggregationRequest.AggregationType has unexpected type %T", x)
	}
	return nil
}

func _AggregationRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AggregationRequest)
	switch tag {
	case 1: // aggregation_type.date_aggregation
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DateAggregation)
		err := b.DecodeMessage(msg)
		m.AggregationType = &AggregationRequest_DateAggregation{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AggregationRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AggregationRequest)
	// aggregation_type
	switch x := m.AggregationType.(type) {
	case *AggregationRequest_DateAggregation:
		s := proto.Size(x.DateAggregation)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*DateAggregation)(nil), "ves.io.schema.log.console_login_events.DateAggregation")
	proto.RegisterType((*AggregationRequest)(nil), "ves.io.schema.log.console_login_events.AggregationRequest")
}
func (this *DateAggregation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DateAggregation)
	if !ok {
		that2, ok := that.(DateAggregation)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Step != that1.Step {
		return false
	}
	return true
}
func (this *AggregationRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AggregationRequest)
	if !ok {
		that2, ok := that.(AggregationRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.AggregationType == nil {
		if this.AggregationType != nil {
			return false
		}
	} else if this.AggregationType == nil {
		return false
	} else if !this.AggregationType.Equal(that1.AggregationType) {
		return false
	}
	return true
}
func (this *AggregationRequest_DateAggregation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AggregationRequest_DateAggregation)
	if !ok {
		that2, ok := that.(AggregationRequest_DateAggregation)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.DateAggregation.Equal(that1.DateAggregation) {
		return false
	}
	return true
}
func (this *DateAggregation) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&console_login_events.DateAggregation{")
	s = append(s, "Step: "+fmt.Sprintf("%#v", this.Step)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *AggregationRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&console_login_events.AggregationRequest{")
	if this.AggregationType != nil {
		s = append(s, "AggregationType: "+fmt.Sprintf("%#v", this.AggregationType)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *AggregationRequest_DateAggregation) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&console_login_events.AggregationRequest_DateAggregation{` +
		`DateAggregation:` + fmt.Sprintf("%#v", this.DateAggregation) + `}`}, ", ")
	return s
}
func valueToGoStringTypes(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *DateAggregation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DateAggregation) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Step) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Step)))
		i += copy(dAtA[i:], m.Step)
	}
	return i, nil
}

func (m *AggregationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AggregationRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AggregationType != nil {
		nn1, err := m.AggregationType.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *AggregationRequest_DateAggregation) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.DateAggregation != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTypes(dAtA, i, uint64(m.DateAggregation.Size()))
		n2, err := m.DateAggregation.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedDateAggregation(r randyTypes, easy bool) *DateAggregation {
	this := &DateAggregation{}
	this.Step = string(randStringTypes(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedAggregationRequest(r randyTypes, easy bool) *AggregationRequest {
	this := &AggregationRequest{}
	oneofNumber_AggregationType := []int32{1}[r.Intn(1)]
	switch oneofNumber_AggregationType {
	case 1:
		this.AggregationType = NewPopulatedAggregationRequest_DateAggregation(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedAggregationRequest_DateAggregation(r randyTypes, easy bool) *AggregationRequest_DateAggregation {
	this := &AggregationRequest_DateAggregation{}
	this.DateAggregation = NewPopulatedDateAggregation(r, easy)
	return this
}

type randyTypes interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTypes(r randyTypes) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTypes(r randyTypes) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneTypes(r)
	}
	return string(tmps)
}
func randUnrecognizedTypes(r randyTypes, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTypes(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTypes(dAtA []byte, r randyTypes, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTypes(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *DateAggregation) Size() (n int) {
	var l int
	_ = l
	l = len(m.Step)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *AggregationRequest) Size() (n int) {
	var l int
	_ = l
	if m.AggregationType != nil {
		n += m.AggregationType.Size()
	}
	return n
}

func (m *AggregationRequest_DateAggregation) Size() (n int) {
	var l int
	_ = l
	if m.DateAggregation != nil {
		l = m.DateAggregation.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *DateAggregation) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DateAggregation{`,
		`Step:` + fmt.Sprintf("%v", this.Step) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AggregationRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AggregationRequest{`,
		`AggregationType:` + fmt.Sprintf("%v", this.AggregationType) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AggregationRequest_DateAggregation) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AggregationRequest_DateAggregation{`,
		`DateAggregation:` + strings.Replace(fmt.Sprintf("%v", this.DateAggregation), "DateAggregation", "DateAggregation", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTypes(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *DateAggregation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DateAggregation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DateAggregation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Step", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Step = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
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
func (m *AggregationRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AggregationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AggregationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DateAggregation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DateAggregation{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.AggregationType = &AggregationRequest_DateAggregation{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipTypes(dAtA[start:])
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
	ErrInvalidLengthTypes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("ves.io/schema/log/console_login_events/types.proto", fileDescriptorTypes)
}

var fileDescriptorTypes = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0x6b, 0xdb, 0x40,
	0x14, 0xc7, 0xf5, 0x5c, 0x53, 0x5c, 0x75, 0xb0, 0xd1, 0x54, 0x5c, 0x73, 0x18, 0x43, 0x4b, 0x17,
	0xdf, 0x81, 0x3b, 0x74, 0xae, 0xe9, 0x50, 0x3a, 0x7a, 0x2c, 0x05, 0x21, 0x59, 0xcf, 0x67, 0x51,
	0x59, 0x4f, 0xd5, 0x9d, 0xd4, 0x76, 0xeb, 0x47, 0xe8, 0xdc, 0xa1, 0x73, 0x3e, 0x42, 0xc0, 0x8b,
	0xc7, 0x8c, 0x1e, 0x3d, 0x46, 0x97, 0x25, 0xa3, 0xc7, 0x8c, 0xc1, 0x67, 0x87, 0x58, 0x26, 0x90,
	0x6c, 0xef, 0xe9, 0xa7, 0xdf, 0xfd, 0x8f, 0x3f, 0xe7, 0x8e, 0x4a, 0x54, 0x3c, 0x26, 0xa1, 0xa6,
	0x73, 0x5c, 0x04, 0x22, 0x21, 0x29, 0xa6, 0x94, 0x2a, 0x4a, 0xd0, 0x4f, 0x48, 0xc6, 0xa9, 0x8f,
	0x25, 0xa6, 0x5a, 0x09, 0xfd, 0x3b, 0x43, 0xc5, 0xb3, 0x9c, 0x34, 0x79, 0x6f, 0xf7, 0x0e, 0xdf,
	0x3b, 0x3c, 0x21, 0xc9, 0x1f, 0x72, 0xba, 0x43, 0x19, 0xeb, 0x79, 0x11, 0xf2, 0x29, 0x2d, 0x84,
	0x24, 0x49, 0xc2, 0xea, 0x61, 0x31, 0xb3, 0x9b, 0x5d, 0xec, 0xb4, 0x3f, 0xb6, 0xfb, 0xba, 0x7e,
	0x15, 0xca, 0x74, 0x4c, 0xe9, 0x21, 0xb3, 0xdb, 0xab, 0xc3, 0x32, 0x48, 0xe2, 0x28, 0xd0, 0x78,
	0xa0, 0xfd, 0x13, 0x1a, 0xe3, 0x4f, 0xbf, 0xe6, 0x0f, 0xde, 0xb8, 0xed, 0x4f, 0x81, 0xc6, 0x8f,
	0x52, 0xe6, 0x28, 0x83, 0x1d, 0xf1, 0x3c, 0xb7, 0xa9, 0x34, 0x66, 0xaf, 0xa0, 0x0f, 0xef, 0x5e,
	0x4c, 0xec, 0x3c, 0xf8, 0x0f, 0xae, 0x77, 0xf4, 0xcf, 0x04, 0x7f, 0x14, 0xa8, 0xb4, 0x17, 0xb9,
	0x9d, 0x5d, 0x9a, 0x1f, 0xdc, 0x23, 0xab, 0xbd, 0x1c, 0x7d, 0xe0, 0x4f, 0x2b, 0x83, 0x9f, 0xa4,
	0x7f, 0x76, 0x26, 0xed, 0xa8, 0xfe, 0x69, 0xdc, 0x73, 0x3b, 0x47, 0x01, 0xfe, 0xae, 0x72, 0xaf,
	0xb5, 0x5a, 0x02, 0xac, 0x97, 0xd0, 0xf8, 0xd2, 0x6c, 0x35, 0x3a, 0xcf, 0xc6, 0xff, 0x60, 0x5d,
	0x31, 0x67, 0x53, 0x31, 0x67, 0x5b, 0x31, 0xb8, 0xa9, 0x18, 0xfc, 0x31, 0x0c, 0xce, 0x0c, 0x83,
	0x73, 0xc3, 0x60, 0x65, 0x18, 0x5c, 0x18, 0x06, 0x6b, 0xc3, 0x60, 0x63, 0x18, 0x5c, 0x1a, 0x06,
	0xd7, 0x86, 0x39, 0x5b, 0xc3, 0xe0, 0xef, 0x15, 0x73, 0xbe, 0x7e, 0x93, 0x94, 0x7d, 0x97, 0xbc,
	0xa4, 0x44, 0x63, 0x9e, 0x07, 0xbc, 0x50, 0xc2, 0x0e, 0x33, 0xca, 0x17, 0xc3, 0x2c, 0xa7, 0x32,
	0x8e, 0x30, 0x1f, 0xde, 0x61, 0x91, 0x85, 0x92, 0x04, 0xfe, 0xd2, 0x87, 0x5e, 0x1f, 0x79, 0x24,
	0xe1, 0x73, 0xdb, 0xf5, 0xfb, 0xdb, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x82, 0xab, 0x72, 0x55,
	0x02, 0x00, 0x00,
}
