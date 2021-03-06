/* SPDX-License-Identifier: GPL-2.0
 *
 * Copyright (C) 2017-2018 WireGuard LLC. All Rights Reserved.
 */

package device

import "errors"

type DummyDatagram struct {
	msg      []byte
	endpoint Endpoint
	world    bool // better type
}

type DummyBind struct {
	in6    chan DummyDatagram
	ou6    chan DummyDatagram
	in4    chan DummyDatagram
	ou4    chan DummyDatagram
	closed bool
}

func (b *DummyBind) SetMark(v uint32) error {
	return nil
}

func (b *DummyBind) ReceiveIPv6(buff []byte) (int, Endpoint, error) {
	datagram, ok := <-b.in6
	if !ok {
		return 0, nil, errors.New("closed")
	}
	copy(buff, datagram.msg)
	return len(datagram.msg), datagram.endpoint, nil
}

func (b *DummyBind) ReceiveIPv4(buff []byte) (int, Endpoint, error) {
	datagram, ok := <-b.in4
	if !ok {
		return 0, nil, errors.New("closed")
	}
	copy(buff, datagram.msg)
	return len(datagram.msg), datagram.endpoint, nil
}

func (b *DummyBind) Close() error {
	close(b.in6)
	close(b.in4)
	b.closed = true
	return nil
}

func (b *DummyBind) Send(buff []byte, end Endpoint) error {
	return nil
}
