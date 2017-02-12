package xattribs

import "syscall"

const ssep string = string(sep)

// basic immediate attribute acccess using []byte values
// Namespace restricted

// gets an attribute value as bytes
func (f FileNS) Get(tag string) ([]byte, error) {
	if size, err := syscall.Getxattr(f.File.Name(), f.Namespace+ssep+tag, nil); err == nil {
		buf := make([]byte, size)
		if read, err := syscall.Getxattr(f.File.Name(), f.Namespace+ssep+tag, buf); err == nil {
			return buf[:read], nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

// sets an attribute's value from bytes.
func (f FileNS) Set(tag string, data []byte) error {
	return syscall.Setxattr(f.File.Name(), f.Namespace+ssep+tag, data, 0)
}

// deletes an attribute.
func (f FileNS) Remove(tag string) error {
	return syscall.Removexattr(f.File.Name(), f.Namespace+ssep+tag)
}

// updates an attribute, returns error if attribute not pre-existing.
func (f FileNS) Update(tag string, data []byte) error {
	if _, err := syscall.Getxattr(f.File.Name(), f.Namespace+ssep+tag, nil); err == nil {
		return f.Set(tag, data)
	} else {
		return err
	}
}

// returns attribute names in bytes, nul char delimitted.
func (f FileNS) list() ([]byte, error) {
	if size, err := syscall.Listxattr(f.File.Name(), nil); err == nil {
		buf := make([]byte, size)
		if read, err := syscall.Listxattr(f.File.Name(), buf); err == nil {
			return buf[0:read], nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
