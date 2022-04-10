<template>
  <v-row>
    <v-col cols="12" sm="6" md="4">
      <v-menu
        ref="menuStart"
        v-model="menuStart"
        :close-on-content-click="false"
        :return-value.sync="dateStart"
        transition="scale-transition"
        offset-y
        min-width="auto"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
            v-model="dateStart"
            label="Start date"
            prepend-icon="mdi-calendar"
            readonly
            v-bind="attrs"
            v-on="on"
          ></v-text-field>
        </template>
        <v-date-picker v-model="dateStart" no-title scrollable>
          <v-spacer></v-spacer>
          <v-btn text color="primary" @click="menuStart = false">
            Cancel
          </v-btn>
          <v-btn text color="primary" @click="$refs.menuStart.save(dateStart)">
            OK
          </v-btn>
        </v-date-picker>
      </v-menu>
    </v-col>
    <v-col cols="12" sm="6" md="4">
      <v-menu
        ref="menuEnd"
        v-model="menuEnd"
        :close-on-content-click="false"
        :return-value.sync="dateEnd"
        transition="scale-transition"
        offset-y
        min-width="auto"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
            v-model="dateEnd"
            label="End date"
            prepend-icon="mdi-calendar"
            readonly
            v-bind="attrs"
            v-on="on"
          ></v-text-field>
        </template>
        <v-date-picker v-model="dateEnd" no-title scrollable>
          <v-spacer></v-spacer>
          <v-btn text color="primary" @click="menuEnd = false"> Cancel </v-btn>
          <v-btn text color="primary" @click="$refs.menuEnd.save(dateEnd)">
            OK
          </v-btn>
        </v-date-picker>
      </v-menu>
    </v-col>
    <v-col cols="12" sm="6" md="4">
      <v-btn
        style="margin-top: 15px"
        text
        color="primary"
        @click="makeReservation()"
      >
        Make reservation
      </v-btn>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: "reservation-form",
  data() {
    return {
      dateStart: new Date(Date.now() - new Date().getTimezoneOffset() * 60000)
        .toISOString()
        .substr(0, 10),
      dateEnd: new Date(Date.now() - new Date().getTimezoneOffset() * 60000)
        .toISOString()
        .substr(0, 10),
      menuStart: false,
      menuEnd: false,
    };
  },
  methods: {
    makeReservation() {
      let timeframe = {
        start: this.dateStart,
        end: this.dateEnd,
      };
      this.$root.$emit("reservation", timeframe);
    },
  },
};
</script>
