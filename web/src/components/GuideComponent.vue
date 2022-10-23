<template>
  <v-card>
    <v-card-title>
      Welkom op de DVB Winterhill ontvanger status pagina van PI6ATV.
    </v-card-title>
    <v-card-text>
      <v-row>
        <v-col>
          Op deze pagina kan je de status van onze Winterhill DVB ontvangers bekijken.<br>
          Kies in de balk bovenaan de ontvanger waarvan je de status wil bekijken.<br>
          Hier kun je ook de symbol rate instellen.<br>
        </v-col>
      </v-row>

      <v-row>
        <v-col>
          <v-img src="receiver.jpg"></v-img>
        </v-col>
      </v-row>
<!--      <v-row>-->
<!--        <v-col>-->
<!--          <h2>Settings</h2>-->
<!--        </v-col>-->
<!--      </v-row>-->

<!--      <v-row>-->
<!--        <v-col>-->
<!--          Gebruik zo veel mogelijk de volgende settings om correct op zowel de-->
<!--          analoge als de DVB zender van PI6ATV te kunnen werken:-->
<!--        </v-col>-->
<!--      </v-row>-->

<!--      <v-row-->
<!--          style="width: 600px"-->
<!--          justify="start"-->
<!--      >-->
<!--        <v-col>-->
<!--          <v-subheader>Modulator</v-subheader>-->
<!--          <v-data-table-->
<!--              dense-->
<!--              hide-default-header-->
<!--              hide-default-footer-->
<!--              :headers="kv_headers"-->
<!--              :items="settings_base"-->
<!--              :items-per-page="1000"-->
<!--              class="elevation-1"-->
<!--          >-->
<!--          </v-data-table>-->
<!--        </v-col>-->
<!--      </v-row>-->


<!--      <v-row-->
<!--          style="width: 600px"-->
<!--          justify="start"-->
<!--      >-->
<!--        <v-col>-->
<!--          <v-subheader>Service</v-subheader>-->
<!--          <v-data-table-->
<!--              dense-->
<!--              hide-default-header-->
<!--              hide-default-footer-->
<!--              :headers="kv_headers"-->
<!--              :items="settings_pid"-->
<!--              :items-per-page="1000"-->
<!--              class="elevation-1"-->
<!--          >-->
<!--          </v-data-table>-->
<!--        </v-col>-->
<!--      </v-row>-->

<!--      <v-row-->
<!--          style="width: 600px"-->
<!--          justify="start"-->
<!--      >-->
<!--        <v-col>-->
<!--          <v-subheader>Video</v-subheader>-->
<!--          <v-data-table-->
<!--              dense-->
<!--              hide-default-header-->
<!--              hide-default-footer-->
<!--              :headers="kv_headers"-->
<!--              :items="settings_video"-->
<!--              :items-per-page="1000"-->
<!--              class="elevation-1"-->
<!--          >-->
<!--          </v-data-table>-->
<!--        </v-col>-->
<!--      </v-row>-->

<!--      <v-row-->
<!--          style="width: 600px"-->
<!--          justify="start"-->
<!--      >-->
<!--        <v-col>-->
<!--          <v-subheader>Audio</v-subheader>-->
<!--          <v-data-table-->
<!--              dense-->
<!--              hide-default-header-->
<!--              hide-default-footer-->
<!--              :headers="kv_headers"-->
<!--              :items="settings_audio"-->
<!--              :items-per-page="1000"-->
<!--              class="elevation-1"-->
<!--          >-->
<!--          </v-data-table>-->
<!--        </v-col>-->
<!--      </v-row>-->

<!--      <v-row>-->
<!--        <v-col>-->
<!--          <h2>Tips</h2>-->
<!--          <v-list dense>-->
<!--            <v-list-item>-->
<!--              <v-list-item-icon>-->
<!--                <v-icon>-->
<!--                  mdi-lightbulb-on-outline-->
<!--                </v-icon>-->
<!--              </v-list-item-icon>-->
<!--              <v-list-item-content>-->
<!--                DVB-S2 is stabieler dan DVB-S-->
<!--              </v-list-item-content>-->
<!--            </v-list-item>-->
<!--            <v-list-item>-->
<!--              <v-list-item-icon>-->
<!--                <v-icon>-->
<!--                  mdi-lightbulb-on-outline-->
<!--                </v-icon>-->
<!--              </v-list-item-icon>-->
<!--              <v-list-item-content>-->
<!--                Probeer eerst DVB-S2 in QPSK met een FEC van 1/4 en als je dan nog ruime SNR Margin-->
<!--                hebt (dus signaal "over" hebt) kan je de FEC hoger zetten en eventuel omschakelen-->
<!--                naar 8PSK om netto meer bandbreedte te krijgen.<br>-->
<!--                Het grijze vlak in de SNR grafiek geeft aan hoeveel je signaal nog kan verzwakken om-->
<!--                toch nog over de ingang te kunnen werken.-->
<!--              </v-list-item-content>-->
<!--            </v-list-item>-->
<!--          </v-list>-->
<!--        </v-col>-->
<!--      </v-row>-->

<!--      <v-row-->
<!--          style="width: 600px"-->
<!--          justify="start"-->
<!--      >-->
<!--        <v-col>-->
<!--          <v-subheader>DVB-S in QPSK</v-subheader>-->
<!--          <v-data-table-->
<!--            dense-->
<!--            hide-default-footer-->
<!--            :headers="mod_headers"-->
<!--            :items="dvbs_qpsk"-->
<!--            :items-per-page="1000"-->
<!--            class="elevation-1"-->
<!--          >-->
<!--          </v-data-table>-->
<!--        </v-col>-->
<!--      </v-row>-->

<!--      <v-row-->
<!--          style="width: 600px"-->
<!--          justify="start"-->
<!--      >-->
<!--        <v-col>-->
<!--          <v-subheader>DVB-S2 in QPSK</v-subheader>-->
<!--          <v-data-table-->
<!--            dense-->
<!--            hide-default-footer-->
<!--            :headers="mod_headers"-->
<!--            :items="dvbs2_qpsk"-->
<!--            :items-per-page="1000"-->
<!--            class="elevation-1"-->
<!--          >-->
<!--          </v-data-table>-->
<!--        </v-col>-->
<!--      </v-row>-->

<!--      <v-row-->
<!--          style="width: 600px"-->
<!--          justify="start"-->
<!--      >-->
<!--        <v-col>-->
<!--          <v-subheader>DVB-S2 in 8PSK</v-subheader>-->
<!--          <v-data-table-->
<!--            dense-->
<!--            hide-default-footer-->
<!--            :headers="mod_headers"-->
<!--            :items="dvbs2_8psk"-->
<!--            :items-per-page="1000"-->
<!--            class="elevation-1"-->
<!--          >-->
<!--          </v-data-table>-->
<!--        </v-col>-->
<!--      </v-row>-->
    </v-card-text>
  </v-card>
</template>
<script>
export default {
  data() {
    return {
      kv_headers: [
        {value: "key"},
        {value: "value"},
      ],
      settings_base: [
        {key: "Standaard", value: "DVB-S of DVB-S2"},
        {key: "Modulatie", value: "QPSK of 8PSK (alleen DVB-S2 ondersteunt 8PSK)"},
        {key: "Symbolrate", value: "1000 (70cm) of 2000 kSym/s (23,13,6,3 en 1.2cm)"},
        {key: "FEC", value: "1/2, 1/3, 1/4, 2/3, 2/5, 3/4, 3/5, 4/5, 5/6, 8/9, 9/10"},
        {key: "Polarisatie", value: "voor ALLE banden horizontaal"},
      ],
      settings_pid: [
        {key: "Service ID", value: "1"},
        {key: "PMT PID", value: "100 (dec) - 0x64 (hex)"},
        {key: "PCR PID", value: "101 (dec) - 0x65 (hex)"},
        {key: "Video PID", value: "101 (dec) - 0x65 (hex)"},
        {key: "Audio PID", value: "102 (dec) - 0x66 (hex)"},
      ],
      settings_video: [
        {key: "Encoding", value: "MPEG2 of MPEG-4 AVC (h.264)"},
        {key: "Resoluties", value: "480i, 576i, 720p of 1080i"},
      ],
      settings_audio: [
        {key: "Encoding", value: "MPEG-1 Layer II of AAC"},
      ],
      mod_headers: [
        {text: "Standaard", value: "standard"},
        {text: "Modulatie", value: "modulation"},
        {text: "FEC", value: "fec"},
        {text: "Minimum signaal", value: "signal"},
        {text: "Netto bandbreedte (1000kS/s)", value: "bandwidth"},
        {text: "Netto bandbreedte (2000kS/s)", value: "bandwidth2"},
      ],
      dvbs_qpsk: [
        {standard: "DVB-S", modulation: "QPSK", fec: "1/2", signal: "2,7 dB", bandwidth: "0,9 Mbps", bandwidth2: "1,8 Mbps"},
        {standard: "DVB-S", modulation: "QPSK", fec: "2/3", signal: "4,4 dB", bandwidth: "1,2 Mbps", bandwidth2: "2,5 Mbps"},
        {standard: "DVB-S", modulation: "QPSK", fec: "3/4", signal: "5,7 dB", bandwidth: "1,4 Mbps", bandwidth2: "2,8 Mbps"},
        {standard: "DVB-S", modulation: "QPSK", fec: "5/6", signal: "6,7 dB", bandwidth: "1,5 Mbps", bandwidth2: "3,1 Mbps"},
        {standard: "DVB-S", modulation: "QPSK", fec: "7/8", signal: "7,2 dB", bandwidth: "1,6 Mbps", bandwidth2: "3,2 Mbps"},
      ],
      dvbs2_qpsk: [
        {standard: "DVB-S2", modulation: "QPSK", fec: "1/4", signal: "-2,4 dB", bandwidth: "0,5 Mbps", bandwidth2: "1 Mbps"},
        {standard: "DVB-S2", modulation: "QPSK", fec: "1/3", signal: "-1,2 dB", bandwidth: "0,65 Mbps", bandwidth2: "1,3 Mbps"},
        {standard: "DVB-S2", modulation: "QPSK", fec: "2/5", signal: "0 dB", bandwidth: "0,8 Mbps", bandwidth2: "1,6 Mbps"},
        {standard: "DVB-S2", modulation: "QPSK", fec: "1/2", signal: "1 dB", bandwidth: "1 Mbps", bandwidth2: "2 Mbps"},
        {standard: "DVB-S2", modulation: "QPSK", fec: "3/5", signal: "2,2 dB", bandwidth: "1,2 Mbps", bandwidth2: "2,4 Mbps"},
        {standard: "DVB-S2", modulation: "QPSK", fec: "2/3", signal: "3,1 dB", bandwidth: "1,3 Mbps", bandwidth2: "2,6 Mbps"},
        {standard: "DVB-S2", modulation: "QPSK", fec: "3/4", signal: "4 dB", bandwidth: "1,5 Mbps", bandwidth2: "3 Mbps"},
        {standard: "DVB-S2", modulation: "QPSK", fec: "4/5", signal: "4,6 dB", bandwidth: "1,58 Mbps", bandwidth2: "3,2 Mbps"},
        {standard: "DVB-S2", modulation: "QPSK", fec: "5/6", signal: "5,2 dB", bandwidth: "1,65 Mbps", bandwidth2: "3,3 Mbps"},
        {standard: "DVB-S2", modulation: "QPSK", fec: "8/9", signal: "6,2 dB", bandwidth: "1,77 Mbps", bandwidth2: "3,5 Mbps"},
        {standard: "DVB-S2", modulation: "QPSK", fec: "9/10", signal: "6,5 dB", bandwidth: "1,79 Mbps", bandwidth2: "3,6 Mbps"},
      ],
      dvbs2_8psk: [
        {standard: "DVB-S2", modulation: "8PSK", fec: "3/5", signal: "5,5 dB", bandwidth: "1,8 Mbps", bandwidth2: "3,6 Mbps"},
        {standard: "DVB-S2", modulation: "8PSK", fec: "2/3", signal: "6,6 dB", bandwidth: "2 Mbps", bandwidth2: "4 Mbps"},
        {standard: "DVB-S2", modulation: "8PSK", fec: "3/4", signal: "7,9 dB", bandwidth: "2,33 Mbps", bandwidth2: "4,5 Mbps"},
        {standard: "DVB-S2", modulation: "8PSK", fec: "5/6", signal: "9,4 dB", bandwidth: "2,48 Mbps", bandwidth2: "5 Mbps"},
        {standard: "DVB-S2", modulation: "8PSK", fec: "8/9", signal: "10,6 dB", bandwidth: "2,64 Mbps", bandwidth2: "5,3 Mbps"},
        {standard: "DVB-S2", modulation: "8PSK", fec: "9/10", signal: "11 dB", bandwidth: "2,67 Mbps", bandwidth2: "5,4 Mbps"},
      ],
    }
  }
}
</script>
