// Code generated by templ DO NOT EDIT.

package teststring

import "github.com/a-h/templ"
import "context"
import "io"

func render(s string) (t templ.Component) {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, templ.EscapeString(s))
		if err != nil {
			return err
		}
		return err
	})
}
