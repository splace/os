# xattribs 
access to [extended file attributes](https://en.wikipedia.org/wiki/Extended_file_attributes), by extending an existing os.File to give it the required methods, methods that are restricted to one particular namespace at a time.

caveat: is supported on virtually all modern FS's, but some have limited size and/or copying a file between fs's may loose them.(check link)

Linux supported, Mac not difficult, Windows possible.

example: reads its own executables dropbox attribute.(so has to be compiled into a dropbox folder.)

	package main

	import "os"
	import "fmt"
	import "github.com/splace/os/xattribs"

	func main(){
		file,err := os.Open(os.Args[0])
		if err==nil {
			//  make a new os.File extended with methods to extended file attribs.
			DropBoxAttribAwareFile:=xattribs.FileNS{file,"user.com.dropbox"}
			attrs,err:=DropBoxAttribAwareFile.Get("attributes") // the only dropbox attribute
			fmt.Println(err,attrs)
		}
	}

installation:

     go get github.com/splace/os/xattribs

uses: (basically about avoiding a separate config file so allowing persistence if not able to add files, and ease of copying whilst retaining persistence, without using a folder.)

persist options to an executable directly.  (see examples)
 
attach expensive to parse file info.

backup a files original name before performing complex batch renaming.

docs: 
     
[![GoDoc](https://godoc.org/github.com/splace/os/xattribs?status.svg)](https://godoc.org/github.com/splace/os/xattribs)  xattribs 

[![GoDoc](https://godoc.org/github.com/splace/os/xattribs/templates?status.svg)](https://godoc.org/github.com/splace/os/xattribs/templates)  xattribs/templates

