package tcp

import (
	"bytes"
	"encoding/binary"

	"github.com/panjf2000/gnet/v2"
	"github.com/pkg/errors"
)

const (
	defaultMagicLen  = 2
	defaultHeaderLen = 4
)

var (
	ErrIncompletePacket = errors.New("incomplete packet")
	ErrTooLargePacket   = errors.New("too large packet")
)

type Codec struct {
	byteOrder  binary.ByteOrder
	magic      []byte
	magicLen   int
	headerLen  int
	maxBodyLen uint32
}

func NewCodec(isBigEndian bool, magic uint16, maxBodyLen uint32) *Codec {
	codec := &Codec{
		magicLen:   defaultMagicLen,
		headerLen:  defaultHeaderLen,
		maxBodyLen: maxBodyLen,
	}
	if isBigEndian {
		codec.byteOrder = binary.BigEndian
	} else {
		codec.byteOrder = binary.LittleEndian
	}
	codec.magic = make([]byte, codec.magicLen)
	codec.byteOrder.PutUint16(codec.magic, magic)
	return codec
}

func (slf *Codec) Encode(data []byte) ([]byte, error) {
	offset := slf.magicLen + slf.headerLen
	length := offset + len(data)
	packet := make([]byte, length)
	copy(packet[0:slf.magicLen], slf.magic)
	slf.byteOrder.PutUint32(packet[slf.magicLen:offset], uint32(len(data)))
	copy(packet[offset:length], data)
	return packet, nil
}

func (slf *Codec) Decode(gn gnet.Conn) ([]byte, error) {
	offset := slf.magicLen + slf.headerLen
	buf, _ := gn.Peek(offset)
	if len(buf) < offset {
		return nil, ErrIncompletePacket
	}
	if !bytes.Equal(slf.magic, buf[:slf.magicLen]) {
		return nil, errors.Errorf("magic number not match, received magic number: %v", buf[:slf.magicLen])
	}
	bodyLen := slf.byteOrder.Uint32(buf[slf.magicLen:offset])
	length := offset + int(bodyLen)
	if gn.InboundBuffered() < length {
		return nil, ErrIncompletePacket
	}
	if slf.maxBodyLen > 0 && bodyLen > slf.maxBodyLen {
		_, _ = gn.Discard(length)
		return nil, errors.WithMessagef(ErrTooLargePacket, "body length: %d, max body length: %d", bodyLen, slf.maxBodyLen)
	}
	buf, _ = gn.Peek(length)
	_, _ = gn.Discard(length)
	return buf[offset:length], nil
}

func (slf *Codec) Unpack(buf []byte) ([]byte, error) {
	offset := slf.magicLen + slf.headerLen
	if len(buf) < offset {
		return nil, ErrIncompletePacket
	}
	if !bytes.Equal(slf.magic, buf[:slf.magicLen]) {
		return nil, errors.Errorf("magic number not match, received magic number: %v", buf[:slf.magicLen])
	}
	bodyLen := slf.byteOrder.Uint32(buf[slf.magicLen:offset])
	length := offset + int(bodyLen)
	if len(buf) < length {
		return nil, ErrIncompletePacket
	}
	return buf[offset:length], nil
}
