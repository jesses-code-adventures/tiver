--- name: CreateGame :one
INSERT INTO game DEFAULT VALUES RETURNING *;

--- name: UpdateEndGame :one
UPDATE game SET ended_at = $2 WHERE id = $1;

--- name: IncrementRequests :exec 
UPDATE game
SET requests = requests + 1
WHERE id = $1;

--- name: DecrementRequests :exec 
UPDATE game
SET requests = requests + 1
WHERE id = $1;

--- name: CreateRequest :one
INSERT INTO request (id, created_at, game_id, colour, origin, speed, width, status) values (DEFAULT, DEFAULT, $1, $2, $3, $4, $5, $6) RETURNING *;

--- name: UpdateEndRequest :one
UPDATE request SET status = $2, ended_at = $3 WHERE id = $1;

--- name: UpdateRequestStatus :exec 
UPDATE request SET status = $2 WHERE id = $1;

--- name: GetActiveGames :many
SELECT * FROM game where ended_at = NULL order by created_at asc;

--- name: GetHangingRequests :manu 
SELECT * FROM game WHERE ended_at IS NOT NULL AND requests > 0;
