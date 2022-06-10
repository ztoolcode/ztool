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
  @note: //specific method
**/
package id

import (
	"strings"
	"time"
)

// RandomUUID The generated UUID is a string with -, similar to:a5c8a5e8-df2b-4706-bea4-08d0939410e3
func (i *Id) RandomUUID() (string, error) {
	v4, err := NewV4()
	if err != nil {
		return "", err
	}
	return v4.String(), nil
}

// SimpleUUID The resulting string is without the -, similar to:a08d1ec9a18e445d833762c3bb21316b
func (i *Id) SimpleUUID() (string, error) {
	v4, err := NewV4()
	if err != nil {
		return "", err
	}
	s := v4.String()
	return strings.Replace(s, "-", "", -1), nil
}

// GenerateSnowflakeId Algorithmically generate IDs
func (i *Id) GenerateSnowflakeId() int64 {
	i.n.mu.Lock()         // Ensure concurrency safety, lock
	defer i.n.mu.Unlock() // Unlock after the method has finished running
	// Get the timestamp of the current time (display in milliseconds)
	now := time.Now().UnixNano() / 1e6
	if i.n.timestamp == now {
		// step step 1
		i.n.step++
		// The current step is used up
		if i.n.step > stepMax {
			// wait for this millisecond to end
			for now <= i.n.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		// Exhausted step in this millisecond
		i.n.step = 0
	}
	i.n.timestamp = now
	// Shift operation to produce final ID
	result := (now-Epoch)<<timeShift | (i.n.node << nodeShift) | (i.n.step)
	return result
}
