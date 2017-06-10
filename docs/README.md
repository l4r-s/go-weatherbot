# sqllite databse in golang

## Links
- https://astaxie.gitbooks.io/build-web-application-with-golang/en/05.3.html


## init Database table
~~~
CREATE TABLE `data` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `devid` VARCHAR(64) NULL,
    `temp` DOUBLE NULL,
    `hum` DOUBLE NULL,
    `timestamp` DOUBLE NULL
);
~~~

## demo data
~~~
INSERT INTO data (
 id,
 devid,
 temp,
 hum,
 timestamp
 )
VALUES
 (
 ?,
 "dev03",
 28.1,
 33.4,
 1257894440);
~~~

## remove table
~~~
DROP TABLE data;
~~~
