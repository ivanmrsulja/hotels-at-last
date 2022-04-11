<template>
  <v-container>
    <v-row align="center" justify="center">
      <v-col class="col-lg-10">
        <br />
        <v-flex class="text-center">
          <h1>Manage Hotels</h1>
        </v-flex>
        <br />
        <v-data-table
          :headers="headers"
          :items="hotels"
          :options.sync="options"
          :items-per-page="5"
          :server-items-length="totalResults"
          class="elevation-1"
        >
          <template v-slot:item="row">
            <tr>
              <td>{{ row.item.Name }}</td>
              <td>{{ row.item.Address }}</td>
              <td>{{ row.item.Description | restrict }}</td>
              <td>
                <v-btn color="blue" dark @click="exploreRooms(row.item.Id)">
                  Explore rooms
                </v-btn>
              </td>
              <td>
                <v-btn color="blue" dark @click="deleteHotel(row.item.Id)">
                  Delete
                </v-btn>
              </td>
              <td>
                <v-btn color="blue" dark @click="updateHotel(row.item.Id)">
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
          @click="createHotel(row.item.Id)"
        >
          Add new Hotel
        </v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import HotelService from "../services/hotelService.js";

export default {
  name: "hotels-admin-view",
  data() {
    return {
      hotels: [],
      page: 0,
      totalResults: 10,
      options: {},
      loading: true,
      headers: [
        {
          text: "Name",
          align: "start",
          sortable: false,
        },
        { text: "Address", align: "start" },
        { text: "Description", align: "start" },
        { text: "Action", align: "start" },
        { text: "Action", align: "start" },
        { text: "Action", align: "start" },
      ],
    };
  },
  watch: {
    options: {
      handler() {
        this.fetchHotels();
      },
      deep: true,
    },
  },
  filters: {
    restrict: function (value) {
      if (!value) return "";
      value = value.toString();
      if (value.length <= 100) {
        return value;
      }
      return value.substring(0, 100) + "...";
    },
  },
  methods: {
    fetchHotels() {
      this.loading = true;
      const { sortBy, sortDesc, page, itemsPerPage } = this.options;
      HotelService.getAllHotels(page - 1, itemsPerPage).then((response) => {
        this.hotels = response.data.Results;
        this.totalResults = response.data.TotalResults;
      });
    },
    exploreRooms(hotelId) {
      this.$router.push("/hotels-admin/" + hotelId);
    },
    deleteHotel(hotelId) {
      HotelService.deleteHotel(hotelId).then((response) => {
        this.fetchHotels();
      });
    },
    updateHotel(hotelId) {},
    createHotel(hotelId) {},
  },
};
</script>
