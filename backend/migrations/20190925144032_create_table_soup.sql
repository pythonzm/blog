-- +goose Up
-- +goose StatementBegin
CREATE TABLE `blog_soup` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `content` tinytext NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS blog_soup;
-- +goose StatementEnd
