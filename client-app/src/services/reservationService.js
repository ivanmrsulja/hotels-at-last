import axios from "axios";
import { BaseService } from "@/services/baseService.js";

class ReservationService extends BaseService {
  async getAllForUser(id) {
    return axios.get(this.basePath + "/reservations/user/" + id);
  }

  async getAllForRoom(id) {
    return axios.get(this.basePath + "/reservations/" + id);
  }

  async cancelReservation(id) {
    return axios.put(this.basePath + "/reservations/" + id + "/cancel");
  }

  async createReservation(requestBody) {
    return axios.post(this.basePath + "/reservations", requestBody);
  }
}

export default new ReservationService();
