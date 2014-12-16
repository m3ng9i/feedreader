/*
RSS 2.0 parser

The elements below are not supported.
channel/category
channel/docs
channel/cloud
channel/rating
channel/textInput
channel/item/category
channel/item/source
*/

package feedreader

import "time"
import "strings"
import "encoding/xml"
import "errors"
import "fmt"
import "github.com/m3ng9i/go-utils/set"

// rss/image node
type Rss20Image struct {
    Url             string              `xml:"url"`
    Title           string              `xml:"title"`
    Link            string              `xml:"link"`
    Width           uint                `xml:"width"`
    Height          uint                `xml:"height"`
    Description     string              `xml:"description"`
}

// rss/item/enclosure node
type Rss20ItemEnclosure struct {
    Url             string              `xml:"url,attr"`
    Length          uint64              `xml:"length,attr"`
    Type            string              `xml:"type,attr"`
}

// rss/item/guid node
type Rss20ItemGuid struct {
    Guid            string              `xml:",chardata"`
    IsPermaLinkRaw  string              `xml:"isPermaLink,attr"`
    IsPermaLink     bool
}

// rss/item node
type Rss20Item struct {
    Title               string              `xml:"title"`
    Link                string              `xml:"link"`
    Description         string              `xml:"description"`
    Author              string              `xml:"author"`
    Comments            string              `xml:"comments"`
    Enclosure           *Rss20ItemEnclosure `xml:"enclosure"`
    Guid                *Rss20ItemGuid      `xml:"guid"`
    PubDateRaw          string              `xml:"pubDate"`
    PubDate             time.Time
}

// whole rss node
type Rss20 struct {
    Version             string              `xml:"version,attr"`
    Title               string              `xml:"channel>title"`
    Link                string              `xml:"channel>link"`
    Description         string              `xml:"channel>description"`
    Language            string              `xml:"channel>language"`
    Copyright           string              `xml:"channel>copyright"`
    ManagingEditor      string              `xml:"channel>managingEditor"`
    WebMaster           string              `xml:"channel>webMaster"`
    PubDateRaw          string              `xml:"channel>pubDate"`
    PubDate             time.Time
    LastBuildDateRaw    string              `xml:"channel>lastBuildDate"`
    LastBuildDate       time.Time
    Generator           string              `xml:"channel>generator"`
    Ttl                 uint                `xml:"channel>ttl"`
    Image               *Rss20Image         `xml:"channel>image"`
    SkipHours           []uint8             `xml:"channel>skipHours>hours"`
    SkipDaysRaw         []string            `xml:"channel>skipDays>days"`
    SkipDays            []time.Weekday
    Item                []*Rss20Item        `xml:"channel>item"`
}


func weekdayToNumber(s string) (week time.Weekday, ok bool) {
    ok = true
    switch strings.ToLower(s) {
        case "monday":
            week = time.Monday
            return
        case "tuesday":
            week = time.Tuesday
            return
        case "wednesday":
            week = time.Wednesday
            return
        case "thursday":
            week = time.Thursday
            return
        case "friday":
            week = time.Friday
            return
        case "saturday":
            week = time.Saturday
            return
        case "sunday":
            week = time.Sunday
            return
    }

    ok = false // error
    return
}


func Rss20Parse(b []byte) (rss *Rss20, err error) {

    err = xml.Unmarshal(b, &rss)
    if err != nil {
        return
    }

    if rss.Version != "2.0" {
        err = errors.New(fmt.Sprintf("RSS version: %s is not supported.",
            rss.Version))
    }

    rss.PubDate, _ = ParseTime(rss.PubDateRaw)
    rss.LastBuildDate, _ = ParseTime(rss.LastBuildDateRaw)

    days := set.New()
    for _, i := range(rss.SkipDaysRaw) {
        n, ok := weekdayToNumber(i)
        if ok {
            days.Add(n)
        }
    }
    daysArray := make([]time.Weekday, 0, days.Len())
    for _, i := range(days.List()) {
        v, _ := i.(time.Weekday)
        daysArray = append(daysArray, v)
    }
    rss.SkipDays = daysArray

    for i := range(rss.Item) {
        if rss.Item[i].Guid != nil {
            rss.Item[i].Guid.IsPermaLink = strings.ToLower(
                rss.Item[i].Guid.IsPermaLinkRaw) == "true"
        }

        rss.Item[i].PubDate, _ = ParseTime(rss.Item[i].PubDateRaw)

    }

    return
}


func Rss20ParseString(xmldata string) (*Rss20, error) {
    return Rss20Parse([]byte(xmldata))
}


