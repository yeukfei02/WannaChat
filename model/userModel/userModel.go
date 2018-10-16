package userModel

import (
  "fmt"
)

func checkErr(err error) {
  if (err != nil) {
    fmt.Println("error = " + err.Error())
  }
}
