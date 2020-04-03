<template>
 <Page>
     <ActionBar>
            <Label text="Groceries"></Label>
        </ActionBar>
     <ScrollView>
     <!-- <Frame ~groceriesFrame id="groceriesFrame">     -->
     <StackLayout>
      <GridLayout rows="*" columns="*, *"  v-for="i in rowCount" :key="i">
      <card-view class="card" col="0" margin="10" elevation="20" radius="1" @tap="seeDetails()">
          <GridLayout rows="280, 40, 40, 60" columns="*, *, *"> 
                  <Mapbox
                    #map
                    v-if="$isIOS"
                    accessToken="pk.eyJ1Ijoicm95YWxzd2lzaCIsImEiOiJjazg2NHhiMHEwOHV3M25tbWo3bXNsaGhsIn0.yH9oQ-IS6McmJ7lBElv4Zw"
                    mapStyle="traffic_day"
                    colSpan="3"
                    :latitude="items[(i-1)* itemsPerRow].location.lat"
                    :longitude="items[(i-1)* itemsPerRow].location.long"
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
                 <Label v-if="items[(i-1)* itemsPerRow].location.district" :text="items[(i-1)* itemsPerRow].location.district" row="1" margin="5" fontSize="18" colSpan="3" rowSpan="2"/>
                 <Label v-else :text="items[(i-1)* itemsPerRow].location.city" row="1" margin="5" fontSize="18" colSpan="3" rowSpan="2"/>
                 <!-- Display the Type of request, Groceries, Petcare... -->
                 <Label :text="items[(i-1)* itemsPerRow].item_categories" row="2" margin="5" fontSize="18" colSpan="3"/>
                  <Button text="Details" row="3" colSpan="3" @tap="seeDetails(items[(i-1)* itemsPerRow])"/> 
          </GridLayout>
      </card-view>
        <card-view class="card" col="1" margin="10" elevation="20" radius="1" @tap="seeDetails()">
          <GridLayout rows="280, 40, 40, 60" columns="*, *, *"> 
                 <Label :text="items[(i-1)* itemsPerRow + 1].location.district" row="0" margin="5" fontSize="15" colSpan="3" rowSpan="2"/>
                    <Mapbox
                    #map
                    v-if="$isIOS"
                    accessToken="pk.eyJ1Ijoicm95YWxzd2lzaCIsImEiOiJjazg2NHhiMHEwOHV3M25tbWo3bXNsaGhsIn0.yH9oQ-IS6McmJ7lBElv4Zw"
                    mapStyle="traffic_day"
                    colSpan="3"
                    :latitude="items[(i-1)* itemsPerRow + 1].location.lat"
                    :longitude="items[(i-1)* itemsPerRow + 1].location.long"
                    hideCompass="true"
                    zoomLevel="14"
                    showUserLocation="true"
                    disableZoom="false"
                    disableRotation="false"
                    disableScroll="false"
                    disableTilt="false"
                    (mapReady)="onMapReady($event)">
                </Mapbox>
                 <Label v-if="items[(i-1)* itemsPerRow + 1].location.district" :text="items[(i-1)* itemsPerRow].location.district" row="1" margin="5" fontSize="18" colSpan="3" rowSpan="2"/>
                 <Label v-else :text="items[(i-1)* itemsPerRow + 1].location.city" row="1" margin="5" fontSize="18" colSpan="3" rowSpan="2"/>
                 <!-- Display the Type of request, Groceries, Petcare... -->
                 <Label :text="items[(i-1)* itemsPerRow + 1].item_categories" row="2" margin="5" fontSize="18" colSpan="3"/>
                 <Button text="Details" row="3" colSpan="3" @tap="seeDetails"/>
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
import { Accuracy } from "ui/enums";
import * as geolocation from "nativescript-geolocation";

export default {
    data: () => {
        return {
            itemsPerRow: 2,
            items: [
                {
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
                        "Zewa"],
                    budget: 100,
                    forSomeoneElse: "false",
                    minimumSupply: "false",
                    inQuarantine: "true",
                    elderly: "false",
                    status: "open"
                },
                {
                    location: {
                        "street": "Adickes-Alee 13",
                        "district": "",
                        "city": "Volpertshausen",
                        "distance": "2.0 km",
                        "lat": "50.5021339",
                        "long": "8.5434490"
                        },
                    item_categories: ["Medicine"],
                    items: [
                        "2x 1.5 l Wasser", 
                        "2x Paprika", 
                        "1x Knäckebrot", 
                        "500g Mehl"]
                },
                {
                    location: {
                        "street": "Adickes-Alee 13",
                        "district": "",
                        "city": "Volpertshausen",
                        "distance": "2.0 km",
                        "lat": "50.5021339",
                        "long": "8.5434490"
                        },
                    item_categories: ["Childcare"]
                },
                {
                    location: {
                        "street": "Adickes-Alee 13",
                        "district": "Bornheim",
                        "city": "Frankfurt a.M.",
                        "distance": "2.0 km",
                        "lat": "50.5021339",
                        "long": "8.5434490"
                        },
                    item_categories: "Food"
                },
                 {
                    location: "Heddernheim, 2.0 km",
                    item_categories: "Food, Beverages"
                },
                {
                    location: "Niedereschbach, 3.1 km",
                    item_categories: "Medicine"
                },
                {
                    location: "Eckenheim, 3.2 km",
                    item_categories: "Childcare"
                },
                {
                    location: "Bornheim, 1.6 km",
                    item_categories: "Food"
                }                
            ]};
    },
    computed: {
        rowCount: function() {
            return Math.ceil(this.items.length / this.itemsPerRow)
        }
    },
    methods: {
        seeDetails (groceryRequest) {
            this.$navigateTo(ItemDetails, {
                    frame: "itemsFrame",
                    props: { groceryRequest: groceryRequest } 
                });
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
                        lat: 50.1584,
                        lng: 8.6399,
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