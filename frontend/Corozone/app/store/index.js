import Vue from 'nativescript-vue';
import Vuex from 'vuex';
//import localStorage from 'nativescript-localstorage';
import { SecureStorage } from "nativescript-secure-storage";

//import firebase from 'nativescript-plugin-firebase'
import { backendService } from '~/app';
Vue.use(Vuex);

const state = {
  isLoggedIn:null,
    shoppingList: [
      // {
      //  name: "Klopapier",
      //  status: "open",     // requested, selfbought, requestbought, requestunavailable, selfunavailable
      //  uid: "",
      //  crypto: "none",
      // }
    ],
    notification_message: "Keine Nachrichten",
    friendListIDs: [
     // {id: "D4YX72AVKrZp4cHONP23llO5omk2", name:"Sebastian"}
    ]
};
//Load local storage after login
//---------------------
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
    console.log("Files loaded from starage")
    console.log(storageStr)
    state.notification_message = storageStr.friendListIDs
   store.replaceState(JSON.parse(storageStr))
   
  //  this.setIsLoggedIn(null)
    console.log(state)
  }

  store.subscribe((mutation, state) => {
   // Suscribe hook.
  // localStorage.setItem('ns-vuex-persistent', JSON.stringify(state));
   console.log("Save files to storage")
   console.log(state)
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
//----------------
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
  },
  getFriendListIDs: state=> {
    return state.friendListIDs
  }
}
const mutations = {
  setIsLoggedIn: (state, value) => {
    state.isLoggedIn = value;
  },
  setShoppingList: (state, shoppingList) => {
    state.shoppingList = shoppingList
  },
  mergeShoppingList: (state,friendsShoppingList) => {

    // Merge two arrays without duplicates
    // https://stackoverflow.com/questions/1584370/how-to-merge-two-arrays-in-javascript-and-de-duplicate-items
    const mergeDedupe = (arr) => {
      return [...new Set([].concat(...arr))];
    }
    //let newShoppingList = mergeDedupe(state.shoppingList, shoppingList) 
    //let newShoppingList =  _.unionBy(shoppingList, state.shoppingList, 'name'); 

//https://stackoverflow.com/questions/22844560/check-if-object-value-exists-within-a-javascript-array-of-objects-and-if-not-add

for(let item of friendsShoppingList) {
    // console.log(item)
       let b = state.shoppingList.find(o => o.name ===item.name)
       //TODO: Change to version with .some instead of undefined
       // https://stackoverflow.com/questions/22844560/check-if-object-value-exists-within-a-javascript-array-of-obje#
       if(b!=undefined){
      console.log(b.name)
       state.shoppingList[state.shoppingList.findIndex(o => o.name ===item.name)].status = item.status
       }
    // console.log(i)
  }
  

    const newShoppingList = state.shoppingList.concat(friendsShoppingList).filter(function(o) {  
      console.log(o)
      return this.has(o.name) ? false : this.add(o.name);
    }, new Set());
    console.log("New Shopping List")
    console.log(newShoppingList)
    state.shoppingList = newShoppingList
    //state.shoppingList = state.shoppingList.concat(shoppingList)
  },
  addItemToShoppingList: (state, item) => {
   // state.shoppingList = shoppingList
    //item.uid = backendService.token
    console.log(item)
    state.shoppingList.push(item);
  },
  addFriendID: (state, friend) => {
     console.log("Added Friend: " + friend.name)
     state.friendListIDs.push(friend);
   },
  delFriendID: (state, id) => {
    console.log("Deleted Friend: " + id)
    // probably need to find here
    state.friendListIDs.splice(state.friendListIDs.indexOf(id), 1);
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
markItemAsHelperBought: (state, item) => {
  item.status = "helperbought"
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