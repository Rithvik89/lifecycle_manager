-- name: LikeTrendingUsers :many
SELECT posts.user_id,COUNT(likes.post_id) as like_count from likes
INNER JOIN posts ON posts.post_id = likes.post_id
WHERE likes.created_at >= DATE_SUB(NOW(),INTERVAL ? MINUTE) 
GROUP BY posts.user_id;

-- name: DeleteTrendingUsers :exec
DELETE from trending_users
where created_at < DATE_SUB(NOW(),INTERVAL ? MINUTE);

-- name: CommentTrendingUsers :many
SELECT user_id,COUNT(post_id) as comment_count from parent_comments 
WHERE created_at >= DATE_SUB(NOW(),INTERVAL ? MINUTE) 
GROUP BY user_id; 

-- name: InsertTrendingUsers :exec
INSERT INTO trending_users (user_id) VALUES (?);
