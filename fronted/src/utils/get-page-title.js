import defaultSettings from "@/settings";

export default function getPageTitle(pageTitle, toPath) {
  var title;
  if (toPath.indexOf("admin") !== -1 || toPath.indexOf("login") !== -1) {
    title = defaultSettings.adminTitle;
  } else {
    title = defaultSettings.blogTitle;
  }
  if (pageTitle) {
    return `${pageTitle} - ${title}`;
  }
  return `${title}`;
}
