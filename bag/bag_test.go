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
package bag

import (
	"math/rand"
	"testing"
)

func TestBag(t *testing.T) {
	// Create some initial data
	size := 1048576
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}
	// Insert the data into the bag, but remove every second
	bag := New()
	for i := 0; i < len(data); i++ {
		bag.Insert(data[i])
		if i%2 == 0 {
			bag.Remove(data[i])
		}
	}
	// Calculate the sum of the elements in and out
	sumBag := int64(0)
	bag.Do(func(val interface{}) {
		sumBag += int64(val.(int))
	})
	sumDat := int64(0)
	for i := 1; i < len(data); i += 2 {
		sumDat += int64(data[i])
	}
	if sumBag != sumDat {
		t.Errorf("sum mismatch after iteration: have %v, want %v", sumBag, sumDat)
	}
	// Verify the contents of the bag
	for i := 1; i < len(data); i += 2 {
		if bag.Count(data[i]) <= 0 {
			t.Errorf("expected data, none found: %v in %v", data[i], bag)
		}
		bag.Remove(data[i])
	}
	if len(bag.data) != 0 {
		t.Errorf("leftovers remained in bag: %v", bag)
	}
}

func TestReset(t *testing.T) {
	// Create some initial data
	size := 1048576
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}
	// Insert the data into the bag, but remove every second
	bag := New()
	for val := range data {
		bag.Insert(val)
	}
	// clear the bag and verify
	bag.Reset()
	if len(bag.data) != 0 {
		t.Errorf("leftovers remained in bag: %v", bag)
	}
}

func BenchmarkInsert(b *testing.B) {
	// Create some initial data
	data := make([]int, b.N)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int()
	}
	// Execute the benchmark
	b.ResetTimer()
	bag := New()
	for val := range data {
		bag.Insert(val)
	}
}

func BenchmarkRemove(b *testing.B) {
	// Create some initial data
	data := make([]int, b.N)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int()
	}
	// Fill the bag with it
	bag := New()
	for val := range data {
		bag.Insert(val)
	}
	// Execute the benchmark
	b.ResetTimer()
	for val := range data {
		bag.Remove(val)
	}
}

func BenchmarkDo(b *testing.B) {
	// Create some initial data
	data := make([]int, b.N)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int()
	}
	// Fill the bag with it
	bag := New()
	for val := range data {
		bag.Insert(val)
	}
	// Execute the benchmark
	b.ResetTimer()
	bag.Do(func(val interface{}) {
		// Do nothing
	})
}
