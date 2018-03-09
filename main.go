package main

import (
  "fmt"

  "github.com/wuryscamp/gopackage-tutorial/config"
  m1 "github.com/wuryscamp/gopackage-tutorial/model"
)

func main() {

  fmt.Println("Package Tutorial")

  pc := config.GetPostgresConnection()

  fmt.Println(pc)

  wury := m1.Person{ID: "1", Name: "Wuriyanto"}

  fmt.Println(wury.ID)
  fmt.Println(wury.Name)
}
