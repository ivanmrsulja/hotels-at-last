<template>
  <v-container>
    <v-flex class="text-center">
      <h1>{{ hotel.Name }}</h1>
      <br />
      <img
        src="https://cdn.vuetifyjs.com/images/cards/sunshine.jpg"
        alt="Hotel"
        height="300px"
        style="border-radius: 5px"
      />
      <h3>{{ hotel.Address }}</h3>
      <br />
      <h3>{{ hotel.Description }}</h3>
    </v-flex>
  </v-container>
</template>

<script>
import HotelService from "../services/hotelService.js";

export default {
  name: "detailed-hotel-view",
  data() {
    return {
      hotel: {},
    };
  },
  mounted() {
    HotelService.getHotel(this.$route.params.id).then((response) => {
      this.hotel = response.data;
      console.log(this.hotel);
    });
  },
};
</script>
