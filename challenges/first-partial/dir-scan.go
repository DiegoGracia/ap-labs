package main

import (
  "fmt"
  "os"
  "path/filepath"
)

var directories int = 0;
var devices int = 0;
var sockets int = 0;
var symbolicLink int = 0;
var othersFiles int = 0;

//se checa cada condicion y si los archivos cumplen con ellas

func visit(dir string, f os.FileInfo, err error) error {

  if f.Mode() & os.ModeSymlink != 0 {
    symbolicLink++;
    return nil;
  }
  if f.IsDir() {
    directories++;
    return nil;
  }
  if f.Mode().IsRegular() {
    othersFiles++;
    return nil;
  }
  if f.Mode()&os.ModeSocket != 0 {
    sockets++;
    return nil;
  }
  if f.Mode()&os.ModeDevice != 0{
    devices++;
    return nil;
  }
  return nil
}

// scanDir stands for the directory scanning implementation
func scanDir(dir string) error {
  err := filepath.Walk(dir, visit);
  return err
}

func main() {

  if len(os.Args) < 2 {
    fmt.Println("Usage: ./dir-scan <directory>")
    os.Exit(1)
  }

  scanDir(os.Args[1])
  fmt.Printf("Directory Scanner Tool\n" +
              "| Directories:\t\t|\t%d\t|\n" +
              "| Symbolic Links:\t|\t%d\t|\n" +
              "| Devices:\t\t|\t%d\t|\n" +
              "| Sockets:\t\t|\t%d\t|\n" +
              "| Symbolic Links:\t|\t%d\t|\n" +
              "| Other files:\t\t|\t%d\t|\n",os.Args[1], directories, symbolicLink, devices, sockets, symbolicLink, othersFiles)
}
