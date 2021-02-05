import request from "@/utils/request";

export function fetchList(id) {
  return request({
    url: `/comments/${id}`,
    method: "get"
  });
}

export function createComment(data) {
  return request({
    url: `/comments`,
    method: "post",
    data
  });
}

export function fetchAll(params) {
  return request({
    url: `/comments`,
    method: "get",
    params
  });
}

export function fetchRecentComments(params) {
  return request({
    url: `/comment/recent`,
    method: "get",
    params
  });
}

export function deleteComment(id) {
  return request({
    url: `/comments/${id}`,
    method: "delete"
  });
}
