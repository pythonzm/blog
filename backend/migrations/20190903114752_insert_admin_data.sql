-- +goose Up
-- +goose StatementBegin
INSERT INTO blog_user
(`username`, `password`, `avatar`, `introduction`, `nickname`, `about`) VALUES
('admin', 'e10adc3949ba59abbe56e057f20f883e', 'https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=2884107401,3797902000&fm=26&gp=0.jpg', 'a gopher', 'gopher', '## hello world');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM blog_user WHERE username='admin';
-- +goose StatementEnd
