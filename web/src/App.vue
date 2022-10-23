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
          {{ receiver.title_bar }}
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
      <v-banner
          class="red"
      >
        <div style="text-align: center;"><b>Deze pagina is nog in ontwerp!</b></div>
      </v-banner>
      <v-tabs-items v-model="tab">
        <v-tab-item value="info">
          <Guide/>
        </v-tab-item>
        <v-tab-item
            v-for="(receiver, index) in receivers"
            :key="index"
            :value="index"
        >
          <ReceiverCard :receiver="receiver" :config="getConfig(index)"/>
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
            this.receivers = data
          })
    },
  },
  async created() {
    this.updateConfig();
    this.updateData();
    setInterval(this.updateData.bind(this), 5000)
  },
};
</script>
