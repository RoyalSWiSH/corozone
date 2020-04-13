<template lang="html">
    <Page>
        <ActionBar>
            <Label text="Corozone"></Label>
        </ActionBar>
      <FlexboxLayout class="page">
			<StackLayout class="form">
       <StackLayout class="input-field">
					<!-- <Label text="Groceries" class="field-title" fontSize="19"/> -->
					<StackLayout class="hr-light" />
		 </StackLayout>
     <GridLayout
            columns="2*,*"
            rows="*,*"
            width="100%"
            height="25%"
          >
            <TextField
              v-model="itemField"
              col="0"
              row="0"
              hint="Milch"
              editable="true"
              @returnPress="onButtonTap"
            />
            <!-- Configures the text field and ensures that pressing Return on the keyboard produces the same result as tapping the button. -->
            <Button
              col="1"
              row="0"
              text="Add Item"
              @tap="onButtonTap"
            />
			<!-- <Button text="Request Groceries" row="1" colSpan="2" @tap="onTapRequestGroceries" class="btn" /> -->
          </GridLayout>
  <!--      <StackLayout class="input-field" marginBottom="25">
					<Label text="Street" class="field-title" fontSize="19"/>
					<TextField class="input" hint="street" keyboardType="street" autocorrect="false" autocapitalizationType="none" v-model="location.street"
					 returnKeyType="next" @returnPress="" fontSize="18" />
					<StackLayout class="hr-light" />
					<Label text="Code" class="field-title" fontSize="19"/>
					<TextField ref="code" class="input" hint="PLZ" keyboardType="plz" autocorrect="false" autocapitalizationType="none" v-model="location.plz"
					 returnKeyType="next" @returnPress="" fontSize="18" />
					<Label text="City" class="field-title" fontSize="19"/>
					<TextField ref="street" class="input" hint="city" keyboardType="street" autocorrect="false" autocapitalizationType="none" v-model="location.city"
					 returnKeyType="next" @returnPress="" fontSize="18" />
					<Label text="Items" class="field-title" fontSize="19"/>
					</StackLayout> -->
					<StackLayout>
					<ListView
            class="list-group"
            for="gitem in requestedItems"
            style="height:75%"
            separator-color="transparent"
			>
            <v-template>
              <!-- <Label
                id="active-task"
                :text="item.name"
                class="list-group-item-heading"
                text-wrap="true"
              /> -->
					<Label ref="items" class="input" hint="items" keyboardType="street" autocorrect="false" autocapitalizationType="none" :text="gitem.name"
					 returnKeyType="next" @returnPress="" fontSize="18" />
            </v-template>
          </ListView>
		  </StackLayout>
		   <GridLayout
            columns="2*,*"
            rows="*,*"
            width="100%"
            height="25%"
          >
            <!-- <TextField
              v-model="itemField"
              col="0"
              row="0"
              hint="Add Item"
              editable="true"
              @returnPress="onButtonTap"
            /> -->
            <!-- Configures the text field and ensures that pressing Return on the keyboard produces the same result as tapping the button. -->
            <!-- <Button
              col="1"
              row="0"
              text="Add Item"
              @tap="onButtonTap"
            /> -->
			<Button text="Request Groceries" row="0" colSpan="2" @tap="onTapRequestGroceries" class="btn" /><
          </GridLayout>
	
			</StackLayout>

	     </FlexboxLayout>
    </Page>
</template>

<script>
import firebase from "nativescript-plugin-firebase";
import axios from "axios/dist/axios"
// A stub for a service that authenticates users
import { backendService } from "../app";

import * as geolocation from "nativescript-geolocation";
import { Accuracy } from "tns-core-modules/ui/enums"; // used to describe at what accuracy the location should be get
var LoadingIndicator = require("@nstudio/nativescript-loading-indicator")
    .LoadingIndicator;
var loader = new LoadingIndicator();

export default {
  data() {
    return {
	  isLoggingIn: true,
	  itemField: "",
      user: {
        email: "foo@foo.com",
        password: "barbar",
        confirmPassword: "barbar"
      },
      location: {
        street: "",
        street_nr: "",
        plz: "",
        district: "",
        city: "",
        country: "",
		    lat: 0.0,
		    long: 0.0
	  },
	  requestedItems: [{name: "Marmelade"}],
	  authHeader: {
    	'Accept': 'application/json',
    	'Content-Type': 'application/json'
		},
    };
  },
  mounted() {
       this.getAddressFromCoord()
    },
  methods: {
    toggleForm() {
      this.isLoggingIn = !this.isLoggingIn;
    },
    getAddressFromCoord() {
          let lat = this.location.lat
          let long = this.location.long
          let locurl = 'https://api.mapbox.com/geocoding/v5/mapbox.places/'+ long + ','+ lat +'.json?types=address&access_token=pk.eyJ1Ijoicm95YWxzd2lzaCIsImEiOiJjazg2NHhiMHEwOHV3M25tbWo3bXNsaGhsIn0.yH9oQ-IS6McmJ7lBElv4Zw' 
          //let locurl = 'https://api.mapbox.com/geocoding/v5/mapbox.places/8.6267734,50.1732813.json?types=address&access_token=pk.eyJ1Ijoicm95YWxzd2lzaCIsImEiOiJjazg2NHhiMHEwOHV3M25tbWo3bXNsaGhsIn0.yH9oQ-IS6McmJ7lBElv4Zw' 
          console.log(locurl)
          let that = this
              this.axios({
                       method: 'get',
                       url: locurl,
                  }).then(resp => {
                // this.groceryRequests = resp.data
                 //console.log(resp.data.features[0]);
                 console.log(resp.data.features[0].text);
                 that.location.street = resp.data.features[0].text
                 that.location.street_nr = resp.data.features[0].address
                 console.log(resp.data.features[0].context[0].text);
                 that.location.district = resp.data.features[0].context[0].text

                 console.log(resp.data.features[0].context[1].text);
                 that.location.plz = resp.data.features[0].context[1].text

                 console.log(resp.data.features[0].context[2].text);
                 that.location.city = resp.data.features[0].context[2].text
                 console.log(resp.data.features[0].context[3].text);
                 console.log(resp.data.features[0].context[4].text);
                 that.location.country = resp.data.features[0].context[4].text
                 //that.groceryRequests = resp.data
              }).catch((error) => {   if (error.response) {
      // The request was made and the server responded with a status code
      // that falls out of the range of 2xx
      console.log(error.response.data);
      console.log(error.response.status);
      console.log(error.response.headers);
    } else if (error.request) {
      // The request was made but no response was received
      // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
      // http.ClientRequest in node.js
      console.log(error.request);
    } else {
      // Something happened in setting up the request that triggered an Error
      console.log(error.message);
    }});
    return this
    },
    enableGeolocation() {

	 geolocation.isEnabled().then(function (isEnabled) {
                    if (!isEnabled) {
                        geolocation.enableLocationRequest(true, true).then(() => {
                            console.log("User Enabled Location Service");
                        }, (e) => {
                            console.log("Error: " + (e.message || e));
                        }).catch(ex => {
                            console.log("Unable to Enable Location", ex);
                        });
                    }
                }, function (e) {
                    console.log("Error: " + (e.message || e));
                });
                return this
	
    },
    getLocationfromGPS() {
	let that = this
        geolocation.enableLocationRequest(true, true).then(() => {
            geolocation.isEnabled().then(value => {
                if (!value) {
                    console.log("NO permissions!");
                    return false;
                } else {
                    console.log("Have location permissions");
                    geolocation
                        .getCurrentLocation({
                            timeout: 20000
                        })
                        .then(location => {
                            if (!location) {
                                console.log("Failed to get location!");
                            } else {
                               that.location.lat = location.latitude
                               that.location.long = location.longitude
                            }
                        });
   			return this;				
			}
			});
        })
        return this
		
    },
    async onTapRequestGroceries() {
    // TODO: Fix this properly to return with promises / async await
    //this.enableGeolocation().then(res => {getLocationfromGPS()}).then(res => {getAddressFromCoord()}).then(res => {postRequestGroceries()})
    loader.show()
    this.enableGeolocation()
     this.getLocationfromGPS()
    await new Promise(r => setTimeout(r, 2500));
       this.getAddressFromCoord()
    await new Promise(r => setTimeout(r, 2500));
       this.postRequestGroceries()
       loader.hide()
    },
    postRequestGroceries() {
    
    //             geolocation.getCurrentLocation({
    //                 desiredAccuracy: Accuracy.high,
    //                 maximumAge: 5000,
    //                 timeout: 10000
    //             }).then(function (loc) {
    //                 if (loc) {
    //                     that.locations.push(loc);
    //                 }
    //             }, function (e) {
    //                 console.log("Error: " + (e.message || e));
	// 			});

        // Send a POST request
        this.axios({
            method: 'post',
            url: 'http://corozone.sebastian-roy.de/api/v1/groceries/create',
            data: {
        createdBy: backendService.token,
        budget: 100.4,
				forSomeoneElse: true,
				inQuarantine: false,
				minimumSupply: false,
				elderly: false,
				requestedItems: this.requestedItems,
				location: {
					city: this.location.city,
          street: this.location.street,
          district: this.location.district,
          country: this.location.country,
					lat: this.location.lat,
					long: this.location.long
				}
				}
        }).then(function (response) {
            console.log(response.data);  //Outputs API response in CL to verify functionality.
		})
		.catch((error) => {   if (error.response) {
      // The request was made and the server responded with a status code
      // that falls out of the range of 2xx
      console.log(error.response.data);
      console.log(error.response.status);
      console.log(error.response.headers);
    } else if (error.request) {
      // The request was made but no response was received
      // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
      // http.ClientRequest in node.js
      console.log(error.request);
    } else {
      // Something happened in setting up the request that triggered an Error
      console.log(error.message);
    }});


	  },
	  onButtonTap() {
      // Prevent users from entering an empty string
      if (!this.itemField) {
        return;
      }
      console.log(`New Grocery Request added: ${this.itemField}.`);
      // Adds tasks in the ToDo array. Newly added tasks are immediately shown on the screen
      this.requestedItems.unshift({
        name: this.itemField,
      });
      // Clear the text field
      this.itemField = '';
    },
    alert(message) {
      return alert({
        title: "Corozone",
        okButtonText: "OK",
        message: message
      });
    },
}
}

</script>

<style scoped>
	.page {
		align-items: center;
		flex-direction: column;
	}

	.form {
		margin-left: 10;
		margin-right: 10;
		flex-grow: 2;
		vertical-align: middle;
	}

	.logo {
		margin-bottom: 12;
		height: 90;
		font-weight: bold;
	}

	.header {
		horizontal-align: center;
		font-size: 25;
		font-weight: 600;
		margin-bottom: 70;
		text-align: center;
		color: #D51A1A;
	}

	.input-field {
		margin-bottom: 25;
	}

	.input {
		font-size: 18;
		placeholder-color: #A8A8A8;
	}

	.input-field .input {
		font-size: 54;
	}

	.btn-primary {
		height: 50;
		margin: 30 5 15 5;
		background-color: #D51A1A;
		border-radius: 5;
		font-size: 20;
		font-weight: 600;
	}

	.login-label {
		horizontal-align: center;
		color: #A8A8A8;
		font-size: 16;
	}

	.sign-up-label {
		margin-bottom: 20;
	}

	.bold {
		color: #000000;
	}
</style>