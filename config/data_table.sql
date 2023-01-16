create table `user`(
    id varchar(20) not null PRIMARY KEY,
    name varchar(20) default NULL COMMENT "用户姓名",
    type int default 1 COMMENT "类型默认为学生 0-管理员 1-学生 2-老师",
    password varchar(16) not null
);