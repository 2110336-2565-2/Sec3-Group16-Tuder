import api from './apiHandler';

export const changeStatus = (course_id,data) => {
    return api.put('api/v1/course/status/'+course_id,data)
}