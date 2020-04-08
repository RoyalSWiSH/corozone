<template>
<!-- <frame ~detailView id="detailView"> -->
<page>
        <ActionBar>
            <NavigationButton @tap="goBack()" android.systemIcon="ic_menu_back"/>
            <!-- <Label :text="item.name"></Label> -->
        </ActionBar>
    <GridLayout rows="50,200,170,170">
        <StackLayout row="0">
        <Label v-if="groceryRequest.location.district" :text="groceryRequest.location.district" row="1" marginTop="10" marginLeft="25" fontSize="25" colSpan="3" rowSpan="2"/>
        <Label v-else :text="groceryRequest.location.city" row="1" marginTop="10" marginLeft="25" fontSize="25" colSpan="3" rowSpan="2"/>
        <Label :text="groceryRequest.location.street" marginLeft="25" fontSize="25"></Label>
        </StackLayout>
        <StackLayout row="1">
         <Mapbox
                    #map
                    accessToken="pk.eyJ1Ijoicm95YWxzd2lzaCIsImEiOiJjazg2NHhiMHEwOHV3M25tbWo3bXNsaGhsIn0.yH9oQ-IS6McmJ7lBElv4Zw"
                    mapStyle="traffic_day"
                    :latitude="groceryRequest.location.lat"
                    :longitude="groceryRequest.location.long"
                    hideCompass="true"
                    zoomLevel="14"
                    margin="20"                    
                    showUserLocation="true"
                    disableZoom="false"
                    disableRotation="false"
                    disableScroll="false"
                    disableTilt="false"
                    (mapReady)="onMapReady($event)">
                </Mapbox> 
        </StackLayout>
		<StackLayout class="hr-light" />

        <ScrollView row="2">
        <ListView for="item in groceryRequest.requestedItems">
  <v-template>
    <CheckBox :text="item.name" checked="false" fontSize="25" class="fontBig"></CheckBox>
  </v-template>
</ListView>
        </ScrollView>
		<StackLayout class="hr-light" />
        <StackLayout row="3">
        <Label :text="groceryRequest.budget" margin="5" marginLeft="25" fontSize="20"></Label>
        <TextView editable="false" v-if="groceryRequest.inQuarantine === 'true' " text="You deliver a request to a person in self-quarantine." margin="5" marginLeft="25" fontSize="20" />
        <TextView editable="false" v-if="groceryRequest.elderly === 'true'" text="You deliver a request for a elderly person" margin="5" marginLeft="25" fontSize="20" />
        <TextFieldTextView editable="false" v-if="groceryRequest.forSomeoneElse === 'true'" text="You deliver groceries to a different person than the one that put in the request" margin="5" marginLeft="25" fontSize="20" />
        <!-- <NavigationButton @tap="goBack" android.systemIcon="ic_menu_back"/> -->
        <!-- <GridLayout> -->
             <Button text="Accept" row="3" colSpan="3" @tap="acceptGroceries()"/> 
             <Button text="Open Google Maps" row="4" colSpan="3" @tap="openGoogleMaps()"/> 
            <!-- <Label class="m-10 h3" :text="item.items[0]" verticalAlignment="top" ></Label> -->
        <!-- </GridLayout> -->
        </StackLayout>
    </GridLayout>
</page>
<!-- </frame> -->
</template>

<script>

import { Directions } from "nativescript-directions";

// instantiate the plugin
let directions = new Directions();

directions.available().then(avail => {
    console.log(avail ? "Yes" : "No");
});

export default {
    props: ["context"],

    computed: {
        // item() {
        //     return this.context || {};
        // }
          },
    methods: {
          goBack(args) {
          this.$navigateBack();
        },
        acceptGroceries() {
            console.log("Accepted Grocery Request")
            console.log(JSON.stringify(this.groceryRequest.requestedItems))
        },
        openGoogleMaps() {
                    directions.navigate({
            from: { // optional, default 'current location'
                lat: 52.215987,
                lng: 5.282764
            },
            to: { // either pass in a single object or an Array (see the TypeScript example below)
                address: "Hof der Kolommen 34, Amersfoort, Netherlands"
            }
            // for iOS-specific options, see the TypeScript example below.
            }).then(
            function() {
                console.log("Maps app launched.");
            },
            function(error) {
                console.log(error);
            }
            );  
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
    },
    props: {
        groceryRequest: {
            type: Object,
            required: true
        }
    },
};
</script>

<style scoped lang="scss">
// Start custom common variables
@import "~@nativescript/theme/scss/variables/blue";

@import '~@nativescript/theme/css/core.css';
@import '~@nativescript/theme/css/default.css';
// End custom common variables
CheckBox {
  vertical-align: center;
  color: white;
  font-size: 20;
}
.fontBig {
  font-size: 25;
}
// Custom styles

</style>