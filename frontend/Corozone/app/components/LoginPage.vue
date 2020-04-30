<template lang="html">
    <Page>
        <ActionBar>
            <Label text="Corozone"></Label>
        </ActionBar>
      <FlexboxLayout class="page">
			<StackLayout class="form">
				<Image class="logo" src="~/images/logo.png" />
				<Label class="header" text="Corozone" />
        	<StackLayout v-show="!isLoggingIn" class="input-field">
           <TextField ref="firstName" class="input" hint="First Name" v-model="user.firstName" returnKeyType="done"
					 fontSize="18" />
            <TextField ref="lastName" class="input" hint="Last Name" v-model="user.lastName" returnKeyType="done"
					 fontSize="18" />
					<StackLayout class="hr-light" />

				</StackLayout>
				<StackLayout class="input-field" marginBottom="25">
					<TextField class="input" hint="Email" keyboardType="email" autocorrect="false" autocapitalizationType="none" v-model="user.email"
					 returnKeyType="next" @returnPress="focusPassword" fontSize="18" />
					<StackLayout class="hr-light" />
				</StackLayout>

				<StackLayout class="input-field" marginBottom="25">
					<TextField ref="password" class="input" hint="Password" secure="true" v-model="user.password" :returnKeyType="isLoggingIn ? 'done' : 'next'"
					 @returnPress="focusConfirmPassword" fontSize="18" />
					<TextField ref="confirmPassword" v-show="!isLoggingIn"  class="input" hint="Confirm password" secure="true" v-model="user.confirmPassword" returnKeyType="done"
					 fontSize="18" />
					<StackLayout class="hr-light" />
				</StackLayout>

				<!-- <StackLayout v-show="!isLoggingIn" class="input-field">
					<StackLayout class="hr-light" />
				</StackLayout> -->

				<Button :text="isLoggingIn ? 'Log In' : 'Sign Up'" @tap="submit" class="btn btn-primary m-t-20" />
				<Label v-show="isLoggingIn" text="Forgot your password?" class="login-label" @tap="forgotPassword" />
			</StackLayout>

			<Label class="login-label sign-up-label" @tap="toggleForm">
	          <FormattedString>
	            <Span :text="isLoggingIn ? 'Don’t have an account? ' : 'Back to Login'" />
	            <Span :text="isLoggingIn ? 'Sign up' : ''" class="bold" />
	          </FormattedString>
	        </Label>
		</FlexboxLayout>
    </Page>
</template>

<script>
import firebase from "nativescript-plugin-firebase"
import App from "./App"

import axios from 'axios'
import VueAxios from 'vue-axios'
// import BackendService from '@/services/BackendService'
// import AuthService from '@/services/AuthService'

import { mapState } from "vuex"


const timerModule = require("tns-core-modules/timer")
var LoadingIndicator = require("@nstudio/nativescript-loading-indicator")
    .LoadingIndicator;
var loader = new LoadingIndicator();
// A stub for a service that authenticates users.
// const userService = {
//   async register(user) {
//     return await firebase.createUser({
//       email: user.email,
//       password: user.password
//     });
//   },
//   async login(user) {
//     return await firebase.login({
//         type: firebase.LoginType.PASSWORD,
//         passwordOptions: {
//           email: user.email,
//           password: user.password
//         }
//     });
//   },
//   async resetPassword(email) {
//     return await firebase.resetPassword({ 
//       email: email
//       });
// }
// }
// A stub for the main page of your app. In a real app you’d put this page in its own .vue file.
const HomePage = {
  template: `
	<Page>
        <Label class="m-20" textWrap="true" text="You have successfully authenticated. This is where you build your core application functionality."></Label>
	</Page>
	`
};
export default {
  data() {
    return {
      isLoggingIn: true,
      isInitialized: false,
      user: {
        email: "",
        password: "",
        confirmPassword: "",
        uid: "",
        firstName: "",
        lastName: ""
      }
    };
  },
  created() {
    // setTimeout(() => {
    //   this.isInitialized = true;
    // }, 1500);
    if(this.$store.state.isLoggedIn!=null){
      this.isInitialized = true;
    }
 },
 mounted() {

   //TODO: Check if this can belong into created()
firebase
  .init({
    // Authentication
    onAuthStateChanged: data => { 
      console.log((data.loggedIn ? "Logged in to firebase" : "Logged out from firebase") + " (firebase.init() onAuthStateChanged callback)");
      if (data.loggedIn) {
        this.$backendService.token = data.user.uid
        console.log("uID: " + data.user.uid)
        this.$store.commit('setIsLoggedIn', true)
      }
      else {     
        this.$store.commit('setIsLoggedIn', false) 
      }
    },

    // Push Notification
      showNotifications: true,
      showNotificationsWhenInForeground: true,

      onPushTokenReceivedCallback: (token) => {
        console.log('[Firebase] onPushTokenReceivedCallback:', { token });

        axios({
          method: 'post',
          url: '/users/'+this.$backendService.token+'/pushtoken',
          data: {
            firebasePushtoken: token
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
      onMessageReceivedCallback: (message) => {
        console.log('[Firebase] onMessageReceivedCallback:', { message });
        this.$store.commit("setNotificationMessage", message.body)
        if(message.data.acceptedItems){
        store.commit("setAcceptedItems",message.data.acceptedItems)


        var notFoundItems = JSON.parse(message.data.acceptedItems).filter(function (item) {
          return item.status == "helpernotavailable" ||
                 item.status == "helpernotfound";
        });

        // Check if some items were not found and create a message
        if(notFoundItems.length>0){
          var arr = new Array(notFoundItems.length);
          for(var i=0; i<arr.length; i++) {
           arr[i]=notFoundItems[i].name;
          }

          var itemsString;
          if(arr.length == 2){
            itemsString = arr.toString().replace(",", " und ") 
          }
          else {
          itemsString = arr.toString().replace(",", ", ") 
          }

          var msg= message.body+ " " + itemsString + " hat er/sie leider nicht bekommen :(."
          // TODO: Send firstname in data part to avoid splitting in message
    axios.get("https://genderapi.io/api/?name=" +  message.body.split(" ")[0])
  .then(function (genderrequest) {
    // handle success
    console.log(genderrequest);
    if(genderrequest.data.gender == "male") {
      msg = msg.replace("er/sie", "er")
    }
    else if(genderrequest.data.gender == "female") {
      msg = msg.replace("er/sie", "sie")
    }
  })
  .catch(function (error) {
    // handle error
    console.log(error);
  })
  .then(function () {
    // always executed
        store.commit("setNotificationMessage", msg)
  });
     }
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
  },
 
  // VueX will synchronize the state of this variable to the page by adding this in the export default { object 
  // What does ... mean?
  computed: {
    ...mapState(["isLoggedIn"])
  },
  watch: {   
    isLoggedIn(val) {
      if (!val) {        
        this.isInitialized = true;        
      }else{
        this.$navigateTo(App, { clearHistory: true });
      }
    }
  },
  methods: {
    toggleForm() {
      this.isLoggingIn = !this.isLoggingIn;
    },
    async postUserProfile(uid) {

        // Send a POST request
        return await this.axios({
            method: 'post',
            url: '/users/createprofile',
            data: {
        user_id: uid,      
        firstName: this.user.firstName,
        lastName: this.user.lastName,
        email: this.user.email
				}})
    //     .then(function (response) {
    //         console.log(response.data);  //Outputs API response in CL to verify functionality.
		// })
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
    submit() {
      if (!this.user.email || !this.user.password) {
        this.alert("Please provide both an email address and password.");
        return;
      }

      loader.show()
      if (this.isLoggingIn) {
        this.login();
      } else {
        this.register();
      }
    },
     login() {
      this.$authService
        .login(this.user)
        .then(() => {
          loader.hide();
          this.$navigateTo(App);
        })
        .catch(err => {
          console.error(err);
          loader.hide();
          this.alert(err);
        });
    },
    register() {
      if (this.user.password != this.user.confirmPassword) {
        loader.hide();
        this.alert("Your passwords do not match.");
        return;
      }
      if (this.user.password.length < 6) {
        loader.hide();
        this.alert("Your password must at least 6 characters.");
        return;
      }
      // TODO: Reactivate validating email adresses later
      // var validator = require("email-validator");
      // if (!validator.validate(this.user.email)) {
      //   loader.hide();
      //   this.alert("Please enter a valid email address.");
      //   return;
      // }
      this.$authService
        .register(this.user)
        .then(uid => {
          //loader.hide();
         // this.alert("Your account was successfully created.");
          this.isLoggingIn = true;
          //console.log(uid)
          return uid
        })
        .then(uid =>  {
          console.log("Update User Profile")
          console.log(uid)
          this.postUserProfile(uid)

//             const db = firebase.firestore
//            const groceriesCollection = db.collection("Groceries");
//             // const shoppingObject = arrayToObject(this.shoppingList)
//             //   console.log(shoppingObject)

//  groceriesCollection.doc(uid).set({exists: true}).then(doc => {
//   console.log(`Shopping List created ${doc}`);
// });

// await groceriesCollection.doc(uid).update({
//   [item.name]: item
// }).then(doc => {
//   console.log(`Shopping List updated ${doc}`);
// });

          loader.hide();
        })
        .catch(err => {
          console.error(err);
          loader.hide();
          this.alert(err);
        });
    },
    forgotPassword() {
      prompt({
        title: "Forgot Password",
        message:
          "Enter the email address you used to register for APP NAME to reset your password.",
        inputType: "email",
        defaultText: "",
        okButtonText: "Ok",
        cancelButtonText: "Cancel"
      }).then(data => {
        if (data.result) {
          loader.show();
          this.$authService
            .resetPassword(data.text.trim())
            .then(() => {
              loader.hide();
              this.alert(
                "Your password was successfully reset. Please check your email for instructions on choosing a new password."
              );
            })
            .catch(err => {
              loader.hide();
              this.alert(err);
            });
        }
      });
    },
    focusPassword() {
      this.$refs.password.nativeView.focus();
    },
    focusConfirmPassword() {
      if (!this.isLoggingIn) {
        this.$refs.confirmPassword.nativeView.focus();
      }
    },
    alert(message) {
      return alert({
        title: "APP NAME",
        okButtonText: "OK",
        message: message
      });
    }
  }
};
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