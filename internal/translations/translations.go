package translations

import (
	"strings"
)

// I18n represents internationalization instance
type I18n struct {
	locale       string
	translations map[string]map[string]string
}

// NewI18n creates a new I18n instance
func NewI18n(locale string) *I18n {
	if locale == "" {
		locale = "en"
	}
	locale = strings.ToLower(locale)

	translations := getTranslations()

	return &I18n{
		locale:       locale,
		translations: translations,
	}
}

// T translates a key
func (i *I18n) T(key string) string {
	if trans, ok := i.translations[key]; ok {
		if val, ok := trans[i.locale]; ok {
			return val
		}
		// Fallback to English
		if val, ok := trans["en"]; ok {
			return val
		}
	}
	return key
}

// IsLocaleAvailable checks if a locale is available
func IsLocaleAvailable(locale string) bool {
	availableLocales := []string{
		"ar", "az", "bn", "bg", "my", "ca", "cn", "zh", "zh-tw", "cs", "nl", "en",
		"fil", "fi", "fr", "de", "el", "he", "hi", "hu", "id", "it", "ja",
		"kr", "ml", "np", "no", "fa", "pl", "pt-br", "pt-pt", "ro", "ru",
		"sa", "sr", "sr-latn", "sk", "es", "sw", "se", "ta", "th", "tr",
		"uk-ua", "ur", "uz", "vi",
	}

	locale = strings.ToLower(locale)
	for _, loc := range availableLocales {
		if loc == locale {
			return true
		}
	}
	return false
}

// getTranslations returns all translations
func getTranslations() map[string]map[string]string {
	return map[string]map[string]string{
		"statcard.totalstars": {
			"en":    "Total Stars",
			"cn":    "获得 Stars",
			"zh":    "获得 Stars",
			"zh-tw": "獲得 Stars",
			"de":    "Sterne insgesamt",
			"it":    "Stelle totali",
			"kr":    "총 스타",
			"ja":    "スター合計",
		},
		"statcard.commits": {
			"en":    "Commits",
			"cn":    "提交数",
			"zh":    "提交数",
			"zh-tw": "提交數",
			"de":    "Commits",
			"it":    "Commit",
			"kr":    "커밋",
			"ja":    "コミット",
		},
		"statcard.prs": {
			"en":    "PRs",
			"cn":    "拉取请求",
			"zh":    "拉取请求",
			"zh-tw": "拉取請求",
			"de":    "PRs",
			"it":    "PR",
			"kr":    "PR",
			"ja":    "PR",
		},
		"statcard.prs-merged": {
			"en":    "Merged PRs",
			"cn":    "已合并 PR",
			"zh":    "已合并 PR",
			"zh-tw": "已合併 PR",
			"de":    "Zusammengeführte PRs",
			"it":    "PR uniti",
			"kr":    "병합된 PR",
			"ja":    "マージ済み PR",
		},
		"statcard.prs-merged-percentage": {
			"en":    "Merged PRs",
			"cn":    "已合并 PR",
			"zh":    "已合并 PR",
			"zh-tw": "已合併 PR",
			"de":    "Zusammengeführte PRs",
			"it":    "PR uniti",
			"kr":    "병합된 PR",
			"ja":    "マージ済み PR",
		},
		"statcard.reviews": {
			"en":    "Reviews",
			"cn":    "代码审查",
			"zh":    "代码审查",
			"zh-tw": "程式碼審查",
			"de":    "Reviews",
			"it":    "Revisioni",
			"kr":    "리뷰",
			"ja":    "レビュー",
		},
		"statcard.issues": {
			"en":    "Issues",
			"cn":    "问题",
			"zh":    "问题",
			"zh-tw": "問題",
			"de":    "Issues",
			"it":    "Issue",
			"kr":    "이슈",
			"ja":    "イシュー",
		},
		"statcard.discussions-started": {
			"en":    "Discussions Started",
			"cn":    "发起的讨论",
			"zh":    "发起的讨论",
			"zh-tw": "發起的討論",
			"de":    "Diskussionen gestartet",
			"it":    "Discussioni avviate",
			"kr":    "시작한 토론",
			"ja":    "開始したディスカッション",
		},
		"statcard.discussions-answered": {
			"en":    "Discussions Answered",
			"cn":    "回答的讨论",
			"zh":    "回答的讨论",
			"zh-tw": "回答的討論",
			"de":    "Diskussionen beantwortet",
			"it":    "Discussioni risposte",
			"kr":    "답변한 토론",
			"ja":    "回答したディスカッション",
		},
		"statcard.contribs": {
			"en":    "Contributed to",
			"cn":    "贡献于",
			"zh":    "贡献于",
			"zh-tw": "貢獻於",
			"de":    "Beigetragen zu",
			"it":    "Contribuito a",
			"kr":    "기여한 저장소",
			"ja":    "コントリビュート",
		},
		"statcard.title": {
			"en":    "GitHub Stats",
			"cn":    "GitHub 统计",
			"zh":    "GitHub 统计",
			"zh-tw": "GitHub 統計",
			"de":    "GitHub Statistiken",
			"it":    "Statistiche GitHub",
			"kr":    "GitHub 통계",
			"ja":    "GitHub 統計",
		},
		"statcard.ranktitle": {
			"en":    "GitHub Rank",
			"cn":    "GitHub 排名",
			"zh":    "GitHub 排名",
			"zh-tw": "GitHub 排名",
			"de":    "GitHub Rang",
			"it":    "Classifica GitHub",
			"kr":    "GitHub 순위",
			"ja":    "GitHub ランク",
		},
		"wakatimecard.lastyear": {
			"en":    "this year",
			"cn":    "今年",
			"zh":    "今年",
			"zh-tw": "今年",
			"de":    "dieses Jahr",
			"it":    "quest'anno",
			"kr":    "올해",
			"ja":    "今年",
		},
		"wakatimecard.last7days": {
			"en":    "last 7 days",
			"cn":    "最近7天",
			"zh":    "最近7天",
			"zh-tw": "最近7天",
			"de":    "letzten 7 Tage",
			"it":    "ultimi 7 giorni",
			"kr":    "최근 7일",
			"ja":    "過去7日間",
		},
		"wakatimecard.nocodingactivity": {
			"en":    "No coding activity this week",
			"cn":    "本周没有编程活动",
			"zh":    "本周没有编程活动",
			"zh-tw": "本週沒有編程活動",
			"de":    "Keine Aktivitäten in dieser Woche",
			"it":    "Nessuna attività in questa settimana",
			"kr":    "이번 주 코딩 활동 없음",
			"ja":    "今週のコーディング活動はありません",
		},
		"wakatimecard.nocodedetails": {
			"en":    "No coding activity this week",
			"cn":    "本周没有编程活动",
			"zh":    "本周没有编程活动",
			"zh-tw": "本週沒有編程活動",
			"de":    "Keine Aktivitäten in dieser Woche",
			"it":    "Nessuna attività in questa settimana",
			"kr":    "이번 주 코딩 활동 없음",
			"ja":    "今週のコーディング活動はありません",
		},
		"wakatimecard.notpublic": {
			"en":    "User doesn't share publicly detailed code statistics",
			"cn":    "用户未公开分享详细的代码统计",
			"zh":    "用户未公开分享详细的代码统计",
			"zh-tw": "用戶未公開分享詳細的程式碼統計",
			"de":    "Benutzer teilt keine detaillierten Codestatistiken öffentlich",
			"it":    "L'utente non condivide pubblicamente statistiche dettagliate del codice",
			"kr":    "사용자가 공개적으로 상세한 코드 통계를 공유하지 않음",
			"ja":    "ユーザーは公開的に詳細なコード統計を共有していません",
		},
		"repocard.archived": {
			"en":    "Archived",
			"cn":    "已归档",
			"zh":    "已归档",
			"zh-tw": "已歸檔",
			"de":    "Archiviert",
			"it":    "Archiviato",
			"kr":    "보관됨",
			"ja":    "アーカイブ済み",
		},
		"repocard.template": {
			"en":    "Template",
			"cn":    "模板",
			"zh":    "模板",
			"zh-tw": "模板",
			"de":    "Vorlage",
			"it":    "Modello",
			"kr":    "템플릿",
			"ja":    "テンプレート",
		},
		"langcard.title": {
			"en":    "Most Used Languages",
			"cn":    "最常用语言",
			"zh":    "最常用语言",
			"zh-tw": "最常用語言",
			"de":    "Meistverwendete Sprachen",
			"it":    "Linguaggi più usati",
			"kr":    "가장 많이 사용한 언어",
			"ja":    "最も使用された言語",
		},
		"wakatimecard.title": {
			"en":    "WakaTime Stats",
			"cn":    "WakaTime 统计",
			"zh":    "WakaTime 统计",
			"zh-tw": "WakaTime 統計",
			"de":    "WakaTime Statistiken",
			"it":    "Statistiche WakaTime",
			"kr":    "WakaTime 통계",
			"ja":    "WakaTime 統計",
		},
	}
}
