package changelog

import (
	"fmt"

	"github.com/gomarkdown/markdown"
	htm "github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/microcosm-cc/bluemonday"
)

func ChangelogToHTML(changelogMD []byte) string {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(changelogMD)
	// create HTML renderer with extensions
	htmlFlags := htm.CommonFlags | htm.HrefTargetBlank
	opts := htm.RendererOptions{Flags: htmlFlags}
	renderer := htm.NewRenderer(opts)
	html := markdown.Render(doc, renderer)
	html = bluemonday.UGCPolicy().SanitizeBytes(html)
	return fmt.Sprintf("%s", html)
}
