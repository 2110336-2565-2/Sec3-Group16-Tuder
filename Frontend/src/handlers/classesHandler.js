import api from "./apiHandler";

export const fetchClassByIdHandler = (id) => {
  return api.get(`api/v1/matches/${id}`);
};
