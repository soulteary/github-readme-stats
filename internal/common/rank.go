package common

import "math"

// CalculateRank calculates the user's rank based on various statistics
type RankParams struct {
	AllCommits bool
	Commits    float64
	PRs        float64
	Issues     float64
	Reviews    float64
	Repos      float64 // unused but kept for compatibility
	Stars      float64
	Followers  float64
}

type Rank struct {
	Level      string
	Percentile float64
}

// ExponentialCDF calculates the exponential cumulative distribution function
func ExponentialCDF(x float64) float64 {
	return 1 - math.Pow(2, -x)
}

// LogNormalCDF calculates the log normal cumulative distribution function (approximation)
func LogNormalCDF(x float64) float64 {
	return x / (1 + x)
}

// CalculateRank calculates the user's rank
func CalculateRank(params RankParams) Rank {
	commitsMedian := 250.0
	commitsWeight := 2.0
	if params.AllCommits {
		commitsMedian = 1000.0
	}

	prsMedian := 50.0
	prsWeight := 3.0

	issuesMedian := 25.0
	issuesWeight := 1.0

	reviewsMedian := 2.0
	reviewsWeight := 1.0

	starsMedian := 50.0
	starsWeight := 4.0

	followersMedian := 10.0
	followersWeight := 1.0

	totalWeight := commitsWeight + prsWeight + issuesWeight + reviewsWeight + starsWeight + followersWeight

	thresholds := []float64{1, 12.5, 25, 37.5, 50, 62.5, 75, 87.5, 100}
	levels := []string{"S", "A+", "A", "A-", "B+", "B", "B-", "C+", "C"}

	rank := 1 - (commitsWeight*ExponentialCDF(params.Commits/commitsMedian)+
		prsWeight*ExponentialCDF(params.PRs/prsMedian)+
		issuesWeight*ExponentialCDF(params.Issues/issuesMedian)+
		reviewsWeight*ExponentialCDF(params.Reviews/reviewsMedian)+
		starsWeight*LogNormalCDF(params.Stars/starsMedian)+
		followersWeight*LogNormalCDF(params.Followers/followersMedian))/totalWeight

	percentile := rank * 100

	// Find the appropriate level
	level := "C" // default
	for i, threshold := range thresholds {
		if percentile <= threshold {
			level = levels[i]
			break
		}
	}

	return Rank{
		Level:      level,
		Percentile: percentile,
	}
}
