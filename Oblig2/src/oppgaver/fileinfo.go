package main

import (
  "os"
  "log"
  "fmt"
//  "io/ioutil"
)

func main() {

  fi, err := os.Lstat("text.txt")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Information about file", fi.Name())
  fmt.Println("Size:", fi.Size(), "bytes." , fi.Size()/1024, "KB.", fi.Size()/1024/1024, "MB.", fi.Size()/1024/1024/1024, "GB.")
  if (fi.Mode().IsDir()==false) {
    fmt.Println("Is not a diretory.")
  } else {
    fmt.Println("Is a directory.")
  }
  if (fi.Mode().IsRegular()==false) {
    fmt.Println("Is not a regular file.")
  } else {
    fmt.Println("Is a regular file.")
  }
  fmt.Println("Has Unix permission bits:", fi.Mode())

  ft := fi.Mode();
  if (ft&os.ModeAppend != 0) {
    fmt.Println("Is append only")
  } else {
    fmt.Println("Is not append only")
  }
  if (ft&os.ModeDevice != 0){
    fmt.Println("Is a device file")
  } else {
    fmt.Println("Is not a device file")
  }
  if (ft&os.ModeSymlink != 0){
    fmt.Println("Is a symbolic link")
  } else {
    fmt.Println("Is not a symbolic link")
  }
}
