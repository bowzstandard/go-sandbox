package main

import(
  "os"
  "os/exec"
  "strconv"
  "fmt"
  "bytes"
  "time"
)

var (
  prefix=""
  suffix=".sql"
  db=""
)

func main(){

  if len(os.Args)!=3{
    fmt.Println("Error -> insufficient args:output path & .conf path")
    return
  }

  err:=os.MkdirAll(os.Args[1],0755)
  if err!=nil{
    fmt.Println("Error -> "+err.Error())
    return
  }
  

  if _,err:=os.Stat(os.Args[2]);err!=nil{
    fmt.Println("Error -> "+err.Error())
    return
  }

  cmd:=exec.Command(
    "mysqldump",
    "--defaults-extra-file="+os.Args[2],
    "--single-transaction",
    db,
  )

  var b bytes.Buffer
  var e bytes.Buffer
  cmd.Stdout = &b
  cmd.Stderr = &e
  err2 := cmd.Run()
  if err2!=nil{
    fmt.Println("Error -> "+err2.Error()+":"+e.String())
    return
  }

  path:=os.Args[1]
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

  filename:=prefix+year+month+day+suffix

  file,err:=os.OpenFile(path+"/"+filename,os.O_RDWR|os.O_CREATE,0600)
  if err!=nil{
    fmt.Println("Error -> "+err.Error())
    return
  }
  defer file.Close()

  info,err:=file.Stat()
  if err!=nil{
    fmt.Println("Error -> "+err.Error())
    return
  }

  file.WriteAt(b.Bytes(),info.Size())

}
