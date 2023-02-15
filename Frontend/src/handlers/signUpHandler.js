import api from './apiHandler';

export async function signUp(signUpData){
    return  await api.post('/api/v1/signUp', signUpData)
}


async function signUpHandler(event) {

    
    // Prevent the default action of the event
    event.preventDefault();
    
    // Get the form data from the event object
    let firstname = event.target.firstname.value;
    let lastname = event.target.lastname.value;
    let username = event.target.username.value;
    let email = event.target.email.value;
    let password = event.target.password.value;
    let confirmpassword = event.target.confirmpassword.value;
    let address = event.target.address.value;
    let contactnumber = event.target.contactnumber.value;
    let gender = event.target.gender.value;
    let birthdate = (new Date(event.target.birthdate.value)).toISOString();
    let school = event.target.school.value;

    // Input validation
    //************************ */
    //

    
    // Post to api
    const response = await api.post('/api/v1/register', { username, password, email ,confirmpassword, firstname,lastname, address, contactnumber, birthdate, gender })
    .then(function(response){
        
        let res = response.data;
        
        // do some response handling
        console.log(res);
        


    }).catch(function(error){
        // if internal error occurs, MOO will return error message
        if (error.response) {
            let res = error.response.data;
            console.log(res.message);
            console.error(res.error)
        }
    });
};

export default signUpHandler;
