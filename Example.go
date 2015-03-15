package main

import (
    pb "github.com/rbrick/Pastebin-Go"
    "fmt"
)


func main() {
    pb.SetAPIKey("25006b3890b4a16b40847018c2c74fd5")
    pb.SetContent("Hello World!")
    pb.SetName("Example")
    pb.SetViewStatus(pb.PUBLIC)

    url, err := pb.Execute()

    if err != nil {
        panic(err)
    }

    fmt.Println(url)
}
