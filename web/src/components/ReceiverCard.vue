<!--
graphs: mer, carrier freq
labels: state, symbol rate, service/provider, modulation, audio type, video type,
-->
<template>
 <v-card>
   <v-card-title>{{receiver.title_bar}}</v-card-title>
   <ul>
     <li v-for="(value,key) in receiver" v-bind:key="key">{{key}}: {{value}}</li>
   </ul>
   <signal-chart
       v-bind:signal="receiver.mer_history"
   />

 </v-card>
</template>

<script>
import SignalChartComponent from "@/components/SignalChartComponent";
  export default {
    name: 'ReceiverCard',
    props: ['receiver'],
    components: {SignalChart: SignalChartComponent},

    data: () => ({
    }),
  }
</script>
