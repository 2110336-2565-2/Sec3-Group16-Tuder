import axios from 'axios';
const api = axios.create({
    baseURL: 'http://localhost:8080' // replace with the URL of your locally running server
  });


async function signInController(username,password) {
        console.log(username)
        console.log(password)
        const response = await api.post('/api/v1/login', { username:username, password:password})
        .then(function(response){
            console.log(response)
            const token = response.data.data.token;
           
            localStorage.setItem('jwtToken', token); 
            const tokencheck = localStorage.getItem('jwtToken');
            console.log(tokencheck);
        }).catch(function(error){
            console.error(error);
        });      
};
export default signInController;