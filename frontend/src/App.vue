<template>
  <img alt="Vue logo" src="./assets/logo.png" />
  <Fortune_2 msg="Welcome to Teonet Fortune-2 Vue.js App" />
</template>

<script>
import Fortune_2 from "./components/Fortune_2.vue";
import { uuid } from "vue-uuid";

export default {
  name: "App",
  components: {
    Fortune_2,
  },
  mounted: function () {
    // Get this browser name from local storage
    if (localStorage.teoname) {
      this.teoname = localStorage.teoname;
      console.debug("teoname:", this.teoname);
    } else {
      this.teoname = uuid.v1();
      localStorage.teoname = this.teoname;
      console.debug("teoname does not exists, new name created:", this.teoname);
    }

    // Connect to Teonet proxy WebRTC server
    this.teoweb.connect(
      "wss://signal.teonet.dev/signal",
      this.teoname,
      "server-1"
    );
  },
  methods: {
    getTeoname: function () {
      return this.teoname;
    },
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
