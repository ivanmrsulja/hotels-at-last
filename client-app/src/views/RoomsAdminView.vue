<template>
  <v-container>
    <v-row align="center" justify="center">
      <v-col class="col-lg-10">
        <br />
        <v-flex class="text-center">
          <h1>Manage Rooms For {{ hotel.Name }}</h1>
        </v-flex>
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
                <v-btn color="blue" dark @click="deleteRoom(row.item.Id)">
                  Delete
                </v-btn>
              </td>
              <td>
                <v-btn color="blue" dark @click="updateRoom(row.item.Id)">
                  Edit
                </v-btn>
              </td>
            </tr>
          </template>
        </v-data-table>
        <br />
        <v-btn
          style="float: right"
          color="blue"
          dark
          @click="createRoom(row.item.Id)"
        >
          Add new Room
        </v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import HotelService from "../services/hotelService.js";

export default {
  name: "rooms-admin-view",
  data() {
    return {
      rooms: [],
      hotel: {},
      page: 0,
      totalResults: 10,
      options: {},
      loading: true,
      headers: [
        {
          text: "Room Number",
          align: "start",
          sortable: false,
        },
        { text: "Beds", align: "start" },
        { text: "Price", align: "start" },
        { text: "Air conditioning", align: "start" },
        { text: "Parking", align: "start" },
        { text: "TV", align: "start" },
        { text: "Action", align: "start" },
        { text: "Action", align: "start" },
      ],
    };
  },
  watch: {
    options: {
      handler() {
        this.fetchRooms();
      },
      deep: true,
    },
  },
  mounted() {
    this.fetchHotel();
  },
  methods: {
    fetchRooms() {
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
    fetchHotel() {
      HotelService.getHotel(this.$route.params.id).then((response) => {
        this.hotel = response.data;
      });
    },
    createRoom(roomId) {},
    updateRoom(roomId) {},
    deleteRoom(roomId) {
      HotelService.deleteRoom(roomId).then((response) => {
        this.fetchRooms();
      });
    },
  },
};
</script>
