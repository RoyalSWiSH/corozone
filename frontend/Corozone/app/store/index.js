import Vue from 'nativescript-vue';
import Vuex from 'vuex';
import firebase from 'nativescript-plugin-firebase'
Vue.use(Vuex);

const state = {
  isLoggedIn:null,
  shoppingList: [{name: "Klopapier"}],
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