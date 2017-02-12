package main

import "os"
import "fmt"
import "github.com/splace/os/xattribs"

func main(){
	file,err := os.Open(os.Args[0])
	if err==nil {
		//  make a new os.File extended with methods for accessing extended file attribs.
		DropBoxAttribAwareFile:=xattribs.FileNS{file,"user.com.dropbox"}
		attrs,err:=DropBoxAttribAwareFile.Get("attributes") // the only dropbox attribute
		fmt.Println(err,attrs)
	}
}


