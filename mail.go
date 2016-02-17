package main

import(
//	"bytes"
	"log"
	"net/smtp"
	"fmt"
)

func main(){
	c,err:=smtp.Dial("")
	if err!=nil{
		log.Fatal(err)
	}
	defer c.Quit()

	if err:=c.Mail("");err!=nil{
		log.Fatal(err)
	}

	if err:=c.Rcpt("");err!=nil{
		log.Fatal(err)
	}

	wc,err := c.Data()
	if err!=nil{
		log.Fatal(err)
	}
	defer wc.Close()

	_,err=fmt.Fprintf(wc,"this is email body")
	if err!=nil{
		log.Fatal(err)
	}

}
