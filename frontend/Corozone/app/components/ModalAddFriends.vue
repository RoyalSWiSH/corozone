<template>
	<StackLayout class="p-20" backgroundColor="white">
        
		<Label class="h2 text-center" text="Scan this Code to add a Friend." />

        <Image :src="friendQRCodeURL" ></Image>
		<Label class="h2 text-center" :text="$backendService.token"/>
		<Label class="h2 text-center" :text="me.firstName"/>

        <TextField
              v-model="me.firstName"
              col="0"
              row="0"
              hint="First Name"
              editable="true"
              @returnPress="onNameTap"
            />
		<Button class="btn btn-outline" text="Close Modal" @tap="$modal.close()" />
	</StackLayout>
</template>

<script>
import { backendService } from '../app';
import { mapState, mapGetters } from "vuex";
export default {
    data() {
        return {
            friendQRCodeURL: ""
        };
    },
    created(){
        this.friendQRCodeURL = this.getFriendQRCode()
        // Get name from server if not stored locally
        //if(this.me.firstName == "") {
            console.log(this.$backendService.token)

        this.axios.get("/users/"+this.$backendService.token).then(me => {

        console.log(me)
        this.$store.commit("setMyUserProfile", me.data)
        })
        
       // }
    },
    computed: {
         ...mapState(["me"]),
    },
    methods: {
        getFriendQRCode() {
            return "https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=app://corozone/"+this.$backendService.token+":"+this.me.firstName
        }
    }
};
</script>

<style>
</style>