<template>
	<StackLayout class="p-20" backgroundColor="white">
        
		<Label class="h2 text-center" text="Scan this Code to add a Friend." />

  
		<Button class="btn btn-outline" text="Open Camera" @tap="scan(false)" />
      

        <TextField
              v-model="receiptCosts"
              col="0"
              row="0"
              hint="z.B. 34"
              editable="true"
              @returnPress="onNameTap"
            />
                     
            <Button class="btn btn-outline" text="Send" @tap="onTapSend"  />
		<Button class="btn btn-outline" text="Close" @tap="$modal.close()" />
	</StackLayout>
</template>

<script>
import { backendService } from '../app';
import { mapState, mapGetters } from "vuex";
import {BarcodeScanner} from "nativescript-barcodescanner";

export default {
    data() {
        return {
            friendQRCodeURL: "",
            addUserIDByHand: false,
            receiptCosts: 0
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
      onTapSend(args) {
          console.log(args)
      },
        getFriendQRCode() {
            return "https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=app://corozone/"+this.generateCode
        },
        onScanResult(scanEvent) {
            //TODO Check if format is QR Code
            console.log(`Scan Result: ${scanEvent.text} (${scanEvent.format})`)
              this.addFriend(scanEvent)            
        },
         subscribeToShoppingList(id) {
 if(id){    
const db = firebase.firestore
const groceriesCollection = db.collection("Groceries");
console.log("Subscribe to shopping list: " + id)
       const unsubscribe2 = groceriesCollection.doc(id).onSnapshot(doc2 => {
  console.log(doc2.data())
  console.log("Subscribed2")

  console.log(Object.values(doc2.data()))
  this.$store.commit("mergeShoppingList", Object.values(doc2.data()) )
});
  return unsubscribe2
  }
  else {
    console.log("Unable to subscribe. Probably no or wrong id set")
  }
    },
        addFriend(scanEvent) {
   const friend = {
                id: scanEvent.text.split("corozone/")[1].split(":")[0],
                name: scanEvent.text.split("corozone/")[1].split(":")[1]
            }
            if(this.$store.getters.getFriendListIDs.some(e => e.id === friend.id)) {
      this.$store.commit("setNotificationMessage", `You share groceries with ${friend.name} already.`) 
  //    this.$modal.close() 
      return
    }
    else {
    this.$store.commit("addFriendID", friend)
      this.$store.commit("setNotificationMessage", "You can now see groceries from " + friend.name + ".")
   //   this.$modal.close() 
      return friend
    }
    this.subscribeToShoppingList(friend.id)
        },
        scan(front) {
        // so this is available when addFriend is called 
        let that = this
        new BarcodeScanner().scan({
          cancelLabel: "EXIT. Also, try the volume buttons!", // iOS only, default 'Close'
          cancelLabelBackgroundColor: "#333333", // iOS only, default '#000000' (black)
          message: "Use the volume buttons for extra light", // Android only, default is 'Place a barcode inside the viewfinder rectangle to scan it.'
          preferFrontCamera: front,     // Android only, default false
          showFlipCameraButton: true,   // default false
          showTorchButton: true,        // iOS only, default false
          torchOn: false,               // launch with the flashlight on (default false)
          resultDisplayDuration: 500,   // Android only, default 1500 (ms), set to 0 to disable echoing the scanned text
          beepOnScan: true,             // Play or Suppress beep on scan (default true)
          openSettingsIfPermissionWasPreviouslyDenied: true, // On iOS you can send the user to the settings app if access was previously denied
          closeCallback: () => {
            console.log("Scanner closed @ " + new Date().getTime());
          }
        }).then(
            function (result) {
              console.log("--- scanned: " + result.text);
              // Note that this Promise is never invoked when a 'continuousScanCallback' function is provided
              setTimeout(async function () {
                // if this alert doesn't show up please upgrade to {N} 2.4.0+
                let friend = await that.addFriend(result)
                alert({
                  title: `${friend.name} hinzugefügt`,
                  message: "Format: " + result.format + ",\nValue: " + result.text,
                  okButtonText: "OK"
                });
              }, 500);
            },
            function (errorMessage) {
              console.log("No scan. " + errorMessage);
            }
        );
      }
    }
};
</script>

<style>
</style>