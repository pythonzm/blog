-- +goose Up
-- +goose StatementBegin
INSERT INTO blog_soup (`content`) VALUES ('你全心全力做到的最好\r\n可能还不如别人的随便搞搞');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE * FROM blog_soup;
-- +goose StatementEnd
