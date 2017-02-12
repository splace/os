package xattribs

import "syscall"

const ssep string = string(sep)

// basic immediate attribute acccess using []byte values
// namespace restricted

// gets an attribute value as bytes
func (f fileNS) Get(tag string) ([]byte, error) {
	if size, err := syscall.Getxattr(f.File.Name(), f.namespace+ssep+tag, nil); err == nil {
		buf := make([]byte, size)
		if read, err := syscall.Getxattr(f.File.Name(), f.namespace+ssep+tag, buf); err == nil {
			return buf[:read], nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

// sets an attribute's value from bytes.
func (f fileNS) Set(tag string, data []byte) error {
	return syscall.Setxattr(f.File.Name(), f.namespace+ssep+tag, data, 0)
}

// deletes an attribute.
func (f fileNS) Remove(tag string) error {
	return syscall.Removexattr(f.File.Name(), f.namespace+ssep+tag)
}

// updates an attribute, returns error if attribute not pre-existing.
func (f fileNS) Update(tag string, data []byte) error {
	if _, err := syscall.Getxattr(f.File.Name(), f.namespace+ssep+tag, nil); err == nil {
		return f.Set(tag, data)
	} else {
		return err
	}
}

// returns attribute names in bytes, nul char delimitted.
func (f fileNS) list() ([]byte, error) {
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
