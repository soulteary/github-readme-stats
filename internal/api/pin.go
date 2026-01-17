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

// PinHandler handles the /api/pin route
func PinHandler(c *gin.Context) {
	// Get query parameters
	username := c.Query("username")
	repo := c.Query("repo")
	hideBorder := c.Query("hide_border")
	titleColor := c.Query("title_color")
	iconColor := c.Query("icon_color")
	textColor := c.Query("text_color")
	bgColor := c.Query("bg_color")
	theme := c.Query("theme")
	showOwner := c.Query("show_owner")
	cacheSeconds := c.Query("cache_seconds")
	locale := c.Query("locale")
	borderRadius := c.Query("border_radius")
	borderColor := c.Query("border_color")
	descriptionLinesCount := c.Query("description_lines_count")

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

	// Fetch repo data
	repoData, err := fetchers.FetchRepo(username, repo)
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
		common.PinCardTTL.Default,
		common.PinCardTTL.Min,
		common.PinCardTTL.Max,
	)
	common.SetCacheHeaders(c.Writer, cacheSecondsInt)

	// Parse options
	borderRadiusFloat := 4.5
	if borderRadius != "" {
		if val, err := strconv.ParseFloat(borderRadius, 64); err == nil {
			borderRadiusFloat = val
		}
	}

	descriptionLinesCountInt := 0
	if descriptionLinesCount != "" {
		descriptionLinesCountInt, _ = strconv.Atoi(descriptionLinesCount)
	}

	// Render card
	cardSVG := cards.RenderRepoCard(repoData, cards.RepoCardOptions{
		HideBorder:            common.ParseBoolean(hideBorder) != nil && *common.ParseBoolean(hideBorder),
		TitleColor:            titleColor,
		IconColor:             iconColor,
		TextColor:             textColor,
		BgColor:               bgColor,
		ShowOwner:             common.ParseBoolean(showOwner) != nil && *common.ParseBoolean(showOwner),
		Theme:                 theme,
		BorderRadius:          borderRadiusFloat,
		BorderColor:           borderColor,
		Locale:                strings.ToLower(locale),
		DescriptionLinesCount: descriptionLinesCountInt,
	})

	c.String(http.StatusOK, cardSVG)
}
