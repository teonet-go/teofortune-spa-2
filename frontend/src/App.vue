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
#app {
  /* font-family: Avenir, Helvetica, Arial, sans-serif; */
  /* -webkit-font-smoothing: antialiased; */
  /* -moz-osx-font-smoothing: grayscale; */

  /* text-align: center; */
  /* color: #2c3e50; */
  /* margin-top: 60px; */
}
img {
  width: 100px;
}
</style>
