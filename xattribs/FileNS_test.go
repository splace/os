package xattribs

import (
	"testing"
	"io/ioutil"
	"os"
)



func TestAccess(t *testing.T) {
	var testFile *os.File
	var err error
	if testFile, err = ioutil.TempFile("", ".testing"); err != nil {
		panic("FileNS: can't make test file")
	}

	FileNS := NewFileNS(testFile, "user")
	if err := FileNS.Set("one", []byte("1?")); err != nil {
		t.Error(FileNS.File.Name() + " Set(\"one\") fromNS:\"" + FileNS.namespace + "\" - " + err.Error())
	}
	if err := FileNS.Set("two", []byte("2")); err != nil {
		t.Error(FileNS.File.Name() + " Set(\"one\") fromNS:\"" + FileNS.namespace + "\" - " + err.Error())
	}
	if err := FileNS.Update("one",[]byte("1")); err != nil {
		t.Error(FileNS.File.Name() + " Update(\"one\") fromNS:\"" + FileNS.namespace + "\" - " + err.Error())
	}
	if v, err := FileNS.Get("one"); string(v) != "1" || err != nil {
		t.Error(FileNS.File.Name() + " Get(\"one\") fromNS:\"" + FileNS.namespace + "\" - " + err.Error())
	}
	if err := FileNS.Update("three",[]byte("1")); err == nil {
		t.Error(FileNS.File.Name() + " Update(\"three\") fromNS:\"" + FileNS.namespace + "\" - " + err.Error())
	}
	if v, err := FileNS.Get("two"); string(v) != "2" || err != nil {
		t.Error(FileNS.File.Name() + " Get(\"one\") fromNS:\"" + FileNS.namespace + "\" - " + err.Error())
	}
	if _, err := FileNS.Get("three"); err == nil {
		t.Error(FileNS.File.Name() + " Get(\"one\") fromNS:\"" + FileNS.namespace + "\" - Success when no attrib should exist")
	}
	if err := FileNS.Set("three", []byte("3")); err != nil {
		t.Error(FileNS.File.Name() + " Set(\"one\") fromNS:\"" + FileNS.namespace + "\" - " + err.Error())
	}
	if v, err := FileNS.Get("three"); string(v) != "3" || err != nil {
		t.Error(FileNS.File.Name() + " Get(\"one\") fromNS:\"" + FileNS.namespace + "\" - " + err.Error())
	}
	if err := FileNS.Remove("one"); err != nil {
		t.Error(FileNS.File.Name() + " Remove(\"one\") fromNS:\"" + FileNS.namespace + "\" - " + err.Error())
	}
	if _, err := FileNS.Get("one"); err == nil {
		t.Error(FileNS.File.Name() + " Get(\"one\") fromNS:\"" + FileNS.namespace + "\" - " + err.Error())
	}

	if err := os.Remove(FileNS.Name()); err != nil {
		panic("FileNS: can't delete test file")
	}
}

func TestAccessNS(t *testing.T) {
	var testFile *os.File
	var err error
	if testFile, err = ioutil.TempFile("", ".testing"); err != nil {
		panic("FileNS: can't make test file")
	}

	FileNS1 := NewFileNS(testFile, "user.test")
	if err := FileNS1.Set("one", []byte("1")); err != nil {
		t.Error(FileNS1.File.Name() + " Set(\"one\") fromNS:\"" + FileNS1.namespace + "\" - " + err.Error())
	}
	FileNS2 := NewFileNS(testFile, "other")
	if _, err := FileNS2.Get("one"); err == nil {
		t.Error(FileNS2.File.Name() + " Get(\"one\") fromNS:\"" + FileNS2.namespace + "\" - " + err.Error())
	}

	if err := os.Remove(testFile.Name()); err != nil {
		panic("FileNS: can't delete test file")
	}
}

