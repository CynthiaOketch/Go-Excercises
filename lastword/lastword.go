package main

import "os"


func main(){
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	
	args := os.Args[1]
	last := ""
	for i:=len(args)-1; i>=0; i--{
		if (args[i]) != ' ' {
           last = string(args[i]) +last
		}else if last!= ""{
			break
		}
	}
	os.Stdout.WriteString(last +"\n")
}