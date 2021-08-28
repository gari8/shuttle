package main

import (
	"fmt"
	"github.com/gari8/shuttle"
)

const (
	multipartText = `----------------------------997351631714338211531432
Content-Disposition: form-data; name="file"; filename="main.go"
Content-Type: application/octet-stream

package main

import (
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
}

----------------------------997351631714338211531432
Content-Disposition: form-data; name="description"

nice to meet you !
----------------------------997351631714338211531432--`
	boundary = `----------------------------997351631714338211531432`
)

func main()  {
	sh := shuttle.New(multipartText, boundary)
	fmt.Printf("========%+v========\n", sh.Launch("description"))
	fmt.Printf("========%+v========\n", sh.Launch("file"))
}
