package cards

import (
	"fmt"
	"strings"

	"github.com/soulteary/github-readme-stats/internal/common"
	"github.com/soulteary/github-readme-stats/internal/fetchers"
	"github.com/soulteary/github-readme-stats/internal/translations"
)

const (
	wakaTimeDefaultCardWidth                = 495
	wakaTimeMinCardWidth                    = 250
	wakaTimeCompactLayoutMinWidth           = 400
	wakaTimeDefaultLineHeight               = 25
	wakaTimeProgressbarPadding              = 130
	wakaTimeHiddenProgressbarPadding        = 170
	wakaTimeCompactLayoutProgressbarPadding = 25
	wakaTimeTotalTextWidth                  = 275
)

// WakaTimeCardOptions represents options for WakaTime card
type WakaTimeCardOptions struct {
	Hide              []string
	HideTitle         bool
	CardWidth         int
	LineHeight        int
	HideProgress      bool
	CustomTitle       string
	Layout            string
	LangsCount        int
	ApiDomain         string
	DisplayFormat     string
	TitleColor        string
	TextColor         string
	BgColor           string
	Theme             string
	BorderRadius      float64
	BorderColor       string
	DisableAnimations bool
	Locale            string
}

// RenderWakaTimeCard renders the WakaTime card
func RenderWakaTimeCard(wakatime *fetchers.WakaTimeData, options WakaTimeCardOptions) string {
	hide := options.Hide
	if hide == nil {
		hide = []string{}
	}

	layout := options.Layout
	if layout == "" {
		layout = "default"
	}

	displayFormat := options.DisplayFormat
	if displayFormat == "" {
		displayFormat = "time"
	}

	lineHeight := options.LineHeight
	if lineHeight == 0 {
		lineHeight = wakaTimeDefaultLineHeight
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

	// Filter languages
	var filteredLangs []fetchers.WakaTimeLanguage
	hideMap := make(map[string]bool)
	for _, h := range hide {
		hideMap[strings.ToLower(h)] = true
	}

	for _, lang := range wakatime.Languages {
		if !hideMap[strings.ToLower(lang.Name)] {
			filteredLangs = append(filteredLangs, lang)
		}
	}

	// Limit languages count
	if options.LangsCount > 0 && len(filteredLangs) > options.LangsCount {
		filteredLangs = filteredLangs[:options.LangsCount]
	}

	// Filter languages that have hours or minutes
	var validLangs []fetchers.WakaTimeLanguage
	for _, lang := range filteredLangs {
		if lang.Hours > 0 || lang.Minutes > 0 {
			validLangs = append(validLangs, lang)
		}
	}
	filteredLangs = validLangs

	// Recalculate percentages
	recalculatePercentages(filteredLangs)

	// Calculate card dimensions
	cardWidth := float64(options.CardWidth)
	if cardWidth <= 0 {
		cardWidth = wakaTimeDefaultCardWidth
	}
	if cardWidth < wakaTimeMinCardWidth {
		cardWidth = wakaTimeMinCardWidth
	}

	height := calculateWakaTimeHeight(layout, len(filteredLangs), lineHeight)

	// Build body based on layout
	var body string
	if layout == "compact" {
		body = renderWakaTimeCompactLayout(filteredLangs, colors, options.HideProgress, displayFormat, cardWidth, lineHeight, wakatime, i18n)
	} else {
		body = renderWakaTimeDefaultLayout(filteredLangs, colors, options.HideProgress, displayFormat, cardWidth, lineHeight, wakatime, i18n)
	}

	// Create card
	cardTitle := options.CustomTitle
	if cardTitle == "" {
		cardTitle = i18n.T("wakatimecard.title")
		// Add range suffix if available
		if wakatime.Range == "last_7_days" {
			cardTitle += " (" + i18n.T("wakatimecard.last7days") + ")"
		} else if wakatime.Range == "last_year" {
			cardTitle += " (" + i18n.T("wakatimecard.lastyear") + ")"
		}
	}

	card := NewCard(CardOptions{
		Width:             cardWidth,
		Height:            height,
		BorderRadius:      options.BorderRadius,
		Colors:            colors,
		CustomTitle:       cardTitle,
		DefaultTitle:      cardTitle,
		HideBorder:        false,
		HideTitle:         options.HideTitle,
		DisableAnimations: options.DisableAnimations,
		CSS:               getWakaTimeStyles(colors),
	})

	return card.Render(body)
}

func calculateWakaTimeHeight(layout string, totalLangs int, lineHeight int) float64 {
	if layout == "compact" {
		return float64(90 + (totalLangs+1)/2*lineHeight)
	}
	return float64(45 + (totalLangs+1)*lineHeight)
}

func renderWakaTimeDefaultLayout(langs []fetchers.WakaTimeLanguage, colors common.CardColors, hideProgress bool, displayFormat string, cardWidth float64, lineHeight int, wakatime *fetchers.WakaTimeData, i18n *translations.I18n) string {
	var items []string

	// Progress bar width should be cardWidth - TOTAL_TEXT_WIDTH (275)
	progressBarWidth := int(cardWidth) - wakaTimeTotalTextWidth

	for i, lang := range langs {
		valueText := lang.Time
		if displayFormat == "percent" {
			valueText = fmt.Sprintf("%.2f %%", lang.Percent)
		}

		// Calculate text x position based on hideProgress
		// When progress bar is shown: PROGRESSBAR_PADDING (130) + progressBarWidth
		// When progress bar is hidden: HIDDEN_PROGRESSBAR_PADDING (170)
		textX := wakaTimeProgressbarPadding + progressBarWidth
		if hideProgress {
			textX = wakaTimeHiddenProgressbarPadding
		}

		staggerDelay := (i + 3) * 150
		progressBar := ""
		if !hideProgress {
			progressBar = common.CreateProgressNode(common.CreateProgressNodeParams{
				X:                          110,
				Y:                          4,
				Width:                      float64(progressBarWidth),
				Color:                      colors.TitleColor,
				Progress:                   lang.Percent,
				ProgressBarBackgroundColor: colors.TextColor,
				Delay:                      staggerDelay + 300,
			})
		}

		item := fmt.Sprintf(`
      <g class="stagger" style="animation-delay: %dms" transform="translate(25, 0)">
        <text class="stat bold" y="12.5">%s:</text>
        <text class="stat" x="%d" y="12.5" data-testid="lang-name">%s</text>
        %s
      </g>
    `, staggerDelay, common.EncodeHTML(lang.Name), textX, valueText, progressBar)

		items = append(items, item)
	}

	// Handle empty languages
	if len(langs) == 0 {
		var message string
		if wakatime.IsCodingActivityVisible {
			if wakatime.IsOtherUsageVisible {
				message = i18n.T("wakatimecard.nocodingactivity")
			} else {
				message = i18n.T("wakatimecard.nocodedetails")
			}
		} else {
			message = i18n.T("wakatimecard.notpublic")
		}
		return fmt.Sprintf(`
      <text x="25" y="11" class="stat bold">%s</text>
    `, common.EncodeHTML(message))
	}

	// Use FlexLayout for proper spacing
	layoutItems := common.FlexLayout(common.FlexLayoutProps{
		Items:     items,
		Gap:       float64(lineHeight),
		Direction: "column",
	})

	return strings.Join(layoutItems, "")
}

func renderWakaTimeCompactLayout(langs []fetchers.WakaTimeLanguage, colors common.CardColors, hideProgress bool, displayFormat string, cardWidth float64, lineHeight int, wakatime *fetchers.WakaTimeData, i18n *translations.I18n) string {
	width := cardWidth - 5

	// Handle empty languages
	if len(langs) == 0 {
		var message string
		if wakatime.IsCodingActivityVisible {
			if wakatime.IsOtherUsageVisible {
				message = i18n.T("wakatimecard.nocodingactivity")
			} else {
				message = i18n.T("wakatimecard.nocodedetails")
			}
		} else {
			message = i18n.T("wakatimecard.notpublic")
		}
		return fmt.Sprintf(`
      <text x="25" y="11" class="stat bold">%s</text>
    `, common.EncodeHTML(message))
	}

	// Build progress bar (only if not hidden)
	var progressBarSection string
	var textYOffset int
	if !hideProgress {
		var progressBars []string
		progressOffset := 0.0
		for _, lang := range langs {
			progress := ((width - float64(wakaTimeCompactLayoutProgressbarPadding)) * lang.Percent) / 100.0
			langColor := common.GetLanguageColor(lang.Name)

			progressBar := fmt.Sprintf(`
          <rect
            mask="url(#rect-mask)"
            data-testid="lang-progress"
            x="%.2f"
            y="0"
            width="%.2f"
            height="8"
            fill="%s"
          />
        `, progressOffset, progress, langColor)

			progressBars = append(progressBars, progressBar)
			progressOffset += progress
		}

		maskWidth := width - 2*float64(wakaTimeCompactLayoutProgressbarPadding)
		progressBarSection = fmt.Sprintf(`
      <mask id="rect-mask">
      <rect x="%d" y="0" width="%.2f" height="8" fill="white" rx="5" />
      </mask>
      %s
    `, wakaTimeCompactLayoutProgressbarPadding, maskWidth, strings.Join(progressBars, ""))
		textYOffset = 25
	} else {
		textYOffset = 0
	}

	// Build language text nodes
	var items []string
	leftX := 25
	rightXBase := 230
	rightOffset := int(cardWidth-float64(wakaTimeDefaultCardWidth)) / 2
	rightX := rightXBase + rightOffset

	for i, lang := range langs {
		langColor := common.GetLanguageColor(lang.Name)

		valueText := lang.Time
		if displayFormat == "percent" {
			valueText = fmt.Sprintf("%.2f %%", lang.Percent)
		}

		isLeft := i%2 == 0
		x := leftX
		if !isLeft {
			x = rightX
		}
		y := lineHeight * (i / 2)

		item := fmt.Sprintf(`
      <g transform="translate(%d, %d)">
        <circle cx="5" cy="6" r="5" fill="%s" />
        <text data-testid="lang-name" x="15" y="10" class='lang-name'>
          %s - %s
        </text>
      </g>
    `, x, y, langColor, common.EncodeHTML(lang.Name), valueText)

		items = append(items, item)
	}

	// Combine progress bar and language text
	return fmt.Sprintf(`
      %s
      <g transform="translate(0, %d)">
        %s
      </g>
    `, progressBarSection, textYOffset, strings.Join(items, ""))
}

func getWakaTimeStyles(colors common.CardColors) string {
	return fmt.Sprintf(`
    .stat {
      font: 600 14px 'Segoe UI', Ubuntu, "Helvetica Neue", Sans-Serif; fill: %s;
    }
    @supports(-moz-appearance: auto) {
      /* Selector detects Firefox */
      .stat { font-size:12px; }
    }
    .lang-name {
      font: 400 11px 'Segoe UI', Ubuntu, Sans-Serif; fill: %s;
    }
    .stagger {
      opacity: 0;
      animation: fadeInAnimation 0.3s ease-in-out forwards;
    }
    .bold { font-weight: 700 }
    .lang-progress {
      animation: progressAnimation 0.6s ease-in-out forwards;
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
    #rect-mask rect{
      animation: slideInAnimation 1s ease-in-out forwards;
    }
  `, colors.TextColor, colors.TextColor)
}

// recalculatePercentages recalculates percentages so that they sum to 100%
// This is needed when languages are filtered or limited
func recalculatePercentages(languages []fetchers.WakaTimeLanguage) {
	totalSum := 0.0
	for _, lang := range languages {
		totalSum += lang.Percent
	}
	if totalSum > 0 {
		weight := 100.0 / totalSum
		for i := range languages {
			languages[i].Percent = languages[i].Percent * weight
		}
	}
}
