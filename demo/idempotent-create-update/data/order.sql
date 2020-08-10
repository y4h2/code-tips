-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  id bigint NOT NULL,
  user_id bigint(11) NOT NULL,
  status varchar(255) DEFAULT "init",
  version int(11) DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;
