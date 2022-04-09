<template>
  <v-container>
    <br />
    <v-flex class="text-center">
      <h1>Reservation list</h1>
    </v-flex>
    <br />
    <br />
    <v-data-table
      :headers="headers"
      :items="reservations"
      :items-per-page="5"
      class="elevation-1"
    >
      <template v-slot:item="row">
        <tr>
          <td>{{ row.item.hotel_name }}</td>
          <td>{{ row.item.room_number }}</td>
          <td>{{ row.item.username }}</td>
          <td>{{ row.item.start }}</td>
          <td>{{ row.item.end }}</td>
          <td v-if="row.item.cancelled"><h2>CANCELLED</h2></td>
          <td
            v-if="
              !row.item.cancelled &&
              new Date(row.item.start).getTime() - 172800000 <=
                new Date().getTime()
            "
          >
            <h2>PASSED</h2>
          </td>
          <td
            v-if="
              !row.item.cancelled &&
              new Date(row.item.start).getTime() - 172800000 >
                new Date().getTime()
            "
          >
            <v-btn color="primary" @click="cancelReservation(row.item)">
              CANCEL
            </v-btn>
          </td>
        </tr>
      </template>
    </v-data-table>
  </v-container>
</template>

<script>
import jwtDecode from "jwt-decode";
import ReservationService from "../services/reservationService.js";

export default {
  name: "reservation-list-view",
  data() {
    return {
      reservations: [],
      headers: [
        {
          text: "Hotel",
          align: "start",
          sortable: false,
          value: "hotel_name",
        },
        { text: "Room", value: "room_number" },
        { text: "User", value: "username" },
        { text: "From", value: "start" },
        { text: "To", value: "end" },
        { text: "STATUS", value: "cancelled" },
      ],
    };
  },
  mounted() {
    this.fetchData();
  },
  methods: {
    cancelReservation(reservation) {
      ReservationService.cancelReservation(reservation.id).then((response) => {
        this.fetchData();
      });
    },
    fetchData() {
      let decodedToken = jwtDecode(localStorage.getItem("jwt"));
      ReservationService.getAllForUser(decodedToken.Id).then((response) => {
        this.reservations = response.data;
      });
    },
  },
};
</script>
