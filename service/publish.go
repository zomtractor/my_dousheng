package service

import (
	"my_dousheng/dao"
	"strconv"
	"time"
)

const Url_pf = "http://139.224.105.6/public/"

// ParseVedio 生成视频和封面
func ParseVedio(title string, uid int, ext string) *dao.Video {
	video := dao.NewVideoOnceInstance().AddVideoToSql(&dao.Video{
		Title:      title,
		AuthorId:   uid,
		CreateTime: time.Now().Unix(),
	})
	videoUrl := strconv.FormatInt(int64(video.Id), 10) + ext
	coverUrl := strconv.FormatInt(int64(video.Id), 10) + ".jpg"
	video.PlayUrl = Url_pf + videoUrl
	video.CoverUrl = Url_pf + coverUrl
	err := dao.NewVideoOnceInstance().UpdateVideo(video)
	if err != nil {
		return nil
	}
	return video
}

// GetPublishList 获取发布列表
func GetPublishList(uid int) []*dao.Video {
	return dao.NewVideoOnceInstance().GetVideosByUid(uid)
}
