import Vue from 'nativescript-vue';
import Vuex from 'vuex';
import firebase from 'nativescript-plugin-firebase'
Vue.use(Vuex);

const state = {
  isLoggedIn:null,
  shoppingList: [{
    name: "Klopapier",
    status: "open"     // requested, selfbought, requestbought, requestunavailable, selfunavailable
    }],
}

const getters = {
  isLoggedIn: state =>{
    return state.isLoggedIn
  },
  shoppingList: state =>{
    return state.shoppingList
  }
}
const mutations = {
  setIsLoggedIn: (state, value) => {
    state.isLoggedIn = value;
  },
  setShoppingList: (state, shoppingList) => {
    state.shoppingList = shoppingList
  },
  addItemToShoppingList: (state, item) => {
   // state.shoppingList = shoppingList
    state.shoppingList.push(item);
  },
  delItemFromShoppingList: (state, item) => {
    // state.shoppingList = shoppingList
    //const item = state.shoppingList.find(item => item.item_name === item.name)
    state.shoppingList.splice(state.shoppingList.indexOf(item), 1) 
    //item.pop();
   },
   markItemAsUnavailable: (state, item) => {
      item.status = "unavailable"
   },
   markItemAsSelfbought: (state, item) => {
      item.status = "selfbought"
 },
 markItemAsOpen: (state, item) => {
  item.status = "open"
}
}

// Asynchronous operations are actions
const actions = {
  // getShoppingList() {
  // Server POST
  // }
  addItem({commit}, item) {
    commit('addItemToShoppingList', item)
  }
}
const storeConf = {
  state,
  getters,
  mutations,
  actions
}
export default new Vuex.Store(storeConf)