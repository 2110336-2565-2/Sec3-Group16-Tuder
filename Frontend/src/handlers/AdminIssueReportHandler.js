import api from "./apiHandler";

export const fetchAdminIssueReportHandler = () => {
  return api.get("api/v1/issuereports", {
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + localStorage.getItem("jwtToken"),
    },
  });
};