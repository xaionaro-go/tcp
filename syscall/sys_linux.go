// Copyright 2014 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

var options = [EndOfSo]Option{
	SoBuffered:  {0, sysSIOCINQ},
	SoAvailable: {0, sysSIOCOUTQ},
}
