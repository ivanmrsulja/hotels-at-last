<template>
  <v-container>
    <v-card class="mx-auto" max-width="600">
      <v-img :src="image" height="200px"></v-img>

      <v-card-title>
        {{ hotel.Name }}
        <v-spacer></v-spacer>
        <v-rating
          color="orange"
          length="5"
          size="16"
          :value="hotel.Stars"
          readonly
        ></v-rating
      ></v-card-title>

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
import HotelService from "../../services/hotelService.js";

export default {
  name: "hotel-entry",
  props: ["hotel"],
  data() {
    return {
      show: false,
      image: "",
    };
  },
  mounted() {
    this.fetchImage();
    this.$root.$on("refreshImages", () => {
      this.fetchImage();
    });
  },
  methods: {
    hotelDetailed(id) {
      this.$router.push("/hotels/" + id);
    },
    fetchImage() {
      HotelService.getImage(this.hotel.Base64Image).then((response) => {
        this.image = response.data.Base64Image;
      });
    },
  },
};
</script>
