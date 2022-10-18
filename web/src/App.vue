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
<!--        <v-tab href="#info">-->
<!--          Info-->
<!--          <v-icon>md-info</v-icon>-->
<!--        </v-tab>-->
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
      <v-tabs-items v-model="tab">
        <v-tab-item
            v-for="(receiver, index) in receivers"
            :key="index"
            :value="index"
        >
          <ReceiverCard :receiver="receiver"/>
        </v-tab-item>
      </v-tabs-items>
<!--      <v-row>-->
<!--        <v-col v-for="index in 4" v-bind:key="index">-->
<!--          <ReceiverCard :receiver="receivers[index-1]"/>-->
<!--        </v-col>-->
<!--      </v-row>-->
    </v-main>
  </v-app>
</template>

<script>
import ReceiverCard from './components/ReceiverCard';

export default {
  name: 'App',

  components: {
    ReceiverCard: ReceiverCard,
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
  }),
  methods: {
    receiverIcon: function (status) {
      if (status === "lost") { return "mdi-wifi-strength-3-alert" }
      if (["header", "DVB-S2", "DVB-S"].indexOf(status) !== -1) {
        return "mdi-wifi"
      }
      return "mdi-wifi-strength-outline"
    },
    updateData() {
      fetch("api/status")
          .then((response) => {
            if (response.ok) {
              this.error = ""
              return response.json();
            } else {
              this.error = "failed to fetch layout status"
              return {}
            }
          })
          .then(data => {
            console.log("DATA", data)
            this.receivers = data
          })
    },
  },
  async created() {
    this.updateData();
    setInterval(this.updateData.bind(this), 5000)
  },
};
</script>
