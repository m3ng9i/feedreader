package feedreader

import "net/http"
import "io/ioutil"
import "github.com/m3ng9i/go-utils/xml"

// fetch feed content from a url, return []byte
func Fetch(url string) ([]byte, error) {
    r, err := http.Get(url)
    if err != nil {
        return []byte{}, err
    }
    defer r.Body.Close()
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return []byte{}, err
    }

    // Remove invalid xml characters
    b = xml.RemoveInvalidChars(b)

    return b, nil
}


// fetch feed content from a url, return string
func FetchString(url string) (string, error) {
    b, err := Fetch(url)
    if err != nil {
        return "", err
    }
    return string(b), nil
}

