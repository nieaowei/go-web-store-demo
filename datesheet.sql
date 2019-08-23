CREATE TABLE `tb_user`(
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(50) NOT NULL COMMENT '用户名',
	`password` VARCHAR(32) NOT NULL COMMENT '密码',
	`phone` VARCHAR(20) DEFAULT NULL COMMENT '手机号',
	`email` VARCHAR(50) DEFAULT NULL COMMENT '邮箱',
	`created` datetime NOT NULL,
	`updated` datetime NOT NULL,
	PRIMARY	KEY (`id`),
	UNIQUE	KEY	`username`	(`username`) USING BTREE,
	UNIQUE	KEY	`phone`	(`phone`) USING	BTREE,
	UNIQUE	KEY	`email`	(`email`) USING	BTREE
)ENGINE=INNODB	AUTO_INCREMENT=37 DEFAULT	CHARSET=utf8	COMMENT='用户表';

CREATE TABLE `tb_item`	(
	`id` BIGINT(20)	NOT NULL	COMMENT	'商品id',
	`title`	VARCHAR(100)	NOT	NULL	COMMENT	'商品标题',
	`sell_point`	VARCHAR(500)	DEFAULT	NULL	COMMENT	'商品卖点',
	`price`	BIGINT(20)	NOT	NULL	COMMENT	'商品价格',
	`num`	BIGINT(20)	NOT	NULL	COMMENT	'商品数量',
	`barcode`	VARCHAR(30)	DEFAULT	NULL	COMMENT	'条形码',
	`image`	VARCHAR(500)	DEFAULT	NULL	COMMENT	'商品图片',
	`cid`	BIGINT(10)	NOT	NULL	COMMENT	'所属类目',
	`status`	TINYINT(4)	NOT	NULL	DEFAULT '1'	COMMENT	'商品状态',
	`created`	datetime	NOT	NULL	COMMENT	'创建时间',
	`updated`	datetime	NOT	NULL	COMMENT	'更新时间',
	
	PRIMARY	KEY	(`id`),
	KEY	`status` (`status`),
	KEY	`cid` (`cid`),
	KEY	`updated` (`updated`)
)ENGINE=INNODB	DEFAULT	CHARSET=utf8	COMMENT='商品表';

CREATE TABLE `tb_item_cat`(
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '类目id',
	`parent_id`	BIGINT(20)	DEFAULT	NULL	COMMENT	'父类id',
	`name` VARCHAR(50) DEFAULT NULL COMMENT '类目名称',
	`status`	INT(1)	DEFAULT '1' COMMENT '状态',
	`sort_order`	INT(4)	DEFAULT	NULL	COMMENT	'排序序号',
	`is_parent`	TINYINT(1)	DEFAULT	'1'	COMMENT	'改类目是否为父类',
	`created`	datetime	DEFAULT	NULL	COMMENT	'创建时间',
	`updated`	datetime	DEFAULT	NULL	COMMENT	'更新时间',
	PRIMARY	KEY (`id`),
	KEY	`parent_id`	(`parent_id`,`status`)	USING	BTREE,
	KEY	`sort_order`	(`sort_order`)
	
)ENGINE=INNODB	AUTO_INCREMENT=1183	DEFAULT	CHARSET=utf8	COMMENT='商品分类表';

CREATE TABLE `tb_item_desc`	(
	`item_id` BIGINT(20)	NOT NULL	COMMENT	'商品id',
	`item_desc`	text	COMMENT	'商品描述',
	`created`	datetime	DEFAULT	NULL	COMMENT	'创建时间',
	`updated`	datetime	DEFAULT	NULL	COMMENT	'更新时间',
	PRIMARY	KEY	(`item_id`)
)ENGINE=INNODB	DEFAULT	CHARSET=utf8	COMMENT='商品描述表';

CREATE TABLE `tb_item_param`	(
	`id` BIGINT(20)	NOT NULL	COMMENT	'商品id',
	`item_cat_id` BIGINT(20)	DEFAULT NULL	COMMENT	'商品类别id',
	`param_data`	text	COMMENT	'参数格式，json',
	`created`	datetime	DEFAULT	NULL	COMMENT	'创建时间',
	`updated`	datetime	DEFAULT	NULL	COMMENT	'更新时间',
	PRIMARY	KEY	(`id`),
	KEY	`item_cat_id`	(`item_cat_id`)
)ENGINE=INNODB	AUTO_INCREMENT=27	DEFAULT	CHARSET=utf8	COMMENT='商品规格参数表';

CREATE TABLE `tb_item_param_item`	(
	`id` BIGINT(20)	NOT NULL	AUTO_INCREMENT COMMENT	' ',
	`item_id` BIGINT(20)	DEFAULT NULL	COMMENT	'商品id',
	`param_data`	text	COMMENT	'参数格式，json',
	`created`	datetime	DEFAULT	NULL	COMMENT	'创建时间',
	`updated`	datetime	DEFAULT	NULL	COMMENT	'更新时间',
	PRIMARY	KEY	(`id`),
	KEY	`item_id`	(`item_id`) USING	BTREE
)ENGINE=INNODB	AUTO_INCREMENT=9	DEFAULT	CHARSET=utf8	COMMENT='商品规格参数表';

CREATE TABLE `tb_content`	(
	`id` BIGINT(20)	NOT NULL	AUTO_INCREMENT COMMENT	' ',
	`category_id` BIGINT(20)	NOT NULL	COMMENT	'内容类目id',
	`title`	VARCHAR(200)	DEFAULT	NULL	COMMENT	'内容标题',
	`sub_title`	VARCHAR(100)	DEFAULT	NULL	COMMENT	'子标题',
	`title_desc`	VARCHAR(500)	DEFAULT	NULL	COMMENT	'标题描述',
	`url`	VARCHAR(500)	DEFAULT	NULL	COMMENT	'连接',
	`pic`	VARCHAR(300)	DEFAULT	NULL	COMMENT	'图片聚堆路径',
	`pic2`	VARCHAR(300)	DEFAULT	NULL	COMMENT	'图片2',
	`content`	text	COMMENT	'内容',
	`created`	datetime	DEFAULT	NULL	COMMENT	'创建时间',
	`updated`	datetime	DEFAULT	NULL	COMMENT	'更新时间',
	PRIMARY	KEY	(`id`),
	KEY	`category_id`	(`category_id`),
	KEY	`updated`	(`updated`)
)ENGINE=INNODB	AUTO_INCREMENT=32	DEFAULT	CHARSET=utf8	COMMENT='内容表';

CREATE TABLE `tb_content_category`	(
	`id` BIGINT(20)	NOT NULL	AUTO_INCREMENT COMMENT	' ',
	`` BIGINT(20)	DEFAULT NULL	COMMENT	'父类id',
	`name`	VARCHAR(50)	DEFAULT	NULL	COMMENT	'分类名称',
	`status`	int(1) DEFAULT '1' COMMENT '状态',
	`sort_order` int(4) DEFAULT NULL	COMMENT '排序序号',
	`is_parent`	TINYINT(4)	DEFAULT	'1' COMMENT '是否为父类目',
	`created`	datetime	DEFAULT	NULL	COMMENT	'创建时间',
	`updated`	datetime	DEFAULT	NULL	COMMENT	'更新时间',
	PRIMARY	KEY	(`id`),
	KEY	`parent_id`	(`parent_id`,`status`) USING BTREE,
	KEY	`sort_order`	(`sort_order`)
)ENGINE=INNODB	AUTO_INCREMENT=98	DEFAULT	CHARSET=utf8	COMMENT='内容分类';
