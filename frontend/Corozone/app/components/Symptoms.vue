<template>
 <Page>
     <ActionBar>            
             <GridLayout rows="*" columns="2*, *, *">
                  <Label text="Symptoms (not functional)" column="0"></Label>
                  <!-- <Label text="Groceries" column="0"></Label> -->
                  <!-- <Button :text="'reload' | L" @tap="getGroceryRequests()" column="1" />
                   <Button :text="'Logout' | L" @tap="logoutApp()" column="2" /> -->
             </GridLayout>
        </ActionBar>

     <!-- <Frame ~groceriesFrame id="groceriesFrame">     -->
     <StackLayout>
          <Label v-if="serverFailure" :text="'groceries.notconnected' | L" />
          <Label v-if="rowCount === 0" :text="'groceries.nogroceryrequests' | L" />
               
                <Label text="Did you visita a risk area in the past two weeks?" row="2" margin="5" fontSize="18" colSpan="3" textWrap="true"/>
                <Label text="Do you have shortness of breath?" row="2" margin="5" fontSize="18" colSpan="3"/>
                <Label text="Have you lost sense of smell and taste?" row="2" margin="5" fontSize="18" colSpan="3" textWrap="true"/>
                  <Button text="Video doctor" row="3" colSpan="3" @tap=""/> 
                  <Button text="Call a doctor" row="3" colSpan="3" @tap=""/> 
                  <Button text="Emergency 116 117" row="3" colSpan="3" @tap=""/> 
                  <Button text="Health Department" row="3" colSpan="3" @tap=""/> 
                  <Button text="Mental health support" row="3" colSpan="3" @tap=""/> 
     </StackLayout>
     <!-- </Frame> -->
</Page>
</template>


<script>
import ItemDetails from "./ItemDetails";
import LoginPage from "./LoginPage";
import {
  Accuracy
} from "ui/enums";
import * as geolocation from "nativescript-geolocation";
var LoadingIndicator = require("@nstudio/nativescript-loading-indicator")
  .LoadingIndicator;
var loader = new LoadingIndicator();
export default {
  data: () => {
    return {
      itemsPerRow: 2,
      serverFailure: true,
      groceryRequests: [{
          location: {
            "street": "Adickes-Alee 13",
            "district": "Heddernheim",
            "city": "Frankfurt a.M",
            "distance": "2.0 km",
            "lat": "50.1584",
            "long": "8.6399"
          },
          item_categories: ["Food", "Beverages", "Medicine"],
          items: [
            "2x 1.5 l Wasser",
            "2x Paprika",
            "1x Knäckebrot",
            "500g Mehl",
            "Schampoo",
            "Müllbeutel 20l",
            "Ginger Ale",
            "Zewa"
          ],
          budget: 100,
          forSomeoneElse: "false",
          minimumSupply: "false",
          inQuarantine: "true",
          elderly: "false",
          status: "open"
        },
        // {
        //     location: {
        //         "street": "Adickes-Alee 13",
        //         "district": "",
        //         "city": "Volpertshausen",
        //         "distance": "2.0 km",
        //         "lat": "50.5021339",
        //         "long": "8.5434490"
        //         },
        //     item_categories: ["Medicine"],
        //     items: [
        //         "2x 1.5 l Wasser", 
        //         "2x Paprika", 
        //         "1x Knäckebrot", 
        //         "500g Mehl"]
        // },
      ]
    };
  },
  created() {
    this.getGroceryRequests()
  },
  computed: {
    rowCount: function () {
      if (this.groceryRequests.length) {
        return Math.ceil(this.groceryRequests.length / this.itemsPerRow)
      } else {
        return 
      }
    }
  },
  //   watch: {
  //   isLoggedIn(val) {
  //     if (!val) {
  //       this.$navigateTo(LoginPage, { clearHistory: true });
  //     }
  //   }
  // },
  methods: {
    logout() {
      this.$authService.logout()
    },
    getGroceryRequests() {
      console.log("Get Grocery Reqests...")
      loader.show()
      let that = this
      this.axios({
        method: 'get',
        url: '/groceries/',
      }).then(resp => {
        // this.groceryRequests = resp.data
        // if(resp.data == "") {
        //     alertConnectionError("Error when connecting to server!")
        // }
        this.groceryRequests = resp.data
        this.serverFailure = false
        //  console.log(resp.data);
        loader.hide()
      }).catch((error) => {
        if (error.response) {
          // The request was made and the server responded with a status code
          // that falls out of the range of 2xx
          console.log(error.response.data);
          console.log(error.response.status);
          console.log(error.response.headers);
          //   alertConnectionError("Error when connecting to server!")
          that.serverFailure = true
        } else if (error.request) {
          // The request was made but no response was received
          // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
          // http.ClientRequest in node.js
          //   alertConnectionError("Error when connecting to server!")
          console.log(error.request);
          that.serverFailure = true
        } else {
          // Something happened in setting up the request that triggered an Error
          //   alertConnectionError("Error when connecting to server!")
          console.log(error.message);
          that.serverConnection = false
        }
      });
    },
    alertConnectionError(msg) {
      return alert({
        title: "Corozone",
        okButtonText: "OK",
        message: msg
      });
    },
    seeDetails(groceryRequest) {
      this.$navigateTo(ItemDetails, {
        frame: "itemsFrame",
        props: {
          groceryRequest: groceryRequest
        }
      });
    },
    // Logout leads to strange behaviour like details page not working
    logoutApp() {
      this.$authService.logout()
      //.then(() => {
      //  this.$navigateTo(LoginPage, {  clearHistory: true });
      // });
    },
    onItemTap(args) {
      const view = args.view;
      const page = view.page;
      const tappedItem = view.bindingContext;

      this.$navigateTo(ItemDetails, {
        props: {
          context: tappedItem,
          animated: true,
          transition: {
            name: "slide",
            duration: 200,
            curve: "ease"
          }
        }
      });
    },
    onMapReady(args) {
      args.map.addMarkers([{
        lat: this.groceryRequests[0].location.lat,
        lng: this.groceryRequests[0].location.lang,
        title: "Tracy, CA",
        subtitle: "Home of The Polyglot Developer!",
        onCalloutTap: () => {
          utils.openUrl("https://www.thepolyglotdeveloper.com");
        }
      }]);
    }
  }
};
</script>

<style scoped lang="scss">
// Start custom common variables
@import "~@nativescript/theme/scss/variables/blue";
// End custom common variables

// Custom styles
.add-button {
    height: 30;
    background-color: rgb(51, 51, 206);
    color: white;
    border-radius: 5;
    font-size: 20;
    font-weight: 600;
}

.card {
    background-color: #fff;
    color: #4d4d4d;
    margin: 15;
}

.card-layout {
    padding: 20;
}
</style>