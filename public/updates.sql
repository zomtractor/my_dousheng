update videos
set favorite_count=(
    select count(*)
    from favorites f
    where videos.id=f.video_id
);
update videos
set comment_count=(
    select count(*)
    from comments
    where videos.id=comments.video_id
);
update users
set favorite_count=(
    select count(*)
    from favorites
    where users.id=favorites.user_id
);
update users
set follow_count=(
    select count(*)
    from followers
    where users.id=followers.follow_id
);
update users
set follower_count=(
    select count(*)
    from followers
    where users.id=followers.follower_id
);
update users
set video_count=(
    select count(*)
    from videos
    where users.id=videos.author_id
);