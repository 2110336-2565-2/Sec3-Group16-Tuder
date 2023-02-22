import api from './apiHandler';

export const fetchCourseHandler = () => {
    return api.get('/courses')
}