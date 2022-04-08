import axios from "axios";
import { BaseService } from "@/services/baseService.js";

class UserService extends BaseService {
  async banUser(id, body) {
    return axios.patch(this.basePath + "/users/" + id + "/ban", body);
  }
}

export default new UserService();
