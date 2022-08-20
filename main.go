// App for uploading image using echo4
package main

import (
	"fmt"
	// "log"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func upload(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	//-----------
	// Read file
	//-----------

	// Source
	// -Single file, err := c.FormFile("file")
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]

	resStr := "<p><pre>"
	resStr += "Upload report:\n"

	for i, file := range files {
		// FF -- For-Single-file ------------------------ ___--\\
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dstFolder := "up/"
		dst, err := os.Create(dstFolder + file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		resStr += fmt.Sprintf("File %d : ", i) + dst.Name() + " upload success!\n"

		// LL __ For-Single-file ________________________ ___--//
	}
	resStr += "</pre></p>"

	return c.HTML(http.StatusOK, fmt.Sprintf(resStr+"<p>Uploaded total  %d files with fields name=%s and email=%s.</p>", len(files), name, email))

}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")
	e.POST("/upload", upload)

	e.Logger.Fatal(e.Start(":1323"))
}
