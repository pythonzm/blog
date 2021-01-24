import request from "@/utils/request";

export function fetchVisitorCount(params) {
  return request({
    url: "/count/visitor",
    method: "get",
    params
  });
}
