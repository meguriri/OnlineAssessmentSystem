create table `user`(
    id varchar(20) not null PRIMARY KEY,
    name varchar(20) default NULL COMMENT "用户姓名",
    type int default 1 COMMENT "类型默认为学生 0-管理员 1-学生 2-老师",
    password varchar(16) not null
);

create table `knowledge_point`(
    id int NOT NULL auto_increment PRIMARY KEY,
    name varchar(20) default null,
    facility_value int not null default 1 COMMENT "难度 1-2-3-4-5 5个档位"
    
);