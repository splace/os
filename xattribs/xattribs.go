// Package xattribs facilitates use of a files extended attributes. (see: https://en.wikipedia.org/wiki/Extended_file_attributes)
// summary: supported on modern FS's, some have limited size and copying a file between fs's may loose them.
package xattribs

import "os"

// FileNS's embed an os.File, and have behaviour to allow access to a particular namespace of extended attributes from that file.
type fileNS struct {
	os.File
	namespace string
}

// NewFileNS returns an extended attrib aware File, with namespace set.
func NewFileNS(f os.File, namespace string) *fileNS {
	xas := fileNS{f, namespace}
	return &xas
}

