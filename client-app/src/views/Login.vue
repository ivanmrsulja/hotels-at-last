<template>
  <v-container>
    <v-row justify="center">
      <v-col lg="4" sm="12">
        <v-form ref="form" v-model="valid" lazy-validation>
          <v-text-field
            v-model="email"
            :rules="emailRules"
            label="Email"
            required
          ></v-text-field>
          <v-text-field
            v-model="password"
            :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
            :rules="passwordRules"
            :type="show ? 'text' : 'password'"
            name="input-10-1"
            label="Insert password"
            counter
            @click:append="show = !show"
          ></v-text-field>
          <v-flex class="text-center">
            <v-btn
              :disabled="!valid"
              color="success"
              class="align-center"
              @click="login"
            >
              Login
            </v-btn>
          </v-flex>
        </v-form>
      </v-col>
    </v-row>
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
import Vue from "vue";
import AuthenticationService from "../services/authenticationService.js";
import jwt_decode from "jwt-decode";

export default Vue.extend({
  name: "login-view",
  data: () => ({
    snackbar: false,
    text: "Wrong email/password combination",
    timeout: 2000,
    valid: true,
    show: false,
    email: "",
    emailRules: [
      (v) => !!v || "You need to insert an email",
      (v) => /.+@.+\..+/.test(v) || "Invalid email address",
    ],
    password: "",
    passwordRules: [(p) => !!p || "You need to insert a password"],
  }),
  mounted: function () {
    if (AuthenticationService.userLoggedIn()) {
      // this.$router.push("/saglasnost");
    }
  },
  methods: {
    login() {
      let that = this;
      let loginRequest = {
        Email: that.email,
        Password: that.password,
      };
      let loginResponse = AuthenticationService.login(loginRequest);
      loginResponse
        .then((res) => {
          localStorage.setItem("jwt", res.data.Token);
          let decodedToken = jwt_decode(res.data.Token);
          that.$root.$emit("loginSuccess", decodedToken.role);
          if (decodedToken.role === 1) {
            that.$router.push("/hotels-admin");
          } else if (decodedToken.role === 0) {
            that.$router.push("/");
          }
        })
        .catch((err) => {
          that.text = err.response.data.Message;
          that.snackbar = true;
        });
    },
  },
});
</script>
