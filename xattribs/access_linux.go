package xattribs

import "syscall"

const ssep string = string(sep)

// basic immediate attribute acccess using []byte values
// namespace restricted

// gets an attribute value as bytes
func (this fileNS) Get(tag string) ([]byte, error) {
	if size, err := syscall.Getxattr(this.File.Name(), this.namespace+ssep+tag, nil); err == nil {
		buf := make([]byte, size)
		if read, err := syscall.Getxattr(this.File.Name(), this.namespace+ssep+tag, buf); err == nil {
			return buf[:read], nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

// sets an attribute's value from bytes.
func (this *fileNS) Set(tag string, data []byte) error {
	return syscall.Setxattr(this.File.Name(), this.namespace+ssep+tag, data, 0)
}

// deletes an attribute.
func (this *fileNS) Remove(tag string) error {
	return syscall.Removexattr(this.File.Name(), this.namespace+ssep+tag)
}

// updates an attribute, returns error if attribute not pre-existing.
func (this *fileNS) Update(tag string, data []byte) error {
	if _, err := syscall.Getxattr(this.File.Name(), this.namespace+ssep+tag, nil); err == nil {
		return this.Set(tag, data)
	} else {
		return err
	}
}

// returns attribute names in bytes, nul char delimitted.
func (this fileNS) list() ([]byte, error) {
	if size, err := syscall.Listxattr(this.File.Name(), nil); err == nil {
		buf := make([]byte, size)
		if read, err := syscall.Listxattr(this.File.Name(), buf); err == nil {
			return buf[0:read], nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
