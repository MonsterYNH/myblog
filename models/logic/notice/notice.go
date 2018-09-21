package notice

import (
	"myblog/models/db"
	"myblog/models/model"
)

func GetNotice(num int) []model.Notice {
	noticeEntry := db.GetObjectEntryByTypeName(db.MONGO_COLL_NOTICE)
	noticeEntry.LRange(db.REDIS_NOTICE_KEY, 0, num)
	notices := make([]model.Notice, 0)
	if objects, ok := noticeEntry.Result.(*[]*model.Notice); ok {
		for _, notice := range *objects {
			notices = append(notices, *notice)
		}
		return notices
	}
	return nil
}
