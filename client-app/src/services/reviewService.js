import axios from "axios";
import { BaseService } from "@/services/baseService.js";

export class ReviewService extends BaseService {
  async getReportedReviews(page, size) {
    return axios.get(
      this.basePath + "/reviews/reported?page=" + page + "&size=" + size + ""
    );
  }

  async getAllForRoom(id, page, size) {
    return axios.get(
      this.basePath + "/reviews/rooms/" + id + "?page=" + page + "&size=" + size
    );
  }

  async dismissReports(id) {
    return axios.patch(this.basePath + "/reviews/" + id + "/dismiss");
  }

  async reportReview(id) {
    return axios.patch(this.basePath + "/reviews/" + id + "/report");
  }

  async deleteReview(id) {
    return axios.delete(this.basePath + "/reviews/" + id);
  }

  async createReview(requestBody) {
    return axios.post(this.basePath + "/reviews", requestBody);
  }
}

export default new ReviewService();
