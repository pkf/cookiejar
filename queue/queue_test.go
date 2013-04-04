// CookieJar - A contestant's algorithm toolbox
// Copyright (c) 2013 Peter Szilagyi. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice,
//       this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright
//       notice, this list of conditions and the following disclaimer in the
//       documentation and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.
//
// Alternatively, the CookieJar toolbox may be used in accordance with the terms
// and conditions contained in a signed written agreement between you and the
// author(s).
//
// Author: peterke@gmail.com (Peter Szilagyi)
package queue

import (
	"math/rand"
	"testing"
)

func TestQueue(t *testing.T) {
	// Create some initial data
	size := 1048576
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}
	queue := New()
	for rep := 0; rep < 2; rep++ {
		// Push all the data into the queue, pop out every second, then the rest
		outs := []int{}
		for i := 0; i < size; i++ {
			queue.Push(data[i])
			if i%2 == 0 {
				outs = append(outs, queue.Pop().(int))
			}
		}
		for !queue.Empty() {
			outs = append(outs, queue.Pop().(int))
		}
		// Make sure the contents of the resulting slices are ok
		for i := 0; i < size; i++ {
			if data[i] != outs[i] {
				t.Errorf("push/pop mismatch: have %v, want %v.", outs[i], data[i])
			}
		}
	}
}

func TestReset(t *testing.T) {
	// Push some stuff into the queue
	size := 1048576
	queue := New()
	for i := 0; i < size; i++ {
		queue.Push(i)
	}
	// Clear and verify
	queue.Reset()
	if !queue.Empty() {
		t.Errorf("queue not empty after reset: %v", queue)
	}
}

func BenchmarkPush(b *testing.B) {
	queue := New()
	for i := 0; i < b.N; i++ {
		queue.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	queue := New()
	for i := 0; i < b.N; i++ {
		queue.Push(i)
	}
	b.ResetTimer()
	for !queue.Empty() {
		queue.Pop()
	}
}
