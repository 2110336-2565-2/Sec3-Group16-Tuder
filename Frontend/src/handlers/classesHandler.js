import api from "./apiHandler";

export const fetchMatchesByIdHandler = (id) => {
  return api.get(`api/v1/matches/${id}`);
};

export const fetchClassDetailById = (id) => {
  return api.get(`api/v1/appointments/${id}`);
};

export const changeStatusHandler = (id, status) => {
  return api.put(`api/v1/appointment/updatestatus/${id}`, { status: status });
};
