import api from "./apiHandler";

export const fetchAdminIssueReportHandler = () => {
  return api.get("api/v1/issuereports", {
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + localStorage.getItem("jwtToken"),
    }
  })
}
export const submitSaveStateHandler = (data) => {
    return api.put('api/v1/issuereport/updatestatus/'+data.IssueReportId, data, {
        headers:
        {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('jwtToken')
        }
    })    
}
export const submitDeleteStateHandler = (data) => {
    return api.delete('api/v1/issuereport/'+data.IssueReportId, {
        headers:
        {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('jwtToken')
        }
    })    
}
