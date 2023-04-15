import api from "./apiHandler";

export const fetchAdminTuitionFeeHandler = () => {
  return api.get("api/v1/admin-tuition-fee", {
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + localStorage.getItem("jwtToken"),
    }
  })
}