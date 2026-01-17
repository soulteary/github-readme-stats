package cards

import (
	"fmt"
	"strings"

	"github.com/soulteary/github-readme-stats/internal/common"
	"github.com/soulteary/github-readme-stats/internal/fetchers"
	"github.com/soulteary/github-readme-stats/internal/translations"
)

const (
	iconSize             = 16
	descriptionLineWidth = 59
	descriptionMaxLines  = 3
)

// RepoCardOptions represents options for repo card
type RepoCardOptions struct {
	HideBorder            bool
	TitleColor            string
	IconColor             string
	TextColor             string
	BgColor               string
	ShowOwner             bool
	Theme                 string
	BorderRadius          float64
	BorderColor           string
	Locale                string
	DescriptionLinesCount int
}

// RenderRepoCard renders the repository card
func RenderRepoCard(repo *fetchers.RepoData, options RepoCardOptions) string {
	lineHeight := 10.0
	header := repo.Name
	if options.ShowOwner {
		header = repo.NameWithOwner
	}

	langName := "Unspecified"
	langColor := "#333"
	if repo.PrimaryLanguage != nil {
		langName = repo.PrimaryLanguage.Name
		langColor = repo.PrimaryLanguage.Color
		if langColor == "" {
			langColor = common.GetLanguageColor(langName)
		}
	}

	descMaxLines := descriptionMaxLines
	if options.DescriptionLinesCount > 0 {
		descMaxLines = int(common.ClampValue(options.DescriptionLinesCount, 1, descriptionMaxLines))
	}

	description := repo.Description
	if description == "" {
		description = "No description provided"
	}

	// Parse emojis in description
	description = common.ParseEmojis(description)

	// Wrap text based on pixel width
	// Card width is 400px, text starts at x=25, so available width is 375px
	// Font size is 13px
	maxTextWidth := 375.0
	fontSize := 13.0
	multiLineDescription := common.WrapTextByPixelWidth(description, maxTextWidth, fontSize, descMaxLines)
	descriptionLinesCount := len(multiLineDescription)
	if options.DescriptionLinesCount > 0 {
		descriptionLinesCount = int(common.ClampValue(options.DescriptionLinesCount, 1, descriptionMaxLines))
	}

	// Build description SVG
	// Note: lines are already HTML encoded in WrapTextByPixelWidth
	var descriptionSvg strings.Builder
	for i, line := range multiLineDescription {
		if i == 0 {
			// First line inherits position from parent <text> element
			descriptionSvg.WriteString(fmt.Sprintf(`<tspan>%s</tspan>`, line))
		} else {
			// Subsequent lines start at x="25" and move down by 1.5em
			descriptionSvg.WriteString(fmt.Sprintf(`<tspan dy="1.5em" x="25">%s</tspan>`, line))
		}
	}

	height := float64(110)
	if descriptionLinesCount > 1 {
		height = 120
	}
	height += float64(descriptionLinesCount) * lineHeight

	// Get translations
	locale := options.Locale
	if locale == "" {
		locale = "en"
	}
	i18n := translations.NewI18n(locale)

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
		colorArgs["theme"] = "default_repocard"
	}
	colors := common.GetCardColors(colorArgs)

	// Language node
	svgLanguage := ""
	if repo.PrimaryLanguage != nil {
		svgLanguage = common.CreateLanguageNode(langName, langColor)
	}

	// Stars and forks
	totalStars := common.KFormatter(float64(repo.Stargazers.TotalCount), nil)
	totalForks := common.KFormatter(float64(repo.Forks.TotalCount), nil)

	svgStars := common.IconWithLabel(
		common.Icons["star"],
		totalStars,
		"stargazers",
		iconSize,
	)
	svgForks := common.IconWithLabel(
		common.Icons["fork"],
		totalForks,
		"forkcount",
		iconSize,
	)

	// Layout stars and forks
	items := []string{svgLanguage, svgStars, svgForks}
	sizes := []float64{
		21.0 + common.MeasureText(langName, 12),                      // circle(6) + spacing(15) + text
		36.0 + common.MeasureText(fmt.Sprintf("%v", totalStars), 12), // icon(16) + gap(20) + text
		36.0 + common.MeasureText(fmt.Sprintf("%v", totalForks), 12), // icon(16) + gap(20) + text
	}

	starAndForkCount := strings.Join(common.FlexLayout(common.FlexLayoutProps{
		Items: items,
		Sizes: sizes,
		Gap:   5,
	}), "")

	// Calculate total content width for right alignment
	totalContentWidth := sizes[0] + 5 + sizes[1] + 5 + sizes[2]
	rightMargin := 15.0
	metadataStartX := 400.0 - totalContentWidth - rightMargin

	// Badges
	badges := []string{}
	if repo.IsArchived {
		badges = append(badges, getBadgeSVG(i18n.T("repocard.archived"), colors.TextColor))
	}
	if repo.IsTemplate {
		badges = append(badges, getBadgeSVG(i18n.T("repocard.template"), colors.TextColor))
	}

	// Create card
	card := NewCard(CardOptions{
		Width:        400,
		Height:       height,
		BorderRadius: options.BorderRadius,
		Colors:       colors,
		DefaultTitle: header,
		HideBorder:   options.HideBorder,
		CSS:          getRepoStyles(colors),
	})

	// Build body
	// Note: Card.Render() will automatically wrap this in <g data-testid="main-card-body">
	// and render the title separately via RenderTitle()
	body := fmt.Sprintf(`
      <g class="description" data-testid="description">
        <text
          class="gray"
          x="25"
          y="15"
          data-testid="description-text"
        >
          %s
        </text>
      </g>
      <g data-testid="metadata" style="display: block">
        <g transform="translate(%.0f, %d)">
          %s
        </g>
      </g>
      %s
  `, descriptionSvg.String(),
		metadataStartX,
		int(height-75),
		starAndForkCount,
		strings.Join(badges, ""))

	return card.Render(body)
}

func getBadgeSVG(label, textColor string) string {
	return fmt.Sprintf(`
  <g data-testid="badge" class="badge" transform="translate(320, -18)">
    <rect stroke="%s" stroke-width="1" width="70" height="20" x="-12" y="-14" ry="10" rx="10"></rect>
    <text
      x="23"
      y="-5"
      alignment-baseline="central"
      dominant-baseline="central"
      text-anchor="middle"
      fill="%s"
    >
      %s
    </text>
  </g>
`, textColor, textColor, label)
}

func getRepoStyles(colors common.CardColors) string {
	return fmt.Sprintf(`
    .header {
      font: 600 18px 'Segoe UI', Ubuntu, "Helvetica Neue", Sans-Serif; fill: %s;
      animation: fadeInAnimation 0.8s ease-in-out forwards;
    }
    @supports(-moz-appearance: auto) {
      /* Selector detects Firefox */
      .header { font-size: 15.5px; }
    }
    .description { font: 400 13px 'Segoe UI', Ubuntu, Sans-Serif; fill: %s; }
    .gray { font: 400 13px 'Segoe UI', Ubuntu, Sans-Serif; fill: %s; }
    .icon {
      fill: %s;
      display: block;
    }
  `, colors.TitleColor, colors.TextColor, colors.TextColor, colors.IconColor)
}
