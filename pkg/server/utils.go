// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package server implements the server
package server

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// SpdkMockServer implements mock function to send data instead of real SPDK app, used in testing
func SpdkMockServer(l net.Listener, toSend []string) {
	for _, spdk := range toSend {
		fd, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		log.Printf("SPDK mockup Server: client connected [%s]", fd.RemoteAddr().Network())
		log.Printf("SPDK ID [%d]", RPCID)

		buf := make([]byte, 512)
		nr, err := fd.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		if strings.Contains(spdk, "%") {
			spdk = fmt.Sprintf(spdk, RPCID)
		}

		log.Printf("SPDK mockup Server: got : %s", string(data))
		log.Printf("SPDK mockup Server: snd : %s", spdk)

		_, err = fd.Write([]byte(spdk))
		if err != nil {
			log.Fatal("Write: ", err)
		}
		err = fd.(*net.UnixConn).CloseWrite()
		if err != nil {
			log.Fatal(err)
		}
	}
}