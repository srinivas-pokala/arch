// Copyright 2021 The Go Authors.  All rights reserved.
// // Use of this source code is governed by a BSD-style
// // license that can be found in the LICENSE file.
// //
// // This file requires gcc and binutils with -march=z16 support.
// // s390xutil runs a series of commands like:
// //   go run map.go -fmt=asm ../s390x.csv > asm.S
// //   /usr/bin/gcc -c asm.S -march=z16
// //   /usr/bin/objdump -d asm.o
// // to create the file decode_generated.txt used to verify the disassembler.
// //
// // Note, the golang disassembler is not expected to support every extended
// // mnemonic, but it should support those which frequently show up in object
// // files compiled by the golang toolchain.

#define R1 8
#define R2 0
#define R3 0

#define X2 2

#define L1 4
#define L2 4

#define B1 2
#define B2 1
#define B3 6
#define B4 8

#define D1 6
#define D2 11
#define D3 182
#define D4 205

#define V1 18
#define V2 3
#define V3 5
#define V4 8



#define I 124
#define I1 12
#define I2 8
#define I3 9
#define I4 105
#define I5 18


#define RI2 -100
#define RI3 294
#define RI4 -168


#define M1 1
#define M3 3
#define M4 2
#define M5 3
#define M6 2
