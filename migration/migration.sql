-- CREATE USER local_user WITH PASSWORD 'local_pwd';
--
-- ALTER USER local_user WITH SUPERUSER;
--
-- CREATE DATABASE my_wallet WITH OWNER = local_user TABLESPACE = pg_default TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'pt_BR.UTF-8' LC_CTYPE = 'pt_BR.UTF-8' CONNECTION LIMIT = -1;
--
-- CREATE SCHEMA my_wallet AUTHORIZATION local_user;
--
-- GRANT ALL ON SCHEMA my_wallet TO local_user WITH GRANT OPTION;

CREATE TABLE wallet (
    id UUID PRIMARY KEY,
    description VARCHAR(255) NOT NULL CHECK (LENGTH(description) >= 3)
);
