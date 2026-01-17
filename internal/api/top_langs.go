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

// TopLangsHandler handles the /api/top-langs route
func TopLangsHandler(c *gin.Context) {
	// Get query parameters
	username := c.Query("username")
	hide := c.Query("hide")
	hideTitle := c.Query("hide_title")
	hideBorder := c.Query("hide_border")
	cardWidth := c.Query("card_width")
	titleColor := c.Query("title_color")
	textColor := c.Query("text_color")
	bgColor := c.Query("bg_color")
	theme := c.Query("theme")
	layout := c.Query("layout")
	langsCount := c.Query("langs_count")
	excludeRepo := c.Query("exclude_repo")
	sizeWeight := c.Query("size_weight")
	countWeight := c.Query("count_weight")
	customTitle := c.Query("custom_title")
	locale := c.Query("locale")
	borderRadius := c.Query("border_radius")
	borderColor := c.Query("border_color")
	disableAnimations := c.Query("disable_animations")
	hideProgress := c.Query("hide_progress")
	statsFormat := c.Query("stats_format")

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
	sizeWeightFloat := 1.0
	if sizeWeight != "" {
		if val, err := strconv.ParseFloat(sizeWeight, 64); err == nil {
			sizeWeightFloat = val
		}
	}

	countWeightFloat := 0.0
	if countWeight != "" {
		if val, err := strconv.ParseFloat(countWeight, 64); err == nil {
			countWeightFloat = val
		}
	}

	// Fetch top languages
	langs, err := fetchers.FetchTopLanguages(
		username,
		common.ParseArray(excludeRepo),
		sizeWeightFloat,
		countWeightFloat,
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
		0,
		common.TopLangsCardTTL.Default,
		common.TopLangsCardTTL.Min,
		common.TopLangsCardTTL.Max,
	)
	common.SetCacheHeaders(c.Writer, cacheSecondsInt)

	// Parse options
	cardWidthInt := 0
	if cardWidth != "" {
		cardWidthInt, _ = strconv.Atoi(cardWidth)
	}

	langsCountInt := 0
	if langsCount != "" {
		langsCountInt, _ = strconv.Atoi(langsCount)
	}

	borderRadiusFloat := 4.5
	if borderRadius != "" {
		if val, err := strconv.ParseFloat(borderRadius, 64); err == nil {
			borderRadiusFloat = val
		}
	}

	// Render card
	cardSVG := cards.RenderTopLanguages(langs, cards.TopLanguagesCardOptions{
		Hide:              common.ParseArray(hide),
		HideTitle:         common.ParseBoolean(hideTitle) != nil && *common.ParseBoolean(hideTitle),
		HideBorder:        common.ParseBoolean(hideBorder) != nil && *common.ParseBoolean(hideBorder),
		CardWidth:         cardWidthInt,
		TitleColor:        titleColor,
		TextColor:         textColor,
		BgColor:           bgColor,
		Theme:             theme,
		Layout:            layout,
		LangsCount:        langsCountInt,
		ExcludeRepo:       common.ParseArray(excludeRepo),
		SizeWeight:        sizeWeightFloat,
		CountWeight:       countWeightFloat,
		CustomTitle:       customTitle,
		Locale:            strings.ToLower(locale),
		BorderRadius:      borderRadiusFloat,
		BorderColor:       borderColor,
		DisableAnimations: common.ParseBoolean(disableAnimations) != nil && *common.ParseBoolean(disableAnimations),
		HideProgress:      common.ParseBoolean(hideProgress) != nil && *common.ParseBoolean(hideProgress),
		StatsFormat:       statsFormat,
	})

	c.String(http.StatusOK, cardSVG)
}
