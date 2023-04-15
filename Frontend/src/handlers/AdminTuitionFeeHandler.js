import api from "./apiHandler";

export const fetchAdminTuitionFeeHandler = () => {
  return api.get("api/v1/admintuitionfees", {
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + localStorage.getItem("jwtToken"),
    }
  })
}