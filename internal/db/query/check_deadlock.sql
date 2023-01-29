-- execute in TablePlus
SELECT
    a.application_name,
    l.relation :: regclass,
    l.transactionid,
    l.mode,
    l.locktype,
    l.GRANTED,
    a.usename,
    a.query,
    a.pid
FROM
    pg_stat_activity a
        JOIN pg_locks l ON l.pid = a.pid
WHERE
        a.application_name = 'psql'
ORDER BY
    a.pid;