<template>
	<StackLayout class="p-20" backgroundColor="white">
        
		<Label class="h2 text-center" text="Scan this Code to add a Friend." />

<BarcodeScanner
    row="1"
    height="200"
    width="200"
    formats="QR_CODE, EAN_13, UPC_A"
    beepOnScan="true"
    reportDuplicates="true"
    preferFrontCamera="false"
    :pause="pause"
    @scanResult="onScanResult"
    v-if="this.$isIOS">
</BarcodeScanner>
    
		<Button class="btn btn-outline" text="Open Camera" @tap="$modal.close()" v-if="this.$isAndroid"/>
        <Image :src="friendQRCodeURL" height="200"></Image>
		<Label class="h4 text-center" :text="this.generateCode" textWrap="true" v-if="addUserIDByHand"/>
		<!-- <Label class="h2 text-center" :text="me.firstName"/> -->

        <!-- <TextField
              v-model="me.firstName"
              col="0"
              row="0"
              hint="First Name"
              editable="true"
              @returnPress="onNameTap"
            />-->
        
            <Button class="btn btn-outline" text="Copy Code to Clipboard" @tap="" />
               <TextField
              col="0"
              row="0"
              hint="Paste code to add friend"
              editable="false"
              @returnPress="onNameTap"
              v-if="addUserIDByHand"
            />
            <Label class="h4 text-center" text="Why do we use QR Codes? We don't want other people have access to your personalized information
            without your permission and don't use our severs if we don't have to. We constantly balance privacy, security and convenience so this may change in the future." textWrap="true"/>
             <Button class="btn btn-outline" text="Add friend" @tap="" v-if="addUserIDByHand" />
		<Button class="btn btn-outline" text="Close" @tap="$modal.close()" />
	</StackLayout>
</template>

<script>
import { backendService } from '../app';
import { mapState, mapGetters } from "vuex";
export default {
    data() {
        return {
            friendQRCodeURL: "",
            addUserIDByHand: false
        };
    },
    created(){
        // Get name from server if not stored locally
        if(!this.me || this.me.user_id != this.$backendService.token) {
            console.log(this.$backendService.token)

        this.axios.get("/users/"+this.$backendService.token).then(me => {

        console.log(me)
        this.$store.commit("setMyUserProfile", me.data)
        })
        
        }

        this.friendQRCodeURL = this.getFriendQRCode()
    },
    computed: {
         ...mapState(["me"]),
         // TODO: Add Hash to check data integrity
         generateCode: function () {return this.$backendService.token+":"+this.me.firstName  }
    },
    methods: {
        getFriendQRCode() {
            return "https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=app://corozone/"+this.generateCode
        },
        onScanResult(scanEvent) {
            console.log(`Scan Result: ${scanEvent.text} (${scanEvent.format})`)
            const friend = {
                uid: scanEvent.text.split("corozone/")[1].split(":")[0],
                name: scanEvent.text.split("corozone/")[1].split(":")[1]
            }
            if(this.$store.getters.getFriendListIDs.some(e => e.uid === friend.uid)) {
      this.$store.commit("setNotificationMessage", "Friend already exists") 
    }
    else {
    this.$store.commit("addFriendID", friend)
      this.$store.commit("setNotificationMessage", "Friend " + friend.name + " added")
      $modal.close() 
    }
           
        },
        // generateCode() {
        //    return this.$backendService.token+":"+this.me.firstName 
        // }
    }
};
</script>

<style>
</style>