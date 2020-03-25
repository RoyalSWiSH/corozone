import Vue from "nativescript-vue";
import VueDevtools from 'nativescript-vue-devtools'
import axios from 'axios'
import VueAxios from 'vue-axios'

import LoginPage from './components/LoginPage'
Vue.use(VueAxios, axios)

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
//  new Vue({ 
//    render: h => h('frame', [h(LoginPage)])
// }).$start();
new Vue({
    template: `
    <Frame>
      <LoginPage />
    </Frame>`,
    components: {
      LoginPage,
     // App
    }
}).$start();
