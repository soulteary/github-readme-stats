package cards

import (
	"fmt"
	"math"
	"strings"

	"github.com/soulteary/github-readme-stats/internal/common"
	"github.com/soulteary/github-readme-stats/internal/fetchers"
	"github.com/soulteary/github-readme-stats/internal/translations"
)

const (
	cardMinWidth             = 287
	cardDefaultWidth         = 287
	rankCardMinWidth         = 420
	rankCardDefaultWidth     = 450
	rankOnlyCardMinWidth     = 290
	rankOnlyCardDefaultWidth = 290
)

// Long locales that need more space for text
var longLocales = []string{
	"az", "bg", "cs", "de", "el", "es", "fil", "fi", "fr", "hu", "id", "ja",
	"ml", "my", "nl", "pl", "pt-br", "pt-pt", "ru", "sr", "sr-latn", "sw",
	"ta", "uk-ua", "uz", "zh-tw",
}

// StatsCardOptions represents options for stats card
type StatsCardOptions struct {
	Hide              []string
	ShowIcons         bool
	HideTitle         bool
	HideBorder        bool
	CardWidth         int
	HideRank          bool
	IncludeAllCommits bool
	CommitsYear       *int
	LineHeight        int
	TitleColor        string
	RingColor         string
	IconColor         string
	TextColor         string
	TextBold          bool
	BgColor           string
	Theme             string
	CustomTitle       string
	BorderRadius      float64
	BorderColor       string
	NumberFormat      string
	NumberPrecision   int
	Locale            string
	DisableAnimations bool
	RankIcon          string
	Show              []string
}

// RenderStatsCard renders the stats card
func RenderStatsCard(stats *fetchers.StatsData, options StatsCardOptions) string {
	hide := options.Hide
	if hide == nil {
		hide = []string{}
	}
	show := options.Show
	if show == nil {
		show = []string{}
	}

	lineHeight := options.LineHeight
	if lineHeight == 0 {
		lineHeight = 25
	}

	numberFormat := options.NumberFormat
	if numberFormat == "" {
		numberFormat = "short"
	}

	theme := options.Theme
	if theme == "" {
		theme = "default"
	}

	locale := options.Locale
	if locale == "" {
		locale = "en"
	}

	rankIcon := options.RankIcon
	if rankIcon == "" {
		rankIcon = "default"
	}

	// Get card colors
	colorArgs := map[string]string{
		"title_color":  options.TitleColor,
		"text_color":   options.TextColor,
		"icon_color":   options.IconColor,
		"bg_color":     options.BgColor,
		"border_color": options.BorderColor,
		"ring_color":   options.RingColor,
		"theme":        theme,
	}
	colors := common.GetCardColors(colorArgs)

	// Get translations
	i18n := translations.NewI18n(locale)

	// Check if locale is long locale (needs more space)
	isLongLocale := false
	for _, longLocale := range longLocales {
		if locale == longLocale {
			isLongLocale = true
			break
		}
	}

	// Calculate shiftValuePos for value positioning
	shiftValuePos := 79.01
	if isLongLocale {
		shiftValuePos += 50
	}

	// Build stats items
	statItems := []string{}

	// Stars
	if !containsString(hide, "stars") {
		statItems = append(statItems, createStatItem(
			common.Icons["star"],
			i18n.T("statcard.totalstars"),
			float64(stats.TotalStars),
			"stars",
			"",
			0,
			options.ShowIcons,
			options.TextBold,
			numberFormat,
			options.NumberPrecision,
			shiftValuePos,
		))
	}

	// Commits
	if !containsString(hide, "commits") {
		commitsLabel := i18n.T("statcard.commits")
		if !options.IncludeAllCommits && options.CommitsYear != nil {
			commitsLabel += fmt.Sprintf(" (%d)", *options.CommitsYear)
		} else if !options.IncludeAllCommits {
			commitsLabel += " (" + i18n.T("wakatimecard.lastyear") + ")"
		}
		statItems = append(statItems, createStatItem(
			common.Icons["commits"],
			commitsLabel,
			float64(stats.TotalCommits),
			"commits",
			"",
			len(statItems),
			options.ShowIcons,
			options.TextBold,
			numberFormat,
			options.NumberPrecision,
			shiftValuePos,
		))
	}

	// PRs
	if !containsString(hide, "prs") {
		statItems = append(statItems, createStatItem(
			common.Icons["prs"],
			i18n.T("statcard.prs"),
			float64(stats.TotalPRs),
			"prs",
			"",
			len(statItems),
			options.ShowIcons,
			options.TextBold,
			numberFormat,
			options.NumberPrecision,
			shiftValuePos,
		))
	}

	// Additional stats from show array
	if containsString(show, "prs_merged") {
		statItems = append(statItems, createStatItem(
			common.Icons["prs_merged"],
			i18n.T("statcard.prs-merged"),
			float64(stats.TotalPRsMerged),
			"prs_merged",
			"",
			len(statItems),
			options.ShowIcons,
			options.TextBold,
			numberFormat,
			options.NumberPrecision,
			shiftValuePos,
		))
	}

	if containsString(show, "prs_merged_percentage") {
		precision := options.NumberPrecision
		if precision < 0 || precision > 2 {
			precision = 2
		}
		statItems = append(statItems, createStatItem(
			common.Icons["prs_merged_percentage"],
			i18n.T("statcard.prs-merged-percentage"),
			stats.MergedPRsPercentage,
			"prs_merged_percentage",
			"%",
			len(statItems),
			options.ShowIcons,
			options.TextBold,
			numberFormat,
			precision,
			shiftValuePos,
		))
	}

	if containsString(show, "reviews") {
		statItems = append(statItems, createStatItem(
			common.Icons["reviews"],
			i18n.T("statcard.reviews"),
			float64(stats.TotalReviews),
			"reviews",
			"",
			len(statItems),
			options.ShowIcons,
			options.TextBold,
			numberFormat,
			options.NumberPrecision,
			shiftValuePos,
		))
	}

	// Issues
	if !containsString(hide, "issues") {
		statItems = append(statItems, createStatItem(
			common.Icons["issues"],
			i18n.T("statcard.issues"),
			float64(stats.TotalIssues),
			"issues",
			"",
			len(statItems),
			options.ShowIcons,
			options.TextBold,
			numberFormat,
			options.NumberPrecision,
			shiftValuePos,
		))
	}

	if containsString(show, "discussions_started") {
		statItems = append(statItems, createStatItem(
			common.Icons["discussions_started"],
			i18n.T("statcard.discussions-started"),
			float64(stats.TotalDiscussionsStarted),
			"discussions_started",
			"",
			len(statItems),
			options.ShowIcons,
			options.TextBold,
			numberFormat,
			options.NumberPrecision,
			shiftValuePos,
		))
	}

	if containsString(show, "discussions_answered") {
		statItems = append(statItems, createStatItem(
			common.Icons["discussions_answered"],
			i18n.T("statcard.discussions-answered"),
			float64(stats.TotalDiscussionsAnswered),
			"discussions_answered",
			"",
			len(statItems),
			options.ShowIcons,
			options.TextBold,
			numberFormat,
			options.NumberPrecision,
			shiftValuePos,
		))
	}

	// Contribs
	if !containsString(hide, "contribs") {
		statItems = append(statItems, createStatItem(
			common.Icons["contribs"],
			i18n.T("statcard.contribs"),
			float64(stats.ContributedTo),
			"contribs",
			"",
			len(statItems),
			options.ShowIcons,
			options.TextBold,
			numberFormat,
			options.NumberPrecision,
			shiftValuePos,
		))
	}

	// Calculate card dimensions
	height := math.Max(
		float64(45+(len(statItems)+1)*lineHeight),
		func() float64 {
			if options.HideRank {
				return 0
			}
			if len(statItems) > 0 {
				return 150
			}
			return 180
		}(),
	)

	progress := 100 - stats.Rank.Percentile

	// Calculate width
	iconWidth := 0.0
	if options.ShowIcons && len(statItems) > 0 {
		iconWidth = 17
	}

	minCardWidth := func() float64 {
		if options.HideRank {
			titleText := options.CustomTitle
			if titleText == "" {
				titleText = fmt.Sprintf("%s's GitHub Stats", stats.Name)
			}
			textWidth := common.MeasureText(titleText, 10)
			return math.Max(cardMinWidth, 50+textWidth*2)
		}
		if len(statItems) > 0 {
			return rankCardMinWidth
		}
		return rankOnlyCardMinWidth
	}() + iconWidth

	defaultCardWidth := func() float64 {
		if options.HideRank {
			return cardDefaultWidth
		}
		if len(statItems) > 0 {
			return rankCardDefaultWidth
		}
		return rankOnlyCardDefaultWidth
	}() + iconWidth

	width := defaultCardWidth
	if options.CardWidth > 0 {
		width = float64(options.CardWidth)
	}
	if width < minCardWidth {
		width = minCardWidth
	}

	// Create card
	cardTitle := fmt.Sprintf("%s's GitHub Stats", stats.Name)
	if options.CustomTitle != "" {
		cardTitle = options.CustomTitle
	} else if len(statItems) == 0 {
		cardTitle = fmt.Sprintf("%s's GitHub Rank", stats.Name)
	}

	card := NewCard(CardOptions{
		Width:             width,
		Height:            height,
		BorderRadius:      options.BorderRadius,
		Colors:            colors,
		CustomTitle:       cardTitle,
		DefaultTitle:      cardTitle,
		HideBorder:        options.HideBorder,
		HideTitle:         options.HideTitle,
		DisableAnimations: options.DisableAnimations,
		CSS:               getStatsStyles(colors, options.ShowIcons, progress),
	})

	// Build body
	rankCircle := ""
	if !options.HideRank {
		rankXTranslation := calculateRankXTranslation(width, len(statItems) > 0, iconWidth, minCardWidth)
		rankYTranslation := height/2 - 50
		rankCircle = fmt.Sprintf(`
      <g data-testid="rank-circle" transform="translate(%.0f, %.0f)">
        <circle class="rank-circle-rim" cx="-10" cy="8" r="40" />
        <circle class="rank-circle" cx="-10" cy="8" r="40" />
        <g class="rank-text">
          %s
        </g>
      </g>`, rankXTranslation, rankYTranslation,
			common.RankIcon(rankIcon, stats.Rank.Level, stats.Rank.Percentile))
	}

	statItemsSVG := common.FlexLayout(common.FlexLayoutProps{
		Items:     statItems,
		Gap:       float64(lineHeight),
		Direction: "column",
	})

	body := fmt.Sprintf(`
    %s
    <svg x="0" y="0">
      %s
    </svg>
  `, rankCircle, strings.Join(statItemsSVG, ""))

	return card.Render(body)
}

// Helper functions

func createStatItem(icon, label string, value float64, id, unitSymbol string, index int, showIcons, bold bool, numberFormat string, numberPrecision int, shiftValuePos float64) string {
	precision := &numberPrecision
	if numberPrecision < 0 || numberPrecision > 2 {
		precision = nil
	}

	kValue := common.KFormatter(value, precision)
	if numberFormat == "long" || id == "prs_merged_percentage" {
		if id == "prs_merged_percentage" {
			precision := numberPrecision
			if precision < 0 || precision > 2 {
				precision = 2
			}
			kValue = fmt.Sprintf("%.*f", precision, value)
		} else {
			kValue = fmt.Sprintf("%.0f", value)
		}
	}

	staggerDelay := (index + 3) * 150
	labelOffset := ""
	if showIcons {
		labelOffset = `x="25"`
	}

	iconSvg := ""
	if showIcons {
		iconSvg = fmt.Sprintf(`
    <svg data-testid="icon" class="icon" viewBox="0 0 16 16" version="1.1" width="16" height="16">
      %s
    </svg>
  `, icon)
	}

	boldClass := "not_bold"
	if bold {
		boldClass = "bold"
	}

	unitStr := ""
	if unitSymbol != "" {
		unitStr = " " + unitSymbol
	}

	valueX := 140.0
	if !showIcons {
		valueX = 120.0
	}
	valueX += shiftValuePos

	return fmt.Sprintf(`
    <g class="stagger" style="animation-delay: %dms" transform="translate(25, 0)">
      %s
      <text class="stat %s" %s y="12.5">%s:</text>
      <text
        class="stat %s"
        x="%.2f"
        y="12.5"
        data-testid="%s"
      >%v%s</text>
    </g>
  `, staggerDelay, iconSvg, boldClass, labelOffset, label, boldClass, valueX, id, kValue, unitStr)
}

func getStatsStyles(colors common.CardColors, showIcons bool, progress float64) string {
	iconDisplay := "none"
	if showIcons {
		iconDisplay = "block"
	}

	progressAnimation := calculateProgressAnimation(progress)

	return fmt.Sprintf(`
    .stat {
      font: 600 14px 'Segoe UI', Ubuntu, "Helvetica Neue", Sans-Serif; fill: %s;
    }
    @supports(-moz-appearance: auto) {
      /* Selector detects Firefox */
      .stat { font-size:12px; }
    }
    .stagger {
      opacity: 0;
      animation: fadeInAnimation 0.3s ease-in-out forwards;
    }
    .rank-text {
      font: 800 24px 'Segoe UI', Ubuntu, Sans-Serif; fill: %s;
      animation: scaleInAnimation 0.3s ease-in-out forwards;
    }
    .rank-percentile-header {
      font-size: 14px;
    }
    .rank-percentile-text {
      font-size: 16px;
    }
    
    .not_bold { font-weight: 400 }
    .bold { font-weight: 700 }
    .icon {
      fill: %s;
      display: %s;
    }

    .rank-circle-rim {
      stroke: %s;
      fill: none;
      stroke-width: 6;
      opacity: 0.2;
    }
    .rank-circle {
      stroke: %s;
      stroke-dasharray: 250;
      fill: none;
      stroke-width: 6;
      stroke-linecap: round;
      opacity: 0.8;
      transform-origin: -10px 8px;
      transform: rotate(-90deg);
      animation: rankAnimation 1s forwards ease-in-out;
      filter: drop-shadow(0 2px 4px rgba(0,0,0,0.1));
    }
    .rank-circle-rim {
      filter: drop-shadow(0 1px 2px rgba(0,0,0,0.1));
    }
    %s
  `, colors.TextColor, colors.TextColor, colors.IconColor, iconDisplay,
		colors.RingColor, colors.RingColor, progressAnimation)
}

func calculateProgressAnimation(progress float64) string {
	radius := 40.0
	c := math.Pi * (radius * 2)

	if progress < 0 {
		progress = 0
	}
	if progress > 100 {
		progress = 100
	}

	progressValue := ((100 - progress) / 100) * c

	return fmt.Sprintf(`
    @keyframes rankAnimation {
      from {
        stroke-dashoffset: %.2f;
      }
      to {
        stroke-dashoffset: %.2f;
      }
    }
  `, c, progressValue)
}

func calculateRankXTranslation(width float64, hasStats bool, iconWidth, minCardWidth float64) float64 {
	if hasStats {
		minXTranslation := rankCardMinWidth + iconWidth - 70
		if width > rankCardDefaultWidth {
			xMaxExpansion := minXTranslation + (450-minCardWidth)/2
			return xMaxExpansion + width - rankCardDefaultWidth
		}
		return minXTranslation + (width-minCardWidth)/2
	}
	return width/2 + 20 - 10
}

func containsString(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
