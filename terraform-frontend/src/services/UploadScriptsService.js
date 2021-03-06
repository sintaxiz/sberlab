import http from "../http-common";

class UploadScriptsService {
  upload(file, onUploadProgress) {
    let formData = new FormData();

    formData.append("myScript", file);

    return http.post("/upload", formData, {
      headers: {
        "Content-Type": "multipart/form-data"
      },
      onUploadProgress
    });
  }
  uploadInfo() {
    return http.get("/upload");
  }
}

export default new UploadScriptsService();
