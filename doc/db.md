# 数据库

## 记录
```
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户表id',
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `password` varchar(50) NOT NULL COMMENT '用户密码，MD5加密',
  `email` varchar(50) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `question` varchar(100) DEFAULT NULL COMMENT '找回密码问题',
  `answer` varchar(100) DEFAULT NULL COMMENT '找回密码答案',
  `role` int(4) NOT NULL COMMENT '角色0-管理员,1-普通用户',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '最后一次更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name_unique` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;
```

**ENGINE=InnoDB不是默认就是这个引擎吗?**
——是的，如果不写也是ok,就会走默认的，在这里写上是因为可以很清楚的看到这个建表语句用了哪些，而且在创建表的时候，写上也是一个很好的习惯

**AUTO_INCREMENT=22,它不是自增的吗？为什么还要设数字?**
——这个是自增的，在这里设置数字的意思是想要让这条语句在增长的时候，从22开始自增。

**utf8不是已经在my.ini里设置过了?**
——这个虽然在my.ini设置过了，但设置的是mysql的的语言编码，而这里创建的时候不设置，就会出现乱码问题，二者的作用域是不一样的，在创建表单的时候，这个charset会作用到这个表上，他代表mysql简历数据库数据表时设定字符集为utf-8

