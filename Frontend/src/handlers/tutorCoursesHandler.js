import api from './apiHandler';

export const fetchTutorCourses = () => {
    return api.get('api/v1/tutor/courses',{
        headers:
        {
            'Content-Type': 'application/json',
            'Authorization' : 'Bearer ' + localStorage.getItem('jwtToken')
        }
    })
}