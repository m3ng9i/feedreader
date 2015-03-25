package feedreader

import "net/http"
import "io/ioutil"
import "github.com/m3ng9i/go-utils/xml"

// Fetch feed content from a url, return []byte
// If returned error is not nil, it will be FetchError.
func FetchByte(url string) ([]byte, error) {
    r, err := http.Get(url)
    if err != nil {
        return []byte{}, &FetchError{Url:url, Err:err}
    }
    defer r.Body.Close()
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return []byte{}, &FetchError{Url:url, Err:err}
    }

    // Remove invalid xml characters
    b = xml.RemoveInvalidChars(b)

    return b, nil
}


// fetch feed content from a url, return string
// If returned error is not nil, it will be FetchError.
func FetchString(url string) (string, error) {
    b, err := FetchByte(url)
    if err != nil {
        return "", err
    }
    return string(b), nil
}

