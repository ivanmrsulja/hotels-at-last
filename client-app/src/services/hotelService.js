import axios from "axios";
import { BaseService } from "@/services/baseService.js";

class HotelService extends BaseService {
  async getAllHotels(page, size) {
    return axios.get(this.basePath + "/hotels?page=" + page + "&size=" + size);
  }

  async getAllRoomsForHotel(page, size, hotelId) {
    return axios.get(
      this.basePath +
        "/hotels/" +
        hotelId +
        "/rooms?page=" +
        page +
        "&size=" +
        size
    );
  }

  async getHotel(id) {
    return axios.get(this.basePath + "/hotels/" + id);
  }

  async getRoom(id) {
    return axios.get(this.basePath + "/rooms/" + id);
  }

  async getImage(image) {
    return axios.get(this.basePath + "/images/" + image);
  }

  async deleteHotel(id) {
    return axios.delete(this.basePath + "/hotels/" + id);
  }

  async deleteRoom(id) {
    return axios.delete(this.basePath + "/rooms/" + id);
  }

  async createRoom(hotelId, body) {
    return axios.post(this.basePath + "/hotels/" + hotelId + "/rooms", body);
  }

  async updateRoom(roomId, body) {
    return axios.put(this.basePath + "/rooms/" + roomId, body);
  }

  async createHotel(body) {
    return axios.post(this.basePath + "/hotels", body);
  }

  async updateHotel(hotelId, body) {
    return axios.put(this.basePath + "/hotels/" + hotelId, body);
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
        params.address +
        "&starsFrom=" +
        params.starsFrom +
        "&starsTo=" +
        params.starsTo
    );
  }
}

export default new HotelService();
