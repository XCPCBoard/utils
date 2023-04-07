package keys

import "fmt"

// @Author: Feng
// @Date: 2022/7/23 20:02

//------------------------------------------//
// 基础
//------------------------------------------//

const (
	// OJ
	NowcoderKey   = "nowcoder"
	AtCoderKey    = "atCoder"
	LuoguKey      = "luogu"
	VJudgeKey     = "vjudge"
	CodeforcesKey = "codeforces"

	// 类型
	PassAmountKey    = "pass_amount"
	RatingKey        = "rating"
	MaxRatingKey     = "max_rating"
	RankingKey       = "ranking"
	MaxRankingKey    = "max_ranking"
	ContestAmountKey = "contest_amount"

	// 时间
	LastAll    = "all"
	Last1days  = "1"
	Last7days  = "7"
	Last14days = "14"
	Last28days = "28"
	Last30days = "30"
	LastMonth  = "month"

	// 难度
	DifficultyEasy     = "easy"
	DifficultyBasic    = "basic"
	DifficultyAdvanced = "advanced"
	DifficultyHard     = "hard"
	DifficultyUnknown  = "unknown"
)

// BuildKeyWithLastSiteKindDifficultyID 时效_网站_类型_难度_用户id
// 难度支持什么类型 取决于 爬虫处理并存了多少信息
func BuildKeyWithLastSiteKindDifficultyID(last, site, kind, difficulty, id string) string {
	return fmt.Sprintf("%v_%v_%v_%v_%v", last, site, kind, difficulty, id)
}

// BuildKeyWithLastSiteKindID 时效_网站_类型_用户id
// 时效支持几天 取决于 爬虫处理并存了多少信息
func BuildKeyWithLastSiteKindID(last, site, kind, id string) string {
	return BuildKeyWithLastSiteKindDifficultyID(last, site, kind, DifficultyUnknown, id)
}

// BuildKeyWithSiteKindID 网站_类型_用户id
func BuildKeyWithSiteKindID(site, kind, id string) string {
	return BuildKeyWithLastSiteKindID(LastAll, site, kind, id)
}

// BuildKeyWithSiteKind 网站_类型
func BuildKeyWithSiteKind(site, kind string) string {
	return fmt.Sprintf("%v_%v", site, kind)
}

// BuildKeyWithSiteID 网站_用户id
func BuildKeyWithSiteID(site, id string) string {
	return fmt.Sprintf("%v_%v", site, id)
}

//------------------------------------------//
// 牛客
//------------------------------------------//

func NowcoderKeyWithKindID(kind, id string) string {
	return BuildKeyWithSiteKindID(NowcoderKey, kind, id)
}

func NowcoderRatingKey(id string) string {
	return NowcoderKeyWithKindID(RatingKey, id)
}

func NowcoderRankingKey(id string) string {
	return NowcoderKeyWithKindID(RankingKey, id)
}

func NowcoderContestAmountKey(id string) string {
	return NowcoderKeyWithKindID(ContestAmountKey, id)
}

func NowcoderPassAmountKey(id string) string {
	return NowcoderKeyWithKindID(PassAmountKey, id)
}

//------------------------------------------//
// codeforces
//------------------------------------------//

func CodeforcesKeyWithKindID(kind, id string) string {
	return BuildKeyWithSiteKindID(CodeforcesKey, kind, id)
}

func CodeforcesRatingKey(id string) string {
	return CodeforcesKeyWithKindID(RatingKey, id)
}

func CodeforcesMaxRatingKey(id string) string {
	return CodeforcesKeyWithKindID(MaxRatingKey, id)
}

func CodeforcesRankingKey(id string) string {
	return CodeforcesKeyWithKindID(RankingKey, id)
}

func CodeforcesMaxRankingKey(id string) string {
	return CodeforcesKeyWithKindID(MaxRankingKey, id)
}

func CodeforcesPassAmountKey(id string) string {
	return CodeforcesKeyWithKindID(PassAmountKey, id)
}

func CodeforcesLastMonthPassAmountKey(id string) string {
	return BuildKeyWithLastSiteKindID(LastMonth, CodeforcesKey, PassAmountKey, id)
}

//------------------------------------------//
// AtCoder
//------------------------------------------//

func AtCoderKeyWithKindID(kind, id string) string {
	return BuildKeyWithSiteKindID(AtCoderKey, kind, id)
}

func AtCoderRatingKey(id string) string {
	return AtCoderKeyWithKindID(RatingKey, id)
}

func AtCoderRankingKey(id string) string {
	return AtCoderKeyWithKindID(RankingKey, id)
}

func AtCoderContestAmountKey(id string) string {
	return AtCoderKeyWithKindID(ContestAmountKey, id)
}

func AtCoderPassAmountKey(id string) string {
	return AtCoderKeyWithKindID(PassAmountKey, id)
}

//------------------------------------------//
// 洛谷
//------------------------------------------//

func LuoguKeyWithKindDifficultyID(kind, difficulty, id string) string {
	return BuildKeyWithLastSiteKindDifficultyID(LastAll, LuoguKey, kind, difficulty, id)
}

func LuoguKeyWithKindID(kind, id string) string {
	return BuildKeyWithSiteKindID(LuoguKey, kind, id)
}

func LuoguEasyPassAmountKey(id string) string {
	return LuoguKeyWithKindDifficultyID(PassAmountKey, DifficultyEasy, id)
}

func LuoguBasicPassAmountKey(id string) string {
	return LuoguKeyWithKindDifficultyID(PassAmountKey, DifficultyBasic, id)
}

func LuoguAdvancedPassAmountKey(id string) string {
	return LuoguKeyWithKindDifficultyID(PassAmountKey, DifficultyAdvanced, id)
}

func LuoguHardPassAmountKey(id string) string {
	return LuoguKeyWithKindDifficultyID(PassAmountKey, DifficultyHard, id)
}

func LuoguUnknownPassAmountKey(id string) string {
	return LuoguKeyWithKindDifficultyID(PassAmountKey, DifficultyUnknown, id)
}

func LuoguRatingKey(id string) string {
	return LuoguKeyWithKindID(RatingKey, id)
}

func LuoguRankingKey(id string) string {
	return LuoguKeyWithKindID(RankingKey, id)
}

func LuoguPassAmountKey(id string) string {
	return LuoguKeyWithKindID(PassAmountKey, id)
}

//------------------------------------------//
// vjudge
//------------------------------------------//

func VJudgeKeyWithKindID(kind, id string) string {
	return BuildKeyWithSiteKindID(VJudgeKey, kind, id)
}

func VJudgeKeyWithLastKindID(last, kind, id string) string {
	return BuildKeyWithLastSiteKindID(last, VJudgeKey, kind, id)
}

func VJudgeLast1dayPassAmountKey(id string) string {
	return VJudgeKeyWithLastKindID(Last1days, PassAmountKey, id)
}

func VJudgeLast7dayPassAmountKey(id string) string {
	return VJudgeKeyWithLastKindID(Last7days, PassAmountKey, id)
}

func VJudgeLast30dayPassAmountKey(id string) string {
	return VJudgeKeyWithLastKindID(Last30days, PassAmountKey, id)
}

func VJudgePassAmountKey(id string) string {
	return VJudgeKeyWithKindID(PassAmountKey, id)
}
