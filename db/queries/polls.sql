-- name: CreatePolls :exec
INSERT INTO polls (poll_by,poll_question,options_count,options,created_at)
VALUES (?,?,?,?,?);

-- name: GetPolls :many
SELECT * FROM polls 
ORDER BY created_at DESC;