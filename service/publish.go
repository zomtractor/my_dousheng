package service

import (
	"my_dousheng/dao"
	"strconv"
	"time"
)

const Url_pf = "http://139.224.105.6/public/"

func ParseVedio(title string, uid int, ext string) *dao.Video {
	video := dao.NewVideoOnceInstance().AddVideoToSql(&dao.Video{
		Title:      title,
		AuthorId:   uid,
		CreateTime: time.Now().Unix(),
	})
	videoUrl := strconv.FormatInt(int64(video.Id), 10) + ext
	video.PlayUrl = Url_pf + videoUrl
	video.CoverUrl = Url_pf + "bk.jpg"
	err := dao.NewVideoOnceInstance().UpdateVideo(video)
	if err != nil {
		return nil
	}
	return video
}
func GetPublishList(uid int) []*dao.Video {
	return dao.NewVideoOnceInstance().GetVideosByUid(uid)
}
