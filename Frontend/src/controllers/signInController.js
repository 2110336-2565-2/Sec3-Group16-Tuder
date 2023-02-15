import axios from 'axios';

const api = axios.create({
    baseURL: process.env.REACT_APP_API_URL 
  });


async function loginHandler(event) {

    // Input validation
    //

    // Prevent the default action of the event
    event.preventDefault();

    let username = event.target.username;
    let password = event.target.password;
    console.log(username);
    // Get the form data from the event object
    const response = await api.post('/api/v1/login', { username:"hee", password:"brightHee"})
    .then(function(response){

        let res = response.data;
        
        console.log(res.message);
        if (res.success === true) {
            // do something from MOO

        } else {
            // do something from MOO
        }
        
        let token = response.data.data.token;
        localStorage.setItem('jwtToken', token); 

    }).catch(function(error){
        // if internal error occurs, MOO will return error message
        console.error(error);
    });
};

export default loginHandler;