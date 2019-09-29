import request from "@/utils/request";

export function login(data) {
  return request({
    url: "/user/login",
    method: "post",
    data
  });
}

export function getInfo() {
  return request({
    url: "/user/info",
    method: "get"
  });
}

export function logout() {
  return request({
    url: "/user/logout",
    method: "post"
  });
}

export function editUser(data, params) {
  return request({
    url: "/user/edit",
    method: "patch",
    data,
    params
  });
}

export function getAbout() {
  return request({
    url: "/user/about",
    method: "get"
  });
}
