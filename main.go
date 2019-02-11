package main

import (
	"log"
	"os/exec"
)


func main(){


	pdf2html("mencius")


}

func pdf2html(name string){
	add:="./pdf/"+name+".pdf"
	aimAdd:="./html/"+name+".html"
	cmd:=exec.Command("pdf2htmlEX", "--zoom","1.35",add,aimAdd)
	_,err:=cmd.Output()
	if err!=nil{
		log.Fatal(err)
	}
}


