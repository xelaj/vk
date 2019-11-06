package methods

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
	types "github.com/xelaj/vk/types"
	"golang.org/x/text/encoding/charmap"
)

type UsersGetRegDateResponse int

func UsersGetRegDate(c types.Client, userId int) (*UsersGetRegDateResponse, error) {
	resp, err := http.Get("https://vk.com/foaf.php?id=" + strconv.Itoa(userId))
	if err != nil {
		return nil, errors.Wrap(err, "request error")
	}

	defer resp.Body.Close()

	f, err := parseFrom(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "foaf parsing error")
	}
	res, err := time.Parse("2018-09-09T17:34:51+03:00", f.Person.Created.Date)
	if err != nil {
		return nil, errors.Wrap(err, "time parsing error")
	}
	d := UsersGetRegDateResponse(res.Unix())
	return &d, nil
}

type object struct {
	XMLName xml.Name `xml:"RDF"`
	Person  struct {
		XMLName xml.Name `xml:"Person"`
		Created struct {
			Date string `xml:"date,attr"`
		} `xml:"created"`
	} `xml:"Person"`
}

func parseFrom(r io.Reader) (*object, error) {
	d := xml.NewDecoder(r)
	d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251", "WINDOWS-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("unknown charset: %s", charset)
		}
	}

	var o object
	err := d.Decode(&o)

	return &o, err
}
