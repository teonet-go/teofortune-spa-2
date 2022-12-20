<template>
  <div class="hello">
    <h1>{{ name }}</h1>
    <div class="address">address: {{ address }}</div>
    <div class="name">teoname: {{ $root.getTeoname() }}</div>
    <div class="uptime">uptime: {{ uptime }}</div>
    <div class="version">{{ version }}</div>
    <div class="clients">clients: {{ clients }}</div>
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

const prefix = "/api/v1/";

const cmdSubscribe = "subscribe";
const cmdClients = "clients";
const cmdList = "list";

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
      clients: "0",
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

    this.teoweb.onconnected = (_, dc) => {
      dc.onopen = () => {
        this.getCommand(cmdList);
        this.getCommand(cmdClients);
        this.getFortuneRTC();
        this.subscribeCommand(cmdList);
        this.subscribeCommand(cmdClients);
      }
      dc.onmessage = (ev) => {
        // The ev.data got bytes array, so convert it to string and pare to
        // gw object. Then base64 decode gw.data to string
        let enc = new TextDecoder("utf-8");
        let msg = enc.decode(ev.data);
        console.debug("dc got answer:", msg);
        let gw = JSON.parse(msg);
        // Teonet proxy responce
        if (gw.address.length) {
          this.fortune_rtc = atob(gw.data);
          this.fortune_rtcTime = new Date().getTime() - this.startRtcTime;
          return;
        }
        // WebRTC server responce
        switch(gw.command) {
        case cmdClients:
          this.clients = atob(gw.data);
          break;
        case cmdList: {
          let listStr = atob(gw.data);
          let list = JSON.parse(listStr);
          console.debug("clients list:", list);
          break;
        }
        }
      }
    };
  },
  methods: {
    getName() {
      this.axios
        .get(prefix + "name")
        .then((response) => (this.name = response.data));
    },
    getUptime() {
      this.axios
        .get(prefix + "uptime")
        .then((response) => (this.uptime = response.data));
    },
    getVersion() {
      this.axios
        .get(prefix + "version")
        .then((response) => (this.version = "ver. " + response.data));
    },
    getAddress() {
      this.axios
        .get(prefix + "address")
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
        .get(prefix + "fortune")
        .then((response) => {
          this.fortune_api = response.data;
          this.fortune_apiTime = new Date().getTime() - startTime;
        });
    },
    /** Get fortune message from WebRTC proxy server */
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
    /** Send request with command to WebRTC server */
    getCommand(cmd, cmdData) {
      let data = null;
      if (cmdData) {
        data = btoa(cmdData);
      }
      let request = {
        id: this.rtc_id++,
        address: "",
        command: cmd,
        data: data,
      }
      let msg = JSON.stringify(request);
      this.teoweb.send(msg);
    },
    /** Send subscribe request with command to WebRTC server */
    subscribeCommand(cmd) {
      this.getCommand(cmdSubscribe, cmd);
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
.uptime,.version,.name,.clients {
  margin-left: 6px;
  font-size: small;
}
</style>
