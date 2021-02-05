import request from "@/utils/request";

export function fetchList(params) {
  return request({
    url: "/articles",
    method: "get",
    params
  });
}

export function fetchArticle(id, params) {
  return request({
    url: `/articles/${id}`,
    method: "get",
    params
  });
}

export function fetchArticleCount(params) {
  return request({
    url: "/count/article",
    method: "get",
    params
  });
}

export function fetchArticleCountByCategory(params) {
  return request({
    url: "/count/article_by_category",
    method: "get",
    params
  });
}

export function fetchArticleCountByTag(params) {
  return request({
    url: "/count/article_by_tag",
    method: "get",
    params
  });
}

export function fetchArticleRank(params) {
  return request({
    url: "/article/rank",
    method: "get",
    params
  });
}

export function createArticle(data) {
  return request({
    url: `/articles`,
    method: "post",
    data
  });
}

export function editArticle(id, data) {
  return request({
    url: `/articles/${id}`,
    method: "put",
    data
  });
}

export function deleteArticle(id) {
  return request({
    url: `/articles/${id}`,
    method: "delete"
  });
}

export function uploadImage(data) {
  return request({
    url: `/upload`,
    method: "post",
    headers: { "Content-Type": "multipart/form-data" },
    data
  });
}
