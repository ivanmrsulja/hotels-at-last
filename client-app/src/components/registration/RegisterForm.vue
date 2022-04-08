<template>
  <v-container fluid>
    <v-row justify="center">
      <v-col lg="4" sm="12">
        <v-form ref="form" v-model="valid">
          <v-text-field
            v-model="name"
            :rules="nameRules"
            label="Name"
            required
          ></v-text-field>

          <v-text-field
            v-model="surname"
            :rules="surnameRules"
            label="Surname"
            required
          ></v-text-field>

          <v-text-field
            v-model="email"
            :rules="emailRules"
            label="Email"
            required
          ></v-text-field>

          <v-text-field
            v-model="username"
            :rules="usernameRules"
            label="Username"
            required
          ></v-text-field>

          <v-text-field
            v-model="password"
            :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
            :type="show ? 'text' : 'password'"
            :rules="passwordRules"
            label="Password"
            @click:append="show = !show"
            counter
            required
          ></v-text-field>

          <v-text-field
            v-model="confirmPassword"
            :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
            :type="show ? 'text' : 'password'"
            :rules="[checkMatching, ...confirmPasswordRules]"
            label="Confirm password"
            @click:append="show = !show"
            counter
            required
          ></v-text-field>

          <v-flex class="text-center">
            <v-btn
              :disabled="!valid"
              color="success"
              class="mr-4"
              @click="register"
            >
              Register
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
import AuthenticationService from "../../services/authenticationService.js";

export default {
  name: "RegisterForm",
  data: () => ({
    snackbar: false,
    text: "Invalid credentials.",
    timeout: 2000,
    valid: true,
    show: false,
    name: "",
    surname: "",
    username: "",
    confirmPassword: "",
    usernameRules: [
      (v) => !!v || "You need to insert a username",
      (v) => (v && v.trim() !== "") || "You need to insert a username",
    ],
    nameRules: [
      (v) => !!v || "You need to insert a name",
      (v) => (v && v.trim() !== "") || "You need to insert a name",
      (v) =>
        (v && v.length <= 20) || "Name cannot be longer than 20 characters",
    ],
    surnameRules: [
      (v) => !!v || "You need to insert a surname",
      (v) => (v && v.trim() !== "") || "You need to insert a surname",
      (v) =>
        (v && v.length <= 20) || "Surname cannot be longer than 20 characters",
    ],
    email: "",
    password: "",
    emailRules: [
      (v) => !!v || "You need to insert an email",
      (v) => /.+@.+\..+/.test(v) || "Email address is not valid",
    ],
    passwordRules: [
      (v) => !!v || "You need to insert a password",
      (v) => (v && v.trim() !== "") || "You need to insert a password",
      (v) => (v && v.length >= 8) || "Password must be 8 characters or longer",
    ],
    confirmPasswordRules: [(v) => !!v || "You need to confirm password"],
  }),

  methods: {
    register() {
      let registrationRequest = {
        Name: this.name,
        Surname: this.surname,
        Email: this.email,
        Password: this.password,
        Username: this.username,
      };
      let that = this;
      let response = AuthenticationService.register(registrationRequest);
      response
        .then(() => {
          that.text = "Registered successfully";
          that.snackbar = true;
          that.$router.push("/login");
        })
        .catch((err) => {
          if (err.message.includes("400")) {
            that.text = "User with that email address allready exists.";
          }
          that.snackbar = true;
        });
    },
    checkMatching(value) {
      if (value !== this.password) {
        return "Passwords must match";
      }
      return true;
    },
  },
};
</script>
