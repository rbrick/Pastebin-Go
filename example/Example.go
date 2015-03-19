package main

import (
    pb "github.com/rbrick/Pastebin-Go"
    "fmt"
    "flag"
)



var apiKey = flag.String("apikey", "", "Your api key")
var content = flag.String("content", "", "The content of the paste")
var name = flag.String("name", "This was made using Pastebin-CLI!", "The name of the paste")
var userKey = flag.String("userkey", "", "Your user key")
var viewStatus = flag.Int("viewstatus", 1, "The view status of the paste, Options are\n0 = Public\n1 = Unlisted\n2=Private")
var expireDate = flag.String("expiredate", "N", "The time in which the paste expires, Options are\nNever = N\nTen minutes = 10M\nOne hour = 1H\nOne day = ")

func main() {

    flag.Parse()

    pb.SetAPIKey(*apiKey)
    pb.SetContent(*content)
    pb.SetName(*name)
    pb.SetViewStatus(*viewStatus)
    pb.SetUserKey(*userKey)
    pb.SetExpireDate(*expireDate)

    fmt.Println(*viewStatus)

    url, err := pb.Execute()

    if err != nil {
        panic(err)
    }

    fmt.Println(url)
}
