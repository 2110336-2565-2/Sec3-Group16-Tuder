import api from './apiHandler';

export const addCourseHandler = (data) => {
    return api.post('api/v1/course', data, {
            headers:
            {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + localStorage.getItem('jwtToken')
            }
    });
}