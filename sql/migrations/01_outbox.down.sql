SELECT remove_retention_policy('outbox');
DROP TABLE IF EXISTS outbox;
DROP TYPE outbox_status;
DROP INDEX outbox_created_dttm_not_completed_idx;
