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
                <add-room-form
                  style="float: left"
                  :edit="true"
                  :roomData="row.item"
                ></add-room-form>
              </td>
            </tr>
          </template>
        </v-data-table>
        <br />
        <add-room-form
          style="float: right"
          :edit="false"
          :roomData="{}"
        ></add-room-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import HotelService from "../services/hotelService.js";
import AddRoomForm from "../components/administration/addRoomForm.vue";

export default {
  name: "rooms-admin-view",
  components: { AddRoomForm },
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
    this.$root.$on("createRoom", (request) => {
      this.createRoom(request);
    });
    this.$root.$on("updateRoom", (request) => {
      this.updateRoom(request);
    });
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
    createRoom(request) {
      HotelService.createRoom(this.hotel.Id, request).then((response) => {
        this.fetchRooms();
      });
    },
    updateRoom(request) {
      HotelService.updateRoom(request.roomId, request.requestBody).then(
        (response) => {
          this.fetchRooms();
        }
      );
    },
    deleteRoom(roomId) {
      HotelService.deleteRoom(roomId).then((response) => {
        this.fetchRooms();
      });
    },
  },
};
</script>
