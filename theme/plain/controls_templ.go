// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.639
package plain

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import (
	"bytes"
	"context"
	"io"

	"github.com/a-h/templ"

	"github.com/cugu/templum"
)

func menu_toggle() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}

		ctx = templ.InitializeContext(ctx)

		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}

		ctx = templ.ClearChildren(ctx)

		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"menu-toggle md:hidden cursor-pointer\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		templ_7745c5c3_Err = hamburger().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}

		return templ_7745c5c3_Err
	})
}

func github_link(c *templum.PageContext) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}

		ctx = templ.InitializeContext(ctx)

		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}

		ctx = templ.ClearChildren(ctx)

		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		var templ_7745c5c3_Var3 templ.SafeURL = templ.URL(c.Config["github_url"])

		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var3)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"hover:bg-gray-100 dark:hover:bg-gray-700 border dark:border-gray-600 font-bold py-2 px-4 rounded-lg flex ml-2 flex-row space-x-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		templ_7745c5c3_Err = github().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span>GitHub</span></a>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}

		return templ_7745c5c3_Err
	})
}
