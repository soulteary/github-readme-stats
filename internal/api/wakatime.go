package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/github-readme-stats/internal/cards"
	"github.com/soulteary/github-readme-stats/internal/common"
	"github.com/soulteary/github-readme-stats/internal/fetchers"
)

// WakaTimeHandler handles the /api/wakatime route
func WakaTimeHandler(c *gin.Context) {
	// Get query parameters
	username := c.Query("username")
	hide := c.Query("hide")
	hideTitle := c.Query("hide_title")
	cardWidth := c.Query("card_width")
	lineHeight := c.Query("line_height")
	hideProgress := c.Query("hide_progress")
	customTitle := c.Query("custom_title")
	layout := c.Query("layout")
	langsCount := c.Query("langs_count")
	apiDomain := c.Query("api_domain")
	displayFormat := c.Query("display_format")
	titleColor := c.Query("title_color")
	textColor := c.Query("text_color")
	bgColor := c.Query("bg_color")
	theme := c.Query("theme")
	borderRadius := c.Query("border_radius")
	borderColor := c.Query("border_color")
	disableAnimations := c.Query("disable_animations")
	locale := c.Query("locale")

	c.Header("Content-Type", "image/svg+xml")

	// Guard access
	colors := map[string]string{
		"title_color":  titleColor,
		"text_color":   textColor,
		"bg_color":     bgColor,
		"border_color": borderColor,
		"theme":        theme,
	}

	accessResult := common.GuardAccess(c.Writer, username, "wakatime", colors)
	if !accessResult.IsPassed {
		c.String(http.StatusForbidden, accessResult.Response)
		return
	}

	// Fetch WakaTime data
	wakatimeData, err := fetchers.FetchWakaTimeStats(username, apiDomain)
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
		common.WakaTimeCardTTL.Default,
		common.WakaTimeCardTTL.Min,
		common.WakaTimeCardTTL.Max,
	)
	common.SetCacheHeaders(c.Writer, cacheSecondsInt)

	// Parse options
	cardWidthInt := 0
	if cardWidth != "" {
		cardWidthInt, _ = strconv.Atoi(cardWidth)
	}

	lineHeightInt := wakaTimeDefaultLineHeight
	if lineHeight != "" {
		lineHeightInt, _ = strconv.Atoi(lineHeight)
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
	cardSVG := cards.RenderWakaTimeCard(wakatimeData, cards.WakaTimeCardOptions{
		Hide:              common.ParseArray(hide),
		HideTitle:         common.ParseBoolean(hideTitle) != nil && *common.ParseBoolean(hideTitle),
		CardWidth:         cardWidthInt,
		LineHeight:        lineHeightInt,
		HideProgress:      common.ParseBoolean(hideProgress) != nil && *common.ParseBoolean(hideProgress),
		CustomTitle:       customTitle,
		Layout:            layout,
		LangsCount:        langsCountInt,
		ApiDomain:         apiDomain,
		DisplayFormat:     displayFormat,
		TitleColor:        titleColor,
		TextColor:         textColor,
		BgColor:           bgColor,
		Theme:             theme,
		BorderRadius:      borderRadiusFloat,
		BorderColor:       borderColor,
		DisableAnimations: common.ParseBoolean(disableAnimations) != nil && *common.ParseBoolean(disableAnimations),
		Locale:            strings.ToLower(locale),
	})

	c.String(http.StatusOK, cardSVG)
}

const wakaTimeDefaultLineHeight = 25
