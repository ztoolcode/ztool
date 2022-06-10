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

/**
  @author: fanyanan
  @date: 2022/6/9
  @note: //id Tool Provides UUID generation method ObjectId generation method Snowflake's Id generation method
**/
package id

import (
	"errors"
	"sync"
)

type Id struct {
	n Node
}

var Epoch int64 = 1288834974657 // timestamp 2006-03-21:20:50:14 GMT

const (
	nodeBits  uint8 = 10                    // The number of bits in the node ID
	stepBits  uint8 = 12                    // Number of digits in the serial number
	nodeMax   int64 = -1 ^ (-1 << nodeBits) // The maximum value of the node ID, used to detect overflow
	stepMax   int64 = -1 ^ (-1 << stepBits) // The maximum value of the sequence number, used to detect overflow
	timeShift uint8 = nodeBits + stepBits   // offset to the left of the timestamp
	nodeShift uint8 = stepBits              // Offset of node ID to the left
)

type Node struct {
	mu        sync.Mutex // Add mutex to ensure concurrency safety
	timestamp int64      // timestamp part
	node      int64      // Node ID section
	step      int64      // Serial number ID part
}

func newNode(node int64) (*Node, error) {
	// If the maximum range of the node is exceeded, an error is generated
	if node < 0 || node > nodeMax {
		return nil, errors.New("Node number must be between 0 and 1023")
	}
	// Generate and return a pointer to a node instance
	return &Node{
		timestamp: 0,
		node:      node,
		step:      0,
	}, nil
}
