<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
      <v-tabs
          background-color="primary"
          center-active
          dark
          icons-and-text
          v-model="tab"
          show-arrows
      >
        <v-tab href="#info">
          Info
          <v-icon>md-info</v-icon>
        </v-tab>
        <v-tab
            v-for="(receiver, index) in receivers"
            v-bind:href="'#' + index"
            v-bind:key="index"
        >
          {{ receiverNames[index] }}
          <v-icon>{{receiverIcon(receiver.state)}}</v-icon>
        </v-tab>
      </v-tabs>
      <v-btn
          icon
          small
          @click="$vuetify.theme.dark = !$vuetify.theme.dark"
      >
        <v-icon>mdi-brightness-6</v-icon>
      </v-btn>
    </v-app-bar>
    <v-main>
      <v-tabs-items v-model="tab">
        <v-tab-item value="info">
          <Guide/>
        </v-tab-item>
        <v-tab-item
            v-for="(receiver, index) in receivers"
            :key="index"
            :value="index"
        >
          <ReceiverCard :receiver="receiver" :config="getConfig(index)" :logMessages="logMessages[receiver.index]"/>
        </v-tab-item>
      </v-tabs-items>
    </v-main>
  </v-app>
</template>

<script>
import ReceiverCard from './components/ReceiverCard';
import GuideComponent from "@/components/GuideComponent";

export default {
  name: 'App',

  components: {
    ReceiverCard: ReceiverCard,
    Guide: GuideComponent,
  },

  computed: {
    tab: {
      set (tab) {
        this.$router.replace({ query: { ...this.$route.query, tab } })
      },
      get () {
        return this.$route.query.tab
      }
    }
  },

  data: () => ({
    receivers: [],
    error: "",
    config: {},
    logMessages: [],
    receiverNames: [
        "2m",
        "70cm 436.00",
        "70cm 437.00",
        "23cm",
    ],
  }),
  methods: {
    getConfig: function(index) {
      try {
        return this.config.receivers[index]
      } catch (e) {
        return null
      }
    },

    receiverIcon: function (status) {
      if (status === "lost") { return "mdi-wifi-strength-3-alert" }
      if (status === "header") return "mdi-wifi"
      if (["DVB-S2", "DVB-S"].indexOf(status) !== -1) {
        return "mdi-wifi-check"
      }
      return "mdi-wifi-strength-outline"
    },
    updateConfig() {
      fetch("api/config")
          .then((response) => {
            if (response.ok) {
              this.error = ""
              return response.json();
            } else {
              this.error = "failed to fetch config"
              return {}
            }
          })
          .then(data => {
            this.config = data
          })
    },
    updateData() {
      fetch("api/status")
          .then((response) => {
            if (response.ok) {
              this.error = ""
              return response.json();
            } else {
              this.error = "failed to fetch receivers statuses"
              return {}
            }
          })
          .then(data => {
            for (let index=0; index<4; index++) {
              data[index].carrier_frequency = [144.6,436,437,1245][index]
            }
            this.receivers = data
          })
    },
    connectLog() {
      console.log("connecting to websocket")
      this.logMessages = [] // reset
      for (let index=1; index<=4; index++) {
        this.logMessages[index] = []
      }
      this.connection = new WebSocket(((location.protocol === "https:") ? "wss://" : "ws://") + location.host + location.pathname + "api/log/ws")
      this.connection.onmessage = function(event) {
        const message = JSON.parse(event.data);
        console.log("LOG:", message)
        if (message.receiver > 0) this.logMessages[message.receiver].push(message)

        // filter and sort
        const now = new Date()
        const history = now.setTime((new Date()).getTime() - 14400 * 1000) // 4 hours of logs
        for (let index=1; index<=4; index++) {
          this.logMessages[index] = this.logMessages[index].filter((item) => new Date(item.time) > history)
          this.logMessages[index] = this.logMessages[index].sort((a, b) => {return new Date(b.time) - new Date(a.time)})
        }
      }.bind(this)
      this.connection.onerror = function (event) {
        console.log(event)
        this.connection.close()
      }.bind(this)
      this.connection.onclose = function (event) {
        console.log(event)
        setTimeout(this.connectLog.bind(this), 1000)
      }.bind(this)
    },
  },
  async created() {
    this.updateConfig();
    this.updateData();
    setInterval(this.updateData.bind(this), 5000)
    this.connectLog()
  },
};
</script>
