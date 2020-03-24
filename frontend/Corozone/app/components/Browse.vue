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
					 returnKeyType="next" @returnPress="code" fontSize="18" />
					<StackLayout class="hr-light" />
					<Label text="Code" class="field-title" fontSize="19"/>
					<TextField ref="code" class="input" hint="plzt" keyboardType="plz" autocorrect="false" autocapitalizationType="none" v-model="adress.plz"
					 returnKeyType="next" @returnPress="street" fontSize="18" />
					<Label text="City" class="field-title" fontSize="19"/>
					<TextField ref="street" class="input" hint="city" keyboardType="street" autocorrect="false" autocapitalizationType="none" v-model="adress.city"
					 returnKeyType="next" @returnPress="focusPassword" fontSize="18" />
					<Label text="Items" class="field-title" fontSize="19"/>
					<TextField ref="items" class="input" hint="items" keyboardType="street" autocorrect="false" autocapitalizationType="none" v-model="adress.items"
					 returnKeyType="next" @returnPress="focusPassword" fontSize="18" />
				<Button text="Request Groceries" @tap="requestGroceries" class="btn" />
				</StackLayout>
		
			</StackLayout>

	     		</FlexboxLayout>
    </Page>
</template>

<script>
import firebase from "nativescript-plugin-firebase";
import axios from "axios/dist/axios"
// A stub for a service that authenticates users.


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
		items: ""
      }
    };
  },
  methods: {
    toggleForm() {
      this.isLoggingIn = !this.isLoggingIn;
    },
    requestGroceries() {
      if (!this.adress.street || !this.adress.plz || !this.adress.city) {
        this.alert("Please provide city, street and adress.");
        return;
      }
      axios.get('http://webcode.me').then(resp => {
        console.log(resp.data);
      });
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