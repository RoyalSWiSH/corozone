<template lang="html">
    <Page>
        <ActionBar>
            <Label text="Corozone"></Label>
            <GridLayout rows="*" columns="2*, *, *">
                  <!-- <Label text="Groceries" column="0"></Label> -->
                  <Label text="Groceries (data is public)" column="0"></Label>
                  <Button text="Add a friend" @tap="showModal" column="1" />
                   <!-- <Button :text="'Logout' | L" @tap="logoutApp()" column="2" /> -->
             </GridLayout>
        </ActionBar>
      <FlexboxLayout class="page">
   
			<StackLayout class="form">
       <Label ref="message" class="input" hint="" :text="notification_message" fontSize="18" textWrap="true"/> 
       <StackLayout class="input-field">
					<!-- <Label text="Groceries" class="field-title" fontSize="19"/> -->
					<StackLayout class="hr-light" />
		 </StackLayout>
     <GridLayout
            columns="2*,*"
            rows="*,*"
            width="100%"
            height="20%"
          >
            <TextField
              v-model="itemField"
              col="0"
              row="0"
              hint="z.B. Milch"
              editable="true"
              @returnPress="onButtonTap"
            />
            <!-- Configures the text field and ensures that pressing Return on the keyboard produces the same result as tapping the button. -->
            <Button
              col="1"
              row="0"
              :text="'groceries.additem' | L"
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
            for="gitem in shoppingList"
            style="height:80%"
            separator-color="transparent"
			>
            <v-template>
              <!-- <Label
                id="active-task"
                :text="item.name"
                class="list-group-item-heading"
                text-wrap="true"
              /> -->
              <GridLayout
            columns="2*,*"
            rows="*,*"
            width="100%"
            height="25%"
          > 
					<Label ref="items" class="input" hint="items" keyboardType="street" autocorrect="false" autocapitalizationType="none" col="0" :text="gitem.name+' '+gitem.status"
					 returnKeyType="next" @returnPress="" @tap="onItemTap(gitem)" fontSize="18" /> 
           <Button
              v-if="myUID === gitem.uid"
              col="1"
              row="0"
              :text="'delete' | L"
              @tap="onButtonTapDelete(gitem)"
            />
         <Label v-if="myUID != gitem.uid" ref="items" class="input" hint="items" keyboardType="street" autocorrect="false" autocapitalizationType="none" col="0" :text="'für: ' +nameFromUID(gitem)"
					 returnKeyType="next" @returnPress="" fontSize="9" row="1"/>  
            </GridLayout>
            </v-template>
          </ListView>
		  </StackLayout>
		   <GridLayout
            columns="2*,*"
            rows="*,*"
            width="100%"
            height="20%"
          >
			<Button automationText="Tap on this button to request groceries" :text="'groceries.requestgroceries' | L" row="0" colSpan="2" @tap="onTapRequestGroceries" class="btn" /><
          </GridLayout>
	
			</StackLayout>

	     </FlexboxLayout>
    </Page>
</template>

<script>
import firebase from "nativescript-plugin-firebase";
import firestore from "nativescript-plugin-firebase/app/firestore"
import axios from "axios/dist/axios"
// A stub for a service that authenticates users
import { backendService } from "../app";
import { mapState, mapGetters } from "vuex";
import * as dialogs from "tns-core-modules/ui/dialogs";

import * as geolocation from "nativescript-geolocation";
import { Accuracy } from "tns-core-modules/ui/enums"; // used to describe at what accuracy the location should be get
import ModalComponent from "./ModalAddFriends";
import {Page, View, ShowModalOptions} from "tns-core-modules/ui/page"

import gql from "graphql-tag";

const LoginQuery = gql`
 query MyQuery {
  delivery_status {
    closed_at
    created_at
    helper_user_id
    receipt_amount
    status
  }
}
`;


var LoadingIndicator = require("@nstudio/nativescript-loading-indicator")
    .LoadingIndicator;
var loader = new LoadingIndicator();

//const db = firebase.firestore
//const groceriesCollection = db.collection("Groceries");


const arrayToObject = (array) =>
   array.reduce((obj, item) => {
     obj[item.name] = item
     return obj
   }, {})

export default {
  data() {
    return {
	  isLoggingIn: true,
    itemField: "",
    myUID: backendService.token,
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
  async created() {
     //  this.requestedItems = O(bject.assign({}, this.$store.state.shoppingList)

       this.getAddressFromCoord()
       const db = firebase.firestore
      const groceriesCollection = db.collection("Groceries");

// Get Groceries from Firebase by ID and if unable, create a List
// TODO:Error in created hook (Promise/async) Possible here is an Error "TypeError: Cannot read property 'doc' of null"
       db.collection("Groceries").doc(backendService.token).get().then( doc4 => {
        console.log("Firebase")
       console.log(`${doc4.id} => ${JSON.stringify(doc4.data())}`);
       this.$store.commit("setShoppingList", Object.values(doc4.data()) ) 
    })
    .then( () => {
 
console.log("Friend List IDS:" + this.$store.getters.getFriendListIDs)
for(const friendID of this.$store.getters.getFriendListIDs) {
  console.log("Friend ID:" + friendID.name)
this.subscribeToShoppingList(friendID.id)     
    }}).catch( () => {
      const shoppingObject = arrayToObject(this.shoppingList)
console.log(shoppingObject)
// console.log("BackendToken:")
// console.log("BackendToken1: " + backendService.token)
 const db = firebase.firestore
      const groceriesCollection = db.collection("Groceries");
  groceriesCollection.doc(backendService.token).set(shoppingObject).then(doc3 => {
  console.log(`Shopping List updated ${doc3}`);
       // this.apolloRequest() 
}); 
 
    });
    
      // This read is actually not required. just push to local list and compare once and a while to save reads
      // Actually it is require if you want to receive status updates in realtime
        const unsubscribe = groceriesCollection.doc(backendService.token).onSnapshot(doc => {
        
  console.log(doc.data())
  console.log("Subscribed")
  console.log(Object.values(doc.data()))
  //this.shoppingList = Object.values(doc.data()) 
  this.$store.commit("mergeShoppingList", Object.values(doc.data()) ) 
          
});
// Loop through firendIDs and add listeners
//var friendID = ""



// const unsubscribe2 = groceriesCollection.doc("wMomJv0pOizSrc2ypeWg").onSnapshot(doc2 => {
//   console.log(doc2.data())
//   console.log("Subscribed2")
//   console.log(Object.values(doc2.data()))
//   this.$store.commit("mergeShoppingList", Object.values(doc2.data()) )
// });


  //  },
//   async  mounted() {
//       console.log("Friend List IDS:" + this.$store.getters.getFriendListIDs)
// for(const friendID of this.$store.getters.getFriendListIDs) {
//   console.log("Friend ID:" + friendID.name)
// await this.subscribeToShoppingList(friendID.id)
// }
    },
  computed: {
    ...mapState(["shoppingList", "notification_message", "friendListIDs"]),
    ...mapGetters(["getShoppingList", "getNotificationMessage", "getFriendListIDs"]),
    showShoppingList: function() {
      return this.getShoppingList
    }
    // nameFromUID(uid) {
    //   return this.friendListIDs.find(x => x.uid === uid)
    // }
    // changeKundeStore: function()
    //     {
    //         return this.$store.state.changeKundeStore
    //     }
   },
  methods: {
    toggleForm() {
      this.isLoggingIn = !this.isLoggingIn;
    },
    apolloRequest() {
      console.log("Apollo Query")
      try {
        this.$apollo
          .queries({
            document: LoginQuery,
            variables: {  }
          })
          .then(({ data: { signinUser } }) => {
            console.log("Apollo Answer")
            console.log(signinUser);
            
          });
      } catch (e) {
        console.error("mutation error", e);
      }
    },
    showModal(args) {
            const view = args.object // as View
            // const opts = new ShowModalOptions()
            // opts = {
            //   context: null,
            //   closeCallBack: () => console.log("Modal closed"),
            //   ios: {
            //     presentationStyle: 5
            //   }
            // }
            this.$showModal(ModalComponent,{
              context: null,
              closeCallBack: () => console.log("Modal closed"),
              ios: {
                presentationStyle: 5
              }});
        },
    // TODO: Potential Privacy issue. Don't let every users see all uid's with names    
    nameFromUID(item) {
      console.log("Name from UID" + item.name)
      console.log(this.$store.getters.getFriendListIDs)
      console.log(this.friendListIDs)
      console.log(this.friendListIDs.find(x => x.id === item.uid))
      if(this.$store.getters.getFriendListIDs) {
      if(this.friendListIDs.find(friend => friend.id === item.uid)){
      return this.friendListIDs.find(friend => friend.id === item.uid).name}
      else {
        console.log("No Friend name found")
        return }
      }
    },
    // Move this somewhere else where it can be accessed by this component and ModalAddFriends.vue
   subscribeToShoppingList(uid) {
 if(uid){    
const db = firebase.firestore
const groceriesCollection = db.collection("Groceries");
console.log("Substibe to shopping list: " + uid)
       const unsubscribe2 = groceriesCollection.doc(uid).onSnapshot(doc2 => {
  console.log(doc2.data())
  console.log("Subscribed2")

  console.log(Object.values(doc2.data()))
  this.$store.commit("mergeShoppingList", Object.values(doc2.data()) )
});
  return unsubscribe2
  }
  else {
    console.log("Unable to subscribe. Probably no or wrong uid set")
  }
    },
    onButtonTapDelete(item) {
        console.log("Deleted Item")
        this.$store.commit("delItemFromShoppingList", item)
        //  this.requestedItems.splice(args.index, 1);
        this.$store.commit("setNotificationMessage", "Item gelöscht")

//    const shoppingObject = await arrayToObject(this.shoppingList)
// console.log(shoppingObject)
  const db = firebase.firestore
      const groceriesCollection = db.collection("Groceries");
 groceriesCollection.doc(backendService.token).update({
   [item.name]: db.FieldValue.delete()
 }).then(doc => {
 // console.log(`Shopping List updated ${doc}`);
});
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
       // TODO: Mark all Items as requested
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
            url: '/groceries/create',
            data: {
        createdBy: backendService.token,
        budget: 100.4,
				forSomeoneElse: true,
				inQuarantine: false,
				minimumSupply: false,
				elderly: false,
			 	requestedItems: this.getshoppingList,  //Don't request groceries from list of friends 
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
          //TODO: Mark all Items as requested
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
      // this.requestedItems.unshift({
      //   name: this.itemField,
      // });

     
      let item = {name: this.itemField, status: "open", crypto: "none", 
      uid: backendService.token, 
      store_category: "none"}
      this.$store.commit("addItemToShoppingList", item);
      // Clear the text field
        console.log("preFirebase")
//note that the options object is optional, but you can use it to specify the source of data ("server", "cache", "default").
// await groceriesCollection.get({ source: "server" }).then(querySnapshot => {
//   querySnapshot.forEach(doc => {
//     console.log("Firebase")
//     console.log(`${doc.id} => ${JSON.stringify(doc.data())}`);
//      if (doc.exists) {
    
//   }
// });
//   });

//console.log(JSON.stringify(this.shoppingList))

  const db = firebase.firestore
  const groceriesCollection = db.collection("Groceries");
// const shoppingObject = arrayToObject(this.shoppingList)
// console.log(shoppingObject)
// let that = this
// groceriesCollection.doc(backendService.token).set(shoppingObject).then(doc => {
//   console.log(`Shopping List updated ${doc}`);
// }).then(() => {that.itemField = '';});

 groceriesCollection.doc(backendService.token).update({
  [item.name]: item
}).then(doc => {
  console.log(`Shopping List updated ${doc}`);
}).then(() => {
 this.itemField = '';
});
this.itemField = '';
// then after a while, to detach the listener:
//unsubscribe();

//unsubscribe2();

  },
    onItemTap(item) {
      if(backendService.token == item.uid) {
        if(item.status == "open" ) {
          console.log(item.name + " bought myself.")
           this.$store.commit("markItemAsSelfbought", item)}
        else if(item.status == "selfbought" || item.status == "helperbought" || item.status == "helpernotavailable") {
          console.log(item.name + " need again.")
          this.$store.commit("markItemAsOpen", item)
          }
      }
      else {
       if(item.status == "open" ) {
          console.log(item.name + " bought myself.")
           this.$store.commit("markItemAsHelperBought", item)}
           // TODO: Maybe remove "selfbought", since should not buy items another person already bought himself
        else if(item.status == "selfbought" || item.status == "helperbought" || item.status == "helpernotavailable") {
          console.log(item.name + " need again.")
          this.$store.commit("markItemAsOpen", item)
          } 
      }

      // Update the status of the item according to whom the item belongs (item.uid)
      const db = firebase.firestore
      const groceriesCollection = db.collection("Groceries");
 groceriesCollection.doc(item.uid).update({
   [item.name]: item
 }).then(doc => {
 // console.log(`Shopping List updated ${doc}`);
});
  
    },
    alert() {
      // return alert({
      //   title: "Corozone",
      //   okButtonText: "OK",
      //   message: message
      // });
    return dialogs.prompt({
    title: "Your title",
    message: "Your UserID is: " +backendService.token,
    okButtonText: "Nutzer hinzufügen",
    cancelButtonText: "Löschen",
    neutralButtonText: "Zurück",
    defaultText: "D4YX72AVKrZp4cHONP23llO5omk2:Sebastian",
    inputType: dialogs.inputType.text
}).then(r => {
    console.log("Dialog result: " + r.result + ", text: " + r.text);
    if(r.result) {
      // TODO: Don't loose the reference to the listener and unregister it when view is changed
    const id = r.text.split(":")[0]
    const name = r.text.split(":")[1]
    const a = this.subscribeToShoppingList(id)
    const friend = {id: id, name: name}
    this.$store.commit("addFriendID", friend ) 
    }
    else {
     this.$store.commit("delFriendID", friend )  
    }    
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