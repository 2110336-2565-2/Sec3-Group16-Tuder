import api from './apiHandler';

export const fetchStudent = (id) => {
    return api.get('api/v1/match/course/'+ id, {
        headers:
        {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('jwtToken')
        }
    })
}

