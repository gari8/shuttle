package shuttle

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	multipartTestText = `----------------------------997351631714338211531432
Content-Disposition: form-data; name="file"; filename="main.go"
Content-Type: application/octet-stream

%s

----------------------------997351631714338211531432
Content-Disposition: form-data; name="description"

%s
----------------------------997351631714338211531432--`
	goText = `package main

import (
        "github.com/gari8/sprinter"
        "tests/internal/tests/infrastructure/database/conf"
        "tests/internal/tests/infrastructure/server"
)

func main() {
        conn, err := conf.NewDatabaseConnection()
        if err != nil {
                panic(err)
        }
        s := server.NewServer(conn)
        sprinter.PrintLogo("GET http://localhost:8080/api/v1/sample", "POST http://localhost:8080/api/v1/sample")
        s.Serve()
}`
	testText = `nice to meet you !`
	boundary = `----------------------------997351631714338211531432`
)

func TestNew(t *testing.T) {
	expected := fmt.Sprintf(multipartTestText, goText, testText)
	sh := New(expected, boundary)
	assert.Equal(t, sh, Shuttle{
		Boundary: boundary,
		Text:     expected,
	})
}

func TestShuttle_Launch(t *testing.T) {
	sh := New(fmt.Sprintf(multipartTestText, goText, testText), boundary)
	assert.Equal(t, goText, sh.Launch("file"))
	assert.Equal(t, testText, sh.Launch("description"))
	assert.Equal(t, "", sh.Launch(""))
}
