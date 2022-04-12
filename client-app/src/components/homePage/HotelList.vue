<template>
  <v-container>
    <hotel-entry
      v-for="hotel in hotels"
      :key="hotel.Id"
      :hotel="hotel"
    ></hotel-entry>
    <v-flex class="text-center">
      <div class="pagination">
        <a href="#" @click.stop.prevent="newPage(pageIndex - 1)">&laquo;</a>
        <a
          href="#"
          v-for="index in totalPages"
          :key="index"
          @click.stop.prevent="newPage(index - 1)"
          :class="pageIndex === index - 1 ? 'active' : ''"
          >{{ index }}</a
        >
        <a href="#" @click.stop.prevent="newPage(pageIndex + 1)">&raquo;</a>
      </div>
    </v-flex>
  </v-container>
</template>

<script>
import HotelService from "../../services/hotelService.js";
import HotelEntry from "./HotelEntry.vue";

export default {
  name: "hotel-list",
  components: { HotelEntry },
  data() {
    return {
      pageSize: 2,
      pageIndex: 0,
      totalPages: 1,
      totalItems: 0,
      hotels: [],
      search: false,
      searchParams: {},
    };
  },
  mounted() {
    this.fetchData();
    this.$root.$on("search", (data) => {
      this.search = true;
      this.searchParams = data;
      this.pageIndex = 0;
      this.fetchData();
    });
    this.$root.$on("resetSearch", () => {
      this.search = false;
      this.pageIndex = 0;
      this.fetchData();
    });
  },
  methods: {
    fetchData() {
      if (this.search) {
        HotelService.searchHotels(
          this.pageIndex,
          this.pageSize,
          this.searchParams
        ).then((response) => {
          this.hotels = response.data.Results;
          this.totalPages = Math.ceil(
            response.data.TotalResults / this.pageSize
          );
          this.$root.$emit("refreshImages");
        });
      } else {
        HotelService.getAllHotels(this.pageIndex, this.pageSize).then(
          (response) => {
            this.hotels = response.data.Results;
            this.totalPages = Math.ceil(
              response.data.TotalResults / this.pageSize
            );
            this.$root.$emit("refreshImages");
          }
        );
      }
    },
    newPage(index) {
      if (index < 0 || index >= this.totalPages) {
        return;
      }
      this.pageIndex = index;
      this.fetchData();
    },
  },
};
</script>

<style scoped>
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
