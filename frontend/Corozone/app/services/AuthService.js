import firebase from "nativescript-plugin-firebase";
import BackendService from "./BackendService";
import { backendService } from "../app";


export default class AuthService extends BackendService {
async login(user) {
     // let that = this

      //loader.show(global.loaderOptions)

      await firebase
        .login({
          type: firebase.LoginType.PASSWORD,
          passwordOptions: {
          email: user.email,
          password: user.password
          }
        })
        .then(async firebaseUser => {
         // alert("Login")
          //this.$navigateTo(App);
          //loader.hide()

          console.log("Firebae UID")
          console.log(firebaseUser.uid)
          //this.user.uid = firebaseUser.uid
          backendService.token = firebaseUser.uid
          return firebaseUser
        })
        // .catch(() => {
        //   loader.hide()
        //   console.error(err);
        //   this.alert("Unfortunately we could not find your account.");
        // });
    }

    async register(user) {
    //   if (this.user.password != this.user.confirmPassword) {
    //     this.alert("Your passwords do not match.");
    //     return;
    //   }
    //   if(this.user.password.length < 6) {
    //     this.alert("Your password must be at least 6 characters.")
    //     return
    //   }

     // loader.show(global.loaderOptions)
      const createdUser = await firebase.createUser({
          email: user.email,
          password: user.password
      });
      return createdUser.uid
    }

   async forgotPassword(email) {
       return await firebase.resetPassword({
           email: email
       })
     // let that = this
    //   prompt({
    //     title: "Forgot Password",
    //     message:
    //       "Enter the email address you used to register for APP NAME to reset your password.",
    //     inputType: "email",
    //     defaultText: "",
    //     okButtonText: "Ok",
    //     cancelButtonText: "Cancel"
    //   }).then(data => {
    //     if (data.result) {
    //       loader.show(global.loaderOptions)
    //       userService
    //         .resetPassword(data.text.trim())
    //         .then(() => {
    //           loader.hide()
    //           this.alert(
    //             "Your password was successfully reset. Please check your email for instructions on choosing a new password."
    //           );
    //         })
    //         .catch(() => {
    //           loader.hide()
    //           this.alert(
    //             "Unfortunately, an error occurred resetting your password."
    //           );
    //         });
    //     }
    //   });
    }
    async logout() {
        backendService.token = "";
        return await firebase.logout();
    }
}