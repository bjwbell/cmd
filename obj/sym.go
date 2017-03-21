package obj

import (
	"os"
	"path/filepath"
)

// WorkingDir returns the current working directory
// (or "/???" if the directory cannot be identified),
// with "/" as separator.
func WorkingDir() string {
	var path string
	path, _ = os.Getwd()
	if path == "" {
		path = "/???"
	}
	return filepath.ToSlash(path)
}

func Linknew(arch *LinkArch) *Link {
	ctxt := new(Link)
	ctxt.Arch = arch
	ctxt.Pathname = WorkingDir()
	ctxt.Framepointer_enabled = true
	return ctxt
}

func Linklookup(ctxt *Link, name string, v int) *LSym {
	s := &LSym{
		Name:    name,
		Version: int16(v),
		Size:    0,
	}
	return s
}

func (ctxt *Link) Lookup(name string, v int) *LSym {
	// TODO
	// s := ctxt.Hash[SymVer{name, v}]
	// if s != nil {
	// 	return s
	// }

	s := &LSym{
		Name:    name,
		Version: int16(v),
		Size:    0,
	}
	// TODO
	// ctxt.Hash[SymVer{name, v}] = s
	return s
}
