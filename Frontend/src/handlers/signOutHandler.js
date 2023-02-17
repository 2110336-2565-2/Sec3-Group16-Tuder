import { setRole } from '../features/role/roleSlice';

export  function signOutAction(navigate) {
    return  (dispatch) => {
        localStorage.removeItem('jwtToken');
        dispatch(setRole('guest'));
        navigate('/signIn')
    }
}
