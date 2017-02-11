# xattribs 
access to [extended file attributes](https://en.wikipedia.org/wiki/Extended_file_attributes), by extending an existing os.File to give it the required methods, methods that are restricted to one particular namespace at a time.

Linux supported, Mac not difficult, Windows possible.

example: dropbox files.(which only have one attribute called "attributes")

    import "os"
    import "github.com/splace/os/xattribs"
  
    file,err := os.Open(<<path>>)
    if err==nil {
  		  //  make a new os.File extended with access to extened attribs.
  		  DropBoxAttribAwareFile:=osx.NewFileNS(*file,"user.com.dropbox")
  		  attrs,err:=DropBoxAttribAwareFile.Get("attributes")
  	}
	
installation:

     go get github.com/splace/os/xattribs

uses: (basically about avoiding a separate config file so allowing persistence if not able to add files, and ease of copying wilst retaining persistence, without using a folder.)

attach expensive to parse file info.

backup a files original name before performing complex batch renaming.

persist options to an executable.  
 
docs: 
     
[![GoDoc](https://godoc.org/github.com/splace/os/xattribs?status.svg)](https://godoc.org/github.com/splace/os/xattribs)  xattribs 

[![GoDoc](https://godoc.org/github.com/splace/os/xattribs/templates?status.svg)](https://godoc.org/github.com/splace/os/xattribs/templates)  xattribs/templates

