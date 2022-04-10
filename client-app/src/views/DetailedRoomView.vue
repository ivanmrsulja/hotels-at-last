<template>
  <v-container>
    <v-flex class="text-center">
      <h1>Room {{ room.RoomNumber }}</h1>
    </v-flex>
    <br />
    <br />
    <v-row align="center" justify="center">
      <v-col class="col-lg-8">
        <table class="info-table">
          <tr>
            <th>Beds</th>
            <th>Price</th>
            <th>Air Conditioning</th>
            <th>Parking</th>
            <th>TV</th>
          </tr>
          <tr>
            <td>{{ room.NumberOfBeds }}</td>
            <td>{{ room.Price }}</td>
            <td>
              <v-icon v-if="room.AirConditioned">mdi-check-circle</v-icon>
              <v-icon v-if="!room.AirConditioned">mdi-close-thick</v-icon>
            </td>
            <td>
              <v-icon v-if="room.HasParkingSpace">mdi-check-circle</v-icon>
              <v-icon v-if="!room.HasParkingSpace">mdi-close-thick</v-icon>
            </td>
            <td>
              <v-icon v-if="room.HasTV">mdi-check-circle</v-icon>
              <v-icon v-if="!room.HasTV">mdi-close-thick</v-icon>
            </td>
          </tr>
        </table>
        <br />
        <br />
        <h2>Reservations:</h2>
        <br />
        <v-data-table
          :headers="headers"
          :items="reservations"
          :items-per-page="5"
          class="elevation-1"
        ></v-data-table>
        <br />
        <br />
        <reservation-form></reservation-form>
        <br />
        <h2>Reviews:</h2>
        <br />
        <div class="comment" v-for="review in reviews" v-bind:key="review.Id">
          <div class="comment-heading">
            <img
              src="@/assets/user.png"
              alt="Avatar"
              style="border-radius: 50%; height: 30px; width: 30px"
            />
            <div class="comment-info">
              <a href="#" class="comment-author">{{ review.UserName }}</a>
              <p>
                {{ review.CreatedAt.toString().split("T")[0] }} (reported
                {{ review.TimesReported }} times)
              </p>
            </div>
          </div>
          <div class="comment-body">
            <p>
              {{ review.Comment }}
              <v-rating
                length="5"
                size="15"
                :value="review.Rating"
                readonly
              ></v-rating>
            </p>

            <button type="button" @click="reportReview(review.Id)">
              Report
            </button>
          </div>
        </div>
        <div class="pagination">
          <a href="#" @click.stop.prevent="newPage(reviewPageIndex - 1)"
            >&laquo;</a
          >
          <a
            href="#"
            v-for="index in reviewTotalPages"
            :key="index"
            @click.stop.prevent="newPage(index - 1)"
            :class="reviewPageIndex === index - 1 ? 'active' : ''"
            >{{ index }}</a
          >
          <a href="#" @click.stop.prevent="newPage(reviewPageIndex + 1)"
            >&raquo;</a
          >
        </div>
        <br />
        <br />
        <div>
          <textarea
            cols="53"
            rows="3"
            v-model="commentText"
            placeholder="Add a review..."
            style="border: solid 1px black"
          ></textarea>
          <div>
            <p style="margin-bottom: -10px">Rate your experience:</p>
            <v-rating
              length="5"
              size="25"
              :value="commentRating"
              @input="setRating($event)"
              hover
            ></v-rating>
          </div>
          <v-btn color="blue" dark @click="postReview()"> Post review </v-btn>
        </div>
      </v-col>
    </v-row>

    <v-snackbar v-model="snackbar" :timeout="timeout">
      {{ text }}

      <template v-slot:action="{ attrs }">
        <v-btn color="blue" dark v-bind="attrs" @click="snackbar = false">
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </v-container>
</template>

<script>
import ReservationForm from "../components/detailedHotel/reservationForm.vue";
import HotelService from "../services/hotelService.js";
import ReservationService from "../services/reservationService.js";
import ReviewService from "../services/reviewService.js";
import jwtDecode from "jwt-decode";

export default {
  name: "detailed-room-view",
  components: { ReservationForm },
  data() {
    return {
      room: {},
      reservations: [],
      reviews: [],
      headers: [
        {
          text: "User",
          align: "start",
          sortable: false,
          value: "username",
        },
        { text: "From", value: "start" },
        { text: "To", value: "end" },
      ],
      reviewPageIndex: 0,
      reviewPageSize: 2,
      reviewTotalPages: 1,
      commentText: "",
      commentRating: 5,
      snackbar: false,
      text: "",
      timeout: 2000,
    };
  },
  mounted() {
    HotelService.getRoom(this.$route.params.roomId).then((response) => {
      this.room = response.data;
    });
    this.fetchReservations();
    this.fetchReviews();

    this.$root.$on("reservation", (timeframe) => {
      HotelService.getHotel(this.$route.params.id).then((response) => {
        let request = {
          user_id: jwtDecode(localStorage.getItem("jwt")).Id,
          room_id: parseInt(this.$route.params.roomId),
          start: timeframe.start,
          end: timeframe.end,
          username: jwtDecode(localStorage.getItem("jwt")).username,
          hotel_name: response.data.Name,
          room_number: parseInt(this.room.RoomNumber),
        };
        ReservationService.createReservation(request).then((response) => {
          console.log(response);
          this.text = response.data.message;
          this.snackbar = true;
          this.fetchReservations();
        });
      });
    });
  },
  methods: {
    newPage(index) {
      if (index < 0 || index >= this.reviewTotalPages) {
        return;
      }
      this.reviewPageIndex = index;
      this.fetchReviews();
    },
    fetchReservations() {
      ReservationService.getAllForRoom(this.$route.params.roomId).then(
        (response) => {
          this.reservations = response.data;
        }
      );
    },
    fetchReviews() {
      ReviewService.getAllForRoom(
        this.$route.params.roomId,
        this.reviewPageIndex,
        this.reviewPageSize
      ).then((response) => {
        this.reviews = response.data.Results;
        this.reviewTotalPages = Math.ceil(
          response.data.TotalResults / this.reviewPageSize
        );
      });
    },
    reportReview(id) {
      ReviewService.reportReview(id).then((response) => {
        this.fetchReviews();
        this.text = "Comment reported";
        this.snackbar = true;
      });
    },
    setRating(value) {
      this.commentRating = value;
    },
    postReview() {
      let decodedToken = jwtDecode(localStorage.getItem("jwt"));
      let createRequest = {
        Comment: this.commentText,
        Rating: this.commentRating,
        UserId: decodedToken.Id,
        RoomId: this.room.Id,
      };
      this.commentText = "";
      this.commentRating = 5;
      ReviewService.createReview(createRequest)
        .then((response) => {
          this.text = "Posted successfully.";
          this.snackbar = true;
          this.fetchReviews();
        })
        .catch((err) => {
          this.text = err.response.data.Message;
          this.snackbar = true;
        });
    },
  },
};
</script>

<style scoped>
.info-table {
  font-family: Arial, Helvetica, sans-serif;
  border-collapse: collapse;
  width: 100%;
}

.info-table td,
.info-table th {
  border: 1px solid #ddd;
  padding: 8px;
}

.info-table tr:nth-child(even) {
  background-color: #f2f2f2;
}

.info-table tr:hover {
  background-color: #ddd;
}

.info-table th {
  padding-top: 12px;
  padding-bottom: 12px;
  text-align: left;
  background-color: #4185f2;
  color: white;
}

button {
  -moz-appearance: none;
  -webkit-appearance: none;
  appearance: none;
  font-size: 14px;
  padding: 4px 8px;
  color: rgba(0, 0, 0, 0.85);
  background-color: #fff;
  border: 1px solid rgba(0, 0, 0, 0.2);
  border-radius: 4px;
}
button:hover,
button:focus,
button:active {
  cursor: pointer;
  background-color: #ecf0f1;
}
.comment-thread {
  width: 700px;
  max-width: 100%;
  margin: auto;
  padding: 0 30px;
  background-color: #fff;
  border: 1px solid transparent; /* Removes margin collapse */
}
.m-0 {
  margin: 0;
}
.sr-only {
  position: absolute;
  left: -10000px;
  top: auto;
  width: 1px;
  height: 1px;
  overflow: hidden;
}

/* Comment */

.comment {
  position: relative;
  margin: 20px auto;
}
.comment-heading {
  display: flex;
  align-items: center;
  height: 50px;
  font-size: 14px;
}
.comment-info {
  color: rgba(0, 0, 0, 0.5);
  margin-left: 10px;
}
.comment-author {
  color: rgba(0, 0, 0, 0.85);
  font-weight: bold;
  text-decoration: none;
}
.comment-author:hover {
  text-decoration: underline;
}

/* PAGINATION */
.pagination {
  display: inline-block;
}

.pagination a {
  color: black;
  float: left;
  padding: 8px 16px;
  text-decoration: none;
}

.pagination a.active {
  background-color: #4caf50;
  color: white;
  border-radius: 5px;
}

.pagination a:hover:not(.active) {
  background-color: #ddd;
  border-radius: 5px;
}
</style>
