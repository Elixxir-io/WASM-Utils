////////////////////////////////////////////////////////////////////////////////
// Copyright © 2022 xx foundation                                             //
//                                                                            //
// Use of this source code is governed by a license that can be found in the  //
// LICENSE file.                                                              //
////////////////////////////////////////////////////////////////////////////////

//go:build !js && !wasm

package exception

// This file is used in testing. Refer to throws.go for more info.

import (
	"unsafe"

	jww "github.com/spf13/jwalterweatherman"
)

func throw(exception, message unsafe.Pointer) {
	jww.FATAL.Panicf("%v: %v", exception, message)
}
