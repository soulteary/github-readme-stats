package cards

import (
	"fmt"
	"strings"

	"github.com/soulteary/github-readme-stats/internal/common"
	"github.com/soulteary/github-readme-stats/internal/fetchers"
)

const (
	gistIconSize         = 16
	gistCardDefaultWidth = 400
	gistHeaderMaxLength  = 35
)

// GistCardOptions represents options for gist card
type GistCardOptions struct {
	TitleColor   string
	IconColor    string
	TextColor    string
	BgColor      string
	Theme        string
	BorderRadius float64
	BorderColor  string
	ShowOwner    bool
	HideBorder   bool
}

// RenderGistCard renders the gist card
func RenderGistCard(gist *fetchers.GistData, options GistCardOptions) string {
	lineWidth := 59
	linesLimit := 10

	description := gist.Description
	if description == "" {
		description = "No description provided"
	}

	// Parse emojis in description
	description = common.ParseEmojis(description)

	// Wrap text
	multiLineDescription := common.WrapTextMultiline(description, lineWidth, linesLimit)
	descriptionLines := len(multiLineDescription)
	descriptionSvg := ""
	for i, line := range multiLineDescription {
		if i == 0 {
			// First line inherits position from parent <text> element
			descriptionSvg += fmt.Sprintf(`<tspan>%s</tspan>`, common.EncodeHTML(line))
		} else {
			// Subsequent lines start at x="25" and move down by 1.2em
			descriptionSvg += fmt.Sprintf(`<tspan dy="1.2em" x="25">%s</tspan>`, common.EncodeHTML(line))
		}
	}

	lineHeight := 10
	if descriptionLines > 3 {
		lineHeight = 12
	}
	height := float64(110)
	if descriptionLines > 1 {
		height = 120
	}
	height += float64(descriptionLines) * float64(lineHeight)

	// Get card colors
	colorArgs := map[string]string{
		"title_color":  options.TitleColor,
		"text_color":   options.TextColor,
		"icon_color":   options.IconColor,
		"bg_color":     options.BgColor,
		"border_color": options.BorderColor,
		"theme":        options.Theme,
	}
	if colorArgs["theme"] == "" {
		colorArgs["theme"] = "default"
	}
	colors := common.GetCardColors(colorArgs)

	// Stars and forks
	totalStars := common.KFormatter(float64(gist.Stargazers.TotalCount), nil)
	totalForks := common.KFormatter(float64(gist.Forks.TotalCount), nil)

	svgStars := common.IconWithLabel(
		common.Icons["star"],
		totalStars,
		"starsCount",
		gistIconSize,
	)
	svgForks := common.IconWithLabel(
		common.Icons["fork"],
		totalForks,
		"forksCount",
		gistIconSize,
	)

	// Language
	languageName := gist.Name
	if languageName == "" {
		languageName = "Unspecified"
	}
	languageColor := common.GetLanguageColor(languageName)

	svgLanguage := common.CreateLanguageNode(languageName, languageColor)

	// Layout stars and forks
	items := []string{svgLanguage, svgStars, svgForks}
	sizes := []float64{
		common.MeasureText(languageName, 12),
		float64(gistIconSize) + common.MeasureText(fmt.Sprintf("%v", totalStars), 12),
		float64(gistIconSize) + common.MeasureText(fmt.Sprintf("%v", totalForks), 12),
	}

	starAndForkCount := strings.Join(common.FlexLayout(common.FlexLayoutProps{
		Items: items,
		Sizes: sizes,
		Gap:   25,
	}), "")

	// Header
	header := gist.Name
	if options.ShowOwner {
		header = fmt.Sprintf("%s/%s", gist.Owner.Login, gist.Name)
	}
	if len(header) > gistHeaderMaxLength {
		header = header[:gistHeaderMaxLength] + "..."
	}

	// Create card
	card := NewCard(CardOptions{
		Width:           gistCardDefaultWidth,
		Height:          height,
		BorderRadius:    options.BorderRadius,
		Colors:          colors,
		DefaultTitle:    header,
		TitlePrefixIcon: common.Icons["gist"],
		HideBorder:      options.HideBorder,
		CSS:             getGistStyles(colors),
	})

	// Build body
	body := fmt.Sprintf(`
    <text class="description" x="25" y="12">
        %s
    </text>

    <g transform="translate(30, %.0f)">
        %s
    </g>
  `, descriptionSvg, height-75, starAndForkCount)

	return card.Render(body)
}

func getGistStyles(colors common.CardColors) string {
	return fmt.Sprintf(`
    .description { font: 400 13px 'Segoe UI', Ubuntu, Sans-Serif; fill: %s }
    .gray { font: 400 12px 'Segoe UI', Ubuntu, Sans-Serif; fill: %s }
    .icon { fill: %s }
  `, colors.TextColor, colors.TextColor, colors.IconColor)
}
