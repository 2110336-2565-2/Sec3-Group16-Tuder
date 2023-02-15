import axios from 'axios';

const api = axios.create({
    baseURL: process.env.REACT_APP_API_URL 
  });


async function signInHandler(event) {

    // Input validation
    //

    // Prevent the default action of the event
    event.preventDefault();

    // Get the form data from the event object
    let username = event.target.username.value;
    let password = event.target.password.value;
    
    // Post to api
    const response = await api.post('/api/v1/login', { username:username, password:password })
    .then(function(response){
        
        let res = response.data;
        
        // if login success, MOO will set token
        if (res.success === true) {
            let token = response.data.data.token;
            localStorage.setItem('jwtToken', token); 
            console.log(res.message);
        }

    }).catch(function(error){
        // if internal error occurs, MOO will return error message
        if (error.response) {
            let res = error.response.data;
            console.log(res.message);
        }
    });
};

export default signInHandler;
