const state = {
  anchors: []
}

const mutations = {
  UPDATE_ANCHORS: (state, payload) => {
    state.anchors = payload
  }
}

const actions = {
  updateAnchors({ commit }, payload) {
    commit('UPDATE_ANCHORS', payload)
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
