package feedreader

import "html"
import "strings"
import myhtml "github.com/m3ng9i/go-utils/html"


// convert content to html
func transformContent(content string) string {
    if strings.ContainsAny(content, "<>") {
        return content
    } else {
        c := html.UnescapeString(content)
        if strings.ContainsAny(c, "<>") {
            return c
        } else {
            return myhtml.Text2Html(c)
        }
    }
}
