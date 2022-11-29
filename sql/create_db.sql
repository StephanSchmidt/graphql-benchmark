DROP DATABASE IF EXISTS graphjin;
CREATE DATABASE graphjin;

DROP OWNED BY graphjin;
DROP ROLE IF EXISTS graphjin;
CREATE ROLE graphjin WITH ENCRYPTED PASSWORD 'graphjin';

GRANT CONNECT ON DATABASE graphjin TO graphjin;
ALTER ROLE "graphjin" WITH LOGIN;
GRANT ALL PRIVILEGES ON DATABASE graphjin TO graphjin;
GRANT ALL ON schema public TO graphjin;

\c graphjin;

DROP TABLE IF EXISTS users CASCADE;

CREATE TABLE users
(
    id  BIGINT unique PRIMARY KEY,
    name TEXT
);

GRANT all ON users TO graphjin;

DROP TABLE IF EXISTS status CASCADE;
CREATE TABLE status
(
    id INT unique,
    status TEXT
);
GRANT all ON status TO graphjin;

DROP TABLE IF EXISTS statuses CASCADE;
CREATE TABLE statuses
(
    id INT unique,
    status TEXT
);
GRANT all ON statuses TO graphjin;

DROP TABLE IF EXISTS tasks CASCADE ;
CREATE TABLE tasks
(
    id  BIGINT PRIMARY KEY,
    title TEXT,
    user_id bigint REFERENCES users(id),
    status_id INT REFERENCES status(id)
);
GRANT all ON tasks TO graphjin;
CREATE INDEX user_index ON tasks (user_id);
