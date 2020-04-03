import Vue from "nativescript-vue";
import VueDevtools from 'nativescript-vue-devtools'
import axios from 'axios'
import VueAxios from 'vue-axios'
import App from './components/App'
import LoginPage from './components/LoginPage'
import Items from './components/Items'
import ItemDetails from './components/ItemDetails'
import { isAndroid, isIOS } from 'tns-core-modules/platform';

// Enable global check if device is android or iOS to disable Maps on Android since mapbox crashes right now
Vue.prototype.$isAndroid = isAndroid;
Vue.prototype.$isIOS = isIOS;

Vue.use(VueAxios, axios)
axios.defaults.baseURL = "http://192.168.1.103:1323";
axios.defaults.headers.post['Content-Type'] = 'application/json';

Vue.registerElement(
  "CardView", 
  () => require("@nstudio/nativescript-cardview").CardView
)

Vue.registerElement("Mapbox", () => require("nativescript-mapbox").MapboxView);

//import App from "./components/App";
if(TNS_ENV !== 'production') {
  Vue.use(VueDevtools)
}
// Prints Vue logs when --env.production is *NOT* set while building
Vue.config.silent = (TNS_ENV === 'production')

var firebase = require("nativescript-plugin-firebase");
firebase
  .init({
    // Optionally pass in properties for database, authentication and cloud messaging,
    // see their respective docs.
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