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
      <h3>Come find us at: {{ hotel.Address }}</h3>
    </v-flex>
    <br />
    <v-row align="center" justify="center">
      <v-col class="col-lg-6">
        <h3>
          {{ hotel.Description }}
        </h3>
      </v-col>
    </v-row>
    <br />
    <br />
    <room-list></room-list>
  </v-container>
</template>

<script>
import HotelService from "../services/hotelService.js";
import RoomList from "../components/detailedHotel/roomList.vue";

export default {
  name: "detailed-hotel-view",
  components: { RoomList },
  data() {
    return {
      hotel: {},
    };
  },
  mounted() {
    HotelService.getHotel(this.$route.params.id).then((response) => {
      this.hotel = response.data;
    });
  },
};
</script>
