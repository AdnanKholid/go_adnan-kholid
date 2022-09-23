## install framework echo-api
1. membuat folder di file explorer laptop dengan nama folder "echo-api-A"
2. membuka folder yang sudah dibuat tadi ke vscode langsung masuk keterminal
3. kemudian ketik "go mod init echo-api-A"
4. dan go get "github.com/labstack/echo/v4"
5. buat file "main.go" lalu ketik 
6. package main

    import (
    "net/http"
 
    "github.com/labstack/echo/v4"
    )

    func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
    })
     e.Logger.Fatal(e.Start(":1323"))
    }
7. terakhir untuk mengecek framework sudah berhasil diinstall dengan browser ketik "http://localhost:1323/"