<template>
  <div class="hello">
    <h1>{{ name }}</h1>
    <div class="address">address: {{ address }}</div>
    <div class="name">teoname: {{ teoname }}</div>
    <div class="uptime">uptime: {{ uptime }}</div>
    <div class="version">{{ version }}</div>
    <div class="clients">clients: {{ clients }}</div>
    <div v-if="online" class="online text-success">online</div>
    <div v-if="!online" class="offline text-danger">offline</div>
    <button v-if="disconnected" type="button" class="btn btn-warning" @click="reconnect()">Reconnect</button>
    <p>
      <!-- An empty line -->
    </p>

    <hr>
    <p>
      This frontend uses Vue.<br />
      <br />
      For a guide and recipes on how to configure and customize this vue 
      frontend project, check out the
      <a href="https://cli.vuejs.org" target="_blank" rel="noopener"
        >vue-cli documentation</a
      >.
    </p>
    <hr>

    <div v-if="fortune != null">
      <h3>Fortune message from Google Function:</h3>
      <button type="button" class="btn btn-primary" @click="getFortune()">Show next</button>
      <p class="fortune">
        <pre>{{ fortune }}</pre>
        <span>Spent time {{ fortuneTime }} ms</span><br />
      </p>
    </div>
    <hr>

    <div v-if="fortune_api != null">
      <h3>Fortune message from this Application RestAPI:</h3>
      <button type="button" class="btn btn-primary" @click="getFortuneApi()">Show next</button>
      <p class="fortune">
        <pre>{{ fortune_api }}</pre>
        <span>Spent time {{ fortune_apiTime }} ms</span><br />
      </p>
    </div>
    <hr>

    <div v-if="fortune_rtc != null">
      <h3>Fortune message from Teonet Proxy by WebRTC:</h3>
      <button type="button" class="btn btn-primary" @click="getFortuneRTC()">Show next</button>
      <p class="fortune">
        <pre>{{ fortune_rtc }}</pre>
        <span>Spent time {{ fortune_rtcTime }} ms</span><br />
      </p>
    </div>
    <hr>
    <div>
      Based on <a href="https://github.com/teonet-go#teonet-v5" target="_blank"
      rel="noopener">Teonet</a>.<br>
      <br>
      Teonet is designed to create client-server systems and build networks for
      server applications operating within a microservice architecture. To do
      this, Teonet creates a network / cloud transport between its members.
      This transport uses UDP for communication between network peers. UDP
      packets are encrypted with unique keys. Teonet uses its own UDP-based
      protocol called Teonet Reliable UDP (TRU) for real-time communication,
      which allows low latency messages to be sent and protocol reliability
      features.
    </div>
    <hr>
    <div class="copyr text-center text-secondary">
      Teonet © 2024
    </div>

  </div>
</template>

<script>
import { uuid } from "vue-uuid";

const prefix = "/api/v1/";

const cmdClients = "clients";
const cmdList = "list";
const cmdForta = "forta";

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
      online: false,
      disconnected: false,
    };
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

    // Get HTTP Api
    this.getName();
    this.getUptime();
    this.getVersion();
    this.getAddress();
    this.getFortune();
    this.getFortuneApi();

    // Connect to Teonet proxy WebRTC server and get WebRTC Api
    let that = this;
    this.teoweb.connect("wss://signal.teonet.dev/signal", this.teoname, "server-1");
    this.teoweb.onOpen(function () {
      console.debug("onOpen");
      that.online = true;
      that.disconnected = false;
      that.teoweb.sendCmd(cmdClients);
      that.teoweb.subscribeCmd(cmdClients);
      that.teoweb.sendCmd(cmdList);
      that.teoweb.subscribeCmd(cmdList);
      that.getFortuneRTC();
    });
    this.teoweb.onClose(function (b) {
      console.debug("onClose");
      that.online = false;
      if (b === true) {
        that.disconnected = true;
      }
    });
    this.teoweb.addReader(function (gw, data) {
      console.debug(
        "execute reader in 'App.vue' with parameters, gw:",
        gw,
        "data:",
        data
      );
      switch (gw.command) {
        case cmdClients:
          that.clients = data;
          break;
        case cmdList:
          that.list = data;
          break;
        case cmdForta:
          that.fortune_rtc = data;
          that.fortune_rtcTime = new Date().getTime() - that.startRtcTime;
          break;
      }
    });
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
        command: cmdForta,
        // data: null,
      }
      let msg = JSON.stringify(request);
      this.teoweb.send(msg);
    },
    /** Reconnect */
    reconnect() {
      window.location.href = './';
    }
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
a {
  color: #42b983;
}
p.fortune span {
  color: gray;
  font-size: small;
}
p.fortune {
  margin-top: 1rem;
}
.address, .uptime,.version,.name,.clients,.online,.offline {
  margin-left: 6px;
  font-size: small;
}
.address{
  margin-top: 1rem;
}
hr { 
  color: gray; 
}
.copyr {
  margin-bottom: 1rem;
}
</style>
