// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

// go test *.go -bench=".*" -count=1

package grpool_test

import (
    "testing"
    "gitee.com/johng/gf/g/os/grpool"
)

var n = 500000

func increment2() {
    for i := 0; i < 1000000; i++ {}
}

func BenchmarkGrpool2(b *testing.B) {
    b.N = n
    for i := 0; i < b.N; i++ {
        grpool.Add(increment)
    }
}

func BenchmarkGoroutine2(b *testing.B) {
    b.N = n
    for i := 0; i < b.N; i++ {
        go increment()
    }
}