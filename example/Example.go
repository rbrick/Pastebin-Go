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
    pb.SetUserKey("99324fa628c831fcbcf8e37b34876ac7")
    pb.SetExpireDate(pb.NEVER)

    url, err := pb.Execute()

    if err != nil {
        panic(err)
    }

    fmt.Println(url)
}
