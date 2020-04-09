import Vue from 'nativescript-vue';
import Vuex from 'vuex';
import firebase from 'nativescript-plugin-firebase'
Vue.use(Vuex);

const state = {
  isLoggedIn:null
}

const getters = {
  isLoggedIn: state =>{
    return state.isLoggedIn
  }
}
const mutations = {
  setIsLoggedIn: (state, value) => {
    state.isLoggedIn = value;
  }
}
const actions = {
}
const storeConf = {
  state,
  getters,
  mutations,
  actions
}
export default new Vuex.Store(storeConf)