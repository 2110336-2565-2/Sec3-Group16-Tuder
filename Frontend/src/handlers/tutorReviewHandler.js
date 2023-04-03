import api from './apiHandler';

export const fetchTutorReviews = (username) => {
    return api.get('api/v1/tutor/'+ username + '/reviews')
}