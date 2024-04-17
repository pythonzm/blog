const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  nickname: state => state.user.nickname,
  introduction: state => state.user.introduction,
  anchors: state => state.anchors.anchors
}
export default getters
