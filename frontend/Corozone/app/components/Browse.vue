<template lang="html">
    <Page>
        <ActionBar>
            <Label text="Corozone"></Label>
        </ActionBar>
      <FlexboxLayout class="page">
			<StackLayout class="form">
       <StackLayout class="input-field">
					<Label text="Groceries" class="field-title" fontSize="19"/>
					<StackLayout class="hr-light" />
				</StackLayout>

        <StackLayout class="input-field" marginBottom="25">
					<Label text="Street" class="field-title" fontSize="19"/>
					<TextField class="input" hint="street" keyboardType="street" autocorrect="false" autocapitalizationType="none" v-model="adress.street"
					 returnKeyType="next" @returnPress="" fontSize="18" />
					<StackLayout class="hr-light" />
					<Label text="Code" class="field-title" fontSize="19"/>
					<TextField ref="code" class="input" hint="plzt" keyboardType="plz" autocorrect="false" autocapitalizationType="none" v-model="adress.plz"
					 returnKeyType="next" @returnPress="" fontSize="18" />
					<Label text="City" class="field-title" fontSize="19"/>
					<TextField ref="street" class="input" hint="city" keyboardType="street" autocorrect="false" autocapitalizationType="none" v-model="adress.city"
					 returnKeyType="next" @returnPress="" fontSize="18" />
					<Label text="Items" class="field-title" fontSize="19"/>
					<TextField ref="items" class="input" hint="items" keyboardType="street" autocorrect="false" autocapitalizationType="none" v-model="requestedItems"
					 returnKeyType="next" @returnPress="" fontSize="18" />
				<Button text="Request Groceries" @tap="requestGroceries" class="btn" />
				</StackLayout>
		
			</StackLayout>

	     		</FlexboxLayout>
    </Page>
</template>

<script>
import firebase from "nativescript-plugin-firebase";
import axios from "axios/dist/axios"
// A stub for a service that authenticates users

import * as geolocation from "nativescript-geolocation";
import { Accuracy } from "tns-core-modules/ui/enums"; // used to describe at what accuracy the location should be get


export default {
  data() {
    return {
      isLoggingIn: true,
      user: {
        email: "foo@foo.com",
        password: "barbar",
        confirmPassword: "barbar"
      },
      adress: {
        street: "",
        plz: "",
		city: "",
		lat: 0.0,
		long: 0.0
	  },
	  requestedItems: "",
	  authHeader: {
    	'Accept': 'application/json',
    	'Content-Type': 'application/json'
		},
    };
  },
  methods: {
    toggleForm() {
      this.isLoggingIn = !this.isLoggingIn;
    },
    requestGroceries() {
    //   if (!this.adress.street || !this.adress.plz || !this.adress.city) {
    //     this.alert("Please provide city, street and adress.");
    //     return;
	//   }
	
    //   axios.get('http://webcode.me').then(resp => {
    //     console.log(resp.data);
	//   });
	 

//var geolocation = require("nativescript-geolocation");

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
	//  let that = this;
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
                               this.adress.lat = location.latitude
                               this.adress.long = location.longitude
                            }
                        });
   			return true;				
			}
			});
        })
		
	//console.log(this.requestedItems.split(","))

        // Send a POST request
        this.axios({
            method: 'post',
            url: 'http://192.168.1.103:1323/groceries/create',
            data: {
				budget: 100.4,
				forSomeoneElse: true,
				inQuarantine: false,
				minimumSupply: false,
				elderly: false,
				location: {
					city: "Frankfurt",
					lat: this.adress.lat,
					long: this.adress.long
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
		margin-left: 30;
		margin-right: 30;
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