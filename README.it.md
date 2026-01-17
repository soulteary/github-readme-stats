# GitHub Readme Stats (Implementazione Go)

[![GitHub](https://img.shields.io/badge/GitHub-soulteary%2Fgithub--readme--stats-blue)](https://github.com/soulteary/github-readme-stats)

![GitHub Readme Stats](.github/assets/banner.jpg)

## Languages / 语言 / Sprachen / Lingue / 언어 / 言語

- [English](README.md)
- [简体中文](README.zh.md)
- [Deutsch](README.de.md)
- [Italiano](README.it.md)
- [한국어](README.kr.md)
- [日本語](README.ja.md)

Questa è un'implementazione in Go del progetto [GitHub Readme Stats (anuraghazra)](https://github.com/anuraghazra/github-readme-stats). Fornisce carte statistiche GitHub dinamiche che possono essere incorporate nei file README per mostrare la tua attività GitHub, informazioni sui repository, utilizzo dei linguaggi di programmazione e altro ancora. Supporta anche l'uso diretto in GitHub Actions.

## Funzionalità

- ✅ Generazione carte GitHub Stats
- ✅ Generazione carte Repository Pin
- ✅ Generazione carte Top Languages
- ✅ Generazione carte Gist
- ✅ Generazione carte WakaTime

## Esempi

Ecco alcuni esempi di ciò che puoi creare con questo progetto:

### Carta GitHub Stats

**Base:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname)
```
![GitHub Stats](.github/assets/stats-basic.svg)

**Tema Scuro:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname&theme=dark)
```
![GitHub Stats Dark](.github/assets/stats-dark.svg)

**Layout Compatto:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname&layout=compact)
```
![GitHub Stats Compact](.github/assets/stats-compact.svg)

**Tema Personalizzato:**
```markdown
![GitHub Stats](http://localhost:9000/api?username=yourname&bg_color=0d1117&title_color=ff6b6b&text_color=c9d1d9&border_color=30363d)
```
![GitHub Stats Custom](.github/assets/stats-custom.svg)

### Carta Repository Pin

**Base:**
```markdown
![Pinned Repo](http://localhost:9000/api/pin?username=yourname&repo=repo-name)
```
![Pinned Repo](.github/assets/repo-basic.svg)

**Tematico:**
```markdown
![Pinned Repo](http://localhost:9000/api/pin?username=yourname&repo=repo-name&theme=dark)
```
![Pinned Repo Themed](.github/assets/repo-themed.svg)

### Carta Top Languages

**Base:**
```markdown
![Top Languages](http://localhost:9000/api/top-langs?username=yourname)
```
![Top Languages](.github/assets/top-langs-basic.svg)

**Layout Compatto:**
```markdown
![Top Languages](http://localhost:9000/api/top-langs?username=yourname&layout=compact&langs_count=6)
```
![Top Languages Compact](.github/assets/top-langs-compact.svg)

### Carta Gist

```markdown
![Gist](http://localhost:9000/api/gist?id=gist_id)
```
![Gist](.github/assets/gist-basic.svg)

### Carta WakaTime

**Base:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname)
```
![WakaTime Basic](.github/assets/wakatime-basic.svg)

**Layout Compatto:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&layout=compact)
```
![WakaTime Compact](.github/assets/wakatime-compact.svg)

**Esempi di Test (Generati con dati di test):**

**Test Base:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname)
```
![WakaTime Test Basic](.github/assets/wakatime-test-basic.svg)

**Test Compatto:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&layout=compact&langs_count=6&theme=dark)
```
![WakaTime Test Compact](.github/assets/wakatime-test-compact.svg)

**Tema:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&theme=radical&langs_count=5)
```
![WakaTime Test Themed](.github/assets/wakatime-test-themed.svg)

**Nascondi Barra di Progresso:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&hide_progress=true)
```
![WakaTime Test Hide Progress](.github/assets/wakatime-test-hide-progress.svg)

**Visualizzazione Percentuale:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&display_format=percent)
```
![WakaTime Test Percent](.github/assets/wakatime-test-percent.svg)

**Linguaggi Limitati:**
```markdown
![WakaTime](http://localhost:9000/api/wakatime?username=yourname&langs_count=3&theme=tokyonight)
```
![WakaTime Test Limited](.github/assets/wakatime-test-limited.svg)

> **Nota:** Gli esempi delle carte WakaTime richiedono un nome utente WakaTime valido con statistiche pubbliche. Le immagini di test sopra sono generate utilizzando dati di test per dimostrare varie opzioni di configurazione.

## Guida Rapida

### Installazione con Go

```bash
go get github.com/soulteary/github-readme-stats
```

### Build dal sorgente

```bash
git clone https://github.com/soulteary/github-readme-stats.git
cd github-readme-stats
go build -o github-readme-stats ./cmd/server
```

## Guida Rapida

### 1. Installazione

```bash
# Installazione con Go
go get github.com/soulteary/github-readme-stats

# O build dal sorgente
git clone https://github.com/soulteary/github-readme-stats.git
cd github-readme-stats
go build -o github-readme-stats ./cmd/server
```

### 2. Configurazione

Crea un file `.env`:

```bash
# GitHub Personal Access Token (richiesto)
PAT_1=your_github_token_here
# Puoi configurare più token per aumentare i limiti dell'API
PAT_2=your_second_token_here

# Porta del server (opzionale, predefinito: 9000)
PORT=9000

# Durata della cache in secondi (opzionale, predefinito: 86400)
CACHE_SECONDS=86400
```

### 3. Esecuzione

```bash
# Metodo 1: Esegui direttamente
go run cmd/server/main.go

# Metodo 2: Usa il binario compilato
./github-readme-stats

# Metodo 3: Usando Docker
docker build -t github-readme-stats .
docker run -p 9000:9000 -e PAT_1=your_token github-readme-stats
```

### 4. Utilizzo

Una volta che il server è in esecuzione, puoi accedere all'API su `http://localhost:9000` e utilizzare gli esempi mostrati sopra!

## GitHub Actions (Consigliato)

Il modo consigliato per utilizzare questo progetto è tramite [GitHub Readme Stats Action](https://github.com/soulteary/github-readme-stats-action), che genera e aggiorna automaticamente le carte statistiche nel tuo repository utilizzando i workflow di GitHub Actions.

**Repository:** [soulteary/github-readme-stats-action](https://github.com/soulteary/github-readme-stats-action)  
**Marketplace:** [Readme Stats Action](https://github.com/marketplace/actions/readme-stats-action)

### Guida Rapida

Crea un file workflow (ad esempio `.github/workflows/update-stats.yml`):

```yaml
name: Update README cards

on:
  schedule:
    - cron: "0 0 * * *" # Esegue una volta al giorno a mezzanotte
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

Quindi incorpora le carte generate nel tuo README:

```markdown
![Stats](./profile/stats.svg)
![Top Languages](./profile/top-langs.svg)
```

Per maggiori dettagli ed esempi, consulta la [documentazione di GitHub Readme Stats Action](https://github.com/soulteary/github-readme-stats-action).

## Utilizzo CLI

Puoi utilizzare la CLI per generare carte localmente senza eseguire un server. Questo è utile per test, sviluppo o generazione di immagini statiche.

### Costruire gli Strumenti CLI

```bash
# Costruisci tutti gli strumenti CLI
make build

# O costruisci individualmente
make build-server    # Costruisci il server
make build-examples  # Costruisci il generatore di esempi

# Demo rapida
./demo.sh
```

### Generare Immagini di Esempio

Il comando `examples` genera immagini di esempio per tutti i tipi di carte. Come altri comandi, legge la configurazione dal file `.env`.

```bash
# Genera tutte le immagini di esempio (richiede token GitHub API in .env)
./bin/examples

# Genera esempi specifici
./bin/examples stats-basic stats-dark repo-basic

# Genera carte di errore di test (nessuna rete richiesta)
./bin/examples --test

# Genera carte di test WakaTime utilizzando dati di test (nessuna rete richiesta)
./bin/examples --wakatime-test

# Visualizza aiuto
./bin/examples --help
```

Questo crea file SVG nella directory `.github/assets/`, inclusi:
- `stats-basic.svg` - Carta statistica di base
- `stats-dark.svg` - Carta statistica tema scuro
- `stats-compact.svg` - Carta statistica compatta
- `repo-basic.svg` - Carta pin repository
- `top-langs-basic.svg` - Carta top languages
- `gist-basic.svg` - Carta pin Gist
- `wakatime-basic.svg` - Carta statistiche WakaTime
- `wakatime-test-*.svg` - Carte di test WakaTime (generate con dati di test)
- E molto altro...

### Generare Carte Singole

Utilizza il binario del server con parametri CLI:

```bash
# Genera carta statistiche
./bin/server "/api?username=anuraghazra&theme=dark" --output=stats.svg

# Genera pin repository
./bin/server "/api/pin?username=anuraghazra&repo=github-readme-stats" --output=repo.svg

# Genera top languages
./bin/server "/api/top-langs?username=anuraghazra&layout=compact" --output=langs.svg

# Genera carta Gist
./bin/server "/api/gist?id=bbfce31e0217a3689c8d961a356cb10d" --output=gist.svg

# Genera carta WakaTime
./bin/server "/api/wakatime?username=anuraghazra" --output=wakatime.svg
```

### Utilizzo CLI Avanzato

Puoi anche utilizzare argomenti in stile flag:

```bash
# Utilizzo con flag
./bin/server --type=stats --username=anuraghazra --theme=dark --output=stats.svg

# Parametri multipli
./bin/server --type=top-langs --username=anuraghazra --layout=compact --langs_count=8 --theme=radical --output=langs.svg
```

## Endpoint API

### Carta GitHub Stats

Genera una carta statistica GitHub per un utente.

```
GET /api?username=USERNAME
```

**Parametri:**
- `username` (richiesto) - Nome utente GitHub
- `hide` - Nascondi statistiche specifiche (separate da virgola, es. `stars,commits`)
- `hide_title` - Nascondi titolo
- `hide_border` - Nascondi bordo
- `hide_rank` - Nascondi classifica
- `show_icons` - Mostra icone
- `include_all_commits` - Includi tutti i commit
- `theme` - Nome del tema (80+ temi disponibili)
- `bg_color` - Colore di sfondo (esadecimale)
- `title_color` - Colore del titolo
- `text_color` - Colore del testo
- `icon_color` - Colore dell'icona
- `border_color` - Colore del bordo
- `border_radius` - Raggio del bordo
- `locale` - Codice lingua (es. `zh`, `en`, `de`, `it`, `kr`, `ja`)
- `cache_seconds` - Durata della cache in secondi

> Vedi esempi sopra per l'utilizzo.

### Carta Repository Pin

Genera una carta repository pin.

```
GET /api/pin?username=USERNAME&repo=REPO
```

**Parametri:**
- `username` (richiesto) - Nome utente GitHub
- `repo` (richiesto) - Nome del repository
- `theme` - Nome del tema
- `show_owner` - Mostra proprietario
- `locale` - Codice lingua

> Vedi esempi sopra per l'utilizzo.

### Carta Top Languages

Genera una carta statistica dei linguaggi di programmazione più utilizzati.

```
GET /api/top-langs?username=USERNAME
```

**Parametri:**
- `username` (richiesto) - Nome utente GitHub
- `hide` - Nascondi linguaggi specifici (separati da virgola)
- `layout` - Tipo di layout (`compact`, `normal`)
- `theme` - Nome del tema
- `locale` - Codice lingua

> Vedi esempi sopra per l'utilizzo.

### Carta Gist

Genera una carta Gist.

```
GET /api/gist?id=GIST_ID
```

**Parametri:**
- `id` (richiesto) - ID Gist
- `theme` - Nome del tema
- `locale` - Codice lingua

> Vedi esempi sopra per l'utilizzo.

### Carta WakaTime

Genera una carta statistica del tempo di programmazione WakaTime.

```
GET /api/wakatime?username=USERNAME
```

**Parametri:**
- `username` (richiesto) - Nome utente WakaTime
- `theme` - Nome del tema
- `hide` - Nascondi statistiche specifiche
- `locale` - Codice lingua

> Vedi esempi sopra per l'utilizzo.

## Struttura del Progetto

```
.
├── cmd/
│   └── server/          # Punto di ingresso del server
│       └── main.go
├── internal/
│   ├── api/             # Gestori API
│   ├── cards/           # Logica di rendering delle carte
│   ├── common/          # Utility comuni e funzioni helper
│   ├── fetchers/        # Recuperatori di dati (GitHub API, WakaTime, ecc.)
│   ├── themes/          # Sistema di temi
│   └── translations/    # Traduzioni di internazionalizzazione
├── pkg/
│   └── svg/             # Utility relative a SVG
├── Dockerfile           # File di build Docker
├── go.mod              # Definizione modulo Go
└── README.md           # Documentazione del progetto
```

## Stato di Sviluppo

✅ **Funzionalità principali completate!**

Tutte le funzionalità principali sono state implementate:

- ✅ Struttura base del progetto
- ✅ Server HTTP (Gin)
- ✅ Integrazione GitHub API (GraphQL + REST)
- ✅ Meccanismo di retry e supporto multi-PAT
- ✅ Gestione cache
- ✅ Sistema di temi (80+ temi)
- ✅ Rendering di tutti i tipi di carte (Stats, Pin, Top Languages, Gist, WakaTime)
- ✅ Calcolo classifica
- ✅ Supporto internazionalizzazione (implementazione base)
- ✅ Tutti gli endpoint API

## Contribuire

I contributi sono benvenuti! Se hai idee o trovi problemi, per favore:

1. Fai un fork di questo progetto
2. Crea il tuo branch delle funzionalità (`git checkout -b feature/AmazingFeature`)
3. Committa le tue modifiche (`git commit -m 'Add some AmazingFeature'`)
4. Pusha al branch (`git push origin feature/AmazingFeature`)
5. Apri una Pull Request

Oppure segnala problemi direttamente in [Issues](https://github.com/soulteary/github-readme-stats/issues).

## Licenza

Questo progetto è concesso in licenza sotto la licenza MIT. Vedi il file [LICENSE](LICENSE) per i dettagli.
