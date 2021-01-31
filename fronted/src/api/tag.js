import request from "@/utils/request";

export function fetchTagList() {
  return request({
    url: "/tags",
    method: "get"
  });
}

export function fetchTagCount(params) {
  return request({
    url: "/count/tag",
    method: "get",
    params
  });
}

export function createTag(data) {
  return request({
    url: `/tags`,
    method: "post",
    data
  });
}

export function editTag(id, data) {
  return request({
    url: `/tags/${id}`,
    method: "put",
    data
  });
}

export function deleteTag(id) {
  return request({
    url: `/tags/${id}`,
    method: "delete"
  });
}
