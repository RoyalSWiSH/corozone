<template lang="html">
  <Page>
    <ActionBar>
      <Label text="Health"></Label>
    </ActionBar>

    <GridLayout class="page__content">
      <Label class="page__content-icon fas" Hello"></Label>
      <Label class="page__content-placeholder" :text="message"></Label>
    </GridLayout>
  </Page>
</template>

<script>
export default {
  data: () => {
    return {
      message: "<!-- Health page content goes here -->"
    };
  }
};
</script>

<style lang="scss" scoped>
// Start custom common variables
@import "~@nativescript/theme/scss/variables/blue";
// End custom common variables

// Custom styles
</style>
