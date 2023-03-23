import api from './apiHandler';



export const fetchCancellingClassHandler = () => {
    return api.get('api/v1/cancelling-classes', {
        headers:
        {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('jwtToken')
        }
    })
}

export const submitApprovmentHandler = (classId) => {
    return api.post('api/v1/audit-class-cancellation', {
        classId: classId,
        approve: true
    }, {
        headers:
        {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('jwtToken')
        }
    })
}

export const submitRejectionHandler = (classId) => {
    return api.post('api/v1/audit-class-cancellation', {
        classId: classId,
        approve: false
    }, {
        headers:
        {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('jwtToken')
        }
    })
}