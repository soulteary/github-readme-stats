package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/github-readme-stats/internal/cards"
	"github.com/soulteary/github-readme-stats/internal/common"
	"github.com/soulteary/github-readme-stats/internal/fetchers"
)

// GistHandler handles the /api/gist route
func GistHandler(c *gin.Context) {
	// Get query parameters
	id := c.Query("id")
	titleColor := c.Query("title_color")
	iconColor := c.Query("icon_color")
	textColor := c.Query("text_color")
	bgColor := c.Query("bg_color")
	theme := c.Query("theme")
	showOwner := c.Query("show_owner")
	cacheSeconds := c.Query("cache_seconds")
	borderRadius := c.Query("border_radius")
	borderColor := c.Query("border_color")
	hideBorder := c.Query("hide_border")

	c.Header("Content-Type", "image/svg+xml")

	// Guard access
	colors := map[string]string{
		"title_color":  titleColor,
		"text_color":   textColor,
		"bg_color":     bgColor,
		"border_color": borderColor,
		"theme":        theme,
	}

	accessResult := common.GuardAccess(c.Writer, id, "gist", colors)
	if !accessResult.IsPassed {
		c.String(http.StatusForbidden, accessResult.Response)
		return
	}

	// Fetch gist data
	gistData, err := fetchers.FetchGist(id)
	if err != nil {
		common.SetErrorCacheHeaders(c.Writer)
		showRepoLink := true
		if _, ok := err.(*common.MissingParamError); ok {
			showRepoLink = false
		}
		errorSVG := common.RenderError(common.ErrorOptions{
			Message:          err.Error(),
			SecondaryMessage: common.RetrieveSecondaryMessage(err),
			RenderOptions:    colors,
			ShowRepoLink:     showRepoLink,
		})
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
		common.GistCardTTL.Default,
		common.GistCardTTL.Min,
		common.GistCardTTL.Max,
	)
	common.SetCacheHeaders(c.Writer, cacheSecondsInt)

	// Parse options
	borderRadiusFloat := 4.5
	if borderRadius != "" {
		if val, err := strconv.ParseFloat(borderRadius, 64); err == nil {
			borderRadiusFloat = val
		}
	}

	// Render card
	cardSVG := cards.RenderGistCard(gistData, cards.GistCardOptions{
		TitleColor:   titleColor,
		IconColor:    iconColor,
		TextColor:    textColor,
		BgColor:      bgColor,
		Theme:        theme,
		BorderRadius: borderRadiusFloat,
		BorderColor:  borderColor,
		ShowOwner:    common.ParseBoolean(showOwner) != nil && *common.ParseBoolean(showOwner),
		HideBorder:   common.ParseBoolean(hideBorder) != nil && *common.ParseBoolean(hideBorder),
	})

	c.String(http.StatusOK, cardSVG)
}
