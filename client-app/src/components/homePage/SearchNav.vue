<template>
  <v-container>
    <div id="mySidenav" class="sidenav">
      <a href="javascript:void(0)" class="closebtn" @click="closeNav()"
        >&times;</a
      >
      <table style="margin-top: 20px">
        <tr>
          <td>Name:</td>
          <td><input type="text" v-model="name" /></td>
        </tr>
        <tr>
          <td>Address:</td>
          <td><input type="text" v-model="address" /></td>
        </tr>
        <tr>
          <td>Beds from:</td>
          <td><input type="number" v-model="bedsFrom" /></td>
        </tr>
        <tr>
          <td>Beds to:</td>
          <td><input type="number" v-model="bedsTo" /></td>
        </tr>
        <tr>
          <td>Stars from:</td>
          <td><input type="number" v-model="starsFrom" /></td>
        </tr>
        <tr>
          <td>Stars to:</td>
          <td><input type="number" v-model="starsTo" /></td>
        </tr>
        <tr>
          <td>Price from:</td>
          <td><input type="number" v-model="priceFrom" /></td>
        </tr>
        <tr>
          <td>Price to:</td>
          <td><input type="number" v-model="priceTo" /></td>
        </tr>
        <tr>
          <td>Air conditioning:</td>
          <td><input type="checkbox" v-model="airCond" /></td>
        </tr>
        <tr>
          <td>Parking:</td>
          <td><input type="checkbox" v-model="parking" /></td>
        </tr>
        <tr>
          <td>TV:</td>
          <td><input type="checkbox" v-model="tv" /></td>
        </tr>
      </table>
      <v-flex class="text-center">
        <v-btn color="white lighten-1" text @click="submitSearch()">
          Search
        </v-btn>
        <v-btn color="white lighten-1" text @click="resetSearch()">
          Reset
        </v-btn>
      </v-flex>
    </div>
    <v-flex class="text-center">
      <span style="font-size: 35px; cursor: pointer" @click="openNav()"
        >&#9776; Search Hotels</span
      >
    </v-flex>
  </v-container>
</template>

<script>
export default {
  name: "search-nav",
  data() {
    return {
      name: "",
      address: "",
      bedsFrom: 0,
      bedsTo: 100,
      priceFrom: 0,
      priceTo: 10000,
      airCond: false,
      parking: false,
      tv: false,
      starsFrom: 1,
      starsTo: 5,
    };
  },
  methods: {
    openNav() {
      document.getElementById("mySidenav").style.width = "300px";
    },
    closeNav() {
      document.getElementById("mySidenav").style.width = "0";
    },
    submitSearch() {
      let searchRequest = {
        name: this.name,
        address: this.address,
        bedsFrom: this.bedsFrom,
        bedsTo: this.bedsTo,
        priceFrom: this.priceFrom,
        priceTo: this.priceTo,
        airCond: this.airCond,
        parking: this.parking,
        tv: this.tv,
        starsFrom: this.starsFrom,
        starsTo: this.starsTo,
      };
      this.$root.$emit("search", searchRequest);
    },
    resetSearch() {
      this.$root.$emit("resetSearch");
    },
  },
};
</script>

<style scoped>
body {
  font-family: "Lato", sans-serif;
}

.sidenav {
  height: 100%;
  width: 0;
  position: fixed;
  z-index: 1;
  top: 0;
  left: 0;
  background-color: #111;
  overflow-x: hidden;
  transition: 0.5s;
  margin-top: 60px;
  padding-top: 60px;
  color: white;
}

.sidenav input {
  background-color: white;
  cursor: pointer;
  color: #818181;
  margin-bottom: 10px;
  margin-left: 10px;
  margin-right: 10px;
}

.sidenav a {
  padding: 8px 8px 8px 32px;
  text-decoration: none;
  font-size: 25px;
  color: #818181;
  display: block;
  transition: 0.3s;
}

.sidenav a:hover {
  color: #f1f1f1;
}

.sidenav .closebtn {
  position: absolute;
  top: 0;
  right: 25px;
  font-size: 36px;
  margin-left: 50px;
}

@media screen and (max-height: 450px) {
  .sidenav {
    padding-top: 15px;
  }
  .sidenav a {
    font-size: 18px;
  }
}
</style>
