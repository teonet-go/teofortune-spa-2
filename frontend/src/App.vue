<template>
  <div>
    <!-- Logo -->
    <!-- <a href="./">
      <img alt="Vue logo" src="./assets/logo.png" />
    </a> -->
    <Fortune_2 msg="Welcome to Teonet Fortune-2 Vue.js App" />
  </div>
</template>

<script>
import Fortune_2 from "./components/Fortune_2.vue";

export default {
  name: "App",
  components: {
    Fortune_2,
  },
  mounted: function () {
  },
};
</script>

<style>
#app { }
/* Telegram collors */
.btn-primary, .btn-primary:hover  {
  background-color: #5288C1 !important;
  border-color:  #5288C1 !important;
}
</style>
