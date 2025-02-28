package util

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func FormatPrice(price int64) (fprice string) {
	chunks := make([]int64, 0, 3)

	for price > 0 {
		chunks = append(chunks, price%1000)
		price = price / 1000
	}

	for i := len(chunks) - 1; i >= 0; i-- {
		if i == 0 {
			fprice += fmt.Sprint(chunks[i])
			break
		}

		fprice += fmt.Sprintf("%d,", chunks[i])
	}

	return
}
