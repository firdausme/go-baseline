create
database go-baseline;

use
database go-baseline;

create table user
(
    id       varchar(100) not null primary key,
    username varchar(100) null,
    password varchar(100) null
);

INSERT INTO user (id, username, password)
VALUES ('4EC9CCC5-A2D2-4140-9EBA-3F525BDFC17C', 'user-test',
        '$2a$10$RayBMrTHdmqL6S6HakIWseXv1uayVtuIVmi/4XJymEzPWdINuU8q.');
