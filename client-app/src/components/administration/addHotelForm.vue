<template>
  <v-row justify="center">
    <v-dialog v-model="dialog" persistent max-width="600px">
      <template v-slot:activator="{ on, attrs }">
        <v-btn v-if="!edit" color="primary" dark v-bind="attrs" v-on="on">
          Add new hotel
        </v-btn>
        <v-btn v-if="edit" color="primary" dark v-bind="attrs" v-on="on">
          Edit
        </v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span v-if="!edit" class="text-h5">Add hotel</span>
          <span v-if="edit" class="text-h5">Edit hotel</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12">
                <v-text-field
                  label="Name"
                  v-model="name"
                  required
                ></v-text-field>
              </v-col>
              <v-col cols="12">
                <v-text-field label="Address" v-model="address" required />
              </v-col>
              <v-col cols="12">
                <v-textarea
                  outlined
                  name="input-7-1"
                  label="Description"
                  v-model="description"
                ></v-textarea>
              </v-col>
              <v-col cols="12">
                <v-text-field
                  label="Stars"
                  v-model="stars"
                  type="number"
                  required
                />
              </v-col>
              <v-file-input
                label="JPG image"
                v-model="image"
                truncate-length="15"
              ></v-file-input>
            </v-row>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" text @click="dialog = false">
            Close
          </v-btn>
          <v-btn v-if="!edit" color="blue darken-1" text @click="createHotel()">
            Add
          </v-btn>
          <v-btn v-if="edit" color="blue darken-1" text @click="updateHotel()">
            Update
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
export default {
  name: "add-hotel-form",
  props: ["edit", "hotelData"],
  data() {
    return {
      dialog: false,
      name: "",
      address: "",
      description: "",
      stars: 1,
      image: null,
    };
  },
  mounted() {
    if (this.edit) {
      this.name = this.hotelData.Name;
      this.address = this.hotelData.Address;
      this.description = this.hotelData.Description;
      this.stars = this.hotelData.Stars;
    }
  },
  methods: {
    createHotel() {
      let reader = new FileReader();
      reader.readAsDataURL(this.image);
      reader.onload = () => {
        this.$root.$emit("createHotel", this.createRequest(reader.result));
        this.dialog = false;
      };
    },
    updateHotel() {
      if (this.image) {
        let reader = new FileReader();
        reader.readAsDataURL(this.image);
        reader.onload = () => {
          this.$root.$emit("updateHotel", {
            requestBody: this.createRequest(reader.result),
            hotelId: this.hotelData.Id,
          });
          this.dialog = false;
        };
      } else {
        this.$root.$emit("updateHotel", {
          requestBody: this.createRequest(""),
          hotelId: this.hotelData.Id,
        });
        this.dialog = false;
      }
    },
    createRequest(base64) {
      return {
        Name: this.name,
        Address: this.address,
        Description: this.description,
        Base64Image: base64,
        Stars: parseInt(this.stars),
      };
    },
  },
};
</script>
