<template lang="html">
    <Page>
        <ActionBar>
            <Label text="Corozone"></Label>
        </ActionBar>
      <FlexboxLayout class="page">
			<StackLayout class="form">
				<Image class="logo" src="~/images/logo.png" />
				<Label class="header" text="Corozone" />

				<StackLayout class="input-field" marginBottom="25">
					<TextField class="input" hint="Email" keyboardType="email" autocorrect="false" autocapitalizationType="none" v-model="user.email"
					 returnKeyType="next" @returnPress="focusPassword" fontSize="18" />
					<StackLayout class="hr-light" />
				</StackLayout>

				<StackLayout class="input-field" marginBottom="25">
					<TextField ref="password" class="input" hint="Password" secure="true" v-model="user.password" :returnKeyType="isLoggingIn ? 'done' : 'next'"
					 @returnPress="focusConfirmPassword" fontSize="18" />
					<StackLayout class="hr-light" />
				</StackLayout>

				<StackLayout v-show="!isLoggingIn" class="input-field">
					<TextField ref="confirmPassword" class="input" hint="Confirm password" secure="true" v-model="user.confirmPassword" returnKeyType="done"
					 fontSize="18" />
					<StackLayout class="hr-light" />
				</StackLayout>

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
import firebase from "nativescript-plugin-firebase";
import App from "./App"
var LoadingIndicator = require("@nstudio/nativescript-loading-indicator")
    .LoadingIndicator;
var loader = new LoadingIndicator();
// A stub for a service that authenticates users.
const userService = {
  async register(user) {
    return await firebase.createUser({
      email: user.email,
      password: user.password
    });
  },
  async login(user) {
    return await firebase.login({
        type: firebase.LoginType.PASSWORD,
        passwordOptions: {
          email: user.email,
          password: user.password
        }
    });
  },
  async resetPassword(email) {
    return await firebase.resetPassword({ 
      email: email
      });
}
}
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
      user: {
        email: "foo@foo.com",
        password: "barbar",
        confirmPassword: "barbar"
      }
    };
  },
  methods: {
    toggleForm() {
      this.isLoggingIn = !this.isLoggingIn;
    },
    submit() {
      if (!this.user.email || !this.user.password) {
        this.alert("Please provide both an email address and password.");
        return;
      }
      if (this.isLoggingIn) {
        this.login();
      } else {
        this.register();
      }
    },
    login() {
      let that = this
      loader.show(global.loaderOptions)
      userService
        .login(this.user)
        .then(() => {
         // alert("Login")
          //this.$navigateTo(App);
          loader.hide()
          this.$navigateTo(App, {clearHistory: true}); //works fine, but not to TabView App.vue :(
        })
        .catch(() => {
          loader.hide()
          console.error(err);
          this.alert("Unfortunately we could not find your account.");
        });
    },
    register() {
      if (this.user.password != this.user.confirmPassword) {
        this.alert("Your passwords do not match.");
        return;
      }
      if(this.user.password.length < 6) {
        this.alert("Your password must be at least 6 characters.")
        return
      }

      loader.show(global.loaderOptions)
      this.userService
        .register(this.user)
        .then(() => {
          loader.hide()
          this.alert("Your account was successfully created.");
          this.isLoggingIn = true;
        })
        .catch(() => {
          loader.hide()
          console.error(err);
          this.alert(err);
        });
    },
    forgotPassword() {
      let that = this
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
          loader.show(global.loaderOptions)
          userService
            .resetPassword(data.text.trim())
            .then(() => {
              loader.hide()
              this.alert(
                "Your password was successfully reset. Please check your email for instructions on choosing a new password."
              );
            })
            .catch(() => {
              loader.hide()
              this.alert(
                "Unfortunately, an error occurred resetting your password."
              );
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