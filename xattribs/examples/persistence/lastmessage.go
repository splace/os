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


