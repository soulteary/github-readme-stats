# GitHub Readme Stats (Go 実装)

[![GitHub](https://img.shields.io/badge/GitHub-soulteary%2Fgithub--readme--stats-blue)](https://github.com/soulteary/github-readme-stats)

![GitHub Readme Stats](.github/assets/banner.jpg)

## Languages / 语言 / Sprachen / Lingue / 언어 / 言語

- [English](README.md)
- [简体中文](README.zh.md)
- [Deutsch](README.de.md)
- [Italiano](README.it.md)
- [한국어](README.kr.md)
- [日本語](README.ja.md)

これは [GitHub Readme Stats (anuraghazra)](https://github.com/anuraghazra/github-readme-stats) プロジェクトの Go 言語実装版です。このプロジェクトは、README ファイルに埋め込むことができる動的な GitHub 統計カードを生成する機能を提供し、GitHub の活動、リポジトリ情報、プログラミング言語の使用状況などを表示します。また、GitHub Action で直接使用することもできます。

## 機能

- ✅ GitHub Stats カード生成
- ✅ Repository Pin カード生成
- ✅ Top Languages カード生成
- ✅ Gist カード生成
- ✅ WakaTime カード生成

## 例

このプロジェクトで作成できる例をいくつか示します：

### GitHub Stats カード

**基本:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname)
```
![GitHub Stats](.github/assets/stats-basic.svg)

**ダークテーマ:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname&theme=dark)
```
![GitHub Stats Dark](.github/assets/stats-dark.svg)

**コンパクトレイアウト:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname&layout=compact)
```
![GitHub Stats Compact](.github/assets/stats-compact.svg)

**カスタムテーマ:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname&bg_color=0d1117&title_color=ff6b6b&text_color=c9d1d9&border_color=30363d)
```
![GitHub Stats Custom](.github/assets/stats-custom.svg)

### Repository Pin カード

**基本:**
```markdown
![Pinned Repo](http://localhost:9000/api/pin?username=yourname&repo=repo-name)
```
![Pinned Repo](.github/assets/repo-basic.svg)

**テーマ適用:**
```markdown
![Pinned Repo](http://localhost:9000/api/pin?username=yourname&repo=repo-name&theme=dark)
```
![Pinned Repo Themed](.github/assets/repo-themed.svg)

### Top Languages カード

**基本:**
```markdown
![Top Languages](http://localhost:9000/api/top-langs?username=yourname)
```
![Top Languages](.github/assets/top-langs-basic.svg)

**コンパクトレイアウト:**
```markdown
![Top Languages](http://localhost:9000/api/top-langs?username=yourname&layout=compact&langs_count=6)
```
![Top Languages Compact](.github/assets/top-langs-compact.svg)

### Gist カード

```markdown
![Gist](http://localhost:9000/api/gist?id=gist_id)
```
![Gist](.github/assets/gist-basic.svg)

### WakaTime カード

**基本:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname)
```
![WakaTime Basic](.github/assets/wakatime-basic.svg)

**コンパクトレイアウト:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&layout=compact)
```
![WakaTime Compact](.github/assets/wakatime-compact.svg)

**テスト例（テストデータで生成）:**

**基本テスト:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname)
```
![WakaTime Test Basic](.github/assets/wakatime-test-basic.svg)

**コンパクトテスト:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&layout=compact&langs_count=6&theme=dark)
```
![WakaTime Test Compact](.github/assets/wakatime-test-compact.svg)

**テーマ:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&theme=radical&langs_count=5)
```
![WakaTime Test Themed](.github/assets/wakatime-test-themed.svg)

**プログレスバー非表示:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&hide_progress=true)
```
![WakaTime Test Hide Progress](.github/assets/wakatime-test-hide-progress.svg)

**パーセント表示:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&display_format=percent)
```
![WakaTime Test Percent](.github/assets/wakatime-test-percent.svg)

**言語数制限:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&langs_count=3&theme=tokyonight)
```
![WakaTime Test Limited](.github/assets/wakatime-test-limited.svg)

> **注意:** WakaTime カードの例は公開統計のある有効な WakaTime ユーザー名が必要です。上記のテスト画像は、さまざまな設定オプションを示すためにテストデータを使用して生成されています。

## クイックスタート

### Go を使用したインストール

```bash
go get github.com/soulteary/github-readme-stats
```

### ソースからビルド

```bash
git clone https://github.com/soulteary/github-readme-stats.git
cd github-readme-stats
go build -o github-readme-stats ./cmd/server
```

## クイックスタート

### 1. インストール

```bash
# Go を使用したインストール
go get github.com/soulteary/github-readme-stats

# またはソースからビルド
git clone https://github.com/soulteary/github-readme-stats.git
cd github-readme-stats
go build -o github-readme-stats ./cmd/server
```

### 2. 設定

`.env` ファイルを作成してください：

```bash
# GitHub Personal Access Token (必須)
PAT_1=your_github_token_here
# API レート制限を増やすために複数のトークンを設定できます
PAT_2=your_second_token_here

# サーバーポート (オプション、デフォルト: 9000)
PORT=9000

# キャッシュ持続時間（秒）(オプション、デフォルト: 86400)
CACHE_SECONDS=86400
```

### 3. 実行

```bash
# 方法 1: 直接実行
go run cmd/server/main.go

# 方法 2: ビルドされたバイナリを使用
./github-readme-stats

# 方法 3: Docker を使用
docker build -t github-readme-stats .
docker run -p 9000:9000 -e PAT_1=your_token github-readme-stats
```

### 4. 使用

サーバーが実行されると、`http://localhost:9000` で API にアクセスし、上記の例を使用できます！

## GitHub Actions (推奨)

このプロジェクトを使用する推奨方法は、[GitHub Readme Stats Action](https://github.com/soulteary/github-readme-stats-action) を通じて、GitHub Actions ワークフローを使用してリポジトリの統計カードを自動的に生成および更新することです。

**リポジトリ:** [soulteary/github-readme-stats-action](https://github.com/soulteary/github-readme-stats-action)  
**マーケットプレイス:** [Readme Stats Action](https://github.com/marketplace/actions/readme-stats-action)

### クイックスタート

ワークフローファイルを作成してください (例: `.github/workflows/update-stats.yml`):

```yaml
name: Update README cards

on:
  schedule:
    - cron: "0 0 * * *" # 毎日深夜に1回実行
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

次に、生成されたカードを README に埋め込みます:

```markdown
![Stats](./profile/stats.svg)
![Top Languages](./profile/top-langs.svg)
```

詳細と例については、[GitHub Readme Stats Action ドキュメント](https://github.com/soulteary/github-readme-stats-action)を参照してください。

## CLI 利用

サーバーを実行せずにローカルでカードを生成するために CLI を使用できます。これはテスト、開発、または静的画像の生成に便利です。

### CLI ツールの構築

```bash
# すべての CLI ツールを構築
make build

# または個別に構築
make build-server    # サーバーを構築
make build-examples  # 例生成器を構築

# クイックデモ
./demo.sh
```

### 例画像の生成

`examples` コマンドはすべてのカードタイプの例画像を生成します。他のコマンドと同様に、`.env` ファイルから設定を読み取ります。

```bash
# すべての例画像を生成 (.env ファイルに GitHub API トークンが必要)
./bin/examples

# 特定の例を生成
./bin/examples stats-basic stats-dark repo-basic

# テストエラーカードを生成 (ネットワーク不要)
./bin/examples --test

# テストデータを使用して WakaTime テストカードを生成 (ネットワーク不要)
./bin/examples --wakatime-test

# ヘルプを表示
./bin/examples --help
```

これにより `.github/assets/` ディレクトリに SVG ファイルが作成され、以下を含みます：
- `stats-basic.svg` - 基本統計カード
- `stats-dark.svg` - ダークテーマ統計カード
- `stats-compact.svg` - コンパクト統計カード
- `repo-basic.svg` - リポジトリ Pin カード
- `top-langs-basic.svg` - Top Languages カード
- `gist-basic.svg` - Gist Pin カード
- `wakatime-basic.svg` - WakaTime 統計カード
- `wakatime-test-*.svg` - WakaTime テストカード（テストデータで生成）
- そしてもっと...

### 個別カードの生成

CLI パラメータとともにサーバーバイナリを使用：

```bash
# 統計カードを生成
./bin/server "/api?username=anuraghazra&theme=dark" --output=stats.svg

# リポジトリ Pin を生成
./bin/server "/api/pin?username=anuraghazra&repo=github-readme-stats" --output=repo.svg

# Top Languages を生成
./bin/server "/api/top-langs?username=anuraghazra&layout=compact" --output=langs.svg

# Gist カードを生成
./bin/server "/api/gist?id=bbfce31e0217a3689c8d961a356cb10d" --output=gist.svg

# WakaTime カードを生成
./bin/server "/api/wakatime?username=anuraghazra" --output=wakatime.svg
```

### 高度な CLI 利用

フラグスタイルの引数を使用することもできます：

```bash
# フラグを使用
./bin/server --type=stats --username=anuraghazra --theme=dark --output=stats.svg

# 複数のパラメータ
./bin/server --type=top-langs --username=anuraghazra --layout=compact --langs_count=8 --theme=radical --output=langs.svg
```

## API エンドポイント

### GitHub Stats カード

ユーザーの GitHub 統計カードを生成します。

```
GET /api?username=USERNAME
```

**パラメータ:**
- `username` (必須) - GitHub ユーザー名
- `hide` - 特定の統計を非表示にする（カンマ区切り、例: `stars,commits`）
- `hide_title` - タイトルを非表示
- `hide_border` - ボーダーを非表示
- `hide_rank` - ランクを非表示
- `show_icons` - アイコンを表示
- `include_all_commits` - すべてのコミットを含める
- `theme` - テーマ名 (80以上のテーマが利用可能)
- `bg_color` - 背景色 (16進数)
- `title_color` - タイトル色
- `text_color` - テキスト色
- `icon_color` - アイコン色
- `border_color` - ボーダー色
- `border_radius` - ボーダー半径
- `locale` - 言語コード (例: `zh`, `en`, `de`, `it`, `kr`, `ja`)
- `cache_seconds` - キャッシュ持続時間（秒）

> 使用例は上記を参照してください。

### Repository Pin カード

リポジトリ Pin カードを生成します。

```
GET /api/pin?username=USERNAME&repo=REPO
```

**パラメータ:**
- `username` (必須) - GitHub ユーザー名
- `repo` (必須) - リポジトリ名
- `theme` - テーマ名
- `show_owner` - 所有者を表示
- `locale` - 言語コード

> 使用例は上記を参照してください。

### Top Languages カード

最も使用されているプログラミング言語の統計カードを生成します。

```
GET /api/top-langs?username=USERNAME
```

**パラメータ:**
- `username` (必須) - GitHub ユーザー名
- `hide` - 特定の言語を非表示にする（カンマ区切り）
- `layout` - レイアウトタイプ (`compact`, `normal`)
- `theme` - テーマ名
- `locale` - 言語コード

> 使用例は上記を参照してください。

### Gist カード

Gist カードを生成します。

```
GET /api/gist?id=GIST_ID
```

**パラメータ:**
- `id` (必須) - Gist ID
- `theme` - テーマ名
- `locale` - 言語コード

> 使用例は上記を参照してください。

### WakaTime カード

WakaTime プログラミング時間統計カードを生成します。

```
GET /api/wakatime?username=USERNAME
```

**パラメータ:**
- `username` (必須) - WakaTime ユーザー名
- `theme` - テーマ名
- `hide` - 特定の統計を非表示にする
- `locale` - 言語コード

> 使用例は上記を参照してください。

## プロジェクト構造

```
.
├── cmd/
│   └── server/          # サーバーエントリーポイント
│       └── main.go
├── internal/
│   ├── api/             # API ハンドラー
│   ├── cards/           # カードレンダリングロジック
│   ├── common/          # 共通ユーティリティとヘルパー関数
│   ├── fetchers/        # データフェッチャー (GitHub API, WakaTime など)
│   ├── themes/          # テーマシステム
│   └── translations/    # 国際化翻訳
├── pkg/
│   └── svg/             # SVG 関連ユーティリティ
├── Dockerfile           # Docker ビルドファイル
├── go.mod              # Go モジュール定義
└── README.md           # プロジェクトドキュメント
```

## 開発状況

✅ **コア機能が完了しました！**

すべての主要機能が実装されました：

- ✅ プロジェクト基本構造
- ✅ HTTP サーバー (Gin)
- ✅ GitHub API 統合 (GraphQL + REST)
- ✅ リトライメカニズムとマルチ PAT サポート
- ✅ キャッシュ処理
- ✅ テーマシステム (80以上のテーマ)
- ✅ すべてのカードタイプレンダリング (Stats, Pin, Top Languages, Gist, WakaTime)
- ✅ ランク計算
- ✅ 国際化サポート (基本実装)
- ✅ すべての API エンドポイント

## 貢献

貢献を歓迎します！アイデアがある場合や問題を見つけた場合は、以下を行ってください：

1. このプロジェクトをフォークしてください
2. 機能ブランチを作成してください (`git checkout -b feature/AmazingFeature`)
3. 変更をコミットしてください (`git commit -m 'Add some AmazingFeature'`)
4. ブランチにプッシュしてください (`git push origin feature/AmazingFeature`)
5. Pull Request を開いてください

または、[Issues](https://github.com/soulteary/github-readme-stats/issues) で直接問題を報告してください。

## ライセンス

このプロジェクトは MIT ライセンスの下でライセンスされています。詳細については [LICENSE](LICENSE) ファイルを参照してください。
