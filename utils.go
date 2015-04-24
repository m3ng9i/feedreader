package feedreader

import "html"
import "strings"


func transformContent(content string) string {
    if strings.ContainsAny(content, "<>") {
        return content
    } else {
        return html.UnescapeString(content)
    }
}
