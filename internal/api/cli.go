package api

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/soulteary/github-readme-stats/internal/cards"
	"github.com/soulteary/github-readme-stats/internal/common"
	"github.com/soulteary/github-readme-stats/internal/fetchers"
	"github.com/soulteary/github-readme-stats/internal/translations"
)

// GuardAccessCLI guards access using whitelist/blacklist (CLI version)
func GuardAccessCLI(id, accessType string, colors map[string]string) (bool, string) {
	if accessType != "username" && accessType != "gist" && accessType != "wakatime" {
		return false, common.RenderError(common.ErrorOptions{
			Message: "Invalid access type",
		})
	}

	isGist := accessType == "gist"
	currentWhitelist := common.GetGistWhitelist()
	notWhitelistedMsg := "This username is not whitelisted"

	if isGist {
		currentWhitelist = common.GetGistWhitelist()
		notWhitelistedMsg = "This gist ID is not whitelisted"
	} else {
		currentWhitelist = common.GetWhitelist()
	}

	// Check whitelist
	if len(currentWhitelist) > 0 && !common.IsWhitelisted(id, isGist) {
		response := common.RenderError(common.ErrorOptions{
			Message:          notWhitelistedMsg,
			SecondaryMessage: "Please deploy your own instance",
			RenderOptions:    colors,
			ShowRepoLink:     false,
		})
		return false, response
	}

	// Check blacklist (only for usernames)
	if accessType == "username" && len(common.GetWhitelist()) == 0 && common.IsBlacklisted(id) {
		response := common.RenderError(common.ErrorOptions{
			Message:          "This username is blacklisted",
			SecondaryMessage: "Please deploy your own instance",
			RenderOptions:    colors,
			ShowRepoLink:     false,
		})
		return false, response
	}

	return true, ""
}

// writeOutput writes SVG content to file or stdout
func writeOutput(content, outputPath string) error {
	if outputPath == "" {
		_, err := os.Stdout.WriteString(content)
		return err
	}
	return os.WriteFile(outputPath, []byte(content), 0644)
}

// getParam gets a parameter from map with default value
func getParam(params map[string]string, key, defaultValue string) string {
	if val, ok := params[key]; ok && val != "" {
		return val
	}
	return defaultValue
}

// GenerateStatsCard generates stats card from parameters
func GenerateStatsCard(params map[string]string, outputPath string) error {
	username := getParam(params, "username", "")
	hide := getParam(params, "hide", "")
	hideTitle := getParam(params, "hide_title", "")
	hideBorder := getParam(params, "hide_border", "")
	cardWidth := getParam(params, "card_width", "")
	hideRank := getParam(params, "hide_rank", "")
	showIcons := getParam(params, "show_icons", "")
	includeAllCommits := getParam(params, "include_all_commits", "")
	commitsYear := getParam(params, "commits_year", "")
	lineHeight := getParam(params, "line_height", "")
	titleColor := getParam(params, "title_color", "")
	ringColor := getParam(params, "ring_color", "")
	iconColor := getParam(params, "icon_color", "")
	textColor := getParam(params, "text_color", "")
	textBold := getParam(params, "text_bold", "")
	bgColor := getParam(params, "bg_color", "")
	theme := getParam(params, "theme", "")
	excludeRepo := getParam(params, "exclude_repo", "")
	customTitle := getParam(params, "custom_title", "")
	locale := getParam(params, "locale", "")
	disableAnimations := getParam(params, "disable_animations", "")
	borderRadius := getParam(params, "border_radius", "")
	numberFormat := getParam(params, "number_format", "")
	numberPrecision := getParam(params, "number_precision", "")
	borderColor := getParam(params, "border_color", "")
	rankIcon := getParam(params, "rank_icon", "")
	show := getParam(params, "show", "")

	// Guard access
	colors := map[string]string{
		"title_color":  titleColor,
		"text_color":   textColor,
		"bg_color":     bgColor,
		"border_color": borderColor,
		"theme":        theme,
	}

	passed, errorSVG := GuardAccessCLI(username, "username", colors)
	if !passed {
		return writeOutput(errorSVG, outputPath)
	}

	// Check locale
	if locale != "" && !translations.IsLocaleAvailable(locale) {
		errorSVG := common.RenderError(common.ErrorOptions{
			Message:          "Something went wrong",
			SecondaryMessage: "Language not found",
			RenderOptions:    colors,
		})
		return writeOutput(errorSVG, outputPath)
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
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return writeOutput(errorSVG, outputPath)
	}

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

	return writeOutput(cardSVG, outputPath)
}

// GeneratePinCard generates pin card from parameters
func GeneratePinCard(params map[string]string, outputPath string) error {
	username := getParam(params, "username", "")
	repo := getParam(params, "repo", "")
	hideBorder := getParam(params, "hide_border", "")
	titleColor := getParam(params, "title_color", "")
	iconColor := getParam(params, "icon_color", "")
	textColor := getParam(params, "text_color", "")
	bgColor := getParam(params, "bg_color", "")
	theme := getParam(params, "theme", "")
	showOwner := getParam(params, "show_owner", "")
	locale := getParam(params, "locale", "")
	borderRadius := getParam(params, "border_radius", "")
	borderColor := getParam(params, "border_color", "")
	descriptionLinesCount := getParam(params, "description_lines_count", "")

	// Guard access
	colors := map[string]string{
		"title_color":  titleColor,
		"text_color":   textColor,
		"bg_color":     bgColor,
		"border_color": borderColor,
		"theme":        theme,
	}

	passed, errorSVG := GuardAccessCLI(username, "username", colors)
	if !passed {
		return writeOutput(errorSVG, outputPath)
	}

	// Check locale
	if locale != "" && !translations.IsLocaleAvailable(locale) {
		errorSVG := common.RenderError(common.ErrorOptions{
			Message:          "Something went wrong",
			SecondaryMessage: "Language not found",
			RenderOptions:    colors,
		})
		return writeOutput(errorSVG, outputPath)
	}

	// Fetch repo data
	repoData, err := fetchers.FetchRepo(username, repo)
	if err != nil {
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
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return writeOutput(errorSVG, outputPath)
	}

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

	return writeOutput(cardSVG, outputPath)
}

// GenerateTopLangsCard generates top languages card from parameters
func GenerateTopLangsCard(params map[string]string, outputPath string) error {
	username := getParam(params, "username", "")
	hide := getParam(params, "hide", "")
	hideTitle := getParam(params, "hide_title", "")
	hideBorder := getParam(params, "hide_border", "")
	cardWidth := getParam(params, "card_width", "")
	titleColor := getParam(params, "title_color", "")
	textColor := getParam(params, "text_color", "")
	bgColor := getParam(params, "bg_color", "")
	theme := getParam(params, "theme", "")
	layout := getParam(params, "layout", "")
	langsCount := getParam(params, "langs_count", "")
	excludeRepo := getParam(params, "exclude_repo", "")
	sizeWeight := getParam(params, "size_weight", "")
	countWeight := getParam(params, "count_weight", "")
	customTitle := getParam(params, "custom_title", "")
	locale := getParam(params, "locale", "")
	borderRadius := getParam(params, "border_radius", "")
	borderColor := getParam(params, "border_color", "")
	disableAnimations := getParam(params, "disable_animations", "")
	hideProgress := getParam(params, "hide_progress", "")
	statsFormat := getParam(params, "stats_format", "")

	// Guard access
	colors := map[string]string{
		"title_color":  titleColor,
		"text_color":   textColor,
		"bg_color":     bgColor,
		"border_color": borderColor,
		"theme":        theme,
	}

	passed, errorSVG := GuardAccessCLI(username, "username", colors)
	if !passed {
		return writeOutput(errorSVG, outputPath)
	}

	// Check locale
	if locale != "" && !translations.IsLocaleAvailable(locale) {
		errorSVG := common.RenderError(common.ErrorOptions{
			Message:          "Something went wrong",
			SecondaryMessage: "Language not found",
			RenderOptions:    colors,
		})
		return writeOutput(errorSVG, outputPath)
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
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return writeOutput(errorSVG, outputPath)
	}

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

	return writeOutput(cardSVG, outputPath)
}

// GenerateGistCard generates gist card from parameters
func GenerateGistCard(params map[string]string, outputPath string) error {
	id := getParam(params, "id", "")
	titleColor := getParam(params, "title_color", "")
	iconColor := getParam(params, "icon_color", "")
	textColor := getParam(params, "text_color", "")
	bgColor := getParam(params, "bg_color", "")
	theme := getParam(params, "theme", "")
	showOwner := getParam(params, "show_owner", "")
	borderRadius := getParam(params, "border_radius", "")
	borderColor := getParam(params, "border_color", "")
	hideBorder := getParam(params, "hide_border", "")

	// Guard access
	colors := map[string]string{
		"title_color":  titleColor,
		"text_color":   textColor,
		"bg_color":     bgColor,
		"border_color": borderColor,
		"theme":        theme,
	}

	passed, errorSVG := GuardAccessCLI(id, "gist", colors)
	if !passed {
		return writeOutput(errorSVG, outputPath)
	}

	// Fetch gist data
	gistData, err := fetchers.FetchGist(id)
	if err != nil {
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
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return writeOutput(errorSVG, outputPath)
	}

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

	return writeOutput(cardSVG, outputPath)
}

// GenerateWakaTimeCard generates wakatime card from parameters
func GenerateWakaTimeCard(params map[string]string, outputPath string) error {
	username := getParam(params, "username", "")
	hide := getParam(params, "hide", "")
	hideTitle := getParam(params, "hide_title", "")
	cardWidth := getParam(params, "card_width", "")
	lineHeight := getParam(params, "line_height", "")
	hideProgress := getParam(params, "hide_progress", "")
	customTitle := getParam(params, "custom_title", "")
	layout := getParam(params, "layout", "")
	langsCount := getParam(params, "langs_count", "")
	apiDomain := getParam(params, "api_domain", "")
	displayFormat := getParam(params, "display_format", "")
	titleColor := getParam(params, "title_color", "")
	textColor := getParam(params, "text_color", "")
	bgColor := getParam(params, "bg_color", "")
	theme := getParam(params, "theme", "")
	borderRadius := getParam(params, "border_radius", "")
	borderColor := getParam(params, "border_color", "")
	disableAnimations := getParam(params, "disable_animations", "")
	locale := getParam(params, "locale", "")
	testDataPath := getParam(params, "test_data_path", "")

	// Guard access (skip if using test data)
	colors := map[string]string{
		"title_color":  titleColor,
		"text_color":   textColor,
		"bg_color":     bgColor,
		"border_color": borderColor,
		"theme":        theme,
	}

	// Fetch WakaTime data (from test file or API)
	var wakatimeData *fetchers.WakaTimeData
	var err error

	if testDataPath != "" {
		// Use test data file
		wakatimeData, err = fetchers.FetchWakaTimeStatsFromFile(testDataPath)
	} else {
		// Use real API
		passed, errorSVG := GuardAccessCLI(username, "wakatime", colors)
		if !passed {
			return writeOutput(errorSVG, outputPath)
		}
		wakatimeData, err = fetchers.FetchWakaTimeStats(username, apiDomain)
	}

	if err != nil {
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
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return writeOutput(errorSVG, outputPath)
	}

	// Parse options
	cardWidthInt := 0
	if cardWidth != "" {
		cardWidthInt, _ = strconv.Atoi(cardWidth)
	}

	lineHeightInt := 25
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

	return writeOutput(cardSVG, outputPath)
}
