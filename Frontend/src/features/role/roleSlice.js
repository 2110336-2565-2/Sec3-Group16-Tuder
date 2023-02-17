import {createSlice} from '@reduxjs/toolkit';

const initialState = {
    role : "guest"
}

export  const roleSlice = createSlice({
    name : 'role',
    initialState,
    reducers:{
        setRole: (state, action) => {
            return { ...state, role: action.payload }
          }
    }
})

export const {setRole} = roleSlice.actions;

export const {roleReducer} = roleSlice.reducer;
