<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <div class="d-flex align-center">
        <v-btn text>
          <h1 class="mr-2">HOTELS AT LAST</h1>
        </v-btn>
      </div>
      <v-spacer></v-spacer>
      <v-btn v-if="!loggedIn || (loggedIn && role === 0)" href="#/" text>
        HOME
      </v-btn>
      <v-btn v-if="loggedIn && role === 0" href="#/reservations" text>
        My reservations
      </v-btn>
      <v-btn v-if="!loggedIn" href="#/register" text> Register </v-btn>
      <v-btn v-if="!loggedIn" href="#/login" text> Login </v-btn>
      <v-btn v-if="loggedIn && role === 1" href="#/reported-comments" text>
        Reported comments
      </v-btn>
      <v-btn v-if="loggedIn && role === 1" href="#/hotels-admin" text>
        Administration
      </v-btn>
      <v-btn v-if="loggedIn" v-on:click="logout" text>Logout</v-btn>
    </v-app-bar>

    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import jwtDecode from "jwt-decode";
import AuthenticationService from "./services/authenticationService.js";

export default {
  name: "App",
  data() {
    return {
      loggedIn: false,
      role: 3,
    };
  },
  mounted() {
    this.loggedIn = AuthenticationService.userLoggedIn();
    if (this.loggedIn) {
      let decodedToken = jwtDecode(localStorage.getItem("jwt"));
      this.role = decodedToken.role;
    }
    this.$root.$on("loginSuccess", (role) => {
      this.loggedIn = AuthenticationService.userLoggedIn();
      this.role = role;
    });
  },
  methods: {
    logout() {
      localStorage.removeItem("jwt");
      this.loggedIn = false;
      this.$router.push("/login");
    },
  },
};
</script>
