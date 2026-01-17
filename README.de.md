# GitHub Readme Stats (Go Implementierung)

[![GitHub](https://img.shields.io/badge/GitHub-soulteary%2Fgithub--readme--stats-blue)](https://github.com/soulteary/github-readme-stats)

![GitHub Readme Stats](.github/assets/banner.jpg)

## Languages / 语言 / Sprachen / Lingue / 언어 / 言語

- [English](README.md)
- [简体中文](README.zh.md)
- [Deutsch](README.de.md)
- [Italiano](README.it.md)
- [한국어](README.kr.md)
- [日本語](README.ja.md)

Dies ist eine Go-Implementierung des [GitHub Readme Stats (anuraghazra)](https://github.com/anuraghazra/github-readme-stats) Projekts. Es bietet dynamische GitHub-Statistik-Karten, die in README-Dateien eingebettet werden können, um Ihre GitHub-Aktivität, Repository-Informationen, Programmiersprachen-Nutzung und mehr zu präsentieren. Es unterstützt auch die direkte Verwendung in GitHub Actions.

## Funktionen

- ✅ GitHub Stats Karten-Generierung
- ✅ Repository Pin Karten-Generierung
- ✅ Top Languages Karten-Generierung
- ✅ Gist Karten-Generierung
- ✅ WakaTime Karten-Generierung

## Beispiele

Hier sind einige Beispiele für das, was Sie mit diesem Projekt erstellen können:

### GitHub Stats Karte

**Basis:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname)
```
![GitHub Stats](.github/assets/stats-basic.svg)

**Dunkles Theme:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname&theme=dark)
```
![GitHub Stats Dark](.github/assets/stats-dark.svg)

**Kompaktes Layout:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname&layout=compact)
```
![GitHub Stats Compact](.github/assets/stats-compact.svg)

**Benutzerdefiniertes Theme:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname&bg_color=0d1117&title_color=ff6b6b&text_color=c9d1d9&border_color=30363d)
```
![GitHub Stats Custom](.github/assets/stats-custom.svg)

### Repository Pin Karte

**Basis:**
```markdown
![Pinned Repo](http://localhost:9000/api/pin?username=yourname&repo=repo-name)
```
![Pinned Repo](.github/assets/repo-basic.svg)

**Thematisiert:**
```markdown
![Pinned Repo](http://localhost:9000/api/pin?username=yourname&repo=repo-name&theme=dark)
```
![Pinned Repo Themed](.github/assets/repo-themed.svg)

### Top Languages Karte

**Basis:**
```markdown
![Top Languages](http://localhost:9000/api/top-langs?username=yourname)
```
![Top Languages](.github/assets/top-langs-basic.svg)

**Kompaktes Layout:**
```markdown
![Top Languages](http://localhost:9000/api/top-langs?username=yourname&layout=compact&langs_count=6)
```
![Top Languages Compact](.github/assets/top-langs-compact.svg)

### Gist Karte

```markdown
![Gist](http://localhost:9000/api/gist?id=gist_id)
```
![Gist](.github/assets/gist-basic.svg)

### WakaTime Karte

**Basis:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname)
```
![WakaTime Basic](.github/assets/wakatime-basic.svg)

**Kompaktes Layout:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&layout=compact)
```
![WakaTime Compact](.github/assets/wakatime-compact.svg)

**Test-Beispiele (Mit Testdaten generiert):**

**Basis-Test:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname)
```
![WakaTime Test Basic](.github/assets/wakatime-test-basic.svg)

**Kompakt-Test:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&layout=compact&langs_count=6&theme=dark)
```
![WakaTime Test Compact](.github/assets/wakatime-test-compact.svg)

**Thema:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&theme=radical&langs_count=5)
```
![WakaTime Test Themed](.github/assets/wakatime-test-themed.svg)

**Fortschrittsbalken ausblenden:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&hide_progress=true)
```
![WakaTime Test Hide Progress](.github/assets/wakatime-test-hide-progress.svg)

**Prozentanzeige:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&display_format=percent)
```
![WakaTime Test Percent](.github/assets/wakatime-test-percent.svg)

**Begrenzte Sprachen:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&langs_count=3&theme=tokyonight)
```
![WakaTime Test Limited](.github/assets/wakatime-test-limited.svg)

> **Hinweis:** WakaTime-Karten-Beispiele erfordern einen gültigen WakaTime-Benutzernamen mit öffentlichen Statistiken. Die obigen Testbilder wurden mit Testdaten generiert, um verschiedene Konfigurationsoptionen zu demonstrieren.

## Schnellstart

### Installation mit Go

```bash
go get github.com/soulteary/github-readme-stats
```

### Build aus Quellcode

```bash
git clone https://github.com/soulteary/github-readme-stats.git
cd github-readme-stats
go build -o github-readme-stats ./cmd/server
```

## Schnellstart

### 1. Installation

```bash
# Installation mit Go
go get github.com/soulteary/github-readme-stats

# Oder Build aus Quellcode
git clone https://github.com/soulteary/github-readme-stats.git
cd github-readme-stats
go build -o github-readme-stats ./cmd/server
```

### 2. Konfiguration

Erstellen Sie eine `.env` Datei:

```bash
# GitHub Personal Access Token (erforderlich)
PAT_1=your_github_token_here
# Sie können mehrere Tokens konfigurieren, um API-Rate-Limits zu erhöhen
PAT_2=your_second_token_here

# Server-Port (optional, Standard: 9000)
PORT=9000

# Cache-Dauer in Sekunden (optional, Standard: 86400)
CACHE_SECONDS=86400
```

### 3. Ausführung

```bash
# Methode 1: Direkt ausführen
go run cmd/server/main.go

# Methode 2: Verwenden Sie die erstellte Binärdatei
./github-readme-stats

# Methode 3: Mit Docker
docker build -t github-readme-stats .
docker run -p 9000:9000 -e PAT_1=your_token github-readme-stats
```

### 4. Verwendung

Sobald der Server läuft, können Sie die API unter `http://localhost:9000` aufrufen und die oben gezeigten Beispiele verwenden!

## GitHub Actions (Empfohlen)

Die empfohlene Verwendung dieses Projekts erfolgt über [GitHub Readme Stats Action](https://github.com/soulteary/github-readme-stats-action), die automatisch Statistik-Karten in Ihrem Repository mithilfe von GitHub Actions-Workflows generiert und aktualisiert.

**Repository:** [soulteary/github-readme-stats-action](https://github.com/soulteary/github-readme-stats-action)  
**Marketplace:** [Readme Stats Action](https://github.com/marketplace/actions/readme-stats-action)

### Schnellstart

Erstellen Sie eine Workflow-Datei (z.B. `.github/workflows/update-stats.yml`):

```yaml
name: Update README cards

on:
  schedule:
    - cron: "0 0 * * *" # Läuft täglich um Mitternacht
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

Dann die generierten Karten in Ihrem README einbetten:

```markdown
![Stats](./profile/stats.svg)
![Top Languages](./profile/top-langs.svg)
```

Weitere Details und Beispiele finden Sie in der [GitHub Readme Stats Action Dokumentation](https://github.com/soulteary/github-readme-stats-action).

## CLI-Verwendung

Du kannst die CLI verwenden, um Karten lokal zu generieren, ohne einen Server auszuführen. Dies ist nützlich für Tests, Entwicklung oder die Generierung statischer Bilder.

### CLI-Tools erstellen

```bash
# Alle CLI-Tools erstellen
make build

# Oder einzeln erstellen
make build-server    # Server erstellen
make build-examples  # Beispiel-Generator erstellen

# Schnelle Demo
./demo.sh
```

### Beispielbilder generieren

Der `examples` Befehl generiert Beispielbilder für alle Kartentypen. Wie andere Befehle liest er die Konfiguration aus der `.env` Datei.

```bash
# Alle Beispielbilder generieren (GitHub API Token in .env erforderlich)
./bin/examples

# Spezifische Beispiele generieren
./bin/examples stats-basic stats-dark repo-basic

# Test-Fehlerkarten generieren (kein Netzwerk erforderlich)
./bin/examples --test

# WakaTime-Testkarten mit Testdaten generieren (kein Netzwerk erforderlich)
./bin/examples --wakatime-test

# Hilfe anzeigen
./bin/examples --help
```

Dies erstellt SVG-Dateien im `.github/assets/` Verzeichnis, einschließlich:
- `stats-basic.svg` - Basis-Statistik-Karte
- `stats-dark.svg` - Dunkles Theme Statistik-Karte
- `stats-compact.svg` - Kompakte Statistik-Karte
- `repo-basic.svg` - Repository Pin Karte
- `top-langs-basic.svg` - Top Languages Karte
- `gist-basic.svg` - Gist Pin Karte
- `wakatime-basic.svg` - WakaTime Statistik-Karte
- `wakatime-test-*.svg` - WakaTime-Testkarten (mit Testdaten generiert)
- Und mehr...

### Einzelne Karten generieren

Verwende die Server-Binärdatei mit CLI-Parametern:

```bash
# Statistik-Karte generieren
./bin/server "/api?username=anuraghazra&theme=dark" --output=stats.svg

# Repository Pin generieren
./bin/server "/api/pin?username=anuraghazra&repo=github-readme-stats" --output=repo.svg

# Top Languages generieren
./bin/server "/api/top-langs?username=anuraghazra&layout=compact" --output=langs.svg

# Gist Karte generieren
./bin/server "/api/gist?id=bbfce31e0217a3689c8d961a356cb10d" --output=gist.svg

# WakaTime Karte generieren
./bin/server "/api/wakatime?username=anuraghazra" --output=wakatime.svg
```

### Erweiterte CLI-Verwendung

Du kannst auch Flag-Style-Argumente verwenden:

```bash
# Mit Flags verwenden
./bin/server --type=stats --username=anuraghazra --theme=dark --output=stats.svg

# Mehrere Parameter
./bin/server --type=top-langs --username=anuraghazra --layout=compact --langs_count=8 --theme=radical --output=langs.svg
```

## API-Endpunkte

### GitHub Stats Karte

Generiert eine GitHub-Statistik-Karte für einen Benutzer.

```
GET /api?username=USERNAME
```

**Parameter:**
- `username` (erforderlich) - GitHub-Benutzername
- `hide` - Bestimmte Statistiken ausblenden (kommagetrennt, z.B. `stars,commits`)
- `hide_title` - Titel ausblenden
- `hide_border` - Rahmen ausblenden
- `hide_rank` - Rang ausblenden
- `show_icons` - Symbole anzeigen
- `include_all_commits` - Alle Commits einbeziehen
- `theme` - Themenname (80+ Themen verfügbar)
- `bg_color` - Hintergrundfarbe (hexadezimal)
- `title_color` - Titelfarbe
- `text_color` - Textfarbe
- `icon_color` - Symbolfarbe
- `border_color` - Rahmenfarbe
- `border_radius` - Rahmenradius
- `locale` - Sprachcode (z.B. `zh`, `en`, `de`, `it`, `kr`, `ja`)
- `cache_seconds` - Cache-Dauer in Sekunden

> Siehe Beispiele oben für die Verwendung.

### Repository Pin Karte

Generiert eine Repository Pin-Karte.

```
GET /api/pin?username=USERNAME&repo=REPO
```

**Parameter:**
- `username` (erforderlich) - GitHub-Benutzername
- `repo` (erforderlich) - Repository-Name
- `theme` - Themenname
- `show_owner` - Besitzer anzeigen
- `locale` - Sprachcode

> Siehe Beispiele oben für die Verwendung.

### Top Languages Karte

Generiert eine Statistik-Karte für die am häufigsten verwendeten Programmiersprachen.

```
GET /api/top-langs?username=USERNAME
```

**Parameter:**
- `username` (erforderlich) - GitHub-Benutzername
- `hide` - Bestimmte Sprachen ausblenden (kommagetrennt)
- `layout` - Layout-Typ (`compact`, `normal`)
- `theme` - Themenname
- `locale` - Sprachcode

> Siehe Beispiele oben für die Verwendung.

### Gist Karte

Generiert eine Gist-Karte.

```
GET /api/gist?id=GIST_ID
```

**Parameter:**
- `id` (erforderlich) - Gist ID
- `theme` - Themenname
- `locale` - Sprachcode

> Siehe Beispiele oben für die Verwendung.

### WakaTime Karte

Generiert eine WakaTime-Programmierzeit-Statistik-Karte.

```
GET /api/wakatime?username=USERNAME
```

**Parameter:**
- `username` (erforderlich) - WakaTime-Benutzername
- `theme` - Themenname
- `hide` - Bestimmte Statistiken ausblenden
- `locale` - Sprachcode

> Siehe Beispiele oben für die Verwendung.

## Projektstruktur

```
.
├── cmd/
│   └── server/          # Server-Einstiegspunkt
│       └── main.go
├── internal/
│   ├── api/             # API-Handler
│   ├── cards/           # Karten-Rendering-Logik
│   ├── common/          # Gemeinsame Utilities und Hilfsfunktionen
│   ├── fetchers/        # Datenabrufer (GitHub API, WakaTime, etc.)
│   ├── themes/          # Themen-System
│   └── translations/    # Internationalisierungs-Übersetzungen
├── pkg/
│   └── svg/             # SVG-bezogene Utilities
├── Dockerfile           # Docker-Build-Datei
├── go.mod              # Go-Modul-Definition
└── README.md           # Projektdokumentation
```

## Entwicklungsstatus

✅ **Kernfunktionen abgeschlossen!**

Alle Hauptfunktionen wurden implementiert:

- ✅ Projekt-Basisstruktur
- ✅ HTTP-Server (Gin)
- ✅ GitHub API-Integration (GraphQL + REST)
- ✅ Wiederholungsmechanismus und Multi-PAT-Unterstützung
- ✅ Cache-Verwaltung
- ✅ Themen-System (80+ Themen)
- ✅ Alle Karten-Typ-Rendering (Stats, Pin, Top Languages, Gist, WakaTime)
- ✅ Rangberechnung
- ✅ Internationalisierungs-Unterstützung (Grundimplementierung)
- ✅ Alle API-Endpunkte

## Beitragen

Beiträge sind willkommen! Wenn Sie Ideen haben oder Probleme finden, bitte:

1. Forken Sie dieses Projekt
2. Erstellen Sie Ihren Feature-Branch (`git checkout -b feature/AmazingFeature`)
3. Committen Sie Ihre Änderungen (`git commit -m 'Add some AmazingFeature'`)
4. Pushen Sie zum Branch (`git push origin feature/AmazingFeature`)
5. Öffnen Sie einen Pull Request

Oder melden Sie Probleme direkt in [Issues](https://github.com/soulteary/github-readme-stats/issues).

## Lizenz

Dieses Projekt ist unter der MIT-Lizenz lizenziert. Siehe die [LICENSE](LICENSE) Datei für Details.
