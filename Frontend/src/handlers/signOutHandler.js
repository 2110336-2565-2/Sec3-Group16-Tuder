

export  function signOutAction(navigate) {
    localStorage.removeItem('jwtToken');
    navigate('/')
    
}
