 <template lang="html">
 <Page actionBarHidden="true" backgroundSpanUnderStatusBar="true">
    <BottomNavigation>
        <TabStrip>
            <TabStripItem class="navigation__item">
                <Label text="Home"></Label>
                <Image src.decode="font://&#xf015;" class="fas t-36"></Image>
            </TabStripItem>
            <TabStripItem class="navigation__item">
                <Label text="Browse"></Label>
                <Image src.decode="font://&#xf1ea;" class="far t-36"></Image>
            </TabStripItem>
            <TabStripItem class="navigation__item">
                <Label text="Search"></Label>
                <Image src.decode="font://&#xf002;" class="fas t-36"></Image>
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
                <Search />
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
import Items from "./Items.vue";
import Browse from "./Browse.vue";
import Search from "./Search.vue";

export default {
  name: "app",
  components: {
    Items,
    Browse,
    Search
  }
};
</script>

<style lang="scss">
// Start custom common variables
@import "~@nativescript/theme/scss/variables/blue";
// End custom common variables

// Custom styles

</style>
