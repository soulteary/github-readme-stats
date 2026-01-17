package themes

// Theme represents a color theme
type Theme struct {
	TitleColor  string
	IconColor   string
	TextColor   string
	BgColor     string
	BorderColor string
	RingColor   string
}

// Themes contains all available themes
var Themes = map[string]Theme{
	"default": {
		TitleColor:  "2f80ed",
		IconColor:   "4c71f2",
		TextColor:   "434d58",
		BgColor:     "fffefe",
		BorderColor: "e4e2e2",
	},
	"default_repocard": {
		TitleColor: "2f80ed",
		IconColor:  "586069",
		TextColor:  "434d58",
		BgColor:    "fffefe",
	},
	"transparent": {
		TitleColor: "006AFF",
		IconColor:  "0579C3",
		TextColor:  "417E87",
		BgColor:    "ffffff00",
	},
	"shadow_red": {
		TitleColor:  "9A0000",
		TextColor:   "444",
		IconColor:   "4F0000",
		BorderColor: "4F0000",
		BgColor:     "ffffff00",
	},
	"shadow_green": {
		TitleColor:  "007A00",
		TextColor:   "444",
		IconColor:   "003D00",
		BorderColor: "003D00",
		BgColor:     "ffffff00",
	},
	"shadow_blue": {
		TitleColor:  "00779A",
		TextColor:   "444",
		IconColor:   "004450",
		BorderColor: "004490",
		BgColor:     "ffffff00",
	},
	"dark": {
		TitleColor: "fff",
		IconColor:  "79ff97",
		TextColor:  "9f9f9f",
		BgColor:    "151515",
	},
	"radical": {
		TitleColor: "fe428e",
		IconColor:  "f8d847",
		TextColor:  "a9fef7",
		BgColor:    "141321",
	},
	"merko": {
		TitleColor: "abd200",
		IconColor:  "b7d364",
		TextColor:  "68b587",
		BgColor:    "0a0f0b",
	},
	"gruvbox": {
		TitleColor: "fabd2f",
		IconColor:  "fe8019",
		TextColor:  "8ec07c",
		BgColor:    "282828",
	},
	"gruvbox_light": {
		TitleColor: "b57614",
		IconColor:  "af3a03",
		TextColor:  "427b58",
		BgColor:    "fbf1c7",
	},
	"tokyonight": {
		TitleColor: "70a5fd",
		IconColor:  "bf91f3",
		TextColor:  "38bdae",
		BgColor:    "1a1b27",
	},
	"onedark": {
		TitleColor: "e4bf7a",
		IconColor:  "8eb573",
		TextColor:  "df6d74",
		BgColor:    "282c34",
	},
	"cobalt": {
		TitleColor: "e683d9",
		IconColor:  "0480ef",
		TextColor:  "75eeb2",
		BgColor:    "193549",
	},
	"synthwave": {
		TitleColor: "e2e9ec",
		IconColor:  "ef8539",
		TextColor:  "e5289e",
		BgColor:    "2b213a",
	},
	"highcontrast": {
		TitleColor: "e7f216",
		IconColor:  "00ffff",
		TextColor:  "fff",
		BgColor:    "000",
	},
	"dracula": {
		TitleColor: "ff6e96",
		IconColor:  "79dafa",
		TextColor:  "f8f8f2",
		BgColor:    "282a36",
	},
	"prussian": {
		TitleColor: "bddfff",
		IconColor:  "38a0ff",
		TextColor:  "6e93b5",
		BgColor:    "172f45",
	},
	"monokai": {
		TitleColor: "eb1f6a",
		IconColor:  "e28905",
		TextColor:  "f1f1eb",
		BgColor:    "272822",
	},
	"vue": {
		TitleColor: "41b883",
		IconColor:  "41b883",
		TextColor:  "273849",
		BgColor:    "fffefe",
	},
	"vue-dark": {
		TitleColor: "41b883",
		IconColor:  "41b883",
		TextColor:  "fffefe",
		BgColor:    "273849",
	},
	"shades-of-purple": {
		TitleColor: "fad000",
		IconColor:  "b362ff",
		TextColor:  "a599e9",
		BgColor:    "2d2b55",
	},
	"nightowl": {
		TitleColor: "c792ea",
		IconColor:  "ffeb95",
		TextColor:  "7fdbca",
		BgColor:    "011627",
	},
	"buefy": {
		TitleColor: "7957d5",
		IconColor:  "ff3860",
		TextColor:  "363636",
		BgColor:    "ffffff",
	},
	"blue-green": {
		TitleColor: "2f97c1",
		IconColor:  "f5b700",
		TextColor:  "0cf574",
		BgColor:    "040f0f",
	},
	"algolia": {
		TitleColor: "00AEFF",
		IconColor:  "2DDE98",
		TextColor:  "FFFFFF",
		BgColor:    "050F2C",
	},
	"great-gatsby": {
		TitleColor: "ffa726",
		IconColor:  "ffb74d",
		TextColor:  "ffd95b",
		BgColor:    "000000",
	},
	"darcula": {
		TitleColor: "BA5F17",
		IconColor:  "84628F",
		TextColor:  "BEBEBE",
		BgColor:    "242424",
	},
	"bear": {
		TitleColor: "e03c8a",
		IconColor:  "00AEFF",
		TextColor:  "bcb28d",
		BgColor:    "1f2023",
	},
	"solarized-dark": {
		TitleColor: "268bd2",
		IconColor:  "b58900",
		TextColor:  "859900",
		BgColor:    "002b36",
	},
	"solarized-light": {
		TitleColor: "268bd2",
		IconColor:  "b58900",
		TextColor:  "859900",
		BgColor:    "fdf6e3",
	},
	"chartreuse-dark": {
		TitleColor: "7fff00",
		IconColor:  "00AEFF",
		TextColor:  "fff",
		BgColor:    "000",
	},
	"nord": {
		TitleColor: "81a1c1",
		TextColor:  "d8dee9",
		IconColor:  "88c0d0",
		BgColor:    "2e3440",
	},
	"gotham": {
		TitleColor: "2aa889",
		IconColor:  "599cab",
		TextColor:  "99d1ce",
		BgColor:    "0c1014",
	},
	"material-palenight": {
		TitleColor: "c792ea",
		IconColor:  "89ddff",
		TextColor:  "a6accd",
		BgColor:    "292d3e",
	},
	"graywhite": {
		TitleColor: "24292e",
		IconColor:  "24292e",
		TextColor:  "24292e",
		BgColor:    "ffffff",
	},
	"vision-friendly-dark": {
		TitleColor: "ffb000",
		IconColor:  "785ef0",
		TextColor:  "ffffff",
		BgColor:    "000000",
	},
	"ayu-mirage": {
		TitleColor: "f4cd7c",
		IconColor:  "73d0ff",
		TextColor:  "c7c8c2",
		BgColor:    "1f2430",
	},
	"midnight-purple": {
		TitleColor: "9745f5",
		IconColor:  "9f4bff",
		TextColor:  "ffffff",
		BgColor:    "000000",
	},
	"calm": {
		TitleColor: "e07a5f",
		IconColor:  "edae49",
		TextColor:  "ebcfb2",
		BgColor:    "373f51",
	},
	"flag-india": {
		TitleColor: "ff8f1c",
		IconColor:  "250E62",
		TextColor:  "509E2F",
		BgColor:    "ffffff",
	},
	"omni": {
		TitleColor: "FF79C6",
		IconColor:  "e7de79",
		TextColor:  "E1E1E6",
		BgColor:    "191622",
	},
	"react": {
		TitleColor: "61dafb",
		IconColor:  "61dafb",
		TextColor:  "ffffff",
		BgColor:    "20232a",
	},
	"jolly": {
		TitleColor: "ff64da",
		IconColor:  "a960ff",
		TextColor:  "ffffff",
		BgColor:    "291B3E",
	},
	"maroongold": {
		TitleColor: "F7EF8A",
		IconColor:  "F7EF8A",
		TextColor:  "E0AA3E",
		BgColor:    "260000",
	},
	"yeblu": {
		TitleColor: "ffff00",
		IconColor:  "ffff00",
		TextColor:  "ffffff",
		BgColor:    "002046",
	},
	"blueberry": {
		TitleColor: "82aaff",
		IconColor:  "89ddff",
		TextColor:  "27e8a7",
		BgColor:    "242938",
	},
	"slateorange": {
		TitleColor: "faa627",
		IconColor:  "faa627",
		TextColor:  "ffffff",
		BgColor:    "36393f",
	},
	"kacho_ga": {
		TitleColor: "bf4a3f",
		IconColor:  "a64833",
		TextColor:  "d9c8a9",
		BgColor:    "402b23",
	},
	"outrun": {
		TitleColor: "ffcc00",
		IconColor:  "ff1aff",
		TextColor:  "8080ff",
		BgColor:    "141439",
	},
	"ocean_dark": {
		TitleColor: "8957B2",
		IconColor:  "FFFFFF",
		TextColor:  "92D534",
		BgColor:    "151A28",
	},
	"city_lights": {
		TitleColor: "5D8CB3",
		IconColor:  "4798FF",
		TextColor:  "718CA1",
		BgColor:    "1D252C",
	},
	"github_dark": {
		TitleColor: "58A6FF",
		IconColor:  "1F6FEB",
		TextColor:  "C3D1D9",
		BgColor:    "0D1117",
	},
	"github_dark_dimmed": {
		TitleColor:  "539bf5",
		IconColor:   "539bf5",
		TextColor:   "ADBAC7",
		BgColor:     "24292F",
		BorderColor: "373E47",
	},
	"discord_old_blurple": {
		TitleColor: "7289DA",
		IconColor:  "7289DA",
		TextColor:  "FFFFFF",
		BgColor:    "2C2F33",
	},
	"aura_dark": {
		TitleColor: "ff7372",
		IconColor:  "6cffd0",
		TextColor:  "dbdbdb",
		BgColor:    "252334",
	},
	"panda": {
		TitleColor: "19f9d899",
		IconColor:  "19f9d899",
		TextColor:  "FF75B5",
		BgColor:    "31353a",
	},
	"noctis_minimus": {
		TitleColor: "d3b692",
		IconColor:  "72b7c0",
		TextColor:  "c5cdd3",
		BgColor:    "1b2932",
	},
	"cobalt2": {
		TitleColor: "ffc600",
		IconColor:  "ffffff",
		TextColor:  "0088ff",
		BgColor:    "193549",
	},
	"swift": {
		TitleColor: "000000",
		IconColor:  "f05237",
		TextColor:  "000000",
		BgColor:    "f7f7f7",
	},
	"aura": {
		TitleColor: "a277ff",
		IconColor:  "ffca85",
		TextColor:  "61ffca",
		BgColor:    "15141b",
	},
	"apprentice": {
		TitleColor: "ffffff",
		IconColor:  "ffffaf",
		TextColor:  "bcbcbc",
		BgColor:    "262626",
	},
	"moltack": {
		TitleColor: "86092C",
		IconColor:  "86092C",
		TextColor:  "574038",
		BgColor:    "F5E1C0",
	},
	"codeSTACKr": {
		TitleColor:  "ff652f",
		IconColor:   "FFE400",
		TextColor:   "ffffff",
		BgColor:     "09131B",
		BorderColor: "0c1a25",
	},
	"rose_pine": {
		TitleColor: "9ccfd8",
		IconColor:  "ebbcba",
		TextColor:  "e0def4",
		BgColor:    "191724",
	},
	"catppuccin_latte": {
		TitleColor: "137980",
		IconColor:  "8839ef",
		TextColor:  "4c4f69",
		BgColor:    "eff1f5",
	},
	"catppuccin_mocha": {
		TitleColor: "94e2d5",
		IconColor:  "cba6f7",
		TextColor:  "cdd6f4",
		BgColor:    "1e1e2e",
	},
	"date_night": {
		TitleColor:  "DA7885",
		TextColor:   "E1B2A2",
		IconColor:   "BB8470",
		BorderColor: "170F0C",
		BgColor:     "170F0C",
	},
	"one_dark_pro": {
		TitleColor:  "61AFEF",
		TextColor:   "E5C06E",
		IconColor:   "C678DD",
		BorderColor: "3B4048",
		BgColor:     "23272E",
	},
	"rose": {
		TitleColor:  "8d192b",
		TextColor:   "862931",
		IconColor:   "B71F36",
		BorderColor: "e9d8d4",
		BgColor:     "e9d8d4",
	},
	"holi": {
		TitleColor:  "5FABEE",
		TextColor:   "D6E7FF",
		IconColor:   "5FABEE",
		BorderColor: "85A4C0",
		BgColor:     "030314",
	},
	"neon": {
		TitleColor:  "00EAD3",
		TextColor:   "FF449F",
		IconColor:   "00EAD3",
		BorderColor: "ffffff",
		BgColor:     "000000",
	},
	"blue_navy": {
		TitleColor:  "82AAFF",
		TextColor:   "82AAFF",
		IconColor:   "82AAFF",
		BorderColor: "ffffff",
		BgColor:     "000000",
	},
	"calm_pink": {
		TitleColor:  "e07a5f",
		TextColor:   "edae49",
		IconColor:   "ebcfb2",
		BorderColor: "e1bc29",
		BgColor:     "2b2d40",
	},
	"ambient_gradient": {
		TitleColor: "ffffff",
		TextColor:  "ffffff",
		IconColor:  "ffffff",
		BgColor:    "35,4158d0,c850c0,ffcc70",
	},
}

// GetTheme returns a theme by name, or default if not found
func GetTheme(name string) Theme {
	if theme, ok := Themes[name]; ok {
		return theme
	}
	return Themes["default"]
}
