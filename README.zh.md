# GitHub Readme Stats (Go 实现)

[![GitHub](https://img.shields.io/badge/GitHub-soulteary%2Fgithub--readme--stats-blue)](https://github.com/soulteary/github-readme-stats)

![GitHub Readme Stats](.github/assets/banner.jpg)

## Languages / 语言 / Sprachen / Lingue / 언어 / 言語

- [English](README.md)
- [简体中文](README.zh.md)
- [Deutsch](README.de.md)
- [Italiano](README.it.md)
- [한국어](README.kr.md)
- [日本語](README.ja.md)

这是 [GitHub Readme Stats (anuraghazra)](https://github.com/anuraghazra/github-readme-stats) 项目的 Go 语言实现版本。该项目提供了动态生成 GitHub 统计卡片的功能，可以嵌入到 README 文件中，展示你的 GitHub 活动、仓库信息、编程语言使用情况等。同时支持在 GitHub Action 中直接使用。

## 功能

- ✅ GitHub Stats 卡片生成
- ✅ Repository Pin 卡片生成
- ✅ Top Languages 卡片生成
- ✅ Gist 卡片生成
- ✅ WakaTime 卡片生成

## 示例展示

以下是一些使用本项目可以创建的示例：

### GitHub Stats 卡片

**基础版本：**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname)
```
![GitHub Stats](.github/assets/stats-basic.svg)

**深色主题：**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname&theme=dark)
```
![GitHub Stats Dark](.github/assets/stats-dark.svg)

**紧凑布局：**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname&layout=compact)
```
![GitHub Stats Compact](.github/assets/stats-compact.svg)

**自定义主题：**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname&bg_color=0d1117&title_color=ff6b6b&text_color=c9d1d9&border_color=30363d)
```
![GitHub Stats Custom](.github/assets/stats-custom.svg)

### Repository Pin 卡片

**基础版本：**
```markdown
![Pinned Repo](http://localhost:9000/api/pin?username=yourname&repo=repo-name)
```
![Pinned Repo](.github/assets/repo-basic.svg)

**主题版本：**
```markdown
![Pinned Repo](http://localhost:9000/api/pin?username=yourname&repo=repo-name&theme=dark)
```
![Pinned Repo Themed](.github/assets/repo-themed.svg)

### Top Languages 卡片

**基础版本：**
```markdown
![Top Languages](http://localhost:9000/api/top-langs?username=yourname)
```
![Top Languages](.github/assets/top-langs-basic.svg)

**紧凑布局：**
```markdown
![Top Languages](http://localhost:9000/api/top-langs?username=yourname&layout=compact&langs_count=6)
```
![Top Languages Compact](.github/assets/top-langs-compact.svg)

### Gist 卡片

```markdown
![Gist](http://localhost:9000/api/gist?id=gist_id)
```
![Gist](.github/assets/gist-basic.svg)

### WakaTime 卡片

**基础版本：**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname)
```
![WakaTime Basic](.github/assets/wakatime-basic.svg)

**紧凑布局：**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&layout=compact)
```
![WakaTime Compact](.github/assets/wakatime-compact.svg)

**测试示例（使用测试数据生成）：**

**基础测试：**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname)
```
![WakaTime Test Basic](.github/assets/wakatime-test-basic.svg)

**紧凑测试：**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&layout=compact&langs_count=6&theme=dark)
```
![WakaTime Test Compact](.github/assets/wakatime-test-compact.svg)

**主题样式：**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&theme=radical&langs_count=5)
```
![WakaTime Test Themed](.github/assets/wakatime-test-themed.svg)

**隐藏进度条：**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&hide_progress=true)
```
![WakaTime Test Hide Progress](.github/assets/wakatime-test-hide-progress.svg)

**百分比显示：**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&display_format=percent)
```
![WakaTime Test Percent](.github/assets/wakatime-test-percent.svg)

**限制语言数量：**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&langs_count=3&theme=tokyonight)
```
![WakaTime Test Limited](.github/assets/wakatime-test-limited.svg)

> **注意：** WakaTime 卡片示例需要有效的 WakaTime 用户名和公开统计数据。上述测试图片使用测试数据生成，用于展示各种配置选项。

## 快速开始

### 使用 Go 安装

```bash
go get github.com/soulteary/github-readme-stats
```

### 从源码构建

```bash
git clone https://github.com/soulteary/github-readme-stats.git
cd github-readme-stats
go build -o github-readme-stats ./cmd/server
```

## 快速开始

### 1. 安装

```bash
# 使用 Go 安装
go get github.com/soulteary/github-readme-stats

# 或从源码构建
git clone https://github.com/soulteary/github-readme-stats.git
cd github-readme-stats
go build -o github-readme-stats ./cmd/server
```

### 2. 配置

创建 `.env` 文件：

```bash
# GitHub Personal Access Token (必需)
PAT_1=your_github_token_here
# 可以配置多个 token 以提高 API 限制
PAT_2=your_second_token_here

# 服务器端口 (可选，默认 9000)
PORT=9000

# 缓存时间（秒）(可选，默认 86400)
CACHE_SECONDS=86400
```

### 3. 运行

```bash
# 方式 1: 直接运行
go run cmd/server/main.go

# 方式 2: 使用构建后的二进制文件
./github-readme-stats

# 方式 3: 使用 Docker
docker build -t github-readme-stats .
docker run -p 9000:9000 -e PAT_1=your_token github-readme-stats
```

### 4. 使用

服务器运行后，你可以访问 `http://localhost:9000` 的 API，并使用上面展示的示例！

## GitHub Actions（推荐方式）

推荐使用 [GitHub Readme Stats Action](https://github.com/soulteary/github-readme-stats-action)，它可以通过 GitHub Actions 工作流自动生成并更新仓库中的统计卡片。

**仓库地址：** [soulteary/github-readme-stats-action](https://github.com/soulteary/github-readme-stats-action)  
**市场地址：** [Readme Stats Action](https://github.com/marketplace/actions/readme-stats-action)

### 快速开始

创建工作流文件（例如 `.github/workflows/update-stats.yml`）：

```yaml
name: Update README cards

on:
  schedule:
    - cron: "0 0 * * *" # 每天午夜运行一次
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

然后在你的 README 中嵌入生成的卡片：

```markdown
![Stats](./profile/stats.svg)
![Top Languages](./profile/top-langs.svg)
```

更多详情和示例，请参阅 [GitHub Readme Stats Action 文档](https://github.com/soulteary/github-readme-stats-action)。

## CLI 使用

你可以使用 CLI 在本地生成卡片，而无需运行服务器。这对于测试、开发或生成静态图像很有用。

### 构建 CLI 工具

```bash
# 构建所有 CLI 工具
make build

# 或单独构建
make build-server    # 构建服务器
make build-examples  # 构建示例生成器

# 快速演示
./demo.sh
```

### 生成示例图像

`examples` 命令会生成所有卡片类型的示例图像。与其他命令一样，它从 `.env` 文件读取配置。

```bash
# 生成所有示例图像（需要 .env 文件中的 GitHub API token）
./bin/examples

# 生成特定示例
./bin/examples stats-basic stats-dark repo-basic

# 生成测试错误卡片（无需网络）
./bin/examples --test

# 使用测试数据生成 WakaTime 测试卡片（无需网络）
./bin/examples --wakatime-test

# 查看帮助
./bin/examples --help
```

这会在 `.github/assets/` 目录下创建 SVG 文件，包括：
- `stats-basic.svg` - 基础统计卡片
- `stats-dark.svg` - 深色主题统计卡片
- `stats-compact.svg` - 紧凑统计卡片
- `repo-basic.svg` - 仓库 Pin 卡片
- `top-langs-basic.svg` - Top Languages 卡片
- `gist-basic.svg` - Gist Pin 卡片
- `wakatime-basic.svg` - WakaTime 统计卡片
- `wakatime-test-*.svg` - WakaTime 测试卡片（使用测试数据生成）
- 以及更多...

### 生成单个卡片

使用服务器二进制文件配合 CLI 参数：

```bash
# 生成统计卡片
./bin/server "/api?username=anuraghazra&theme=dark" --output=stats.svg

# 生成仓库 Pin
./bin/server "/api/pin?username=anuraghazra&repo=github-readme-stats" --output=repo.svg

# 生成 Top Languages
./bin/server "/api/top-langs?username=anuraghazra&layout=compact" --output=langs.svg

# 生成 Gist 卡片
./bin/server "/api/gist?id=bbfce31e0217a3689c8d961a356cb10d" --output=gist.svg

# 生成 WakaTime 卡片
./bin/server "/api/wakatime?username=anuraghazra" --output=wakatime.svg
```

### 高级 CLI 使用

你也可以使用标志式参数：

```bash
# 使用标志
./bin/server --type=stats --username=anuraghazra --theme=dark --output=stats.svg

# 多个参数
./bin/server --type=top-langs --username=anuraghazra --layout=compact --langs_count=8 --theme=radical --output=langs.svg
```

## API 端点

### GitHub Stats 卡片

生成用户 GitHub 统计信息卡片。

```
GET /api?username=USERNAME
```

**参数：**
- `username` (必需) - GitHub 用户名
- `hide` - 隐藏指定统计项（逗号分隔，如：`stars,commits`）
- `hide_title` - 隐藏标题
- `hide_border` - 隐藏边框
- `hide_rank` - 隐藏排名
- `show_icons` - 显示图标
- `include_all_commits` - 包含所有提交
- `theme` - 主题名称（80+ 主题可选）
- `bg_color` - 背景颜色（十六进制）
- `title_color` - 标题颜色
- `text_color` - 文本颜色
- `icon_color` - 图标颜色
- `border_color` - 边框颜色
- `border_radius` - 边框圆角
- `locale` - 语言代码（如：`zh`, `en`, `de`, `it`, `kr`, `ja`）
- `cache_seconds` - 缓存时间（秒）

> 使用示例请参考上方。

### Repository Pin 卡片

生成仓库 Pin 卡片。

```
GET /api/pin?username=USERNAME&repo=REPO
```

**参数：**
- `username` (必需) - GitHub 用户名
- `repo` (必需) - 仓库名称
- `theme` - 主题名称
- `show_owner` - 显示所有者
- `locale` - 语言代码

> 使用示例请参考上方。

### Top Languages 卡片

生成最常用编程语言统计卡片。

```
GET /api/top-langs?username=USERNAME
```

**参数：**
- `username` (必需) - GitHub 用户名
- `hide` - 隐藏指定语言（逗号分隔）
- `layout` - 布局类型（`compact`, `normal`）
- `theme` - 主题名称
- `locale` - 语言代码

> 使用示例请参考上方。

### Gist 卡片

生成 Gist 卡片。

```
GET /api/gist?id=GIST_ID
```

**参数：**
- `id` (必需) - Gist ID
- `theme` - 主题名称
- `locale` - 语言代码

> 使用示例请参考上方。

### WakaTime 卡片

生成 WakaTime 编程时间统计卡片。

```
GET /api/wakatime?username=USERNAME
```

**参数：**
- `username` (必需) - WakaTime 用户名
- `theme` - 主题名称
- `hide` - 隐藏指定统计项
- `locale` - 语言代码

> 使用示例请参考上方。

## 项目结构

```
.
├── cmd/
│   └── server/          # 服务器入口
│       └── main.go
├── internal/
│   ├── api/             # API 处理器
│   ├── cards/           # 卡片渲染逻辑
│   ├── common/          # 通用工具和辅助函数
│   ├── fetchers/        # 数据获取器（GitHub API, WakaTime 等）
│   ├── themes/          # 主题系统
│   └── translations/    # 国际化翻译
├── pkg/
│   └── svg/             # SVG 相关工具
├── Dockerfile           # Docker 构建文件
├── go.mod              # Go 模块定义
└── README.md           # 项目文档
```

## 开发状态

✅ **核心功能已完成！**

所有主要功能都已实现：

- ✅ 项目基础结构
- ✅ HTTP 服务器 (Gin)
- ✅ GitHub API 集成 (GraphQL + REST)
- ✅ 重试机制和多 PAT 支持
- ✅ 缓存处理
- ✅ 主题系统 (80+ 主题)
- ✅ 所有卡片类型渲染 (Stats, Pin, Top Languages, Gist, WakaTime)
- ✅ 排名计算
- ✅ 国际化支持 (基础实现)
- ✅ 所有 API 端点

## 贡献

欢迎贡献代码！如果你有任何想法或发现问题，请：

1. Fork 本项目
2. 创建你的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交你的更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启一个 Pull Request

或者直接在 [Issues](https://github.com/soulteary/github-readme-stats/issues) 中报告问题。

## 许可证

本项目采用 MIT 许可证。详情请参阅 [LICENSE](LICENSE) 文件。
