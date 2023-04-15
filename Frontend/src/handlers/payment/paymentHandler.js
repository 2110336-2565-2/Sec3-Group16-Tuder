import api from '../apiHandler'

export const fetchQRCodeForCoursePayment = (data) => {
    return api.post('api/v1/payment/getQRCodeForCoursePayment', data)
}

export const fetchQRCodeForTuitionFee = (data) => {
    return api.post('api/v1/payment/getQRCodeForTuitionFee', data)
}

export const changePaymentStatus = (data) => {
    return api.post('api/v1/payment/webhookChargeHandler', data)
}