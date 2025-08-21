package consts

const (
	TaskWait    = 0  //任务等待
	TaskSuccess = 1  // 任务完成
	TaskFail    = -1 // 任务失败
)

const (
	CrawlerRating = "Rating"
)

var RatingCountryList = []string{"tgstat.com", "tgstat.ru/en", "uk.tgstat.com/en", "by.tgstat.com/en", "uz.tgstat.com/en", "kaz.tgstat.com/en", "kg.tgstat.com/en", "ir.tgstat.com", "cn.tgstat.com", "in.tgstat.com", "et.tgstat.com"}

var PrivateDisclosure = []string{"", "/public", "/private"}

var Sort = []string{"members", "members_t", "members_y", "members_7d", "members_30d", "reach", "ci"}
