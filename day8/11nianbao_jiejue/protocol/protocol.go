package protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

//Encode 将消息编码
func Encode(message string) ([]byte, error) {
	//读取消息的长度，转换为int32
	length := int32(len(message))
	pkg := new(bytes.Buffer)
	//写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	//写入消息体
	err1 := binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err1 != nil {
		return nil, err1
	}
	return pkg.Bytes(), err
}

//Decode 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	//读取消息的长度
	peek, err := reader.Peek(4)
	if err != nil {
		return "", err
	} //读取前四个字节的数据
	buffer := bytes.NewBuffer(peek)
	var length int32
	err1 := binary.Read(buffer, binary.LittleEndian, &length)
	if err1 != nil {
		return "", err1
	}
	//Buffered 返回缓冲区中现有的可读取的字节数
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	//读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err2 := reader.Read(pack)
	if err2 != nil {
		return "", err2
	}

	return string(pack[4:]), nil

}
