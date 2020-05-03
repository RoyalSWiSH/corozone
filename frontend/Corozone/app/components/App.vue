 <template lang="html">
 <Page actionBarHidden="true" backgroundSpanUnderStatusBar="true">
    <BottomNavigation selectedIndex="1">

        <TabStrip selectedItemColor="#3d99ff" unSelectedItemColor="#5c5c5c">
            <TabStripItem class="navigation__item">
                <Label text="Helping"></Label>
                <Image src.decode="font://&#xf0c0;" class="fas t-36"></Image>
            </TabStripItem>
             <TabStripItem class="navigation__item">
                <Label text="Groceries"></Label>
                <Image src.decode="font://&#xf290;" class="fas t-36"></Image>
            </TabStripItem>
            <!-- <TabStripItem class="navigation__item">
                <Label text="Groceries"></Label>
                <Image src.decode="font://&#xf218;" class="far t-36"></Image>
            </TabStripItem> -->
            <TabStripItem class="navigation__item">
                <Label text="Symptoms"></Label>
                <Image src.decode="font://&#xf21e;" class="fas t-36"></Image>
            </TabStripItem>
            <TabStripItem class="navigation__item">
                <Label text="Contacts"></Label>
                <Image src.decode="font://&#xf2a1;" class="fas t-36"></Image>
            </TabStripItem>
            <!-- <TabStripItem class="navigation__item">
                <Label text="Hospitals"></Label>
                <Image src.decode="font://&#xf0f8;" class="fas t-36"></Image>
            </TabStripItem> -->
            <TabStripItem class="navigation__item">
                <Label text="Privacy"></Label>
                <Image src.decode="font://&#xf023;" class="fas t-36"></Image>
            </TabStripItem>
        </TabStrip>

        <TabContentItem>
            <Frame ~itemsFrame id="itemsFrame">
                <Items />
            </Frame>
        </TabContentItem>

        <TabContentItem>
            <Frame>
                <Browse />
            </Frame>
        </TabContentItem>

        <TabContentItem>
            <Frame>
                <Symptoms />
            </Frame>
        </TabContentItem>

        <TabContentItem>
            <Frame>
                <Contacts />
            </Frame>
        </TabContentItem>
        <!-- <TabContentItem>
            <Frame>
                <Hospitals />
            </Frame>
        </TabContentItem> -->
        <TabContentItem>
            <Frame>
                <Privacy />
            </Frame>
        </TabContentItem>

    </BottomNavigation>
    </Page>
</template> 

<!--<template lang="html">
<Page>
        <Label class="m-20" textWrap="true" text="You have successfully authenticated. This is where you build your core application functionality."></Label>
</Page>
</template>-->

<script>
import { Frame } from 'tns-core-modules/ui/frame'
import { mapState } from "vuex";
import Items from "./Items.vue";
import Browse from "./Browse.vue";
import Search from "./Search.vue";
import LoginPage from "./LoginPage.vue";

import Contacts from "./Contacts"
import Hospitals from "./Hospitals"
import Privacy from "./Privacy"
import Symptoms from "./Symptoms"


export default {
  name: "app",
  components: {
    Items,
    Browse,
    Search,
    Privacy,
    Symptoms,
    Hospitals,
    Contacts,
    LoginPage
  },
  computed: {
    ...mapState(["isLoggedIn"])    
  },
 watch: {
     // Logout leads to strange behaviour like details page not working
    isLoggedIn(val) {
      if (!val) {
          console.log("Watch is LoggedIn")
            this.$navigateTo(LoginPage, {  clearHistory: true, frame: "mainFrame" });
         // Frame.topmost().currentPage.$navigateTo(LoginPage)
       //   Frame.topmost().navigate(LoginPage)
        // this.$navigateTo(LoginPage, {   
        //      frame: "mainFrame" 
        // //clearHistory: true 
        // });
      }
    }
  } 
};
</script>

<style lang="scss">
// Start custom common variables
@import "~@nativescript/theme/scss/variables/blue";
// End custom common variables

// Custom styles

</style>
