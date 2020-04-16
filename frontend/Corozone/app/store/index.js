import Vue from 'nativescript-vue';
import Vuex from 'vuex';
//import localStorage from 'nativescript-localstorage';
import { SecureStorage } from "nativescript-secure-storage";

import firebase from 'nativescript-plugin-firebase'
import { backendService } from '~/app';
Vue.use(Vuex);

const state = {
  isLoggedIn:null,
    shoppingList: [{
       name: "Klopapier",
       status: "open",     // requested, selfbought, requestbought, requestunavailable, selfunavailable
       uid: ""
      }]
}
//Load local storage after login
let secureStorage = new SecureStorage();
const VuexPersistent = store => {
  // Init hook.
console.log("Vues Persistent")
 // let storageStr = localStorage.getItem('ns-vuex-persistent');
 //TODO: Also encrypt data
  var storageStr = secureStorage.getSync({
    key: "vuex-persistent-secure", //+ backendService.token
  });
  if (storageStr) {
   store.replaceState(JSON.parse(storageStr))
  }
  store.subscribe((mutation, state) => {
   // Suscribe hook.
  // localStorage.setItem('ns-vuex-persistent', JSON.stringify(state));
   secureStorage.setSync({
    key: "vuex-persistent-secure", //+ backendService.token,
    value: JSON.stringify(state) 
  });
  })
 };

//  secureStorage.set({
//   key: "vuex-persistent-secure",
//   value: JSON.stringify(state)
// }).then(success => console.log("Successfully set a value? " + success));

// sync
// const success = secureStorage.setSync({
//   key: "vuex-persistent-secure",
//   value: JSON.stringify(state) 
// });

const getters = {
  isLoggedIn: state =>{
    return state.isLoggedIn
  },
  getShoppingList: state =>{
    return state.shoppingList.filter(item => item.uid === backendService.token)
    //return state.shoppingList
    //return shoppingList
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
    item.uid = backendService.token
    console.log(item)
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
  actions,
  plugins:[VuexPersistent]
}
export default new Vuex.Store(storeConf)