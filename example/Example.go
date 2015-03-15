package main

import (
    pb "github.com/rbrick/Pastebin-Go"
    "fmt"
)


func main() {
    pb.SetAPIKey("REDACTED")
    pb.SetContent("Hello World!")
    pb.SetName("Example")
    pb.SetViewStatus(pb.PUBLIC)
    pb.SetUserKey("REDACTED")
    pb.SetExpireDate(pb.NEVER)

    url, err := pb.Execute()

    if err != nil {
        panic(err)
    }

    fmt.Println(url)
}
