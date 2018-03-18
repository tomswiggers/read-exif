package main

import (
  "fmt"
  "strings"
  "os"
  "io/ioutil"
  "github.com/xiam/exif"
)
func check(e error) {
  if e != nil {
    panic(e)
  }
}


func main() {
  var dirname, filename, newfilename string
  var model string
  var timestamp string

  args := os.Args

  if len(args) > 1 {
    dirname = args[1]
  }

  files, err := ioutil.ReadDir(dirname)
  check(err)

  for _, f := range files {
    filename = fmt.Sprintf("%s/%s", dirname, f.Name())
    fmt.Printf("Filename: %s\n", filename)

    data, err := exif.Read(filename)
    check(err)

    timestamp = strings.Replace(data.Tags["Date and Time"], ":", "", -1)
    timestamp = strings.Replace(timestamp, " ", "", -1)
    model = strings.Replace(data.Tags["Model"], " ", "_", -1)

    newfilename = fmt.Sprintf("%s_%s", timestamp, model)

    fmt.Printf("Date and Time: %s\n", timestamp)
    fmt.Printf("Model: %s\n", model)
    fmt.Printf("Model: %s\n", newfilename)
  }
}
