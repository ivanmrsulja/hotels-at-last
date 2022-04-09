import axios from "axios";
import { BaseService } from "@/services/baseService.js";

class HotelService extends BaseService {
  async getAllHotels(page, size) {
    return axios.get(this.basePath + "/hotels?page=" + page + "&size=" + size);
  }

  async getHotel(id) {
    return axios.get(this.basePath + "/hotels/" + id);
  }

  async getRoom(id) {
    return axios.get(this.basePath + "/rooms/" + id);
  }

  async searchHotels(page, size, params) {
    return axios.get(
      this.basePath +
        "/hotels?page=" +
        page +
        "&size=" +
        size +
        "&bedsFrom=" +
        params.bedsFrom +
        "&bedsTo=" +
        params.bedsTo +
        "&priceFrom=" +
        params.priceFrom +
        "&priceTo=" +
        params.priceTo +
        "&airCond=" +
        params.airCond +
        "&parking=" +
        params.parking +
        "&tv=" +
        params.tv +
        "&name=" +
        params.name +
        "&address=" +
        params.address
    );
  }
}

export default new HotelService();
