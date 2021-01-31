import request from "@/utils/request";

export function fetchVisitorCount(params) {
  return request({
    url: "/count/visitor",
    method: "get",
    params
  });
}

export function fetchCountByDate(params) {
  return request({
    url: "/count/date",
    method: "get",
    params
  });
}

export function fetchCountByUa(params) {
  return request({
    url: "/count/ua",
    method: "get",
    params
  });
}
