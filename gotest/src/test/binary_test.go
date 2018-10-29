package test

import (
	"testing"
	"bytes"
	"math"
	"encoding/binary"
)

func TestBinary00(t *testing.T) {
	buf := new(bytes.Buffer)
	var pi float64 = math.Pi
	err := binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		t.Log("binary.Write failed:", err)
	}
	t.Logf("% x", buf.Bytes())
}

func TestBinary01(t *testing.T) {
	buf := new(bytes.Buffer)
	var pi float64 = math.Pi
	err := binary.Write(buf, binary.BigEndian, pi)
	if err != nil {
		t.Log("binary.Write failed:", err)
	}
	t.Logf("% x", buf.Bytes())
}

func TestBinary1(t *testing.T) {
	buf := new(bytes.Buffer)
	var data = []interface{}{
		uint16(61374),
		int8(-54),
		uint8(254),
	}
	for _, v := range data {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			t.Log("binary.Write failed:", err)
		}
	}
	arr := buf.Bytes()
	t.Logf("%s, %x, %q", arr, arr, arr)
}

func TestBinary2(t *testing.T) {
	buf := new(bytes.Buffer)
	var data = []interface{}{
		uint16(61374),
		int8(-54),
		uint8(254),
	}
	for _, v := range data {
		err := binary.Write(buf, binary.BigEndian, v)
		if err != nil {
			t.Log("binary.Write failed:", err)
		}
	}
	arr := buf.Bytes()
	t.Logf("%s, %x, %q", arr, arr, arr)
}

func TestBinary3(t *testing.T) {
	t.Log(binary.BigEndian.Uint16([]byte{0x01, 0x1f}))
	t.Log(binary.LittleEndian.Uint16([]byte{0x1f, 0x01}))
	t.Log(binary.BigEndian.Uint16([]byte{0x1f, 0x01}))
	t.Log(binary.LittleEndian.Uint16([]byte{0x01, 0x1f}))
}