package sofortpay

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPrepareError(t *testing.T) {
	j := `{"code":400,"message":"Bad Request"}`

	b := ioutil.NopCloser(bytes.NewReader([]byte(j)))
	defer b.Close()

	r := &http.Response{
		StatusCode: 200,
		Body:       b,
	}

	err := prepareError(r)

	if err == nil {
		t.Fail()
	}

	if err.Error() != "sofort pay client: Bad Request" {
		t.Fail()
	}
}

func BenchmarkPrepareError(b *testing.B) {
	j := `{"code":400,"message":"Bad Request"}`

	for n := 0; n < b.N; n++ {
		prepareError(&http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(j))),
		})
	}
}
