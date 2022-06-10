//MIT License
//
//Copyright (c) 2022 druidcaesa
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

package net

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
)

func Export(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, b, "", "\t")
	if err != nil {
		return ""
	}
	return buf.String()
}

func Json(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

func IntByte(v interface{}) []byte {
	b := bytes.NewBuffer([]byte{})
	switch v.(type) {
	case int:
		binary.Write(b, binary.BigEndian, int64(v.(int)))
	case int8:
		binary.Write(b, binary.BigEndian, v.(int8))
	case int16:
		binary.Write(b, binary.BigEndian, v.(int16))
	case int32:
		binary.Write(b, binary.BigEndian, v.(int32))
	case int64:
		binary.Write(b, binary.BigEndian, v.(int64))
	case uint:
		binary.Write(b, binary.BigEndian, uint64(v.(uint)))
	case uint8:
		binary.Write(b, binary.BigEndian, v.(uint8))
	case uint16:
		binary.Write(b, binary.BigEndian, v.(uint16))
	case uint32:
		binary.Write(b, binary.BigEndian, v.(uint32))
	case uint64:
		binary.Write(b, binary.BigEndian, v.(uint64))
	}
	return b.Bytes()
}
