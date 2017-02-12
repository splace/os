package xattribs

// methods to save/load file attributes to/from a variable, used for multiple/cached/templated/non-immediate manipulation
// store type: map[string(NAME)]string(VALUE)
// namespace restricted
// also conveniently uses string values unlike simple access that used []byte
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
