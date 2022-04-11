<template>
  <v-container>
    <v-data-table
      :headers="headers"
      :items="reviews"
      :options.sync="options"
      :items-per-page="1"
      :server-items-length="totalResults"
      class="elevation-1"
    >
      <template v-slot:item="row">
        <tr>
          <td>{{ row.item.Comment }}</td>
          <td>{{ row.item.Rating }}</td>
          <td>{{ row.item.TimesReported }}</td>
          <td>{{ row.item.UserName }}</td>
          <td align="right">
            <v-btn
              class="mx-2"
              dark
              small
              color="primary"
              @click="dismiss(row.item)"
            >
              Dismiss
            </v-btn>
          </td>
          <td align="right">
            <ban-form :review="row.item"></ban-form>
          </td>
          <td align="right">
            <v-btn
              class="mx-2"
              dark
              small
              color="primary"
              @click="deleteReview(row.item)"
            >
              Delete
            </v-btn>
          </td>
        </tr>
      </template>
    </v-data-table>
    <v-snackbar v-model="snackbar" :timeout="timeout">
      {{ text }}

      <template v-slot:action="{ attrs }">
        <v-btn color="blue" text v-bind="attrs" @click="snackbar = false">
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </v-container>
</template>

<script>
import ReviewService from "../../services/reviewService.js";
import UserService from "../../services/userService.js";
import BanForm from "./banForm.vue";

export default {
  name: "reported-comments-list",
  components: { BanForm },
  data() {
    return {
      reviews: [],
      page: 0,
      totalResults: 10,
      totalPages: 1,
      options: {},
      loading: true,
      text: "",
      snackbar: false,
      timeout: 2000,
      headers: [
        {
          text: "Comment",
          align: "start",
          sortable: false,
          value: "Comment",
        },
        { text: "Rating", value: "Rating" },
        { text: "Times reported", value: "TimesReported" },
        { text: "User", value: "UserName" },
        { text: "Action", align: "end" },
        { text: "Action", align: "end" },
        { text: "Action", align: "end" },
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
  mounted() {
    this.$root.$on("banUser", (request) => {
      this.banUser(request);
    });
  },
  methods: {
    getDataFromApi() {
      this.loading = true;
      const { sortBy, sortDesc, page, itemsPerPage } = this.options;
      ReviewService.getReportedReviews(page - 1, itemsPerPage).then(
        (response) => {
          this.reviews = response.data.Results;
          this.totalResults = response.data.TotalResults;
        }
      );
    },
    dismiss(review) {
      ReviewService.dismissReports(review.Id).then((response) => {
        this.getDataFromApi();
      });
    },
    banUser(request) {
      console.log(request);
      UserService.banUser(request.userId, { EndOfBan: request.endOfBan }).then(
        () => {
          this.text =
            "User " + request.username + " banned until " + request.endOfBan;
          this.snackbar = true;
        }
      );
    },
    deleteReview(review) {
      ReviewService.deleteReview(review.Id).then((response) => {
        this.getDataFromApi();
      });
    },
  },
};
</script>
