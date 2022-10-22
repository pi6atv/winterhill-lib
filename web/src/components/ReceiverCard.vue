<!--
graphs: mer, carrier freq
labels: state, symbol rate, service/provider, modulation, audio type, video type,
-->
<template>
  <v-row>
    <v-col>
      <v-row>
        <signal-chart
            v-bind:signal="receiver.mer_history"
        />
      </v-row>
      <v-row>
        <v-card v-if="config !== null">
          <v-card-title>settings</v-card-title>
<!--          set SR -->
          {{ config }}
        </v-card>
      </v-row>
    </v-col>
    <v-col>
      <v-card>
<!--        <v-expansion-panels>-->
<!--          <v-expansion-panel>-->
<!--            <v-expansion-panel-header>-->
              <v-card-title>{{receiver.title_bar}}</v-card-title>
<!--            </v-expansion-panel-header>-->
<!--            <v-expansion-panel-content>-->
              <v-list dense>
                <v-list-item
                    v-for="(value,key) in non_empty_values()"
                    v-bind:key="key"
                    style="min-height: 30px"
                >
                  <v-list-item-content>
                    <span>
                      <b>{{ key }}:</b>
                    </span>
                  </v-list-item-content>
                  <v-list-item-content class="align-end">
                    <span>
                    {{ value }}
                  </span>
                  </v-list-item-content>
                </v-list-item>
              </v-list>
<!--              <ul>-->
<!--                <li v-for="(value,key) in receiver" v-bind:key="key">{{key}}: {{value}}</li>-->
<!--              </ul>-->
<!--            </v-expansion-panel-content>-->
<!--          </v-expansion-panel>-->
<!--        </v-expansion-panels>-->
      </v-card>
    </v-col>
  </v-row>
 </template>

<script>
import SignalChartComponent from "@/components/SignalChartComponent";
  export default {
    name: 'ReceiverCard',
    props: ['receiver', 'config'],
    components: {SignalChart: SignalChartComponent},

    data: () => ({
    }),
    methods: {
      non_empty_values () {
        let result = {}
        for (const [key, value] of Object.entries(this.receiver)) {
          console.log(`${key}: ${value}`);
          if (value === null) continue
          if (key.endsWith("_history")) continue
          // if ([].indexOf(key) !== -1) continue
          result[key] = value
        }
        return result
      },
    }
  }
</script>
