<template>
 <Page>
     <!-- <ScrollView>
     <StackLayout> -->
         <!-- <CardView class="card"> -->
                 <Label text="Location: Frankfur" />
         <!-- </CardView> -->
     <!-- </StackLayout>
     </ScrollView> -->
</Page>
</template>

<script>
import ItemDetails from "./ItemDetails";

export default {
    data: () => {
        return {
            items: [
                {
                    name: "Item 1",
                    description: "Description for Item 1"
                },
                {
                    name: "Item 2",
                    description: "Description for Item 2"
                },
                {
                    name: "Item 3",
                    description: "Description for Item 3"
                },
                {
                    name: "Item 4",
                    description: "Description for Item 4"
                },
                {
                    name: "Item 5",
                    description: "Description for Item 5"
                },
                {
                    name: "Item 6",
                    description: "Description for Item 6"
                },
                {
                    name: "Item 7",
                    description: "Description for Item 7"
                },
                {
                    name: "Item 8",
                    description: "Description for Item 8"
                },
                {
                    name: "Item 9",
                    description: "Description for Item 9"
                },
                {
                    name: "Item 10",
                    description: "Description for Item 10"
                },
                {
                    name: "Item 11",
                    description: "Description for Item 11"
                },
                {
                    name: "Item 12",
                    description: "Description for Item 12"
                },
                {
                    name: "Item 13",
                    description: "Description for Item 13"
                },
                {
                    name: "Item 14",
                    description: "Description for Item 14"
                },
                {
                    name: "Item 15",
                    description: "Description for Item 15"
                },
                {
                    name: "Item 16",
                    description: "Description for Item 16"
                }
            ]};
    },
    methods: {
        onItemTap (args) {
            const view = args.view;
            const page = view.page;
            const tappedItem = view.bindingContext;

            this.$navigateTo(ItemDetails, {
                props: { 
                    context: tappedItem,
                    animated: true,
                    transition: {
                        name: "slide",
                        duration: 200,
                        curve: "ease"
                    }}});
        }
    }
};
</script>

<style scoped lang="scss">
// Start custom common variables
@import "~@nativescript/theme/scss/variables/blue";
// End custom common variables

// Custom styles
.add-button {
    height: 30;
    background-color: rgb(51, 51, 206);
    color: white;
    border-radius: 5;
    font-size: 20;
    font-weight: 600;
}

.card {
    background-color: #fff;
    color: #4d4d4d;
    margin: 15;
}

.card-layout {
    padding: 20;
}
</style>