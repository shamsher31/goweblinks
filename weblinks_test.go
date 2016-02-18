package weblinks

import "testing"
import "strings"

func TestGet(t *testing.T) {

	reader := strings.NewReader(`<p>
    This webpage contains lot of links
    <a href="http://google.com"></a>
    <a href="http://github.com/shamsher31"></a>
    <a href="http://about.me/shamsher"></a>
  </p>`)

	links := Get(reader)

	if len(links) != 3 {
		t.Error("Incorrect number of URL returned")
	}

}
