package main

import(
	"os"
	"os/exec"
	"strings"
	"fmt"
	"bytes"
)

var (
	suffix=".go"
	osArch=map[string][]string{
		"linux":[]string{"amd64"},
		"darwin":[]string{"amd64"},
	}
)

func main(){

	if len(os.Args)!=2{
		return
	}

	if !strings.HasSuffix(os.Args[1],suffix){
		return
	}

	for key,val:=range osArch{
		for _,val2:=range val{
			execBuild(key,val2,os.Args[1])
		}
	}

}


func execBuild(oss,arch,src string){

	os.Setenv("GOOS",oss)
	os.Setenv("GOARCH",arch)

	packageName:=strings.TrimSuffix(src,suffix)
	osArch:="_"+oss+"_"+arch

	cmd:=exec.Command(
		"go",
		"build",
		"-o",
		packageName+osArch,
		src,
	)

	fmt.Println("Build on "+oss+"/"+arch+"...")

	var b bytes.Buffer
	cmd.Stderr=&b
	err := cmd.Run()
	if err!=nil{
		fmt.Println("Error -> "+err.Error()+":"+b.String())
		return
	}
	
	fmt.Println("Completed! -> "+packageName+osArch)
}
