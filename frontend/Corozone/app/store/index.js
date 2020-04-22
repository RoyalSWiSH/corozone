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
      }],
    notification_message: "Keine Nachrichten"
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

 secureStorage.set({
  key: "vuex-persistent-secure",
  value: JSON.stringify(state)
}).then(success => console.log("Successfully set a value? " + success));

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
  },
  getNotificationMessage: state=> {
    return state.notification_message
  }
}
const mutations = {
  setIsLoggedIn: (state, value) => {
    state.isLoggedIn = value;
  },
  setShoppingList: (state, shoppingList) => {
    state.shoppingList = shoppingList
  },
  mergeShoppingList: (state,shoppingList) => {
    // Merge two arrays without duplicates
    // https://stackoverflow.com/questions/1584370/how-to-merge-two-arrays-in-javascript-and-de-duplicate-items
    const mergeDedupe = (arr) => {
      return [...new Set([].concat(...arr))];
    }
    //let newShoppingList = mergeDedupe(state.shoppingList, shoppingList) 
    //let newShoppingList =  _.unionBy(shoppingList, state.shoppingList, 'name'); 
    const newShoppingList = shoppingList.concat(state.shoppingList).filter(function(o) {  
      return this.has(o.name) ? false : this.add(o.name);
    }, new Set());
    console.log("New Shopping LIst")
    console.log(newShoppingList)
    state.shoppingList = newShoppingList
    //state.shoppingList = state.shoppingList.concat(shoppingList)
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
},
setNotificationMessage: (state, msg) => {
  state.notification_message = msg
},
setAcceptedItems: (state, items) => {
  //TODO Just change some items, dont override
 state.shoppingList = JSON.parse(items)
 console.log(items)
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