-- name: UserCredentials :many
SELECT
    c.id as "credential_id",
    c.title as "title",
    c.login_name as "login_name",
    c.secret as "secret",
    c.description as "description",
    ss.show_immediately as "show_immediately",
    ss.send_to_email as "send_to_email",
    ss.send_to_phone as "send_to_phone"
FROM credentials c
JOIN show_strategies ss ON ss.credential_id = c.id
WHERE c.customer_id = $1;