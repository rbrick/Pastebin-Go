package pastebin

import (
    "net/url"
    "strconv"
    "net/http"
    "io/ioutil"
)

// The parameters that will be written when making the request
var api_params = make(map[string]string)

// The view status of the paste
const (
    PUBLIC = iota
    UNLISTED
    PRIVATE
)

// The expire times value, defaults to Never if not provided
const (
    NEVER = "N"
    TEN_MINUTES = "10M"
    ONE_HOUR = "1H"
    ONE_DAY = "1D"
    ONE_WEEK = "1W"
    TWO_WEEK = "2W"
    ONE_MONTH = "1M"
)

func addParam(k, v string) {
    api_params[k] = v
}

// Sets the API key in the params
func SetAPIKey(api_key string) {
    addParam("api_dev_key", api_key)
}

func SetContent(content string) {
    addParam("api_paste_code", content)
}

func SetName(name string) {
    addParam("api_paste_name", name)
}

func SetViewStatus(status int) {
    addParam("api_paste_private", strconv.Itoa(status))
}

func SetUserKey(user_key string) {
    addParam("api_user_key", user_key)
}

func SetExpireDate(expire_date string) {
     addParam("api_paste_expire_date", expire_date)
}


func Execute() (string, error) {
    data := url.Values{}
    // This sets the API option to paste so we can paste stuff
    data.Add("api_option", "paste")
    for k, v := range api_params {
        data.Add(k, v)
    }

    resp, err := http.PostForm("http://pastebin.com/api/api_post.php", data)

    if err != nil {
        panic(err)
        return "", err
    }

    ra, err1 := ioutil.ReadAll(resp.Body)

    if err1 != nil {
        panic(err1)
        return "", err1
    }

    return string(ra), nil
}