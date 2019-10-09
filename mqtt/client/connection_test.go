package client

import (
	"crypto/tls"
	"testing"

	"time"
)

const testAddress = "iot.eclipse.org:1883"

func Test_newConnection_tlsErr(t *testing.T) {
	if _, err := newConnection("", "", &tls.Config{}, time.Second); err == nil {
		notNilErrorExpected(t)
	}
}

func Test_newConnection(t *testing.T) {
	if _, err := newConnection("tcp", testAddress, nil, 3*time.Second); err != nil {
		nilErrorExpected(t, err)
	}
}

func notNilErrorExpected(t *testing.T) {
	t.Error("err => nil, want => not nil")
}

func nilErrorExpected(t *testing.T, err error) {
	t.Errorf("err => %q, want => nil", err)
}
