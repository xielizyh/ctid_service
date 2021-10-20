# 开发

## 数据库

1. 创建数据库

   ```mysql
   CREATE DATABASE
   IF
   	NOT EXISTS ctid_service DEFAULT CHARACTER 
   	SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;
   ```

2. 添加用户

   ```mysql
   GRANT ALL PRIVILEGES ON ctid_service.* TO ctid@localhost IDENTIFIED BY '12345678';
   ```

3. 创建订单表

   ```mysql
   CREATE TABLE `ctid_order` (
     `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
     `username` varchar(100) DEFAULT '' COMMENT '姓名',
     `phone` int(11) unsigned DEFAULT '0' COMMENT '手机',
     `cert_type` varchar(100) DEFAULT '身份证' COMMENT '证件类型',
     `cert_number` int(18) unsigned DEFAULT '0' COMMENT '证件号码',
     `room_number` int(4) unsigned DEFAULT '0' COMMENT '入住房号',
     `checkin_time` int(10) unsigned DEFAULT '0' COMMENT '入住时间',
     `checkout_time` int(10) unsigned DEFAULT '0' COMMENT '退房时间',
     `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
     `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
     `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
     `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
     `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
     `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
     `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0 为禁用、1 为启用',
     PRIMARY KEY (`id`)
   ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单管理';
   ```

   

# 总结

1. swag安装后找不到？

   需要将“GOPATH“如"/home/xieli/go/bin"加入到PATH路径上

   