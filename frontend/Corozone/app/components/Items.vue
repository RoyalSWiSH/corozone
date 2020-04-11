<template>
 <Page>
     <ActionBar>            
             <GridLayout rows="*" columns="3*, *, *">
                  <!-- <Label text="Groceries" column="0"></Label> -->
                  <Button text="Reload" @tap="getGroceryRequests()" column="1" />
                   <Button text="Logout" @tap="logoutApp()" column="2" />
             </GridLayout>
        </ActionBar>
     <ScrollView>
     <!-- <Frame ~groceriesFrame id="groceriesFrame">     -->
     <StackLayout>
          <Label v-if="serverFailure" text="Not Connected" />
      <GridLayout rows="*" columns="*, *"  v-for="i in rowCount" :key="i">
      <card-view class="card" col="0" margin="10" elevation="20" radius="1" @tap="seeDetails()">
          <GridLayout rows="280, 40, 40, 60" columns="*, *, *"> 
                  <Mapbox
                    #map
                    accessToken="pk.eyJ1Ijoicm95YWxzd2lzaCIsImEiOiJjazg2NHhiMHEwOHV3M25tbWo3bXNsaGhsIn0.yH9oQ-IS6McmJ7lBElv4Zw"
                    mapStyle="traffic_day"
                    colSpan="3"
                    :latitude="groceryRequests[(i-1)* itemsPerRow].location.lat"
                    :longitude="groceryRequests[(i-1)* itemsPerRow].location.long"
                    hideCompass="true"
                    zoomLevel="14"
                    showUserLocation="true"
                    disableZoom="false"
                    disableRotation="false"
                    disableScroll="false"
                    disableTilt="false"
                    (mapReady)="onMapReady($event)">
                </Mapbox>
                 <!-- Check if a district for larger cities is present and if not display the city -->
                 <Label v-if="groceryRequests[(i-1)* itemsPerRow].location.district" :text="groceryRequests[(i-1)* itemsPerRow].location.district" row="1" margin="5" fontSize="18" colSpan="3" rowSpan="2"/>
                 <Label v-else :text="groceryRequests[(i-1)* itemsPerRow].location.city" row="1" margin="5" fontSize="18" colSpan="3" rowSpan="2"/>
                 <Label :text="groceryRequests[(i-1)* itemsPerRow].createdBy" row="2" margin="5" fontSize="18" colSpan="3" rowSpan="2"/>
                 <!-- Display the Type of request, Groceries, Petcare... -->
                 <Label :text="groceryRequests[(i-1)* itemsPerRow].item_categories" row="2" margin="5" fontSize="18" colSpan="3"/>
                  <Button text="Details" row="3" colSpan="3" @tap="seeDetails(groceryRequests[(i-1)* itemsPerRow])"/> 
          </GridLayout>
      </card-view>
        <card-view class="card" col="1" margin="10" elevation="20" radius="1" @tap="seeDetails()">
          <GridLayout rows="280, 40, 40, 60" columns="*, *, *"> 
                 <Label :text="groceryRequests[(i-1)* itemsPerRow + 1].location.district" row="0" margin="5" fontSize="15" colSpan="3" rowSpan="2"/>
                    <Mapbox
                    #map
                    accessToken="pk.eyJ1Ijoicm95YWxzd2lzaCIsImEiOiJjazg2NHhiMHEwOHV3M25tbWo3bXNsaGhsIn0.yH9oQ-IS6McmJ7lBElv4Zw"
                    mapStyle="traffic_day"
                    colSpan="3"
                    :latitude="groceryRequests[(i-1)* itemsPerRow + 1].location.lat"
                    :longitude="groceryRequests[(i-1)* itemsPerRow + 1].location.long"
                    hideCompass="true"
                    zoomLevel="14"
                    showUserLocation="true"
                    disableZoom="false"
                    disableRotation="false"
                    disableScroll="false"
                    disableTilt="false"
                    (mapReady)="onMapReady($event)">
                </Mapbox>
                 <Label v-if="groceryRequests[(i-1)* itemsPerRow + 1].location.district" :text="groceryRequests[(i-1)* itemsPerRow+1].location.district" row="1" margin="5" fontSize="18" colSpan="3" rowSpan="2"/>
                 <Label v-else :text="groceryRequests[(i-1)* itemsPerRow + 1].location.city" row="3" margin="5" fontSize="18" colSpan="3" rowSpan="2"/>
                 <Label :text="groceryRequests[(i-1)* itemsPerRow + 1].createdBy" row="2" margin="5" fontSize="18" colSpan="3" rowSpan="2"/>
                 <!-- Display the Type of request, Groceries, Petcare... -->
                 <Label :text="groceryRequests[(i-1)* itemsPerRow + 1].item_categories" row="2" margin="5" fontSize="18" colSpan="3"/>
                 <Button text="Details" row="3" colSpan="3" @tap="seeDetails(groceryRequests[(i-1)* itemsPerRow+1])"/>
          </GridLayout>
      </card-view>
      </GridLayout>
     </StackLayout>
     <!-- </Frame> -->
     </ScrollView>
</Page>
</template>

<script>
import ItemDetails from "./ItemDetails";
import LoginPage from "./LoginPage";
import { Accuracy } from "ui/enums";
import * as geolocation from "nativescript-geolocation";
var LoadingIndicator = require("@nstudio/nativescript-loading-indicator")
    .LoadingIndicator;
var loader = new LoadingIndicator();
export default {
    data: () => {
        return {
            itemsPerRow: 2,
            serverFailure: true,
            groceryRequests: [
                // {
                //     location: {
                //         "street": "Adickes-Alee 13",
                //         "district": "Heddernheim", 
                //         "city": "Frankfurt a.M",
                //         "distance": "2.0 km",
                //         "lat": "50.1584",
                //         "long": "8.6399"
                //         },
                //     item_categories: ["Food", "Beverages", "Medicine"],
                //     items: [
                //         "2x 1.5 l Wasser", 
                //         "2x Paprika", 
                //         "1x Knäckebrot", 
                //         "500g Mehl",
                //         "Schampoo",
                //         "Müllbeutel 20l",
                //         "Ginger Ale",
                //         "Zewa"],
                //     budget: 100,
                //     forSomeoneElse: "false",
                //     minimumSupply: "false",
                //     inQuarantine: "true",
                //     elderly: "false",
                //     status: "open"
                // },
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
                // {
                //     location: {
                //         "street": "Adickes-Alee 13",
                //         "district": "",
                //         "city": "Volpertshausen",
                //         "distance": "2.0 km",
                //         "lat": "50.5021339",
                //         "long": "8.5434490"
                //         },
                //     item_categories: ["Childcare"]
                // },
                // {
                //     location: {
                //         "street": "Adickes-Alee 13",
                //         "district": "Bornheim",
                //         "city": "Frankfurt a.M.",
                //         "distance": "2.0 km",
                //         "lat": "50.5021339",
                //         "long": "8.5434490"
                //         },
                //     item_categories: "Food"
                // },
                //  {
                //     location: "Heddernheim, 2.0 km",
                //     item_categories: "Food, Beverages"
                // },
                // {
                //     location: "Niedereschbach, 3.1 km",
                //     item_categories: "Medicine"
                // },
                // {
                //     location: "Eckenheim, 3.2 km",
                //     item_categories: "Childcare"
                // },
                // {
                //     location: "Bornheim, 1.6 km",
                //     item_categories: "Food"
                // }                
            ]};
    },
    created() {
       this.getGroceryRequests()
    },
    computed: {
        rowCount: function() {
            return Math.ceil(this.groceryRequests.length / this.itemsPerRow)
          //  return 6
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
        logout () {
            this.$authService.logout()
        },
        getGroceryRequests () {
            console.log("Get Grocery Reqests...")
            loader.show()
            let that = this
              this.axios({
                       method: 'get',
                       url: 'http://corozone.sebastian-roy.de/groceries/getgroceries',
                  }).then(resp => {
                // this.groceryRequests = resp.data
                // if(resp.data == "") {
                //     alertConnectionError("Error when connecting to server!")
                // }
                //Bug: apparantly this doesnt work when there are more than 17 grocery requests
                 that.groceryRequests = resp.data
                 that.serverFailure = false
                 console.log(resp.data);
                 loader.hide()
              }).catch((error) => {   if (error.response) {
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
    }});
        },
        alertConnectionError (msg) {
          return  alert({
        title: "Corozone",
        okButtonText: "OK",
        message: msg
      });
        },
        seeDetails (groceryRequest) {
            this.$navigateTo(ItemDetails, {
                    frame: "itemsFrame",
                    props: { groceryRequest: groceryRequest } 
                });
        },
        // Logout leads to strange behaviour like details page not working
        logoutApp() {
      this.$authService.logout()
      //.then(() => {
      //  this.$navigateTo(LoginPage, {  clearHistory: true });
     // });
    },
        onItemTap (args) {
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
                    }}});
        },
          onMapReady(args) {
                args.map.addMarkers([
                    {
                        lat: this.groceryRequests[0].location.lat,
                        lng: this.groceryRequests[0].location.lang,
                        title: "Tracy, CA",
                        subtitle: "Home of The Polyglot Developer!",
                        onCalloutTap: () => {
                            utils.openUrl("https://www.thepolyglotdeveloper.com");
                        }
                    }
                ]);
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