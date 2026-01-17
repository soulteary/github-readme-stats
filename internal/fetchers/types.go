package fetchers

// StatsData represents the stats data structure
type StatsData struct {
	Name                     string
	TotalPRs                 int
	TotalPRsMerged           int
	MergedPRsPercentage      float64
	TotalReviews             int
	TotalCommits             int
	TotalIssues              int
	TotalStars               int
	TotalDiscussionsStarted  int
	TotalDiscussionsAnswered int
	ContributedTo            int
	Rank                     RankData
}

// RankData represents rank information
type RankData struct {
	Level      string
	Percentile float64
}

// RepoData represents repository data
type RepoData struct {
	Name            string
	NameWithOwner   string
	Description     string
	PrimaryLanguage *LanguageData
	Stargazers      StargazersData
	Forks           ForksData
	IsPrivate       bool
	IsFork          bool
	IsArchived      bool
	IsTemplate      bool
}

// LanguageData represents language information
type LanguageData struct {
	Name  string
	Color string
}

// StargazersData represents stargazers count
type StargazersData struct {
	TotalCount int
}

// ForksData represents forks count
type ForksData struct {
	TotalCount int
}

// TopLanguagesData represents top languages data
type TopLanguagesData struct {
	Languages []LanguageStats
}

// LanguageStats represents language statistics
type LanguageStats struct {
	Name      string
	Size      int64
	Color     string
	RepoCount int
}

// GistData represents gist data
type GistData struct {
	Name        string
	Description string
	Files       []GistFile
	Owner       GistOwner
	Stargazers  StargazersData
	Forks       ForksData
	IsPublic    bool
}

// GistFile represents a gist file
type GistFile struct {
	Name     string
	Language string
	Size     int
}

// GistOwner represents gist owner
type GistOwner struct {
	Login string
}

// WakaTimeData represents WakaTime statistics
type WakaTimeData struct {
	Languages               []WakaTimeLanguage
	TotalTime               string
	IsCodingActivityVisible bool
	IsOtherUsageVisible     bool
	Range                   string
}

// WakaTimeLanguage represents WakaTime language statistics
type WakaTimeLanguage struct {
	Name    string
	Percent float64
	Time    string
	Hours   int
	Minutes int
}
