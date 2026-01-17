package common

import (
	"fmt"
	"strings"
)

const ErrorCardLength = 576.5

var upstreamAPIErrors = []string{
	TryAgainLater,
	SecondaryErrorMessages[ErrorTypeMaxRetry],
}

// FlexLayoutProps represents properties for flex layout
type FlexLayoutProps struct {
	Items     []string
	Gap       float64
	Direction string // "column" or "row"
	Sizes     []float64
}

// FlexLayout auto layout utility, allows us to layout things vertically or horizontally
func FlexLayout(props FlexLayoutProps) []string {
	var result []string
	lastSize := 0.0

	for i, item := range props.Items {
		if item == "" {
			continue
		}

		size := 0.0
		if i < len(props.Sizes) {
			size = props.Sizes[i]
		}

		transform := fmt.Sprintf("translate(%.0f, 0)", lastSize)
		if props.Direction == "column" {
			transform = fmt.Sprintf("translate(0, %.0f)", lastSize)
		}

		result = append(result, fmt.Sprintf(`<g transform="%s">%s</g>`, transform, item))
		lastSize += size + props.Gap
	}

	return result
}

// CreateLanguageNode creates a node to display the primary programming language
func CreateLanguageNode(langName, langColor string) string {
	return fmt.Sprintf(`
    <g data-testid="primary-lang">
      <circle data-testid="lang-color" cx="0" cy="-5" r="6" fill="%s" />
      <text data-testid="lang-name" class="gray" x="15">%s</text>
    </g>
    `, langColor, langName)
}

// CreateProgressNodeParams represents parameters for creating a progress node
type CreateProgressNodeParams struct {
	X                          float64
	Y                          float64
	Width                      float64
	Color                      string
	Progress                   float64
	ProgressBarBackgroundColor string
	Delay                      int
}

// CreateProgressNode creates a node to indicate progress in percentage
func CreateProgressNode(params CreateProgressNodeParams) string {
	progressPercentage := ClampValue(params.Progress, 2, 100)

	return fmt.Sprintf(`
    <svg width="%.0f" x="%.0f" y="%.0f">
      <rect rx="5" ry="5" x="0" y="0" width="%.0f" height="8" fill="%s"></rect>
      <svg data-testid="lang-progress" width="%.2f%%">
        <rect
            height="8"
            fill="%s"
            rx="5" ry="5" x="0" y="0"
            class="lang-progress"
            style="animation-delay: %dms;"
        />
      </svg>
    </svg>
  `, params.Width, params.X, params.Y, params.Width, params.ProgressBarBackgroundColor,
		progressPercentage, params.Color, params.Delay)
}

// IconWithLabel creates an icon with label to display repository/gist stats
func IconWithLabel(icon string, label interface{}, testid string, iconSize int) string {
	var labelValue interface{} = label
	if num, ok := label.(float64); ok && num <= 0 {
		return ""
	}
	if num, ok := label.(int); ok && num <= 0 {
		return ""
	}

	iconSvg := fmt.Sprintf(`
      <svg
        class="icon"
        y="-12"
        viewBox="0 0 16 16"
        version="1.1"
        width="%d"
        height="%d"
      >
        %s
      </svg>
    `, iconSize, iconSize, icon)

	labelStr := fmt.Sprintf("%v", labelValue)
	text := fmt.Sprintf(`<text data-testid="%s" class="gray">%s</text>`, testid, labelStr)

	items := []string{iconSvg, text}
	layout := FlexLayout(FlexLayoutProps{
		Items: items,
		Gap:   20,
	})

	return strings.Join(layout, "")
}

// ErrorOptions represents options for rendering error
type ErrorOptions struct {
	Message          string
	SecondaryMessage string
	RenderOptions    map[string]string
	ShowRepoLink     bool
}

// RenderError renders error message on the card
func RenderError(options ErrorOptions) string {
	renderOpts := options.RenderOptions
	if renderOpts == nil {
		renderOpts = make(map[string]string)
	}

	theme := renderOpts["theme"]
	if theme == "" {
		theme = "default"
	}

	showRepoLink := options.ShowRepoLink
	if !showRepoLink {
		showRepoLink = true
	}

	// Get card colors
	colorArgs := map[string]string{
		"title_color":  renderOpts["title_color"],
		"text_color":   renderOpts["text_color"],
		"bg_color":     renderOpts["bg_color"],
		"border_color": renderOpts["border_color"],
		"theme":        theme,
	}
	colors := GetCardColors(colorArgs)

	// Check if secondary message is in upstream API errors
	isUpstreamError := false
	for _, err := range upstreamAPIErrors {
		if err == options.SecondaryMessage {
			isUpstreamError = true
			break
		}
	}

	repoLinkText := ""
	if !isUpstreamError && showRepoLink {
		repoLinkText = " file an issue at https://tiny.one/readme-stats"
	}

	messageHTML := EncodeHTML(options.Message)
	secondaryHTML := EncodeHTML(options.SecondaryMessage)

	return fmt.Sprintf(`
    <svg width="%.1f"  height="120" viewBox="0 0 %.1f 120" fill="%s" xmlns="http://www.w3.org/2000/svg">
    <style>
    .text { font: 600 16px 'Segoe UI', Ubuntu, Sans-Serif; fill: %s }
    .small { font: 600 12px 'Segoe UI', Ubuntu, Sans-Serif; fill: %s }
    .gray { fill: #858585 }
    </style>
    <rect x="0.5" y="0.5" width="%.1f" height="99%%" rx="4.5" fill="%s" stroke="%s"/>
    <text x="25" y="45" class="text">Something went wrong!%s</text>
    <text data-testid="message" x="25" y="55" class="text small">
      <tspan x="25" dy="18">%s</tspan>
      <tspan x="25" dy="18" class="gray">%s</tspan>
    </text>
    </svg>
  `, ErrorCardLength, ErrorCardLength, colors.BgColor, colors.TitleColor, colors.TextColor,
		ErrorCardLength-1, colors.BgColor, colors.BorderColor, repoLinkText,
		messageHTML, secondaryHTML)
}

// MeasureText retrieves text length
// See https://stackoverflow.com/a/48172630/10629172
func MeasureText(str string, fontSize float64) float64 {
	if fontSize == 0 {
		fontSize = 10
	}

	widths := []float64{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0.2796875, 0.2765625,
		0.3546875, 0.5546875, 0.5546875, 0.8890625, 0.665625, 0.190625,
		0.3328125, 0.3328125, 0.3890625, 0.5828125, 0.2765625, 0.3328125,
		0.2765625, 0.3015625, 0.5546875, 0.5546875, 0.5546875, 0.5546875,
		0.5546875, 0.5546875, 0.5546875, 0.5546875, 0.5546875, 0.5546875,
		0.2765625, 0.2765625, 0.584375, 0.5828125, 0.584375, 0.5546875,
		1.0140625, 0.665625, 0.665625, 0.721875, 0.721875, 0.665625,
		0.609375, 0.7765625, 0.721875, 0.2765625, 0.5, 0.665625,
		0.5546875, 0.8328125, 0.721875, 0.7765625, 0.665625, 0.7765625,
		0.721875, 0.665625, 0.609375, 0.721875, 0.665625, 0.94375,
		0.665625, 0.665625, 0.609375, 0.2765625, 0.3546875, 0.2765625,
		0.4765625, 0.5546875, 0.3328125, 0.5546875, 0.5546875, 0.5,
		0.5546875, 0.5546875, 0.2765625, 0.5546875, 0.5546875, 0.221875,
		0.240625, 0.5, 0.221875, 0.8328125, 0.5546875, 0.5546875,
		0.5546875, 0.5546875, 0.3328125, 0.5, 0.2765625, 0.5546875,
		0.5, 0.721875, 0.5, 0.5, 0.5, 0.3546875, 0.259375, 0.353125, 0.5890625,
	}

	avg := 0.5279276315789471
	total := 0.0

	for _, char := range str {
		code := int(char)
		if code < len(widths) {
			total += widths[code]
		} else {
			total += avg
		}
	}

	return total * fontSize
}
