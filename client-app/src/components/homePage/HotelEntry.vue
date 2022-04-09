<template>
  <v-container>
    <v-card class="mx-auto" max-width="600">
      <v-img
        src="https://cdn.vuetifyjs.com/images/cards/sunshine.jpg"
        height="200px"
      ></v-img>

      <v-card-title> {{ hotel.Name }} </v-card-title>

      <v-card-subtitle> {{ hotel.Address }} </v-card-subtitle>

      <v-card-actions>
        <v-btn color="orange lighten-2" text @click="hotelDetailed(hotel.Id)">
          Explore
        </v-btn>

        <v-spacer></v-spacer>

        <v-btn icon @click="show = !show">
          <v-icon>{{ show ? "mdi-chevron-up" : "mdi-chevron-down" }}</v-icon>
        </v-btn>
      </v-card-actions>

      <v-expand-transition>
        <div v-show="show">
          <v-divider></v-divider>

          <v-card-text>
            {{ hotel.Description }}
          </v-card-text>
        </div>
      </v-expand-transition>
    </v-card>
    <br />
    <br />
  </v-container>
</template>

<script>
export default {
  name: "hotel-entry",
  props: ["hotel"],
  data() {
    return {
      show: false,
    };
  },
  methods: {
    hotelDetailed(id) {
      this.$router.push("/hotels/" + id);
    },
  },
};
</script>
