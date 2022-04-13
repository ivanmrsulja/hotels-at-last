<template>
  <v-row justify="center">
    <v-dialog v-model="dialog" persistent max-width="600px">
      <template v-slot:activator="{ on, attrs }">
        <v-btn v-if="!edit" color="primary" dark v-bind="attrs" v-on="on">
          Add new room
        </v-btn>
        <v-btn v-if="edit" color="primary" dark v-bind="attrs" v-on="on">
          Edit
        </v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span v-if="!edit" class="text-h5">Add room</span>
          <span v-if="edit" class="text-h5">Edit room</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12">
                <v-text-field
                  label="Room number"
                  v-model="number"
                  required
                ></v-text-field>
              </v-col>
              <v-col cols="12">
                <v-text-field
                  label="Beds"
                  v-model="beds"
                  type="number"
                  required
                />
              </v-col>
              <v-col cols="12">
                <v-text-field
                  label="Price"
                  type="number"
                  v-model="price"
                  required
                />
              </v-col>
              <v-col cols="12" sm="4">
                <v-checkbox
                  label="Air conditioned"
                  v-model="airCond"
                  required
                ></v-checkbox>
              </v-col>
              <v-col cols="12" sm="4">
                <v-checkbox
                  label="Has parking"
                  v-model="parking"
                  required
                ></v-checkbox>
              </v-col>
              <v-col cols="12" sm="4">
                <v-checkbox label="Has TV" v-model="tv" required></v-checkbox>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" text @click="dialog = false">
            Close
          </v-btn>
          <v-btn v-if="!edit" color="blue darken-1" text @click="createRoom()">
            Add
          </v-btn>
          <v-btn v-if="edit" color="blue darken-1" text @click="updateRoom()">
            Update
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
export default {
  name: "add-room-form",
  props: ["edit", "roomData"],
  data() {
    return {
      dialog: false,
      number: "",
      beds: 0,
      price: 0,
      airCond: false,
      parking: false,
      tv: false,
    };
  },
  mounted() {
    if (this.edit) {
      this.number = this.roomData.RoomNumber;
      this.beds = this.roomData.NumberOfBeds;
      this.price = this.roomData.Price;
      this.airCond = this.roomData.AirConditioned;
      this.parking = this.roomData.HasParkingSpace;
      this.tv = this.roomData.HasTV;
    }
  },
  methods: {
    createRoom() {
      this.$root.$emit("createRoom", this.createRequest());
      this.dialog = false;
    },
    updateRoom() {
      this.$root.$emit("updateRoom", {
        requestBody: this.createRequest(),
        roomId: this.roomData.Id,
      });
      this.dialog = false;
    },
    createRequest() {
      return {
        RoomNumber: this.number,
        NumberOfBeds: parseInt(this.beds),
        Price: parseFloat(this.price),
        AirConditioned: this.airCond,
        HasParkingSpace: this.parking,
        HasTV: this.tv,
      };
    },
  },
};
</script>
