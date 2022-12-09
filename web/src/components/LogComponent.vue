<template>
  <v-card width="100%">
    <v-card-title>Log</v-card-title>
    <v-card-subtitle>Let op: Na het veranderen van een symbolrate wordt deze automatisch na 30 minuten gereset.</v-card-subtitle>
    <v-card-text>
      <v-list>
        <v-list-item v-for="(message, index) in messages" :key="index">
          <v-list-item-content>
            <v-list-item-title>
              {{ new Date(message.time).toLocaleString('nl-NL') }} &rang;
              {{ message.user }} stelde <strong>{{ message.setting }} in op {{ message.value }}ks</strong>
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: "LogComponent",
  props: ["receiver"],
  data: () => ({
    connection: null,
    messages: [],
  }),
  methods: {
    connect() {
      console.log("connecting to websocket")
      this.messages = [] // reset
      this.connection = new WebSocket(((location.protocol === "https:") ? "wss://" : "ws://") + location.host + location.pathname + "api/log/ws")
      this.connection.onmessage = function(event) {
        const message = JSON.parse(event.data);
        if (message.receiver !== this.receiver) return // filter out what is not for this view
        this.messages.push(message)

        const now = new Date()
        const history = now.setTime((new Date()).getTime() - 14400 * 1000) // 4 hours of logs
        this.messages = this.messages.filter((item) => new Date(item.time) > history)
        this.messages = this.messages.sort((a, b) => {return new Date(b.time) - new Date(a.time)})
      }.bind(this)
      this.connection.onerror = function () {
        this.connection.close()
      }.bind(this)
      this.connection.onclose = function () {
        setTimeout(this.connect.bind(this), 1000)
      }.bind(this)
    },
  },
  async created() {
    this.connect()
  },
}
</script>

<style scoped>

</style>