package cards

import (
	"fmt"
	"os"
	"strings"

	"github.com/soulteary/github-readme-stats/internal/common"
)

// Card represents a base card structure
type Card struct {
	Width           float64
	Height          float64
	BorderRadius    float64
	Colors          common.CardColors
	Title           string
	TitlePrefixIcon string
	HideBorder      bool
	HideTitle       bool
	PaddingX        float64
	PaddingY        float64
	Animations      bool
	A11yTitle       string
	A11yDesc        string
	CSS             string
}

// NewCard creates a new card instance
func NewCard(options CardOptions) *Card {
	width := 100.0
	if options.Width > 0 {
		width = options.Width
	}

	height := 100.0
	if options.Height > 0 {
		height = options.Height
	}

	borderRadius := 4.5
	if options.BorderRadius > 0 {
		borderRadius = options.BorderRadius
	}

	title := options.DefaultTitle
	if options.CustomTitle != "" {
		title = options.CustomTitle
	}

	return &Card{
		Width:           width,
		Height:          height,
		BorderRadius:    borderRadius,
		Colors:          options.Colors,
		Title:           common.EncodeHTML(title),
		TitlePrefixIcon: options.TitlePrefixIcon,
		HideBorder:      options.HideBorder,
		HideTitle:       options.HideTitle,
		PaddingX:        25,
		PaddingY:        35,
		Animations:      !options.DisableAnimations,
		CSS:             options.CSS,
	}
}

// CardOptions represents card options
type CardOptions struct {
	Width             float64
	Height            float64
	BorderRadius      float64
	Colors            common.CardColors
	CustomTitle       string
	DefaultTitle      string
	TitlePrefixIcon   string
	HideBorder        bool
	HideTitle         bool
	DisableAnimations bool
	CSS               string
}

// RenderTitle renders the card title
func (c *Card) RenderTitle() string {
	titleText := fmt.Sprintf(`
      <text
        x="0"
        y="0"
        class="header"
        data-testid="header"
      >%s</text>
    `, c.Title)

	var items []string
	if c.TitlePrefixIcon != "" {
		prefixIcon := fmt.Sprintf(`
      <svg
        class="icon"
        x="0"
        y="-13"
        viewBox="0 0 16 16"
        version="1.1"
        width="16"
        height="16"
      >
        %s
      </svg>
    `, c.TitlePrefixIcon)
		items = []string{prefixIcon, titleText}
	} else {
		items = []string{titleText}
	}
	layout := common.FlexLayout(common.FlexLayoutProps{
		Items: items,
		Gap:   25,
	})

	return fmt.Sprintf(`
      <g
        data-testid="card-title"
        transform="translate(%.0f, %.0f)"
      >
        %s
      </g>
    `, c.PaddingX, c.PaddingY, strings.Join(layout, ""))
}

// RenderGradient renders the card gradient
func (c *Card) RenderGradient() string {
	bgColor := c.Colors.BgColor
	gradient, ok := bgColor.([]string)
	if !ok || len(gradient) < 2 {
		return ""
	}

	gradients := gradient[1:]
	angle := gradient[0]

	var stops []string
	for i, grad := range gradients {
		offset := float64(i*100) / float64(len(gradients)-1)
		stops = append(stops, fmt.Sprintf(`<stop offset="%.0f%%" stop-color="#%s" />`, offset, grad))
	}

	return fmt.Sprintf(`
        <defs>
          <linearGradient
            id="gradient"
            gradientTransform="rotate(%s)"
            gradientUnits="userSpaceOnUse"
          >
            %s
          </linearGradient>
        </defs>
        `, angle, strings.Join(stops, "\n            "))
}

// GetAnimations returns CSS animations
func (c *Card) GetAnimations() string {
	if os.Getenv("NODE_ENV") == "test" {
		return ""
	}
	return `
      /* Animations */
      @keyframes scaleInAnimation {
        from {
          transform: translate(-5px, 5px) scale(0);
        }
        to {
          transform: translate(-5px, 5px) scale(1);
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
      @keyframes slideInAnimation {
        from {
          transform: translateX(-10px);
          opacity: 0;
        }
        to {
          transform: translateX(0);
          opacity: 1;
        }
      }
      @keyframes bounceInAnimation {
        0% {
          transform: scale(0.3);
          opacity: 0;
        }
        50% {
          transform: scale(1.05);
        }
        70% {
          transform: scale(0.9);
        }
        100% {
          transform: scale(1);
          opacity: 1;
        }
      }
    `
}

// Render renders the card
func (c *Card) Render(body string) string {
	bgColorStr := ""
	if _, ok := c.Colors.BgColor.([]string); ok {
		bgColorStr = "url(#gradient)"
	} else if str, ok := c.Colors.BgColor.(string); ok {
		bgColorStr = str
	} else {
		bgColorStr = "#fffefe"
	}

	strokeOpacity := 1.0
	if c.HideBorder {
		strokeOpacity = 0
	}

	animationsCSS := ""
	if !c.Animations {
		animationsCSS = `* { animation-duration: 0s !important; animation-delay: 0s !important; }`
	}

	titleSection := ""
	if !c.HideTitle {
		titleSection = c.RenderTitle()
	}

	bodyTransformY := c.PaddingX
	if !c.HideTitle {
		bodyTransformY = c.PaddingY + 20
	}

	return fmt.Sprintf(`
      <svg
        width="%.0f"
        height="%.0f"
        viewBox="0 0 %.0f %.0f"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        role="img"
        aria-labelledby="descId"
      >
        <title id="titleId">%s</title>
        <desc id="descId">%s</desc>
        <style>
          .header {
            font: 600 18px 'Segoe UI', Ubuntu, "Helvetica Neue", Sans-Serif;
            fill: %s;
            animation: fadeInAnimation 0.8s ease-in-out forwards;
          }
          @supports(-moz-appearance: auto) {
            /* Selector detects Firefox */
            .header { font-size: 15.5px; }
          }
          %s

          %s
          %s
        </style>

        %s

        <rect
          data-testid="card-bg"
          x="0.5"
          y="0.5"
          rx="%.1f"
          height="99%%"
          stroke="%s"
          width="%.0f"
          fill="%s"
          stroke-opacity="%.1f"
        />

        %s

        <g
          data-testid="main-card-body"
          transform="translate(0, %.0f)"
        >
          %s
        </g>
      </svg>
    `, c.Width, c.Height, c.Width, c.Height,
		c.A11yTitle, c.A11yDesc,
		c.Colors.TitleColor,
		c.CSS,
		c.GetAnimations(),
		animationsCSS,
		c.RenderGradient(),
		c.BorderRadius,
		c.Colors.BorderColor,
		c.Width-1,
		bgColorStr,
		strokeOpacity,
		titleSection,
		bodyTransformY,
		body)
}
