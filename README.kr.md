# GitHub Readme Stats (Go 구현)

[![GitHub](https://img.shields.io/badge/GitHub-soulteary%2Fgithub--readme--stats-blue)](https://github.com/soulteary/github-readme-stats)

![GitHub Readme Stats](.github/assets/banner.jpg)

## Languages / 语言 / Sprachen / Lingue / 언어 / 言語

- [English](README.md)
- [简体中文](README.zh.md)
- [Deutsch](README.de.md)
- [Italiano](README.it.md)
- [한국어](README.kr.md)
- [日本語](README.ja.md)

이것은 [GitHub Readme Stats (anuraghazra)](https://github.com/anuraghazra/github-readme-stats) 프로젝트의 Go 언어 구현 버전입니다. 이 프로젝트는 README 파일에 임베드할 수 있는 동적 GitHub 통계 카드를 생성하는 기능을 제공하여 GitHub 활동, 저장소 정보, 프로그래밍 언어 사용량 등을 보여줍니다. 또한 GitHub Action에서 직접 사용할 수 있습니다.

## 기능

- ✅ GitHub Stats 카드 생성
- ✅ Repository Pin 카드 생성
- ✅ Top Languages 카드 생성
- ✅ Gist 카드 생성
- ✅ WakaTime 카드 생성

## 예제

이 프로젝트로 만들 수 있는 몇 가지 예제입니다:

### GitHub Stats 카드

**기본:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname)
```
![GitHub Stats](.github/assets/stats-basic.svg)

**다크 테마:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname&theme=dark)
```
![GitHub Stats Dark](.github/assets/stats-dark.svg)

**컴팩트 레이아웃:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname&layout=compact)
```
![GitHub Stats Compact](.github/assets/stats-compact.svg)

**사용자 정의 테마:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname&bg_color=0d1117&title_color=ff6b6b&text_color=c9d1d9&border_color=30363d)
```
![GitHub Stats Custom](.github/assets/stats-custom.svg)

### Repository Pin 카드

**기본:**
```markdown
![Pinned Repo](http://localhost:9000/api/pin?username=yourname&repo=repo-name)
```
![Pinned Repo](.github/assets/repo-basic.svg)

**테마 적용:**
```markdown
![Pinned Repo](http://localhost:9000/api/pin?username=yourname&repo=repo-name&theme=dark)
```
![Pinned Repo Themed](.github/assets/repo-themed.svg)

### Top Languages 카드

**기본:**
```markdown
![Top Languages](http://localhost:9000/api/top-langs?username=yourname)
```
![Top Languages](.github/assets/top-langs-basic.svg)

**컴팩트 레이아웃:**
```markdown
![Top Languages](http://localhost:9000/api/top-langs?username=yourname&layout=compact&langs_count=6)
```
![Top Languages Compact](.github/assets/top-langs-compact.svg)

### Gist 카드

```markdown
![Gist](http://localhost:9000/api/gist?id=gist_id)
```
![Gist](.github/assets/gist-basic.svg)

### WakaTime 카드

**기본:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname)
```
![WakaTime Basic](.github/assets/wakatime-basic.svg)

**컴팩트 레이아웃:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&layout=compact)
```
![WakaTime Compact](.github/assets/wakatime-compact.svg)

**테스트 예제 (테스트 데이터로 생성):**

**기본 테스트:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname)
```
![WakaTime Test Basic](.github/assets/wakatime-test-basic.svg)

**컴팩트 테스트:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&layout=compact&langs_count=6&theme=dark)
```
![WakaTime Test Compact](.github/assets/wakatime-test-compact.svg)

**테마:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&theme=radical&langs_count=5)
```
![WakaTime Test Themed](.github/assets/wakatime-test-themed.svg)

**진행 표시줄 숨기기:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&hide_progress=true)
```
![WakaTime Test Hide Progress](.github/assets/wakatime-test-hide-progress.svg)

**백분율 표시:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&display_format=percent)
```
![WakaTime Test Percent](.github/assets/wakatime-test-percent.svg)

**제한된 언어:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&langs_count=3&theme=tokyonight)
```
![WakaTime Test Limited](.github/assets/wakatime-test-limited.svg)

> **참고:** WakaTime 카드 예제는 공개 통계가 있는 유효한 WakaTime 사용자 이름을 필요로 합니다. 위의 테스트 이미지는 다양한 구성 옵션을 보여주기 위해 테스트 데이터를 사용하여 생성되었습니다.

## 빠른 시작

### Go를 사용한 설치

```bash
go get github.com/soulteary/github-readme-stats
```

### 소스에서 빌드

```bash
git clone https://github.com/soulteary/github-readme-stats.git
cd github-readme-stats
go build -o github-readme-stats ./cmd/server
```

## 빠른 시작

### 1. 설치

```bash
# Go를 사용한 설치
go get github.com/soulteary/github-readme-stats

# 또는 소스에서 빌드
git clone https://github.com/soulteary/github-readme-stats.git
cd github-readme-stats
go build -o github-readme-stats ./cmd/server
```

### 2. 구성

`.env` 파일을 생성하세요:

```bash
# GitHub Personal Access Token (필수)
PAT_1=your_github_token_here
# API 제한을 늘리기 위해 여러 토큰을 구성할 수 있습니다
PAT_2=your_second_token_here

# 서버 포트 (선택 사항, 기본값: 9000)
PORT=9000

# 캐시 지속 시간(초) (선택 사항, 기본값: 86400)
CACHE_SECONDS=86400
```

### 3. 실행

```bash
# 방법 1: 직접 실행
go run cmd/server/main.go

# 방법 2: 빌드된 바이너리 사용
./github-readme-stats

# 방법 3: Docker 사용
docker build -t github-readme-stats .
docker run -p 9000:9000 -e PAT_1=your_token github-readme-stats
```

### 4. 사용

서버가 실행되면 `http://localhost:9000`에서 API에 액세스하고 위에 표시된 예제를 사용할 수 있습니다!

## GitHub Actions (권장)

이 프로젝트를 사용하는 권장 방법은 [GitHub Readme Stats Action](https://github.com/soulteary/github-readme-stats-action)을 통해 GitHub Actions 워크플로를 사용하여 저장소의 통계 카드를 자동으로 생성하고 업데이트하는 것입니다.

**저장소:** [soulteary/github-readme-stats-action](https://github.com/soulteary/github-readme-stats-action)  
**마켓플레이스:** [Readme Stats Action](https://github.com/marketplace/actions/readme-stats-action)

### 빠른 시작

워크플로 파일을 생성하세요 (예: `.github/workflows/update-stats.yml`):

```yaml
name: Update README cards

on:
  schedule:
    - cron: "0 0 * * *" # 매일 자정에 한 번 실행
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

그런 다음 생성된 카드를 README에 임베드하세요:

```markdown
![Stats](./profile/stats.svg)
![Top Languages](./profile/top-langs.svg)
```

자세한 내용과 예제는 [GitHub Readme Stats Action 문서](https://github.com/soulteary/github-readme-stats-action)를 참조하세요.

## CLI 사용

서버를 실행하지 않고 로컬에서 카드를 생성하기 위해 CLI를 사용할 수 있습니다. 이는 테스트, 개발 또는 정적 이미지 생성에 유용합니다.

### CLI 도구 빌드

```bash
# 모든 CLI 도구 빌드
make build

# 또는 개별적으로 빌드
make build-server    # 서버 빌드
make build-examples  # 예제 생성기 빌드

# 빠른 데모
./demo.sh
```

### 예제 이미지 생성

`examples` 명령은 모든 카드 유형에 대한 예제 이미지를 생성합니다. 다른 명령과 마찬가지로 `.env` 파일에서 구성을 읽습니다.

```bash
# 모든 예제 이미지 생성 (.env 파일에 GitHub API 토큰 필요)
./bin/examples

# 특정 예제 생성
./bin/examples stats-basic stats-dark repo-basic

# 테스트 오류 카드 생성 (네트워크 필요 없음)
./bin/examples --test

# 테스트 데이터를 사용하여 WakaTime 테스트 카드 생성 (네트워크 필요 없음)
./bin/examples --wakatime-test

# 도움말 보기
./bin/examples --help
```

이렇게 하면 `.github/assets/` 디렉토리에 SVG 파일이 생성되며 다음을 포함합니다:
- `stats-basic.svg` - 기본 통계 카드
- `stats-dark.svg` - 다크 테마 통계 카드
- `stats-compact.svg` - 컴팩트 통계 카드
- `repo-basic.svg` - 저장소 Pin 카드
- `top-langs-basic.svg` - Top Languages 카드
- `gist-basic.svg` - Gist Pin 카드
- `wakatime-basic.svg` - WakaTime 통계 카드
- `wakatime-test-*.svg` - WakaTime 테스트 카드 (테스트 데이터로 생성)
- 그리고 더 많은...

### 개별 카드 생성

CLI 매개변수와 함께 서버 바이너리 사용:

```bash
# 통계 카드 생성
./bin/server "/api?username=anuraghazra&theme=dark" --output=stats.svg

# 저장소 Pin 생성
./bin/server "/api/pin?username=anuraghazra&repo=github-readme-stats" --output=repo.svg

# Top Languages 생성
./bin/server "/api/top-langs?username=anuraghazra&layout=compact" --output=langs.svg

# Gist 카드 생성
./bin/server "/api/gist?id=bbfce31e0217a3689c8d961a356cb10d" --output=gist.svg

# WakaTime 카드 생성
./bin/server "/api/wakatime?username=anuraghazra" --output=wakatime.svg
```

### 고급 CLI 사용

플래그 스타일 인수를 사용할 수도 있습니다:

```bash
# 플래그 사용
./bin/server --type=stats --username=anuraghazra --theme=dark --output=stats.svg

# 여러 매개변수
./bin/server --type=top-langs --username=anuraghazra --layout=compact --langs_count=8 --theme=radical --output=langs.svg
```

## API 엔드포인트

### GitHub Stats 카드

사용자에 대한 GitHub 통계 카드를 생성합니다.

```
GET /api?username=USERNAME
```

**매개변수:**
- `username` (필수) - GitHub 사용자 이름
- `hide` - 특정 통계 숨기기 (쉼표로 구분, 예: `stars,commits`)
- `hide_title` - 제목 숨기기
- `hide_border` - 테두리 숨기기
- `hide_rank` - 순위 숨기기
- `show_icons` - 아이콘 표시
- `include_all_commits` - 모든 커밋 포함
- `theme` - 테마 이름 (80개 이상의 테마 사용 가능)
- `bg_color` - 배경색 (16진수)
- `title_color` - 제목 색상
- `text_color` - 텍스트 색상
- `icon_color` - 아이콘 색상
- `border_color` - 테두리 색상
- `border_radius` - 테두리 반경
- `locale` - 언어 코드 (예: `zh`, `en`, `de`, `it`, `kr`, `ja`)
- `cache_seconds` - 캐시 지속 시간(초)

> 사용 예제는 위를 참조하세요.

### Repository Pin 카드

저장소 Pin 카드를 생성합니다.

```
GET /api/pin?username=USERNAME&repo=REPO
```

**매개변수:**
- `username` (필수) - GitHub 사용자 이름
- `repo` (필수) - 저장소 이름
- `theme` - 테마 이름
- `show_owner` - 소유자 표시
- `locale` - 언어 코드

> 사용 예제는 위를 참조하세요.

### Top Languages 카드

가장 많이 사용된 프로그래밍 언어 통계 카드를 생성합니다.

```
GET /api/top-langs?username=USERNAME
```

**매개변수:**
- `username` (필수) - GitHub 사용자 이름
- `hide` - 특정 언어 숨기기 (쉼표로 구분)
- `layout` - 레이아웃 유형 (`compact`, `normal`)
- `theme` - 테마 이름
- `locale` - 언어 코드

> 사용 예제는 위를 참조하세요.

### Gist 카드

Gist 카드를 생성합니다.

```
GET /api/gist?id=GIST_ID
```

**매개변수:**
- `id` (필수) - Gist ID
- `theme` - 테마 이름
- `locale` - 언어 코드

> 사용 예제는 위를 참조하세요.

### WakaTime 카드

WakaTime 프로그래밍 시간 통계 카드를 생성합니다.

```
GET /api/wakatime?username=USERNAME
```

**매개변수:**
- `username` (필수) - WakaTime 사용자 이름
- `theme` - 테마 이름
- `hide` - 특정 통계 숨기기
- `locale` - 언어 코드

> 사용 예제는 위를 참조하세요.

## 프로젝트 구조

```
.
├── cmd/
│   └── server/          # 서버 진입점
│       └── main.go
├── internal/
│   ├── api/             # API 핸들러
│   ├── cards/           # 카드 렌더링 로직
│   ├── common/          # 공통 유틸리티 및 헬퍼 함수
│   ├── fetchers/        # 데이터 페처 (GitHub API, WakaTime 등)
│   ├── themes/          # 테마 시스템
│   └── translations/    # 국제화 번역
├── pkg/
│   └── svg/             # SVG 관련 유틸리티
├── Dockerfile           # Docker 빌드 파일
├── go.mod              # Go 모듈 정의
└── README.md           # 프로젝트 문서
```

## 개발 상태

✅ **핵심 기능 완료!**

모든 주요 기능이 구현되었습니다:

- ✅ 프로젝트 기본 구조
- ✅ HTTP 서버 (Gin)
- ✅ GitHub API 통합 (GraphQL + REST)
- ✅ 재시도 메커니즘 및 다중 PAT 지원
- ✅ 캐시 처리
- ✅ 테마 시스템 (80개 이상의 테마)
- ✅ 모든 카드 유형 렌더링 (Stats, Pin, Top Languages, Gist, WakaTime)
- ✅ 순위 계산
- ✅ 국제화 지원 (기본 구현)
- ✅ 모든 API 엔드포인트

## 기여

기여를 환영합니다! 아이디어가 있거나 문제를 발견한 경우:

1. 이 프로젝트를 포크하세요
2. 기능 브랜치를 만드세요 (`git checkout -b feature/AmazingFeature`)
3. 변경 사항을 커밋하세요 (`git commit -m 'Add some AmazingFeature'`)
4. 브랜치에 푸시하세요 (`git push origin feature/AmazingFeature`)
5. Pull Request를 열어주세요

또는 [Issues](https://github.com/soulteary/github-readme-stats/issues)에서 직접 문제를 보고하세요.

## 라이선스

이 프로젝트는 MIT 라이선스 하에 있습니다. 자세한 내용은 [LICENSE](LICENSE) 파일을 참조하세요.
