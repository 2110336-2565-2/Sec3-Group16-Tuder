import api from './apiHandler';

export default async function signUpHandler(signUpData , navigate){
    try{
        const response =  await api.post('/api/v1/signUp', signUpData)
        navigate('/sign-in')
        return response;
    }catch(error){
        if (error.response) {
            let res = error.response.data;
            throw Error(res.message);
        }
    }

};