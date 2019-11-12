import Cookies from "js-cookie";

const TokenKey = "blog_token";
const IsAuthorKey = "is_author";
const inEightHours = new Date(new Date().getTime() + 8 * 60 * 60 * 1000); // cookie超时时间，8小时

export function getToken() {
  return Cookies.get(TokenKey);
}

export function setToken(token) {
  return Cookies.set(TokenKey, token, { expires: inEightHours });
}

export function removeToken() {
  return Cookies.remove(TokenKey);
}

export function getAuthor() {
  return Cookies.get(IsAuthorKey);
}

export function setAuthor(is_author) {
  return Cookies.set(IsAuthorKey, is_author, { expires: inEightHours });
}

export function removeAuthor() {
  return Cookies.remove(IsAuthorKey);
}
