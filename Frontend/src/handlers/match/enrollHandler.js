import api from '../apiHandler'

export const submitEnrollFormHandler = ( data) => {
    return api.post('api/v1/match', data, {
        headers:
        {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('jwtToken')
        }
    })
}
