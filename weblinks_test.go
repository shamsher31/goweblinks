package weblinks

import (
	"strings"
	"testing"
)

var reader = strings.NewReader(`<p>
  This webpage contains lot of links
  <a href="http://google.com"></a>
  <a href="http://github.com/shamsher31"></a>
  <a href="http://about.me/shamsher"></a>
</p>`)

// go test -v
func TestGet(t *testing.T) {

	links := Get(reader)

	if len(links) != 3 {
		t.Error("Incorrect number of URL returned")
	}

}

// go test -test.bench=".*"
func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get(reader)
	}
}
