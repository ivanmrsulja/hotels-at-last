import axios from "axios";
import { BaseService } from "@/services/baseService.js";

class AuthenticationService extends BaseService {
  async login(loginRequest) {
    return axios.post(this.basePath + "/users/login", loginRequest);
  }

  async register(registrationRequest) {
    return axios.post(this.basePath + "/users/register", registrationRequest);
  }

  userLoggedIn() {
    let jwt = localStorage.getItem("jwt");
    if (jwt) {
      return true;
    } else {
      return false;
    }
  }
}

export default new AuthenticationService();
