# Go-Blog
A golang implemented Blog backend

> 请先创建名为blog_service的数据库，并执行如下的sql建表语句
# 创建博客标签表
CREATE TABLE `blog_tag` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`name` varchar(100) DEFAULT '' COMMENT '标签名称',
`created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
`created_by` varchar(255) DEFAULT '' COMMENT '创建人',
`modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
`modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
`deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
`is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除，0为未删除，1为删除',
`state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态0为禁用，1为启用',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签管理';

# 创建博客文章表
CREATE TABLE `blog_article` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`title` varchar(100) DEFAULT '' COMMENT '文章标题',
`desc` varchar(255) DEFAULT '' COMMENT '文章简述',
`cover_image_url` varchar(255) DEFAULT '' COMMENT '图片链接',
`content` longtext COMMENT '文章内容',
`created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
`created_by` varchar(255) DEFAULT '' COMMENT '创建人',
`modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
`modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
`deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
`is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除，0为未删除，1为删除',
`state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态0为禁用，1为启用',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';

# 文章标签关联表
CREATE TABLE `blog_artile_tag` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`article_id` int(11) NOT NULL COMMENT '文章标题',
`tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '标签ID',
`desc` varchar(255) DEFAULT '' COMMENT '文章ID',
`cover_image_url` varchar(255) DEFAULT '' COMMENT '图片链接',
`content` longtext COMMENT '文章内容',
`created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
`created_by` varchar(255) DEFAULT '' COMMENT '创建人',
`modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
`modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
`deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
`is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除，0为未删除，1为删除',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签关联';