package server

import (
	"io/ioutil"
	"net/url"
	"strings"
	"testing"
)

func testMakeProperties() interface{} {
	return &struct {
		Annotators   string `json:"annotators"`
		OutputFormat string `json:"outputFormat"`
	}{
		Annotators:   "tokenize,ssplit,pos",
		OutputFormat: "json",
	}

}

func TestCoreNlpClient_MakeUrl(t *testing.T) {
	port := 9000

	client := CoreNlpClient(NewCoreNlpParameter("./local-server", port))

	exist, err := client.makeUrl(testMakeProperties())
	if err != nil {
		t.Error("failed to make URL with", err)
	}

	expected := "http://localhost:9000/?properties=" + url.QueryEscape(`{"annotators":"tokenize,ssplit,pos","outputFormat":"json"}`)

	if strings.Compare(exist, expected) != 0 {
		t.Error("failed to make URL. Got\n'", exist, "', but expected is\n'", expected, "'")
	}
}

func TestCoreNlpClient_DoString(t *testing.T) {
	resp, err := CoreNlpClient(LocalServer()).DoString(testMakeProperties(), "the quick brown fox jumped over the lazy dog")
	if err != nil {
		t.Error("failed to execute request with", err)
	}

	if resp == nil {
		t.Error("failed to get response. The body is nil")
	} else {
		expected := `{"sentences":`

		b, _ := ioutil.ReadAll(resp.Body)
		if body := string(b); !strings.HasPrefix(body, expected) {
			t.Error("failed to get expected result. Got '", body[:16], "...', but expected should have a prefix '", expected, "'")
		}
	}
}
