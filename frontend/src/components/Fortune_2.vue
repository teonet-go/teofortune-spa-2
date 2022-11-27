<template>
  <div class="hello">
    <h1>{{ name }}</h1>
    <div class="version">{{ version }}</div>
    <p>
      For a guide and recipes on how to configure / customize this project,
      check out the
      <a href="https://cli.vuejs.org" target="_blank" rel="noopener"
        >vue-cli documentation</a
      >.
    </p>

    <div v-if="fortune != null">
      <h3>Fortune function message:</h3>
      <p class="fortune">
        <pre>{{ fortune }}</pre>
        <span>Spent time {{ fortuneTime }} ms</span><br />
        <br />
        <button type="button" @click="getFortune()">Show next</button>
      </p>
    </div>

    <div v-if="fortune_api != null">
      <h3>Fortune this application api message:</h3>
      <p class="fortune">
        <pre>{{ fortune_api }}</pre>
        <span>Spent time {{ fortune_apiTime }} ms</span><br />
        <br />
        <button type="button" @click="getFortuneApi()">Show next</button>
      </p>
    </div>

  </div>
</template>

<script>
export default {
  name: "Fortune_2",
  props: {
    msg: String,
  },
  data() {
    return {
      name: null,
      version: null,
      fortune: null,
      fortune_api: null,
    };
  },
  mounted: function () {
    this.getName();
    this.getVersion();
    this.getFortune();
    this.getFortuneApi();
  },
  methods: {
    getName() {
      this.axios
        .get("/api/v1/name")
        .then((response) => (this.name = response.data));
    },
    getVersion() {
      this.axios
        .get("/api/v1/version")
        .then((response) => (this.version = "ver. " + response.data));
    },
    getFortune() {
      const startTime = new Date().getTime();
      this.fortune = "loading ...";
      this.axios
        .get("https://fortune-2.teonet.app")
        .then((response) => {
          this.fortune = response.data; 
          this.fortuneTime = new Date().getTime() - startTime; 
        });
    },
    getFortuneApi() {
      const startTime = new Date().getTime();
      this.fortune_api = "loading ...";
      this.axios
        .get("/api/v1/fortune")
        .then((response) => {
          this.fortune_api = response.data;
          this.fortune_apiTime = new Date().getTime() - startTime; 
        });
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
p.fortune span {
  color: gray;
  font-size: small;
}
.version {
  margin-top: -20px;
  margin-left: 6px;
  font-size: small;
}
</style>
