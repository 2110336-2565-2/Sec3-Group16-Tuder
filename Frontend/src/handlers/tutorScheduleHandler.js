import api from './apiHandler';
import { fetchCourseByIdHandler } from './searchCourseHandler';

export const fetchTutorSchedule = (courseID) => {
    return api.get('api/v1/tutor/schedule/'+ courseID)
}