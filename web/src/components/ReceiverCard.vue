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
                  <b>{{ item.name }}:</b><br/>
                  <span>{{ item.sub_name }}</span>
                </span>
              </v-list-item-content>
              <v-list-item-content class="align-end">
                <v-row>
                  <v-col>
                    <div v-if="item.key==='antenna'">{{config.antenna}} - {{ {"BOT": "Bottom", "TOP": "Top"}[receiver[item.key]] }}</div>
                    <div v-else>{{ receiver[item.key] }}</div>
                  </v-col>
                  <v-col v-if="item.key === 'symbol_rate'">
                      <v-select
                          :items="symbolRates"
                          v-model="wantedSymbolRate"
                          outlined
                          dense
                      ></v-select>
                  </v-col>
                 <v-col v-if="item.key === 'symbol_rate'">
                      <v-btn
                          @click="send_symbolrate"
                          :color="receiver.symbol_rate !== setSymbolRate?'red':'blue'"
                      >
                        <span>Set</span>
                      </v-btn>
                  </v-col>
                </v-row>
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
            </v-timeline-item>
          </v-timeline>
        </v-card>
      </v-col>
      <v-col>
        <v-card
            style="min-height: 364px"
        >
          <v-card-title>Log</v-card-title>
          <v-card-subtitle>Let op: Na het veranderen van een symbolrate wordt deze automatisch na 30 minuten gereset.</v-card-subtitle>
          <v-timeline
              dense
              class="overflow-y-auto"
              style="max-height: 300px"
          >
            <v-timeline-item
                v-for="message in logMessages"
                :key="message.time"
                color="blue"
                small
                fill-dot
            >
              <strong>{{ message.time | formatDate }}</strong>
              <div class="text-caption">
                {{ message.user }} stelde {{ message.setting }} in op {{ message.value }}ks
              </div>
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
    props: ['receiver', 'config', 'logMessages'],
    components: {SignalChart: SignalChartComponent},
    computed: {
      call_log () {
        return this.receiver.service_history.filter(item => item!==null ).reverse()
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
        let rates = [25, 35, 66, 125, 250, 333, 360, 500, 1000, 1200, 1500, 2000, 3000, 4000, 4167, 22000, 27500]
        if (this.receiver.carrier_frequency <= 1300) rates = [25, 35, 66, 125, 250, 333, 360, 500, 1000, 1200, 1500, 2000, 3000, 4000, 4167] // 23cm
        if (this.receiver.carrier_frequency <= 440) rates = [25, 35, 66, 125, 250, 333, 360, 500, 1000, 1200, 1500, 2000] // 70cm
        if (this.receiver.carrier_frequency <= 146) rates = [25, 35, 66, 125] // 2m
        return rates.map(value => {
          return {value: value, text: value===this.receiver['symbol_rate']?"->"+value:value}
        })
      },
      cards () {
        return [
          { title: 'RF', items: [
              {name: "Status", key: "state"},
              {name: "Frequentie", key: "carrier_frequency"},
              {name: "MER", key: "mer"},
              {name: "D-nummer", key: "d_number"},
              {name: "Modulatie", key: "modulation_code"},
              {name: "Symbol rate", sub_name: "(default = " + this.config.symbol_rate + ")", key: "symbol_rate"},
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
        ]
      }
    },
    data: () => ({
      wantedSymbolRate: "--",
      setSymbolRate: null,
    }),
    methods: {
      send_symbolrate() {
        this.setSymbolRate = this.wantedSymbolRate
        fetch("api/set/srate/"+this.receiver.index+"/"+this.wantedSymbolRate, {method: "POST"})
            .then((response) => {
              if (response.ok) {
                this.error = ""
              } else {
                this.error = "failed to set symbol rate"
                console.log(response)
              }
            })
      },
    },
    async created() {
      this.setSymbolRate = this.receiver.symbol_rate
      this.wantedSymbolRate = this.receiver.symbol_rate
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