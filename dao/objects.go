package dao

import "sync"

// User 用户
type User struct {
	Id              int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`           // 主键id
	Name            string `gorm:"column:name;NOT NULL" json:"name"`                         // 用户昵称
	Avatar          string `gorm:"column:avatar;NOT NULL" json:"avatar"`                     // 头像
	Signature       string `gorm:"column:signature;NOT NULL" json:"signature"`               // 头像
	BackgroundImage string `gorm:"column:background_image;NOT NULL" json:"background_image"` // 头像
	Hash            string `gorm:"column:hash;NOT NULL" json:"hash"`
	Salt            string `gorm:"column:salt;NOT NULL" json:"salt"`
	VideoCount      int    `gorm:"column:video_count;NOT NULL" json:"video_count"`
	FollowCount     int    `gorm:"column:follow_count;NOT NULL" json:"follow_count"`
	FollowerCount   int    `gorm:"column:follower_count;NOT NULL" json:"follower_count"`
	FavoriteCount   int    `gorm:"column:favorite_count;NOT NULL" json:"favorite_count"`
	IsFollow        bool   `gorm:"-" json:"is_follow"`
}
type userDao struct {
}

// Video 视频
type Video struct {
	Id            int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Title         string `gorm:"column:title" json:"title"`
	PlayUrl       string `gorm:"column:play_url" json:"play_url"`
	AuthorId      int    `gorm:"column:author_id" json:"-"`
	CoverUrl      string `gorm:"column:cover_url" json:"cover_url"`
	FavoriteCount int    `gorm:"column:favorite_count" json:"favorite_count"`
	CommentCount  int    `gorm:"column:comment_count" json:"comment_count"`
	CreateTime    int64  `gorm:"column:create_time" json:"create_time"`
	IsFavorite    bool   `gorm:"-"  json:"is_favorite"`
	User          *User  `gorm:"-" json:"author"`
}

type videoDao struct {
}

// Comment 评论
type Comment struct {
	Id         int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"` // 主键id
	VideoID    int    `gorm:"column:video_id" json:"-"`
	UserID     int    `gorm:"column:user_id" json:"-"`
	User       *User  `gorm:"-" json:"user"`
	Content    string `gorm:"column:content" json:"content"`
	CreateDate string `gorm:"-" json:"create_date"`
	CreateTime int64  `gorm:"column:create_time" json:"-"`
}

type CommentDao struct {
}

// Follower 关注关系
type Follower struct {
	ID         int `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"-"` // 主键id
	FollowID   int `gorm:"column:follow_id" json:"follow_id"`
	FollowerID int `gorm:"column:follower_id" json:"follower_id"`
}

type followerDao struct {
}

//当前环境
type nowAccount struct {
	NowUser map[int]*User
	Token   map[string]int
}

// Favorite 点赞关系
type Favorite struct {
	UserID  int `gorm:"column:user_id;primary_key;" json:"user_id"`
	VideoID int `gorm:"column:video_id;primary_key;" json:"video_id"`
}

type favoriteDao struct {
}

//单例
var (
	uDao        *userDao
	userOnce    sync.Once
	vDao        *videoDao
	videoOnce   sync.Once
	cDao        *CommentDao
	cDaoOnce    sync.Once
	account     *nowAccount
	accountOnce sync.Once
	fDao        *followerDao
	fDaoOnce    sync.Once
	faDao       *favoriteDao
	faDaoOnce   sync.Once
)
