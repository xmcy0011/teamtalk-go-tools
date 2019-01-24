package im_protocol

import (
	"bytes"
	"encoding/binary"
)

// PduHeader 协议头
type PduHeader struct {
	length    uint32 // the whole pdu length
	version   uint16 // pdu version number
	flag      uint16 // not used
	serviceID uint16 // 模块
	commandID uint16 // 消息ID
	seqNum    uint16 // 包序号
	reversed  uint16 // 保留
}

// GetPacket 添加数据包头部信息
func GetPacket(serverID uint16, msgID uint16, seqNum uint16, content []byte) []byte {
	length := (uint32)(len(content))
	var buf [16 + length]byte

	copy(buf, UInt32ToBytes(length))

	buf := UInt32ToBytes(length)            // 总长度
	buf = bytes.Join(buf, UInt16ToBytes(0)) //version
	buf = bytes.Join(buf, UInt16ToBytes(0)) //flag
	buf = bytes.Join(buf, UInt16ToBytes(serverID))
	buf = bytes.Join(buf, UInt16ToBytes(msgID))
	buf = bytes.Join(buf, UInt16ToBytes(seqNum))
	buf = bytes.Join(buf, UInt16ToBytes(0)) //reversed

}

// UInt32ToBytes 转换
func UInt32ToBytes(i uint32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, i)
	return buf
}

// UInt16ToBytes 转换
func UInt16ToBytes(i uint16) []byte {
	var buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, i)
	return buf
}

// BytesToUInt32 转换
func BytesToUInt32(buf []byte) uint32 {
	return binary.BigEndian.Uint32(buf)
}

// BytesToUInt16 转换
func BytesToUInt16(buf []byte) uint16 {
	return binary.BigEndian.Uint16(buf)
}
