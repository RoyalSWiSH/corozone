<template>
<!-- <frame ~detailView id="detailView"> -->
<page>
        <ActionBar>
            <NavigationButton @tap="gohereBack()" android.systemIcon="ic_menu_back"/>
            <!-- <Label :text="item.name"></Label> -->
        </ActionBar>
    <StackLayout>
        <Label text="DetailView" margin="25" fontSize="25"></Label>
        <!-- <NavigationButton @tap="goBack" android.systemIcon="ic_menu_back"/> -->
        <!-- <GridLayout> -->
            <!-- <Label class="m-10 h3" :text="item.items[0]" verticalAlignment="top" ></Label> -->
        <!-- </GridLayout> -->
    </StackLayout>
</page>
<!-- </frame> -->
</template>

<script>

export default {
    props: ["context"],

    computed: {
        // item() {
        //     return this.context || {};
        // }
          },
    methods: {
          gohereBack(args) {
           //const getFrameById = require("tns-core-modules/ui/frame").getFrameById
          // const myFrame = getFrameById("itemsFrame");
          this.$navigateBack();
        }

    }
};
</script>

<style scoped lang="scss">
// Start custom common variables
@import "~@nativescript/theme/scss/variables/blue";
// End custom common variables

// Custom styles

</style>