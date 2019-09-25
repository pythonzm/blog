import request from "@/utils/request";

export function fetchSoupList(params) {
  return request({
    url: "/soups",
    method: "get",
    params
  });
}

export function createSoup(data) {
  return request({
    url: `/soups`,
    method: "post",
    data
  });
}

export function editSoup(id, data) {
  return request({
    url: `/soups/${id}`,
    method: "put",
    data
  });
}

export function deleteSoup(id) {
  return request({
    url: `/soups/${id}`,
    method: "delete"
  });
}

export function fetchRandSoup() {
  return request({
    url: "/soup/random",
    method: "get"
  });
}
