package main

import (
    pb "github.com/rbrick/Pastebin-Go"
    "fmt"
)

func main() {
    pb.SetAPIKey("SOME API KEY")
    pb.SetContent("Hello World!")
    pb.SetName("Example")
    pb.SetViewStatus(pb.PUBLIC)

    url, err := pb.Execute()

    if err != nil {
        panic(err)
    }

    fmt.Println(url)
}
