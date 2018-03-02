# xattribs 
access to [extended file attributes](https://en.wikipedia.org/wiki/Extended_file_attributes), by extending an existing os.File to give it the required methods, methods that are restricted to one particular namespace at a time.

caveat: is supported on virtually all modern FS's, but some have limited size and/or copying a file between fs's may loose them.(check link)

Linux supported, Mac not difficult, Windows possible.

example: persists a flag option into an attribute on the executable.

	package main

	import "flag"
	import "fmt"
	import "os"
	import "log"
	import "github.com/splace/os/xattribs"

	const attrib="message"

	func main() {
		thisFile,err := os.Open(os.Args[0])
		if err!=nil {panic(err)}
		xf:=xattribs.FileNS{thisFile,"user.com.github"}  // use github as test namespace
		bs,err:=xf.Get(attrib) 
		if err!=nil {log.Printf("\"%s\" attrib in \"%s\" not got.(%s)\n",attrib,xf.Name(),err)}
		mess:= flag.String(attrib, string(bs), "persisted "+attrib)
		flag.Parse()
		fmt.Println(*mess)
		if *mess!=string(bs){
			err=xf.Set(attrib,[]byte(*mess))
			if err!=nil{
				log.Fatal(err)
			}
		}
	}

installation:

     go get github.com/splace/os/xattribs

uses:


*	not needing a separate config file

*	persistence if not able to add files

*	retaining persistence accross copy

*	persist options directly on an executable directly.  (see examples)
 
*	persist to a file some expensive to parse info.

*	backup a files original name before performing some complex batch renaming.


docs: 
     
[![GoDoc](https://godoc.org/github.com/splace/os/xattribs?status.svg)](https://godoc.org/github.com/splace/os/xattribs)  xattribs 

[![GoDoc](https://godoc.org/github.com/splace/os/xattribs/templates?status.svg)](https://godoc.org/github.com/splace/os/xattribs/templates)  xattribs/templates

