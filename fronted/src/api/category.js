import request from "@/utils/request";

export function fetchCategoryList() {
  return request({
    url: "/categories",
    method: "get"
  });
}

export function fetchCategoryCount(params) {
  return request({
    url: "/count/category",
    method: "get",
    params
  });
}

export function createCategory(data) {
  return request({
    url: `/categories`,
    method: "post",
    data
  });
}

export function editCategory(id, data) {
  return request({
    url: `/categories/${id}`,
    method: "put",
    data
  });
}

export function deleteCategory(id) {
  return request({
    url: `/categories/${id}`,
    method: "delete"
  });
}
