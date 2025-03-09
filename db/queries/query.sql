-- name: CreateAd :exec
insert into ad (id, title, description, status, genre, target_audiences, visual_elements,call_to_action,
                duration, priority, created_at)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);

-- name: UpdateAdStatus :exec
update ad
set status = $1
where id = $2;

-- name: UpdateAdAnalysis :exec
update ad
set status = $1, analysis = $2, completed_at = $3
where id = $4;

-- name: UpdateAdRetry :exec
update ad
set retried_at = $1, retry_time = $2
where id = $3;

-- name: GetAdByID :one
select *
from ad
where id = $1;