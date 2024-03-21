// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.639
package blog

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import (
	"bytes"
	"context"
	"io"

	"github.com/a-h/templ"

	"github.com/cugu/templum"
)

func html(c *templum.PageContext, data string) templ.Component {
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

		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html><head><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		var templ_7745c5c3_Var2 string

		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(c.Page.Title())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `theme/blog/html.templ`, Line: 8, Col: 26}
		}

		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><link href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		var templ_7745c5c3_Var3 string

		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(c.BaseURL + "favicon.ico")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `theme/blog/html.templ`, Line: 11, Col: 41}
		}

		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" rel=\"icon\"><link rel=\"stylesheet\" href=\"https://unpkg.com/normalize.css@8.0.1/normalize.css\"><link rel=\"stylesheet\" href=\"https://unpkg.com/concrete.css@2.1.1/concrete.css\"><link href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		var templ_7745c5c3_Var4 string

		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(c.BaseURL + "style.css")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `theme/blog/html.templ`, Line: 14, Col: 39}
		}

		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" rel=\"stylesheet\"></head><body><main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		if c.Page.Order() == 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<header>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}

			if title, ok := c.Config["title"]; ok {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}

				var templ_7745c5c3_Var5 string

				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(title)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `theme/blog/html.templ`, Line: 21, Col: 18}
				}

				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}

				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}

			if description, ok := c.Config["description"]; ok {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}

				templ_7745c5c3_Err = templ.Raw(description).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}

				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}

			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</header>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}

			templ_7745c5c3_Err = links(c.Pages).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			if title, ok := c.Config["title"]; ok {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a href=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}

				var templ_7745c5c3_Var6 templ.SafeURL = templ.URL(c.BaseURL)

				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var6)))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}

				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><h1>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}

				var templ_7745c5c3_Var7 string

				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(title)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `theme/blog/html.templ`, Line: 33, Col: 18}
				}

				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}

				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1></a>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}

			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}

			var templ_7745c5c3_Var8 templ.SafeURL = templ.URL(c.BaseURL)

			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var8)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}

			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Home</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}

		templ_7745c5c3_Err = templ.Raw(data).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		if c.Page.Order() != 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<footer><hr><p><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}

			var templ_7745c5c3_Var9 templ.SafeURL = templ.URL(c.BaseURL)

			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var9)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}

			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Go back</a></p></footer>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}

		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</main></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}

		return templ_7745c5c3_Err
	})
}

func links(pages []*templum.Page) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}

		ctx = templ.InitializeContext(ctx)

		templ_7745c5c3_Var10 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var10 == nil {
			templ_7745c5c3_Var10 = templ.NopComponent
		}

		ctx = templ.ClearChildren(ctx)

		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		for _, page := range pages {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}

			var templ_7745c5c3_Var11 templ.SafeURL = templ.URL(page.Link())

			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var11)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}

			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}

			var templ_7745c5c3_Var12 string

			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(page.Title())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `theme/blog/html.templ`, Line: 61, Col: 19}
			}

			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}

			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}

			if len(page.Children()) > 0 {
				templ_7745c5c3_Err = links(page.Children()).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}

			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}

		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}

		return templ_7745c5c3_Err
	})
}
