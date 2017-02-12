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
// has methods to save/load file attributes to/from a variable, used for multiple/cached/templated/non-immediate manipulation
// store type: map[string(NAME)]string(VALUE)
func NewFileNS(f os.File, namespace string) *fileNS {
	xas := fileNS{f, namespace}
	return &xas
}

const sep byte = '.'

// Attribs returns a map with keys set to existing attribute names, values not set.(see Populate())
func (f fileNS) Attribs() (map[string]string, error) {
	buf, err := f.list()
	return f.parse(buf), err
}

// parse takes attribute name listing in bytes, as returned by list(), and returns a map with keys set to attribute name strings.
func (f fileNS) parse(raw []byte) map[string]string {
	attribs := make(map[string]string)
	var itemStart int
	for pos, b := range raw {
		if b == 0 {
			item := string(raw[itemStart:pos])
			if item[len(f.namespace)] == sep && item[:len(f.namespace)] == f.namespace {
				attribs[item[len(f.namespace)+1:]] = ""
			}
			itemStart = pos + 1
		}
	}
	return attribs
}

// Populate fills in attribute map values
func (f *fileNS) Populate(attribs map[string]string) error {
	for tag:= range attribs {
		if attrib, err := f.Get(tag);err == nil {
			attribs[tag] = string(attrib)
		}
	}
	return nil
}

// Flush updates all attributes from map, deletes any not present, within same namespace
func (f *fileNS) Flush(attribs map[string]string) error {
	var err error
	temp, err := f.Attribs()
	if err != nil {
		return err
	}
	for k, v := range attribs {
		if err := f.Set(k, []byte(v));err != nil {
			return err
		}
	}
	for k := range temp {
		if _, ok := attribs[k]; !ok {
			if err := f.Remove(k);err != nil {
				return err
			}
		}
	}
	return nil
}
