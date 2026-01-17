package cards

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/soulteary/github-readme-stats/internal/common"
	"github.com/soulteary/github-readme-stats/internal/fetchers"
	"github.com/soulteary/github-readme-stats/internal/translations"
)

const (
	defaultCardWidth                     = 300
	minCardWidth                         = 280
	defaultLangColor                     = "#858585"
	cardPadding                          = 25
	compactLayoutBaseHeight              = 90
	maximumLangsCount                    = 20
	normalLayoutDefaultLangsCount        = 5
	compactLayoutDefaultLangsCount       = 6
	donutLayoutDefaultLangsCount         = 5
	pieLayoutDefaultLangsCount           = 6
	donutVerticalLayoutDefaultLangsCount = 6
)

// TopLanguagesCardOptions represents options for top languages card
type TopLanguagesCardOptions struct {
	Hide              []string
	HideTitle         bool
	HideBorder        bool
	CardWidth         int
	TitleColor        string
	TextColor         string
	BgColor           string
	Theme             string
	Layout            string
	LangsCount        int
	ExcludeRepo       []string
	SizeWeight        float64
	CountWeight       float64
	CustomTitle       string
	Locale            string
	BorderRadius      float64
	BorderColor       string
	DisableAnimations bool
	HideProgress      bool
	StatsFormat       string
}

// LanguageWithPercent represents language with percentage
type LanguageWithPercent struct {
	Name      string
	Size      int64
	Color     string
	Percent   float64
	RepoCount int
}

// RenderTopLanguages renders the top languages card
func RenderTopLanguages(langs map[string]*fetchers.LanguageStats, options TopLanguagesCardOptions) string {
	hide := options.Hide
	if hide == nil {
		hide = []string{}
	}

	layout := options.Layout
	if layout == "" {
		layout = "normal"
	}

	langsCount := options.LangsCount
	if langsCount <= 0 {
		switch layout {
		case "normal", "donut":
			langsCount = normalLayoutDefaultLangsCount
		default:
			langsCount = compactLayoutDefaultLangsCount
		}
	}
	if langsCount > maximumLangsCount {
		langsCount = maximumLangsCount
	}

	// Convert to slice, filter, and sort
	var langList []LanguageWithPercent
	hideMap := make(map[string]bool)
	for _, h := range hide {
		hideMap[strings.ToLower(h)] = true
	}

	// Collect languages (excluding hidden ones)
	for _, lang := range langs {
		if !hideMap[strings.ToLower(lang.Name)] {
			langList = append(langList, LanguageWithPercent{
				Name:      lang.Name,
				Size:      lang.Size,
				Color:     lang.Color,
				RepoCount: lang.RepoCount,
			})
		}
	}

	// Sort by size
	sort.Slice(langList, func(i, j int) bool {
		return langList[i].Size > langList[j].Size
	})

	// Limit to langsCount (trim after sorting)
	if len(langList) > langsCount {
		langList = langList[:langsCount]
	}

	// Calculate totalSize from trimmed languages
	totalSize := int64(0)
	for i := range langList {
		totalSize += langList[i].Size
	}

	// Calculate percentages based on trimmed totalSize
	for i := range langList {
		if totalSize > 0 {
			langList[i].Percent = float64(langList[i].Size) / float64(totalSize) * 100
		}
	}

	// Get card colors
	colorArgs := map[string]string{
		"title_color":  options.TitleColor,
		"text_color":   options.TextColor,
		"bg_color":     options.BgColor,
		"border_color": options.BorderColor,
		"theme":        options.Theme,
	}
	if colorArgs["theme"] == "" {
		colorArgs["theme"] = "default"
	}
	colors := common.GetCardColors(colorArgs)

	// Get translations
	locale := options.Locale
	if locale == "" {
		locale = "en"
	}
	i18n := translations.NewI18n(locale)

	// Calculate card dimensions
	cardWidth := float64(options.CardWidth)
	if cardWidth <= 0 {
		cardWidth = defaultCardWidth
	}
	if cardWidth < minCardWidth {
		cardWidth = minCardWidth
	}

	height := calculateLayoutHeight(layout, len(langList))

	// Build body based on layout
	var body string
	switch layout {
	case "compact":
		body = renderCompactLayout(langList, cardWidth, totalSize, colors, options.HideProgress, options.StatsFormat)
	case "donut":
		body = renderDonutLayout(langList, cardWidth, colors, options.HideProgress)
	case "donut-vertical":
		body = renderDonutVerticalLayout(langList, cardWidth, colors, options.HideProgress)
	case "pie":
		body = renderPieLayout(langList, cardWidth, colors, options.HideProgress)
	default: // normal
		body = renderNormalLayout(langList, cardWidth, colors, options.HideProgress, options.StatsFormat)
	}

	// Create card
	cardTitle := options.CustomTitle
	if cardTitle == "" {
		cardTitle = i18n.T("langcard.title")
	}

	card := NewCard(CardOptions{
		Width:             cardWidth,
		Height:            height,
		BorderRadius:      options.BorderRadius,
		Colors:            colors,
		CustomTitle:       cardTitle,
		DefaultTitle:      cardTitle,
		HideBorder:        options.HideBorder,
		HideTitle:         options.HideTitle,
		DisableAnimations: options.DisableAnimations,
		CSS:               getTopLanguagesStyles(colors),
	})

	// Wrap body in SVG with padding to match original JS implementation
	wrappedBody := fmt.Sprintf(`
    <svg data-testid="lang-items" x="%d">
      %s
    </svg>
  `, cardPadding, body)

	return card.Render(wrappedBody)
}

func calculateLayoutHeight(layout string, totalLangs int) float64 {
	switch layout {
	case "compact":
		return float64(compactLayoutBaseHeight + (totalLangs+1)/2*25)
	case "donut":
		return 215 + math.Max(float64(totalLangs-5), 0)*32
	case "donut-vertical":
		return 300 + float64((totalLangs+1)/2)*25
	case "pie":
		return 300 + float64((totalLangs+1)/2)*25
	default: // normal
		return 45 + float64(totalLangs+1)*40
	}
}

func renderNormalLayout(langs []LanguageWithPercent, cardWidth float64, colors common.CardColors, hideProgress bool, statsFormat string) string {
	var items []string
	// Calculate progress bar width and text position based on card width
	// Matching original JS implementation: paddingRight = 95
	paddingRight := 95.0
	progressWidth := cardWidth - paddingRight
	progressTextX := cardWidth - paddingRight + 10

	for i, lang := range langs {
		langName := lang.Name
		if lang.Color == "" {
			lang.Color = common.GetLanguageColor(langName)
			if lang.Color == "" {
				lang.Color = defaultLangColor
			}
		}

		staggerDelay := (i + 3) * 150

		progressBar := ""
		if !hideProgress {
			progressBar = common.CreateProgressNode(common.CreateProgressNodeParams{
				X:                          0,
				Y:                          25,
				Width:                      progressWidth,
				Color:                      lang.Color,
				Progress:                   lang.Percent,
				ProgressBarBackgroundColor: "#ddd",
				Delay:                      staggerDelay + 300,
			})
		}

		valueText := fmt.Sprintf("%.1f%%", lang.Percent)
		if statsFormat == "bytes" {
			formatted, _ := common.FormatBytes(lang.Size)
			valueText = formatted
		}

		item := fmt.Sprintf(`
      <g class="stagger" style="animation-delay: %dms">
        <text data-testid="lang-name" x="2" y="15" class="lang-name">%s</text>
        <text x="%.0f" y="34" class="lang-name" data-testid="lang-name">%s</text>
        %s
      </g>
    `, staggerDelay, common.EncodeHTML(langName), progressTextX, valueText, progressBar)

		items = append(items, item)
	}

	// Use FlexLayout to arrange items vertically with gap of 40
	layoutItems := common.FlexLayout(common.FlexLayoutProps{
		Items:     items,
		Gap:       40,
		Direction: "column",
	})

	return strings.Join(layoutItems, "")
}

// createCompactLangNode creates a compact language text node
func createCompactLangNode(lang LanguageWithPercent, totalSize int64, hideProgress bool, statsFormat string, index int) string {
	percentages := float64(lang.Size) / float64(totalSize) * 100
	var displayValue string
	if statsFormat == "bytes" {
		formatted, _ := common.FormatBytes(lang.Size)
		displayValue = formatted
	} else {
		displayValue = fmt.Sprintf("%.2f%%", percentages)
	}

	staggerDelay := (index + 3) * 150
	color := lang.Color
	if color == "" {
		color = defaultLangColor
	}

	textContent := common.EncodeHTML(lang.Name)
	if !hideProgress {
		textContent += " " + displayValue
	}

	return fmt.Sprintf(`
    <g class="stagger" style="animation-delay: %dms">
      <circle cx="5" cy="6" r="5" fill="%s" />
      <text data-testid="lang-name" x="15" y="10" class="lang-name">%s</text>
    </g>
  `, staggerDelay, color, textContent)
}

// createLanguageTextNode creates two-column layout for language text nodes
func createLanguageTextNode(langs []LanguageWithPercent, totalSize int64, hideProgress bool, statsFormat string) string {
	// Find longest language name
	longestLang := langs[0]
	for _, lang := range langs {
		if len(lang.Name) > len(longestLang.Name) {
			longestLang = lang
		}
	}

	// Split languages into two columns
	perChunk := len(langs) / 2
	if perChunk == 0 {
		perChunk = 1
	}
	chunked := common.ChunkArray(langs, perChunk)

	// Create layouts for each column
	var layouts []string
	for colIdx, column := range chunked {
		var items []string
		for i, lang := range column {
			item := createCompactLangNode(lang, totalSize, hideProgress, statsFormat, colIdx*perChunk+i)
			items = append(items, item)
		}
		// Use FlexLayout to arrange items vertically with gap of 25
		layoutItems := common.FlexLayout(common.FlexLayoutProps{
			Items:     items,
			Gap:       25,
			Direction: "column",
		})
		layouts = append(layouts, strings.Join(layoutItems, ""))
	}

	// Calculate gap between columns based on longest language name
	percent := (float64(longestLang.Size) / float64(totalSize) * 100)
	var displayValue string
	if statsFormat == "bytes" {
		formatted, _ := common.FormatBytes(longestLang.Size)
		displayValue = formatted
	} else {
		displayValue = fmt.Sprintf("%.2f%%", percent)
	}
	longestText := longestLang.Name + " " + displayValue
	textWidth := common.MeasureText(longestText, 11)

	minGap := 150.0
	maxGap := 20.0 + textWidth
	gap := maxGap
	if gap < minGap {
		gap = minGap
	}

	// For two-column layout, manually position columns
	// First column at x=0, second column at x=gap
	var result strings.Builder
	if len(layouts) > 0 {
		result.WriteString(fmt.Sprintf(`<g transform="translate(0, 0)">%s</g>`, layouts[0]))
		if len(layouts) > 1 {
			result.WriteString(fmt.Sprintf(`<g transform="translate(%.0f, 0)">%s</g>`, gap, layouts[1]))
		}
	}

	return result.String()
}

func renderCompactLayout(langs []LanguageWithPercent, cardWidth float64, totalSize int64, colors common.CardColors, hideProgress bool, statsFormat string) string {
	paddingRight := 50.0
	offsetWidth := cardWidth - paddingRight

	// Create horizontal stacked progress bar
	var compactProgressBar strings.Builder
	if !hideProgress {
		progressOffset := 0.0
		for _, lang := range langs {
			percentage := (float64(lang.Size) / float64(totalSize)) * offsetWidth
			progress := percentage
			if progress < 10 {
				progress = percentage + 10
			}

			compactProgressBar.WriteString(fmt.Sprintf(`
        <rect
          mask="url(#rect-mask)"
          data-testid="lang-progress"
          x="%.2f"
          y="0"
          width="%.2f"
          height="8"
          fill="%s"
        />
      `, progressOffset, progress, lang.Color))
			progressOffset += percentage
		}
	}

	progressBarSection := ""
	if !hideProgress {
		progressBarSection = fmt.Sprintf(`
      <mask id="rect-mask">
        <rect x="0" y="0" width="%.0f" height="8" fill="white" rx="5"/>
      </mask>
      %s
    `, offsetWidth, compactProgressBar.String())
	}

	// Create language text nodes with two-column layout
	textYOffset := 0.0
	if !hideProgress {
		textYOffset = 25.0
	}

	languageTextNode := createLanguageTextNode(langs, totalSize, hideProgress, statsFormat)

	return fmt.Sprintf(`
  %s
  <g transform="translate(0, %.0f)">
    %s
  </g>
  `, progressBarSection, textYOffset, languageTextNode)
}

func renderDonutLayout(langs []LanguageWithPercent, cardWidth float64, colors common.CardColors, hideProgress bool) string {
	// Simplified donut layout - full implementation would require SVG path calculations
	return renderNormalLayout(langs, cardWidth, colors, hideProgress, "percentages")
}

func renderDonutVerticalLayout(langs []LanguageWithPercent, cardWidth float64, colors common.CardColors, hideProgress bool) string {
	// Simplified donut vertical layout
	return renderNormalLayout(langs, cardWidth, colors, hideProgress, "percentages")
}

func renderPieLayout(langs []LanguageWithPercent, cardWidth float64, colors common.CardColors, hideProgress bool) string {
	// Simplified pie layout
	return renderNormalLayout(langs, cardWidth, colors, hideProgress, "percentages")
}

func getTopLanguagesStyles(colors common.CardColors) string {
	return fmt.Sprintf(`
    .lang-name {
      font: 400 11px 'Segoe UI', Ubuntu, Sans-Serif; fill: %s;
    }
    .stagger {
      opacity: 0;
      animation: fadeInAnimation 0.3s ease-in-out forwards;
    }
    .lang-progress {
      animation: growWidthAnimation 0.6s ease-in-out forwards;
    }
    #rect-mask rect {
      animation: slideInAnimation 1s ease-in-out forwards;
    }
    @keyframes progressAnimation {
      from {
        width: 0;
      }
    }
    @keyframes slideInAnimation {
      from {
        width: 0;
      }
      to {
        width: calc(100%%-100px);
      }
    }
    @keyframes growWidthAnimation {
      from {
        width: 0;
      }
      to {
        width: 100%%;
      }
    }
    @keyframes fadeInAnimation {
      from {
        opacity: 0;
      }
      to {
        opacity: 1;
      }
    }
  `, colors.TextColor)
}
