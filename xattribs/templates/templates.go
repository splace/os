// pre-defined attribute names templates.
// sub-packages are used for each templates data.
package templates


func RemoveKeysWithEmptyStringValue(attribs *map[string]string) {
	for k, v := range *attribs {
		if v == "" {
			delete(*attribs, k)
		}
	}
}

func ParseTags(dotTerminated string) map[string]string{
	attribs := make(map[string]string)
	var itemStart int
	for pos, b := range dotTerminated {
		if b == '.' {
			attribs[dotTerminated[itemStart:pos]] = ""
			itemStart = pos + 1
		}
	}
	return attribs
}

