package main

import "fmt"

const Id_NAME int = 10

func main() {
   var a string
   var b int
   a = "Hello World"
   b = 1111
   fmt.Println(a)
   fmt.Println(b)
   fmt.Println(Id_NAME)

   const (
      x = iota
      y = 3 << iota
      z
      w
   )

   fmt.Println(x,y,z,w)
}