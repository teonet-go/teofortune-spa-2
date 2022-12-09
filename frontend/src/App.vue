<template>
  <img alt="Vue logo" src="./assets/logo.png" />
  <Fortune_2 msg="Welcome to Teonet Fortune-2 Vue.js App" />
</template>

<script>
import Fortune_2 from "./components/Fortune_2.vue";

export default {
  name: "App",
  components: {
    Fortune_2,
  },
  mounted: function () {
    // Connect to Teonet proxy WebRTC server
    this.teoweb.connect(
      "wss://signal.teonet.dev/signal",
      "teofortune-spa-2",
      "server-1"
    );
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /* text-align: center; */
  color: #2c3e50;
  /* margin-top: 60px; */
}
img {
  width: 100px;
}
</style>
