create table dou_sheng.comment
(
	id int auto_increment
		primary key,
	user_id int null,
	video_id int null,
	content varchar(1024) null,
	create_date datetime null,
	constraint comment_ibfk_1
		foreign key (user_id) references dou_sheng.user (id),
	constraint comment_ibfk_2
		foreign key (video_id) references dou_sheng.video (id)
)
charset=utf8;

create index user_id
	on dou_sheng.comment (user_id);

create index video_id
	on dou_sheng.comment (video_id);

