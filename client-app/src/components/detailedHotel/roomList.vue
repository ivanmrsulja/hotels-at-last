<template>
  <v-container>
    <v-row justify="center">
      <v-col class="col-lg-8 col-sm-12">
        <h2>Room list:</h2>
        <br />
        <v-data-table
          :headers="headers"
          :items="rooms"
          :options.sync="options"
          :items-per-page="5"
          :server-items-length="totalResults"
          class="elevation-1"
        >
          <template v-slot:item="row">
            <tr>
              <td>{{ row.item.RoomNumber }}</td>
              <td>{{ row.item.NumberOfBeds }}</td>
              <td>{{ row.item.Price }}</td>
              <td>
                <v-icon v-if="row.item.AirConditioned">mdi-check-circle</v-icon>
                <v-icon v-if="!row.item.AirConditioned">mdi-close-thick</v-icon>
              </td>
              <td>
                <v-icon v-if="row.item.HasParkingSpace"
                  >mdi-check-circle</v-icon
                >
                <v-icon v-if="!row.item.HasParkingSpace"
                  >mdi-close-thick</v-icon
                >
              </td>
              <td>
                <v-icon v-if="row.item.HasTV">mdi-check-circle</v-icon>
                <v-icon v-if="!row.item.HasTV">mdi-close-thick</v-icon>
              </td>
              <td>
                <v-btn color="blue" text @click="exploreRoom(row.item.Id)">
                  Explore
                </v-btn>
              </td>
            </tr>
          </template>
        </v-data-table>
      </v-col>
    </v-row>
    <br />
    <br />
    <br />
    <br />
  </v-container>
</template>

<script>
import HotelService from "../../services/hotelService.js";

export default {
  name: "room-list",
  data() {
    return {
      rooms: [],
      page: 0,
      totalResults: 10,
      totalPages: 1,
      options: {},
      loading: true,
      headers: [
        {
          text: "Room number",
          align: "start",
          sortable: false,
          value: "RoomNumber",
        },
        { text: "Number of beds", align: "start" },
        { text: "Price", align: "v" },
        { text: "Air Conditioning", align: "start" },
        { text: "Parking", align: "start" },
        { text: "TV", align: "start" },
        { text: "Action", align: "start" },
      ],
    };
  },
  watch: {
    options: {
      handler() {
        this.getDataFromApi();
      },
      deep: true,
    },
  },
  mounted() {},
  methods: {
    getDataFromApi() {
      this.loading = true;
      const { sortBy, sortDesc, page, itemsPerPage } = this.options;
      HotelService.getAllRoomsForHotel(
        page - 1,
        itemsPerPage,
        this.$route.params.id
      ).then((response) => {
        this.rooms = response.data.Results;
        this.totalResults = response.data.TotalResults;
      });
    },
    exploreRoom(id) {
      this.$router.push("/hotels/" + this.$route.params.id + "/rooms/" + id);
    },
  },
};
</script>
