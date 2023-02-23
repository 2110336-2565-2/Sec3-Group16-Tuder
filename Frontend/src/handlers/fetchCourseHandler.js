import api from './apiHandler';

export const fetchCourseHandler = () => {
    return api.get('api/v1/courses')
}