<template>
  <v-card>
    <v-card-title>
      Dit is de ruwe data van de winterhill module
    </v-card-title>
    <v-card-text>
      <pre>{{ summary }}</pre>
    </v-card-text>
  </v-card>
</template>
<script>

export default {
  data() {
    return {
      summary:"waiting for data...",
    }
  },
  methods: {
    updateData() {
      fetch("api/summary")
          .then((response) => {
            if (response.ok) {
              this.error = ""
              return response.json();
            } else {
              this.error = "failed to fetch  summary"
              return {}
            }
          })
          .then(data => {
            this.summary = data
          })
    },
  },
  async created() {
    setInterval(this.updateData.bind(this), 5000)
  }
}
</script>
