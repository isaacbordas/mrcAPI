CREATE DATABASE IF NOT EXISTS myretrocollectiondb CHARACTER SET utf8;

CREATE USER myretrouser@'%' IDENTIFIED BY 'root';

GRANT ALL PRIVILEGES ON myretrocollectiondb.* TO myretrouser@'%';