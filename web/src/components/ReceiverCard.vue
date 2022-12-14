<!--
graphs: mer, carrier freq
labels: state, symbol rate, service/provider, modulation, audio type, video type,
-->
<template>
  <v-container fluid>
    <v-row>
      <v-col :cols="cols">
        <v-card>
          <v-card-title>MER</v-card-title>
            <signal-chart
                v-bind:signal="receiver.mer_history"
            />
        </v-card>
      </v-col>
      <v-col
          v-for="card in cards"
          :key="card.title"
          :cols="cols"
      >
        <v-card>
          <v-card-title v-text="card.title"></v-card-title>
          <v-list dense>
            <v-list-item
                v-for="item in card.items"
                v-bind:key="item.key"
                style="min-height: 30px"
            >
              <v-list-item-content>
                <span>
                  <b>{{ item.name }}:</b>
                </span>
              </v-list-item-content>
              <v-list-item-content class="align-end">
              <span v-if="item.key === 'symbol_rate'">
                <v-row>
                  <v-col>
                    <v-select
                        :items="symbolRates"
                        v-model="wantedSymbolRate"
                        outlined
                        dense
                    ></v-select>
                  </v-col>
                  <v-col>
                    <v-btn
                        @click="send_symbolrate"
                        disabled
                    >
                      <v-progress-circular
                          v-if="config.symbol_rate !== setSymbolRate"
                          indeterminate
                          color="primary"
                      >
                      </v-progress-circular>
                      <span v-else>Set</span>
                    </v-btn>
                  </v-col>
                </v-row>
              </span>
              <span v-else>
                  <span v-if="item.key==='antenna'">{{config.antenna}} - {{ {"BOT": "Bottom", "TOP": "Top"}[receiver[item.key]] }}</span>
                  <span v-else>{{ receiver[item.key] }}</span>
                </span>
            </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-card
            style="min-height: 364px"
        >
          <v-card-title>Last seen</v-card-title>
          <v-timeline
              dense
              class="overflow-y-auto"
              style="max-height: 300px"
          >
            <v-timeline-item
                v-for="event in call_log"
                :key="event.time"
                color="blue"
                small
                fill-dot
            >
              <strong>{{ event.value }}</strong>
              <div class="text-caption">
                start: {{ event.time | formatDate }}
              </div>
<!--              <div class="text-caption" v-if="event.end !== 0">-->
<!--                end: {{ event.end | formatDate }}-->
<!--              </div>-->
            </v-timeline-item>
          </v-timeline>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import SignalChartComponent from "@/components/SignalChartComponent";
import moment from 'moment';
import Vue from "vue";

Vue.filter('formatDate', function(value) {
  if (value) {
    return moment(String(value)).format('HH:mm:ss')
  }
});

  export default {
    name: 'ReceiverCard',
    props: ['receiver', 'config'],
    components: {SignalChart: SignalChartComponent},
    computed: {
      call_log () {
        return this.receiver.service_history.filter(item => item!==null ).reverse()
        // let logs = this.receiver.service_history
        // if (logs !== null)
        //   return logs.reverse()
        // return null
      },
      cols () {
        console.log(this.$vuetify.breakpoint.name)
        // there are 12 cols on the screen, so return 3 means 4 cards per row
        switch (this.$vuetify.breakpoint.name) {
          case 'xs': return 12
          case 'sm': return 6
          case 'md': return 4
          case 'lg': return 4
          case 'xl': return 3
        }
        return 1
      },
      symbolRates () {
        return [25, 35, 66, 125, 250, 333, 360, 500, 1000, 1200, 1500, 2000, 3000, 4000, 4167, 22000, 27500].map(value => {
          return {value: value, text: value===this.config['symbol_rate']?"*"+value:value}
        })
      }
    },
    data: () => ({
      cards:
          [
            { title: 'RF', items: [
                {name: "Status", key: "state"},
                {name: "Frequentie", key: "carrier_frequency"},
                {name: "MER", key: "mer"},
                {name: "D-nummer", key: "d_number"},
                {name: "Modulatie", key: "modulation_code"},
                {name: "Symbol rate", key: "symbol_rate"},
                {name: "Antenne input", key: "antenna"},
              ], flex: 3 },
            { title: 'Transport Stream', items: [
                {name: "null percentage", key: "ts_null_percentage"},
                {name: "Service", key: "service_name"},
                {name: "Provider", key: "service_provider_name"},
                {name: "Frame type", key: "frame_type"},
                {name: "Pilots", key: "pilots"},
                {name: "Roll off", key: "roll_off"},
              ], flex: 3 },
            { title: 'Video', items: [
                {name: "Coding type", key: "video_type"},
              ], flex: 3 },
            { title: 'Audio', items: [
                {name: "Coding type", key: "audio_type"},
              ], flex: 3 },
          ],
      wantedSymbolRate: "--",
      setSymbolRate: null,
    }),
    methods: {
      send_symbolrate() {
        this.setSymbolRate = this.wantedSymbolRate
        console.log("SETTING SYMBOL RATE TO", this.setSymbolRate)
      },
    },
    async created() {
      this.setSymbolRate = this.config.symbol_rate
      this.wantedSymbolRate = this.config.symbol_rate
    },
  }
</script>

<style lang="scss">
.md-table-cell {
  height: 30px;
}

.md-table-cell#value {
  text-align: right;
}

/* https://getflywheel.com/layout/flexbox-create-modern-card-design-layout/ */
.v-row {
  margin: 4px;
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
}

@media screen and (min-width: 20em) {
  .v-card {
    flex: 0 1 calc(100% - 1em);
  }
}

@media screen and (min-width: 40em) {
  .v-card {
    flex: 0 1 calc(50% - 1em);
  }
}

@media screen and (min-width: 60em) {
  .v-card {
    flex: 0 1 calc(33% - 1em);
  }
}

</style>