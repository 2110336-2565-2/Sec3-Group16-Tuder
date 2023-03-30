import api from './apiHandler';



export const fetchCancellingRequestsHandler = () => {
    return api.get('api/v1/cancelling-requests', {
        headers:
        {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('jwtToken')
        }
    })
}

export const submitAudittingHandler = ( data) => {
    return api.post('api/v1/audit-request', data, {
        headers:
        {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('jwtToken')
        }
    })
}
