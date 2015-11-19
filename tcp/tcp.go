// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tcp

import (
	"math/rand"
	"net"

	"github.com/mikioh/tcp"
)

func Fuzz(data []byte) int {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0
	}
	defer ln.Close()
	c, err := net.Dial(ln.Addr().Network(), ln.Addr().String())
	if err != nil {
		return 0
	}
	defer c.Close()
	tc, err := tcp.NewConn(c)
	if err != nil {
		return 0
	}
	level, name := rand.Int(), rand.Int()
	if _, err := tc.Option(level, name, data); err != nil {
		return 0
	}
	return 1
}
