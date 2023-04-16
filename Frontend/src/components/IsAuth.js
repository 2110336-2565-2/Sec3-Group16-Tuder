import { useDataContext } from "../App";
import {verify} from "../utils/jwtGet";

export function IsTutor({children}) {
    const [role, handleRole] = useDataContext();
    if(!verify()){
        window.location.href = "/sign-in";
        return
    }

    if( role.role === "tutor"){
        return children;
    }else{
        window.location.href = "/sign-in";
        return
    }
}

export function IsStudent({children}) {
    const [role, handleRole] = useDataContext();

    if(!verify()){
        window.location.href = "/sign-in";
        return
    }

    if( role.role === "student"){
        return children;
    }else{
        window.location.href = "/sign-in";
        return
    }
}

export function IsAdmin({children}) {
    const [role, handleRole] = useDataContext();

    if(!verify()){
        window.location.href = "/sign-in";
        return
    }
    if(role.role === "admin"){
        return children;
    }else{
        window.location.href = "/sign-in";
        return
    }
}

export function IsGuest({children}) {
    const [role, handleRole] = useDataContext();
    

    if( role.role === "guest"){
        return children;
    }else{

        window.location.href = "/";
        return
    }
}

export function IsUser({children}) {
    const [role, handleRole] = useDataContext();

    if(!verify()){
        window.location.href = "/sign-in";
        return
    }

    if( (role.role === "student" || role.role === "tutor")){
        return children;
    }else{
        window.location.href = "/sign-in";
      
        return
    }
}

export function IsEnroll({children}) {
    const [role, handleRole] = useDataContext();

    if(!verify()){
        alert("You need to login to enroll this course")
        window.location.href = "/sign-in";
        return
    }

    if( role.role === "student"){
        return children;
    }else{
        window.location.href = "/";
        return
    }
}

