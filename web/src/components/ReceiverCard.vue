<!--
graphs: mer, carrier freq
labels: state, symbol rate, service/provider, modulation, audio type, video type,
-->
<template>
  <v-container fluid>
    <v-row>
      <v-col :cols="cols">
        <v-card>
          <v-card-title>Signal</v-card-title>
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
              <v-list-item-icon v-if="config.symbol_rate !== symbolRate && item.key === 'symbol_rate'">
                <v-progress-circular
                  indeterminate
                  color="primary"
                >
                </v-progress-circular>
              </v-list-item-icon>
              <v-list-item-content class="align-end">
                <v-combobox
                    :items="symbolRates"
                    v-model="symbolRate"
                    label="aanpassen (werkt nog niet)"
                    outlined
                    dense
                    @change="send_symbolrate"
                    v-if="item.key === 'symbol_rate'"
                ></v-combobox>
                <span v-else>
                    {{ receiver[item.key] }}
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
            <span>(Werkt nog niet)</span>
            <v-timeline-item
                v-for="event in call_log"
                :key="event.start"
                color="blue"
                small
                fill-dot
            >
              <strong>{{ event.value }}</strong>
              <div class="text-caption">
                start: {{ event.start | formatDate }}
              </div>
              <div class="text-caption" v-if="event.end !== 0">
                end: {{ event.end | formatDate }}
              </div>
            </v-timeline-item>
          </v-timeline>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<!--  <v-row>-->
<!--    <v-col>-->
<!--      <signal-chart-->
<!--          v-bind:signal="receiver.mer_history"-->
<!--      />-->
<!--      <v-card v-if="config !== null">-->
<!--        <v-card-title>settings</v-card-title>-->
<!--        <v-list>-->
<!--          <v-list-item>-->
<!--            <v-list-item-icon v-if="config.symbol_rate !== symbolRate">-->
<!--              <v-progress-circular-->
<!--                indeterminate-->
<!--                color="primary"-->
<!--              >-->
<!--              </v-progress-circular>-->
<!--            </v-list-item-icon>-->
<!--            <v-list-item-content>-->
<!--&lt;!&ndash;            <v-list-item-title>Symbol rate</v-list-item-title>&ndash;&gt;-->
<!--&lt;!&ndash;            <v-list-item-subtitle>&ndash;&gt;-->
<!--              <v-combobox-->
<!--                  :items="symbolRates"-->
<!--                  v-model="symbolRate"-->
<!--                  label="Symbol Rate"-->
<!--                  outlined-->
<!--                  dense-->
<!--                  @change="send_symbolrate"-->
<!--              ></v-combobox>-->
<!--&lt;!&ndash;            </v-list-item-subtitle>&ndash;&gt;-->
<!--            </v-list-item-content>-->
<!--          </v-list-item>-->
<!--        </v-list>-->
<!--      </v-card>-->
<!--    </v-col>-->
<!--    <v-col>-->
<!--      <v-card>-->
<!--&lt;!&ndash;        <v-expansion-panels>&ndash;&gt;-->
<!--&lt;!&ndash;          <v-expansion-panel>&ndash;&gt;-->
<!--&lt;!&ndash;            <v-expansion-panel-header>&ndash;&gt;-->
<!--              <v-card-title>{{receiver.title_bar}}</v-card-title>-->
<!--&lt;!&ndash;            </v-expansion-panel-header>&ndash;&gt;-->
<!--&lt;!&ndash;            <v-expansion-panel-content>&ndash;&gt;-->
<!--              <v-list dense>-->
<!--                <v-list-item v-for="item in infoItems" v-bind:key="item.header">-->
<!--                  <v-list-item-content>-->
<!--                    <span><b>{{ item.header }}</b></span>-->
<!--                  </v-list-item-content>-->
<!--                  <v-list-item-content class="align-end">-->
<!--                    <span>{{item.value}}</span>-->
<!--                  </v-list-item-content>-->
<!--                </v-list-item>-->
<!--                <v-list-item-->
<!--                    v-for="(value,key) in non_empty_values()"-->
<!--                    v-bind:key="key"-->
<!--                    style="min-height: 30px"-->
<!--                >-->
<!--                  <v-list-item-content>-->
<!--                    <span>-->
<!--                      <b>{{ key }}:</b>-->
<!--                    </span>-->
<!--                  </v-list-item-content>-->
<!--                  <v-list-item-content class="align-end">-->
<!--                    <span>-->
<!--                    {{ value }}-->
<!--                  </span>-->
<!--                  </v-list-item-content>-->
<!--                </v-list-item>-->
<!--              </v-list>-->
<!--&lt;!&ndash;              <ul>&ndash;&gt;-->
<!--&lt;!&ndash;                <li v-for="(value,key) in receiver" v-bind:key="key">{{key}}: {{value}}</li>&ndash;&gt;-->
<!--&lt;!&ndash;              </ul>&ndash;&gt;-->
<!--&lt;!&ndash;            </v-expansion-panel-content>&ndash;&gt;-->
<!--&lt;!&ndash;          </v-expansion-panel>&ndash;&gt;-->
<!--&lt;!&ndash;        </v-expansion-panels>&ndash;&gt;-->
<!--      </v-card>-->
<!--    </v-col>-->
<!--  </v-row>-->
<!-- </template>-->

<script>
import SignalChartComponent from "@/components/SignalChartComponent";
  export default {
    name: 'ReceiverCard',
    props: ['receiver', 'config'],
    components: {SignalChart: SignalChartComponent},
    computed: {
      call_log () {
        let logs = null // this.receiver.metrics['ProviderName'].history?.events
        if (logs !== null)
          return logs.reverse()
        return null
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
                {name: "Antenne", key: "antenna"},
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
      symbolRates: [25, 35, 66, 125, 250, 333, 360, 500, 1000, 1200, 1500, 2000, 3000, 4000, 4167, 22000, 27500],
      symbolRate: "select",
    }),
    methods: {
      send_symbolrate() {
        console.log("SETTING SYMBOL RATE TO", this.symbolRate)
      },
      non_empty_values() {
        let result = {}
        for (const [key, value] of Object.entries(this.receiver)) {
          // console.log(`${key}: ${value}`);
          if (value === null) continue
          if (key.endsWith("_history")) continue
          // if ([].indexOf(key) !== -1) continue
          result[key] = value
        }
        return result
      },
    },
    async created() {
      this.symbolRate = this.config.symbol_rate
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