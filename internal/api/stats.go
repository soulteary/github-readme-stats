package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/github-readme-stats/internal/cards"
	"github.com/soulteary/github-readme-stats/internal/common"
	"github.com/soulteary/github-readme-stats/internal/fetchers"
	"github.com/soulteary/github-readme-stats/internal/translations"
)

// StatsHandler handles the /api route
func StatsHandler(c *gin.Context) {
	// Get query parameters
	username := c.Query("username")
	hide := c.Query("hide")
	hideTitle := c.Query("hide_title")
	hideBorder := c.Query("hide_border")
	cardWidth := c.Query("card_width")
	hideRank := c.Query("hide_rank")
	showIcons := c.Query("show_icons")
	includeAllCommits := c.Query("include_all_commits")
	commitsYear := c.Query("commits_year")
	lineHeight := c.Query("line_height")
	titleColor := c.Query("title_color")
	ringColor := c.Query("ring_color")
	iconColor := c.Query("icon_color")
	textColor := c.Query("text_color")
	textBold := c.Query("text_bold")
	bgColor := c.Query("bg_color")
	theme := c.Query("theme")
	cacheSeconds := c.Query("cache_seconds")
	excludeRepo := c.Query("exclude_repo")
	customTitle := c.Query("custom_title")
	locale := c.Query("locale")
	disableAnimations := c.Query("disable_animations")
	borderRadius := c.Query("border_radius")
	numberFormat := c.Query("number_format")
	numberPrecision := c.Query("number_precision")
	borderColor := c.Query("border_color")
	rankIcon := c.Query("rank_icon")
	show := c.Query("show")

	// Set content type
	c.Header("Content-Type", "image/svg+xml")

	// Guard access
	colors := map[string]string{
		"title_color":  titleColor,
		"text_color":   textColor,
		"bg_color":     bgColor,
		"border_color": borderColor,
		"theme":        theme,
	}

	accessResult := common.GuardAccess(c.Writer, username, "username", colors)
	if !accessResult.IsPassed {
		c.String(http.StatusForbidden, accessResult.Response)
		return
	}

	// Check locale
	if locale != "" && !translations.IsLocaleAvailable(locale) {
		errorSVG := common.RenderError(common.ErrorOptions{
			Message:          "Something went wrong",
			SecondaryMessage: "Language not found",
			RenderOptions:    colors,
		})
		c.String(http.StatusBadRequest, errorSVG)
		return
	}

	// Parse parameters
	showStats := common.ParseArray(show)
	includeMergedPRs := containsString(showStats, "prs_merged") || containsString(showStats, "prs_merged_percentage")
	includeDiscussions := containsString(showStats, "discussions_started")
	includeDiscussionsAnswers := containsString(showStats, "discussions_answered")

	var commitsYearInt *int
	if commitsYear != "" {
		year, err := strconv.Atoi(commitsYear)
		if err == nil {
			commitsYearInt = &year
		}
	}

	// Fetch stats
	stats, err := fetchers.FetchStats(
		username,
		common.ParseBoolean(includeAllCommits) != nil && *common.ParseBoolean(includeAllCommits),
		common.ParseArray(excludeRepo),
		includeMergedPRs,
		includeDiscussions,
		includeDiscussionsAnswers,
		commitsYearInt,
	)

	if err != nil {
		common.SetErrorCacheHeaders(c.Writer)
		errorSVG := common.RenderError(common.ErrorOptions{
			Message:          err.Error(),
			SecondaryMessage: common.RetrieveSecondaryMessage(err),
			RenderOptions:    colors,
			ShowRepoLink:     true,
		})
		if _, ok := err.(*common.MissingParamError); ok {
			errorSVG = common.RenderError(common.ErrorOptions{
				Message:          err.Error(),
				SecondaryMessage: common.RetrieveSecondaryMessage(err),
				RenderOptions:    colors,
				ShowRepoLink:     false,
			})
		}
		c.String(http.StatusOK, errorSVG)
		return
	}

	// Resolve cache seconds
	cacheSecondsInt := common.ResolveCacheSeconds(
		func() int {
			if cacheSeconds == "" {
				return 0
			}
			val, _ := strconv.Atoi(cacheSeconds)
			return val
		}(),
		common.StatsCardTTL.Default,
		common.StatsCardTTL.Min,
		common.StatsCardTTL.Max,
	)
	common.SetCacheHeaders(c.Writer, cacheSecondsInt)

	// Parse options
	hideArray := common.ParseArray(hide)
	cardWidthInt := 0
	if cardWidth != "" {
		cardWidthInt, _ = strconv.Atoi(cardWidth)
	}
	lineHeightInt := 25
	if lineHeight != "" {
		lineHeightInt, _ = strconv.Atoi(lineHeight)
	}
	numberPrecisionInt := 0
	if numberPrecision != "" {
		numberPrecisionInt, _ = strconv.Atoi(numberPrecision)
	}
	borderRadiusFloat := 4.5
	if borderRadius != "" {
		if val, err := strconv.ParseFloat(borderRadius, 64); err == nil {
			borderRadiusFloat = val
		}
	}

	// Render card
	cardSVG := cards.RenderStatsCard(stats, cards.StatsCardOptions{
		Hide:              hideArray,
		ShowIcons:         common.ParseBoolean(showIcons) != nil && *common.ParseBoolean(showIcons),
		HideTitle:         common.ParseBoolean(hideTitle) != nil && *common.ParseBoolean(hideTitle),
		HideBorder:        common.ParseBoolean(hideBorder) != nil && *common.ParseBoolean(hideBorder),
		CardWidth:         cardWidthInt,
		HideRank:          common.ParseBoolean(hideRank) != nil && *common.ParseBoolean(hideRank),
		IncludeAllCommits: common.ParseBoolean(includeAllCommits) != nil && *common.ParseBoolean(includeAllCommits),
		CommitsYear:       commitsYearInt,
		LineHeight:        lineHeightInt,
		TitleColor:        titleColor,
		RingColor:         ringColor,
		IconColor:         iconColor,
		TextColor:         textColor,
		TextBold:          common.ParseBoolean(textBold) != nil && *common.ParseBoolean(textBold),
		BgColor:           bgColor,
		Theme:             theme,
		CustomTitle:       customTitle,
		BorderRadius:      borderRadiusFloat,
		BorderColor:       borderColor,
		NumberFormat:      numberFormat,
		NumberPrecision:   numberPrecisionInt,
		Locale:            strings.ToLower(locale),
		DisableAnimations: common.ParseBoolean(disableAnimations) != nil && *common.ParseBoolean(disableAnimations),
		RankIcon:          rankIcon,
		Show:              showStats,
	})

	c.String(http.StatusOK, cardSVG)
}

func containsString(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
