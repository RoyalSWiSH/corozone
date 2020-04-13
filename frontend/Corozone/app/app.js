import Vue from "nativescript-vue";
// store needs to imported before VueDevtools
import store from './store'
import VueDevtools from 'nativescript-vue-devtools'
import axios from 'axios'
import VueAxios from 'vue-axios'
import App from './components/App'
import LoginPage from './components/LoginPage'
import Items from './components/Items'
import ItemDetails from './components/ItemDetails'
import firebase from "nativescript-plugin-firebase"
import BackendService from './services/BackendService'
import AuthService from './services/AuthService'
import { isAndroid, isIOS } from 'tns-core-modules/platform';

import { localize } from "nativescript-localize";

// require('globals')
// require('nativescript-i18n')

// Enable global check if device is android or iOS to disable Maps on Android since mapbox crashes right now
Vue.prototype.$isAndroid = isAndroid;
Vue.prototype.$isIOS = isIOS;

Vue.use(VueAxios, axios)
axios.defaults.baseURL = "http://corozone.sebastian-roy.de/api/v1";
axios.defaults.headers.post['Content-Type'] = 'application/json';


Vue.registerElement(
  "CardView", 
  () => require("@nstudio/nativescript-cardview").CardView
)
Vue.registerElement(
  'CheckBox',
  () => require('@nstudio/nativescript-checkbox').CheckBox,
)
Vue.registerElement("Mapbox", () => require("nativescript-mapbox").MapboxView);

export const backendService = new BackendService()
export const authService = new AuthService()

Vue.prototype.$authService = authService
Vue.prototype.$backendService = backendService
// What does this import? Everything? Just index.js?

//import App from "./components/App";
if(TNS_ENV !== 'production') {
  Vue.use(VueDevtools)
}
// Prints Vue logs when --env.production is *NOT* set while building
Vue.config.silent = (TNS_ENV === 'production')

Vue.filter("L", localize);

firebase
  .init({
    onAuthStateChanged: data => { 
      console.log((data.loggedIn ? "Logged in to firebase" : "Logged out from firebase") + " (firebase.init() onAuthStateChanged callback)");
      if (data.loggedIn) {
        backendService.token = data.user.uid
        console.log("uID: " + data.user.uid)
        store.commit('setIsLoggedIn', true)
      }
      else {     
        store.commit('setIsLoggedIn', false) 
      }
    }
  })
  .then(
    function(instance) {
      console.log("firebase.init done");
    },
    function(error) {
      console.log("firebase.init error: " + error);
    }
  );


// Used for loading indicator
  global.loaderOptions = {
    android: {
        margin: 100,
        dimBackground: true,
        color: "#4B9ED6", 
        hideBezel: true, 
        mode: 3 
    },
    ios: {
        dimBackground: true,
        color: "#FFFFFF", 
        hideBezel: true, 
        mode: 3 
    }
};
//   new Vue({ 
//     render: h => h('frame', [h(LoginPage)])
//  }).$start();
new Vue({
    store,
    template: `
    <Frame>
      <LoginPage />
    </Frame>`,
    components: {
     LoginPage
     // App
     //Items
    }
}).$start();