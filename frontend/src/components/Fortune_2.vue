<template>
  <div class="hello">
    <h1>{{ name }}</h1>
    <div class="address">address: {{ address }}</div>
    <div class="name">teoname: {{ $root.getTeoname() }}</div>
    <div class="uptime">uptime: {{ uptime }}</div>
    <div class="version">{{ version }}</div>
    <p>
      For a guide and recipes on how to configure / customize this project,
      check out the
      <a href="https://cli.vuejs.org" target="_blank" rel="noopener"
        >vue-cli documentation</a
      >.
    </p>

    <div v-if="fortune != null">
      <h3>Fortune message from Google Function:</h3>
      <p class="fortune">
        <pre>{{ fortune }}</pre>
        <span>Spent time {{ fortuneTime }} ms</span><br />
        <br />
        <button type="button" @click="getFortune()">Show next</button>
      </p>
    </div>

    <div v-if="fortune_api != null">
      <h3>Fortune message from this Application RestAPI:</h3>
      <p class="fortune">
        <pre>{{ fortune_api }}</pre>
        <span>Spent time {{ fortune_apiTime }} ms</span><br />
        <br />
        <button type="button" @click="getFortuneApi()">Show next</button>
      </p>
    </div>

    <div v-if="fortune_rtc != null">
      <h3>Fortune message from Teonet Proxy by WebRTC:</h3>
      <p class="fortune">
        <pre>{{ fortune_rtc }}</pre>
        <span>Spent time {{ fortune_rtcTime }} ms</span><br />
        <br />
        <button type="button" @click="getFortuneRTC()">Show next</button>
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
      uptime: null,
      version: null,
      address: null,
      fortune: null,
      fortune_api: null,
      fortune_rtc: null,
      rtc_id: 0,
    };
  },
  mounted: function () {

    this.getName();
    this.getUptime();
    this.getVersion();
    this.getAddress();
    this.getFortune();
    this.getFortuneApi();

    let that = this;
    this.teoweb.onconnected = (_, dc) => {
      dc.onopen = that.getFortuneRTC;
      dc.onmessage = (ev) => {
        // The ev.data got bytes array, so convert it to string and pare to
        // gw object. Then base64 decode gw.data to string
        let enc = new TextDecoder("utf-8");
        let msg = enc.decode(ev.data);
        console.debug("dc got answer:", msg);
        let gw = JSON.parse(msg);
        that.fortune_rtc = atob(gw.data);
        that.fortune_rtcTime = new Date().getTime() - that.startRtcTime;
      }
    };
  },
  methods: {
    getName() {
      this.axios
        .get("/api/v1/name")
        .then((response) => (this.name = response.data));
    },
    getUptime() {
      this.axios
        .get("/api/v1/uptime")
        .then((response) => (this.uptime = response.data));
    },
    getVersion() {
      this.axios
        .get("/api/v1/version")
        .then((response) => (this.version = "ver. " + response.data));
    },
    getAddress() {
      this.axios
        .get("/api/v1/address")
        .then((response) => (this.address = response.data));
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
    getFortuneRTC() {
      this.startRtcTime = new Date().getTime();
      this.fortune_rtc = "loading ...";
      let request = {
        id: this.rtc_id++,
        address: "8agv3IrXQk7INHy5rVlbCxMWVmOOCoQgZBF", 
        command: "forta", 
        // data: null,
      }
      let msg = JSON.stringify(request);
      this.teoweb.send(msg);
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
.address {
  margin-top: -20px;
  margin-left: 6px;
  font-size: small;
}
.uptime,.version,.name {
  margin-left: 6px;
  font-size: small;
}
</style>
