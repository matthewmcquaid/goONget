package main
import (
  "fmt"
  "log"
  r "github.com/dancannon/gorethink"
)

func main() {

  fmt.Println("====================================================================================")

  fmt.Println("G'Day RethinkDB.... using Go\n")

  session, err := r.Connect(r.ConnectOpts{
    Address:  "127.0.0.1:28015",
  })

  if err != nil {
    log.Fatal(err)
  }

  res, err := r.Table("records").Run(session)

  var value interface{}

  if err != nil {
    log.Fatalln(err)
  }

  for res.Next(&value) {
    fmt.Println(value)
  }

  fmt.Println("====================================================================================")

}
