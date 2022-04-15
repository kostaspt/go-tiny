DROP TABLE IF EXISTS dummy;

CREATE TABLE dummy
(
    id         BIGSERIAL PRIMARY KEY NOT NULL,
    data       JSON                  NOT NULL,
    created_at TIMESTAMP without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

DROP TRIGGER IF EXISTS dummy_updated_at_update ON dummy;

CREATE TRIGGER dummy_updated_at_update
    BEFORE UPDATE
    ON dummy
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();