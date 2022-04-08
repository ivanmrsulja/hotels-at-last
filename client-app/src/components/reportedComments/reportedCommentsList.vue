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
            <v-dialog v-model="dialog" persistent max-width="600px">
              <template v-slot:activator="{ on, attrs }">
                <v-btn color="primary" dark small v-bind="attrs" v-on="on">
                  Ban User
                </v-btn>
              </template>
              <v-card>
                <v-card-title>
                  <span class="text-h5">Ban Until</span>
                </v-card-title>
                <v-card-text>
                  <v-container>
                    <v-row>
                      <v-col cols="12">
                        <v-menu
                          ref="menu"
                          v-model="menu"
                          :close-on-content-click="false"
                          :return-value.sync="date"
                          transition="scale-transition"
                          offset-y
                          min-width="auto"
                        >
                          <template v-slot:activator="{ on, attrs }">
                            <v-text-field
                              v-model="date"
                              label="Picker in menu"
                              prepend-icon="mdi-calendar"
                              readonly
                              v-bind="attrs"
                              v-on="on"
                            ></v-text-field>
                          </template>
                          <v-date-picker v-model="date" no-title scrollable>
                            <v-spacer></v-spacer>
                            <v-btn text color="primary" @click="menu = false">
                              Cancel
                            </v-btn>
                            <v-btn
                              text
                              color="primary"
                              @click="$refs.menu.save(date)"
                            >
                              OK
                            </v-btn>
                          </v-date-picker>
                        </v-menu>
                      </v-col>
                    </v-row>
                  </v-container>
                </v-card-text>
                <v-card-actions>
                  <v-spacer></v-spacer>
                  <v-btn color="blue darken-1" text @click="dialog = false">
                    Close
                  </v-btn>
                  <v-btn
                    color="blue darken-1"
                    text
                    @click="
                      dialog = false;
                      banUser(row.item);
                    "
                  >
                    Save
                  </v-btn>
                </v-card-actions>
              </v-card>
            </v-dialog>
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

export default {
  name: "reported-comments-list",
  data() {
    return {
      date: new Date(Date.now() - new Date().getTimezoneOffset() * 60000)
        .toISOString()
        .substr(0, 10),
      menu: false,
      reviews: [],
      page: 0,
      totalResults: 10,
      totalPages: 1,
      options: {},
      loading: true,
      dialog: false,
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
  mounted() {},
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
    banUser(review) {
      let banRequest = {
        EndOfBan: this.date,
      };
      UserService.banUser(review.UserId, banRequest).then(() => {
        this.text = "User " + review.UserName + " banned until " + this.date;
        this.snackbar = true;
      });
    },
    deleteReview(review) {
      ReviewService.deleteReview(review.Id).then((response) => {
        this.getDataFromApi();
      });
    },
  },
};
</script>
