# GitHub Readme Stats (Go Implementation)

[![GitHub](https://img.shields.io/badge/GitHub-soulteary%2Fgithub--readme--stats-blue)](https://github.com/soulteary/github-readme-stats)

![GitHub Readme Stats](.github/assets/banner.jpg)

## Languages / 语言 / Sprachen / Lingue / 언어 / 言語

- [English](README.md)
- [简体中文](README.zh.md)
- [Deutsch](README.de.md)
- [Italiano](README.it.md)
- [한국어](README.kr.md)
- [日本語](README.ja.md)

This is a Go implementation of the [GitHub Readme Stats (anuraghazra)](https://github.com/anuraghazra/github-readme-stats) project. It provides dynamic GitHub statistics cards that can be embedded in README files to showcase your GitHub activity, repository information, programming language usage, and more. It also supports direct use in GitHub Actions.

## Features

- ✅ GitHub Stats card generation
- ✅ Repository Pin card generation
- ✅ Top Languages card generation
- ✅ Gist card generation
- ✅ WakaTime card generation

## Examples / 示例展示

Here are some examples of what you can create with this project:

### GitHub Stats Card

**Basic:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname)
```
![GitHub Stats](.github/assets/stats-basic.svg)

**Dark Theme:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname&theme=dark)
```
![GitHub Stats Dark](.github/assets/stats-dark.svg)

**Compact Layout:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname&layout=compact)
```
![GitHub Stats Compact](.github/assets/stats-compact.svg)

**Custom Theme:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname&bg_color=0d1117&title_color=ff6b6b&text_color=c9d1d9&border_color=30363d)
```
![GitHub Stats Custom](.github/assets/stats-custom.svg)

### Repository Pin Card

**Basic:**
```markdown
![Pinned Repo](http://localhost:9000/api/pin?username=yourname&repo=repo-name)
```
![Pinned Repo](.github/assets/repo-basic.svg)

**Themed:**
```markdown
![Pinned Repo](http://localhost:9000/api/pin?username=yourname&repo=repo-name&theme=dark)
```
![Pinned Repo Themed](.github/assets/repo-themed.svg)

### Top Languages Card

**Basic:**
```markdown
![Top Languages](http://localhost:9000/api/top-langs?username=yourname)
```
![Top Languages](.github/assets/top-langs-basic.svg)

**Compact Layout:**
```markdown
![Top Languages](http://localhost:9000/api/top-langs?username=yourname&layout=compact&langs_count=6)
```
![Top Languages Compact](.github/assets/top-langs-compact.svg)

### Gist Card

```markdown
![Gist](http://localhost:9000/api/gist?id=gist_id)
```
![Gist](.github/assets/gist-basic.svg)

### WakaTime Card

**Basic:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname)
```
![WakaTime Basic](.github/assets/wakatime-basic.svg)

**Compact Layout:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&layout=compact)
```
![WakaTime Compact](.github/assets/wakatime-compact.svg)

**Test Examples (Generated with test data):**

**Basic Test:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname)
```
![WakaTime Test Basic](.github/assets/wakatime-test-basic.svg)

**Compact Test:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&layout=compact&langs_count=6&theme=dark)
```
![WakaTime Test Compact](.github/assets/wakatime-test-compact.svg)

**Themed:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&theme=radical&langs_count=5)
```
![WakaTime Test Themed](.github/assets/wakatime-test-themed.svg)

**Hide Progress:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&hide_progress=true)
```
![WakaTime Test Hide Progress](.github/assets/wakatime-test-hide-progress.svg)

**Percent Display:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&display_format=percent)
```
![WakaTime Test Percent](.github/assets/wakatime-test-percent.svg)

**Limited Languages:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&langs_count=3&theme=tokyonight)
```
![WakaTime Test Limited](.github/assets/wakatime-test-limited.svg)

> **Note:** WakaTime card examples require a valid WakaTime username with public stats. The test images above are generated using test data to demonstrate various configuration options.

## Quick Start

### 1. Install

```bash
# Install using Go
go get github.com/soulteary/github-readme-stats

# Or build from source
git clone https://github.com/soulteary/github-readme-stats.git
cd github-readme-stats
go build -o github-readme-stats ./cmd/server
```

### 2. Configure

Create a `.env` file:

```bash
# GitHub Personal Access Token (required)
PAT_1=your_github_token_here
# You can configure multiple tokens to increase API rate limits
PAT_2=your_second_token_here

# Server port (optional, default: 9000)
PORT=9000

# Cache duration in seconds (optional, default: 86400)
CACHE_SECONDS=86400
```

### 3. Run

```bash
# Method 1: Run directly
go run cmd/server/main.go

# Method 2: Use the built binary
./github-readme-stats

# Method 3: Using Docker
docker build -t github-readme-stats .
docker run -p 9000:9000 -e PAT_1=your_token github-readme-stats
```

### 4. Use

Once the server is running, you can access the API at `http://localhost:9000` and use the examples shown above!

## GitHub Actions (Recommended)

The recommended way to use this project is through [GitHub Readme Stats Action](https://github.com/soulteary/github-readme-stats-action), which automatically generates and updates stats cards in your repository using GitHub Actions workflows.

**Repository:** [soulteary/github-readme-stats-action](https://github.com/soulteary/github-readme-stats-action)  
**Marketplace:** [Readme Stats Action](https://github.com/marketplace/actions/readme-stats-action)

### Quick Start

Create a workflow file (e.g., `.github/workflows/update-stats.yml`):

```yaml
name: Update README cards

on:
  schedule:
    - cron: "0 0 * * *" # Runs once daily at midnight
  workflow_dispatch:

jobs:
  build:
    runs-on: ubuntu-latest

    permissions:
      contents: write

    steps:
      - uses: actions/checkout@v4

      - name: Generate stats card
        uses: soulteary/github-readme-stats-action@v1.0.0
        with:
          card: stats
          options: 'username=${{ github.repository_owner }}&show_icons=true'
          path: profile/stats.svg
          token: ${{ secrets.GITHUB_TOKEN }}

      - name: Generate top languages card
        uses: soulteary/github-readme-stats-action@v1.0.0
        with:
          card: top-langs
          options: 'username=${{ github.repository_owner }}&layout=compact&langs_count=6'
          path: profile/top-langs.svg
          token: ${{ secrets.GITHUB_TOKEN }}

      - name: Commit cards
        run: |
          git config user.name "github-actions[bot]"
          git config user.email "41898282+github-actions[bot]@users.noreply.github.com"
          git add profile/*.svg
          git commit -m "Update README cards" || exit 0
          git push
```

Then embed the generated cards in your README:

```markdown
![Stats](./profile/stats.svg)
![Top Languages](./profile/top-langs.svg)
```

For more details and examples, see the [GitHub Readme Stats Action documentation](https://github.com/soulteary/github-readme-stats-action).

## CLI Usage

You can use the CLI to generate cards locally without running a server. This is useful for testing, development, or generating static images.

### Build CLI Tools

```bash
# Build all CLI tools
make build

# Or build individually
make build-server    # Build the server
make build-examples  # Build the examples generator

# Quick demo
./demo.sh
```

### Generate Example Images

The `examples` command generates example images for all card types. Like other commands, it reads the `.env` file for configuration.

```bash
# Generate all example images (requires GitHub API token in .env)
./bin/examples

# Generate specific examples
./bin/examples stats-basic stats-dark repo-basic

# Generate test error cards (no network required)
./bin/examples --test

# Generate WakaTime test cards using test data (no network required)
./bin/examples --wakatime-test

# View help
./bin/examples --help
```

This creates SVG files in the `.github/assets/` directory, including:
- `stats-basic.svg` - Basic stats card
- `stats-dark.svg` - Dark theme stats card
- `stats-compact.svg` - Compact stats card
- `repo-basic.svg` - Repository pin card
- `top-langs-basic.svg` - Top languages card
- `gist-basic.svg` - Gist pin card
- `wakatime-basic.svg` - WakaTime stats card
- `wakatime-test-*.svg` - WakaTime test cards (generated with test data)
- And more...

### Generate Individual Cards

Use the server binary with CLI parameters:

```bash
# Generate stats card
./bin/server "/api?username=anuraghazra&theme=dark" --output=stats.svg

# Generate repository pin
./bin/server "/api/pin?username=anuraghazra&repo=github-readme-stats" --output=repo.svg

# Generate top languages
./bin/server "/api/top-langs?username=anuraghazra&layout=compact" --output=langs.svg

# Generate gist card
./bin/server "/api/gist?id=bbfce31e0217a3689c8d961a356cb10d" --output=gist.svg

# Generate WakaTime card
./bin/server "/api/wakatime?username=anuraghazra" --output=wakatime.svg
```

### Advanced CLI Usage

You can also use flag-style arguments:

```bash
# Using flags
./bin/server --type=stats --username=anuraghazra --theme=dark --output=stats.svg

# Multiple parameters
./bin/server --type=top-langs --username=anuraghazra --layout=compact --langs_count=8 --theme=radical --output=langs.svg
```

## API Endpoints

### GitHub Stats Card

Generate a GitHub statistics card for a user.

```
GET /api?username=USERNAME
```

**Parameters:**
- `username` (required) - GitHub username
- `hide` - Hide specific stats (comma-separated, e.g., `stars,commits`)
- `hide_title` - Hide title
- `hide_border` - Hide border
- `hide_rank` - Hide rank
- `show_icons` - Show icons
- `include_all_commits` - Include all commits
- `theme` - Theme name (80+ themes available)
- `bg_color` - Background color (hexadecimal)
- `title_color` - Title color
- `text_color` - Text color
- `icon_color` - Icon color
- `border_color` - Border color
- `border_radius` - Border radius
- `locale` - Language code (e.g., `zh`, `en`, `de`, `it`, `kr`, `ja`)
- `cache_seconds` - Cache duration in seconds

> See examples above for usage.

### Repository Pin Card

Generate a repository pin card.

```
GET /api/pin?username=USERNAME&repo=REPO
```

**Parameters:**
- `username` (required) - GitHub username
- `repo` (required) - Repository name
- `theme` - Theme name
- `show_owner` - Show owner
- `locale` - Language code

> See examples above for usage.

### Top Languages Card

Generate a most used programming languages statistics card.

```
GET /api/top-langs?username=USERNAME
```

**Parameters:**
- `username` (required) - GitHub username
- `hide` - Hide specific languages (comma-separated)
- `layout` - Layout type (`compact`, `normal`)
- `theme` - Theme name
- `locale` - Language code

> See examples above for usage.

### Gist Card

Generate a Gist card.

```
GET /api/gist?id=GIST_ID
```

**Parameters:**
- `id` (required) - Gist ID
- `theme` - Theme name
- `locale` - Language code

> See examples above for usage.

### WakaTime Card

Generate a WakaTime programming time statistics card.

```
GET /api/wakatime?username=USERNAME
```

**Parameters:**
- `username` (required) - WakaTime username
- `theme` - Theme name
- `hide` - Hide specific stats
- `locale` - Language code

> See examples above for usage.

## Project Structure

```
.
├── cmd/
│   └── server/          # Server entry point
│       └── main.go
├── internal/
│   ├── api/             # API handlers
│   ├── cards/           # Card rendering logic
│   ├── common/          # Common utilities and helper functions
│   ├── fetchers/        # Data fetchers (GitHub API, WakaTime, etc.)
│   ├── themes/          # Theme system
│   └── translations/    # Internationalization translations
├── pkg/
│   └── svg/             # SVG related utilities
├── Dockerfile           # Docker build file
├── go.mod              # Go module definition
└── README.md           # Project documentation
```

## Development Status

✅ **Core features completed!**

All major features have been implemented:

- ✅ Project base structure
- ✅ HTTP server (Gin)
- ✅ GitHub API integration (GraphQL + REST)
- ✅ Retry mechanism and multi-PAT support
- ✅ Cache handling
- ✅ Theme system (80+ themes)
- ✅ All card type rendering (Stats, Pin, Top Languages, Gist, WakaTime)
- ✅ Rank calculation
- ✅ Internationalization support (basic implementation)
- ✅ All API endpoints

## Contributing

Contributions are welcome! If you have any ideas or find issues, please:

1. Fork this project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

Or report issues directly in [Issues](https://github.com/soulteary/github-readme-stats/issues).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
