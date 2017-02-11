package templates

import (
	"testing"
	//"fmt"
	"io/ioutil"
	"os"
)
import "./dublincore"
import "../../xattribs"

func TestTemplate(t *testing.T) {
	var testFile *os.File
	var err error
	if testFile, err = ioutil.TempFile("", ".testing"); err != nil {
		panic("FileNS: can't make test file")
	}
	FileNS1 := xattribs.NewFileNS(*testFile, dublincore.Namespace)
	DublinCoreAttribs:=ParseTags(dublincore.Tags)
	if len(DublinCoreAttribs) !=15{
		t.Error("dublin core wrong count")
	}

	DublinCoreAttribs["title"] = "my book"
	RemoveKeysWithEmptyStringValue(&DublinCoreAttribs)
	if len(DublinCoreAttribs) != 1 {
		t.Error("no")
	}

	FileNS1.Flush(DublinCoreAttribs)

	if err := os.Remove(testFile.Name()); err != nil {
		panic("FileNS: can't delete test file")
	}

}

func TestFlushAttribs(t *testing.T) {
	var testFile *os.File
	var err error
	if testFile, err = ioutil.TempFile("", ".testing"); err != nil {
		panic("FileNS: can't make test file")
	}
	
	// add one dublin core atrib to test file
	FileNS1 := xattribs.NewFileNS(*testFile, dublincore.Namespace)
	DublinCoreAttribs:=ParseTags(dublincore.Tags)
	DublinCoreAttribs["title"] = "my book"
	RemoveKeysWithEmptyStringValue(&DublinCoreAttribs)
	FileNS1.Flush(DublinCoreAttribs)
	
	// check its there
	FileNS2 := xattribs.NewFileNS(*testFile, dublincore.Namespace)
	if attribs2, err:=FileNS2.Attribs(); err != nil || len(attribs2) !=1{
		panic("attrib not saved")
	}
	// remove it
	FileNS2.Flush(nil)

	// check its gone
	FileNS3 := xattribs.NewFileNS(*testFile, dublincore.Namespace)
	if attribs3, err:=FileNS3.Attribs(); err != nil || len(attribs3) !=0{
		panic("attrib not deleted")
	}

	if err := os.Remove(testFile.Name()); err != nil {
		panic("FileNS: can't delete test file")
	}


}

