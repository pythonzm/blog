const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  articleHtml: state => state.app.articleHtml,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  nickname: state => state.user.nickname,
  introduction: state => state.user.introduction
}
export default getters
