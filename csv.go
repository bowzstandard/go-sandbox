package main

import(
	"os"
	"time"
	"log"
	"strconv"
	"strings"
)

type TestData struct{
	time int64
	body string
}

func(t *TestData)New(){
	t.time = time.Now().Unix()
	t.body = "this ,\"is\" `test`"
}

func(t *TestData)DumpCSV()[]byte{
	str:=[]string{
		strconv.FormatInt(t.time,10),
		t.body,
	}
	return []byte(ConcatStr(str))
}

func ConcatStr(arr []string)string{
	for key,val:=range arr{
		a:=strings.Split(val,"\"")
		b:=strings.Join(a,"\"\"")
		arr[key]="\""+b+"\""
	}
	return strings.Join(arr,",")
}

type FileParser struct{
	file *os.File
}

func (p *FileParser)New(f *os.File){
	p.file = f
}

func (p *FileParser)WriteData(v []byte){
	info,err:=p.file.Stat()
	if err!=nil{
		log.Fatal(err)
	}
	p.file.WriteAt(v,info.Size())
}

func main(){

	tmp:=&TestData{}
	tmp.New()

	now:=time.Now()
	
	year:=strconv.Itoa(now.Year())
	
	month:=strconv.Itoa(int(now.Month()))
	if len(month)<2{
		month="0"+month
	}

	day:=strconv.Itoa(now.Day())
	if len(day)<2{
		day="0"+day
	}

	path:="logs/"+year+"/"+month+"/"+day+"/"
	filepath:=path+"data.csv"

	err:=os.MkdirAll(path,0755)
	if err!=nil{
		log.Fatal(err)
	}
	
	file,err:=os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0600)
	if err!=nil{
		log.Fatal(err)
	}
	defer file.Close()

	p:=&FileParser{}
	p.New(file)

	data:=tmp.DumpCSV()

	p.WriteData(data)
	p.WriteData([]byte("\n"))

}

