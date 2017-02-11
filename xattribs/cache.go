package xattribs

// methods to save/load file attributes to/from a variable, used for multiple/cached/templated/non-immediate manipulation
// store type: map[string(NAME)]string(VALUE)
// namespace restricted
// also conveniently uses string values unlike simple access that used []byte
const sep byte = '.'

// Attribs returns a map with keys set to existing attribute names, values not set.(see Populate())
func (this fileNS) Attribs() (map[string]string, error) {
	buf, err := this.list()
	return this.parse(buf), err
}

// parse takes attribute name listing in bytes, as returned by list(), and returns a map with keys set to attribute name strings.
func (this fileNS) parse(raw []byte) map[string]string {
	attribs := make(map[string]string)
	var itemStart int
	for pos, b := range raw {
		if b == 0 {
			item := string(raw[itemStart:pos])
			if item[len(this.namespace)] == sep && item[:len(this.namespace)] == this.namespace {
				attribs[item[len(this.namespace)+1:]] = ""
			}
			itemStart = pos + 1
		}
	}
	return attribs
}

// Populate fills in attribute map values
func (this *fileNS) Populate(attribs map[string]string) error {
	for tag:= range attribs {
		if attrib, err := this.Get(tag);err == nil {
			attribs[tag] = string(attrib)
		}
	}
	return nil
}

// Flush updates all attributes from map, deletes any not present, within same namespace
func (this *fileNS) Flush(attribs map[string]string) error {
	var err error
	temp, err := this.Attribs()
	if err != nil {
		return err
	}
	for k, v := range attribs {
		if err := this.Set(k, []byte(v));err != nil {
			return err
		}
	}
	for k := range temp {
		if _, ok := attribs[k]; !ok {
			if err := this.Remove(k);err != nil {
				return err
			}
		}
	}
	return nil
}
