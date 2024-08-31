// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func App() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Your Web App Title</title><meta name=\"description\" content=\"A brief description of your web application\"><!-- Scripts --><script src=\"static/js/telegram-web-app.js\"></script><script defer src=\"static/js/alpine.min.js\"></script><!-- Styles --><link rel=\"stylesheet\" href=\"styles.css\"><!-- Favicon --><link rel=\"icon\" type=\"image/png\" href=\"path/to/favicon.png\"><!-- Open Graph Meta Tags --><meta property=\"og:title\" content=\"Your Web App Title\"><meta property=\"og:description\" content=\"A brief description of your web application\"><meta property=\"og:image\" content=\"path/to/og-image.jpg\"><meta property=\"og:url\" content=\"https://yourwebsite.com\"><!-- Twitter Card Meta Tags --><meta name=\"twitter:card\" content=\"summary_large_image\"></head><body><header x-data=\"{ message: &#39;Hello, World! 🤖 &#39; }\"><h1 class=\"text-3xl\">Your Web App Title</h1><h1 x-text=\"message\"></h1><input type=\"text\" x-model=\"message\"><!-- Your header content --></header><nav><!-- Your navigation menu --></nav><main x-data x-init=\"console.log(&#39;Step 1: Initial window.Telegram object&#39;);\n                    console.log(window.Telegram);\"><!-- Your main content --></main><footer><!-- Your footer content --></footer></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}