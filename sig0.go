// Copyright 2011 Miek Gieben. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// SIG(0)
//
// From RFC 2931:
//
//     SIG(0) provides protection for DNS transactions and requests ....
//     ... protection for glue records, DNS requests, protection for message headers 
//     on requests and responses, and protection of the overall integrity of a response.
//
//
package dns

// SIG0Sign will sign the ...
func (dns *Msg) SIG0Sign() ([]byte, error) {
	return nil, nil
}

// SIG0Verify will bla bla
func SIG0Verify([]byte) bool {
	return false
}