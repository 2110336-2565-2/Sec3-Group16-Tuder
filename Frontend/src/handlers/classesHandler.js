import api from "./apiHandler";

export const fetchClassByStudentIdHandler = (student_id) => {
  return api.get(`api/v1/appointment/${student_id}`);
};
