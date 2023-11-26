package tcppkt

import (
	"bytes"
	"encoding/binary"
	"io"
)

/**
 * @Author: admin
 * @Description:  封包格式及发送
 * @File: packet
 * @Version: 1.0.0
 * @Date: 2023/11/26 20:27
 */

// 二进制封包格式

type Packet struct {
	Size uint16 // 包体大小
	Body []byte // 包体数据
}

// 将数据写入dataWrWriter
func writePacket(dataWriter io.Writer, data []byte) error {

	// 准备一个字节数组缓冲
	var buffer bytes.Buffer

	// 将Size写入缓冲
	if err := binary.Write(&buffer, binary.LittleEndian, uint16(len(data))); err != nil {
		return err
	}

	// 写入包体数据
	if _, err := buffer.Write(data); err != nil {
		return err
	}

	// 获得写入的完整数据
	out := buffer.Bytes()

	// 写入dataWriter
	if _, err := dataWriter.Write(out); err != nil {
		return err
	}
	return nil
}
