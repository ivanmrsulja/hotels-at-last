<template>
  <v-container>
    <v-flex class="text-center">
      <h1>HOME PAGE</h1>
    </v-flex>
    <search-nav></search-nav>
    <br />
    <hotel-list></hotel-list>
  </v-container>
</template>

<script>
import HotelList from "@/components/homePage/HotelList.vue";
import SearchNav from "@/components/homePage/SearchNav.vue";
export default {
  components: { HotelList, SearchNav },
  name: "home-view",
  data() {
    return {
      //
    };
  },
};
</script>
