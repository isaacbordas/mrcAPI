CREATE DATABASE IF NOT EXISTS myretrocollection CHARACTER SET utf8;

CREATE USER myretrousern@'%' IDENTIFIED BY 'root';

GRANT ALL PRIVILEGES ON myretrocollection.* TO myretrousern@'%';