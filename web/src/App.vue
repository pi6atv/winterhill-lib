<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
    </v-app-bar>

    <v-main>
      <ReceiverCard v-for="index in 4" :receiver="receivers[index-1]" v-bind:key="index"/>
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

  data: () => ({
    receivers: [],
    error: "",
  }),
  methods: {
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
    setInterval(this.updateData.bind(this), 1000)
  },
};
</script>
