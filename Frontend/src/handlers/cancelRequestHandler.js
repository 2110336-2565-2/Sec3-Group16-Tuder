import api from './apiHandler';



export const fetchCancellingRequestHandler = (id) => {
    return api.get('api/v1/cancelling-request/'+ id, {
        headers:
        {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('jwtToken')
        }
    })
}