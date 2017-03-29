// Derived from Inferno utils/6l/l.h and related files.
// http://code.google.com/p/inferno-os/source/browse/utils/6l/l.h
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package obj

// An As denotes an assembler opcode.
// There are some portable opcodes, declared here in package obj,
// that are common to all architectures.
// However, the majority of opcodes are arch-specific
// and are declared in their respective architecture's subpackage.
type As int16

// Prog.as opcodes.
// These are the portable opcodes, common to all architectures.
// Each architecture defines many more arch-specific opcodes,
// with values starting at A_ARCHSPECIFIC.
// Each architecture adds an offset to this so each machine has
// distinct space for its instructions. The offset is a power of
// two so it can be masked to return to origin zero.
// See the definitions of ABase386 etc.
const (
	AXXX = 0 + iota
	ACALL
	ACHECKNIL
	ADATA
	ADUFFCOPY
	ADUFFZERO
	AEND
	AFUNCDATA
	AGLOBL
	AJMP
	ANOP
	APCDATA
	ARET
	ATEXT
	ATYPE
	AUNDEF
	AUSEFIELD
	AVARDEF
	AVARKILL
	A_ARCHSPECIFIC
)

// An LSym is the sort of symbol that is written to an object file.
type LSym struct {
	Name    string
	Version int16
	Size    int
}

// The compiler needs LSym to satisfy fmt.Stringer, because it stores
// an LSym in ssa.ExternSymbol.
func (s *LSym) String() string {
	return s.Name
}

// Link holds the context for writing object code from a compiler
// to be linker input or for reading that input into the linker.
type Link struct {
	Arch                 *LinkArch
	Pathname             string
	Flag_shared          bool
	Flag_dynlink         bool
	Framepointer_enabled bool
	DiagFunc             func(string, ...interface{})
	Errors               int
}

// LinkArch is a stubbed struct
type LinkArch struct {
	IntSize int
}

// The smallest possible offset from the hardware stack pointer to a local
// variable on the stack. Architectures that use a link register save its value
// on the stack in the function prologue and so always have a pointer between
// the hardware stack pointer and the local variable area.
func (ctxt *Link) FixedFrameSize() int64 {
	// TODO: for now, assume amd64
	return 0
}

func (ctxt *Link) Diag(format string, args ...interface{}) {
	ctxt.Errors++
	ctxt.DiagFunc(format, args...)
}
