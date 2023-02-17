import { Outlet, Link } from 'react-router-dom';
import navbarContent from '../datas/Navbar.role.js';
import styled from 'styled-components';
import {signOutAction} from '../handlers/signOutHandler';
import {useNavigate} from 'react-router-dom';
import { Fragment } from 'react';
import getRole from '../utils/jwtGet';



export default function Navbar(props){
    // choose Navbar contents array from role 
    
    const role = getRole()
    const navigate = useNavigate();

    
    function signOutHandler(e) {
        e.preventDefault();
        signOutAction(navigate);
    }

    let navbarRole = null
    if(role === 'guest'){
        navbarRole = navbarContent[0]
    }else if(role === 'tutor'){
        navbarRole = navbarContent[1]
    }else if(role === 'student'){
        navbarRole = navbarContent[2]
    }else{
        navbarRole = navbarContent[3]
    }
    // raw data contents JSON
    const contents = navbarRole.content;
    // change to component for use in JSX  --> Generate NavItem for each content
    const contentElement = contents.map((content, index) => {
        if(content === 'Home'){
            return <NavbarItem><TuderLinkNav to='/' key={index}>{content}</TuderLinkNav></NavbarItem>
        }else if(content === 'Sign Up'){
            return <NavbarItem><TuderButton type='tudor-button' to="/SignUp" key={index}>{content}</TuderButton></NavbarItem>
        }else if(content === 'Sign In'){
            return <NavbarItem><TuderLinkNav to='/SignIn' key={index}>{content}</TuderLinkNav></NavbarItem>
        }else if(content === 'Sign Out'){
            return <NavbarItem><TuderButton type='red-button' onClick={signOutHandler} key={index}>{content}</TuderButton></NavbarItem>
        }else{
            let urlLink = "/" + content
            return <NavbarItem><TuderLinkNav to= {urlLink} key={index}>{content}</TuderLinkNav></NavbarItem>
        }
    });

    return (
        <Fragment>
            <NavbarSection>
                <NavbarHeader>Tuder</NavbarHeader>
                <NavbarItems>
                    {contentElement}
                </NavbarItems>
            </NavbarSection>   
            <Outlet />    
        </Fragment>
    )
}






// styled-components for Navbar components
const NavbarSection = styled.nav`
    height: 50px;
    display: flex;
    padding: 10px 30px;
    box-shadow: 0px 2.98px 7.45px rgba(0, 0, 0, 0.1);
    justify-content: space-between;
    align-items: center;
`;

const NavbarHeader = styled.div`
    color:rgb(227, 105, 12); 
    font-size: 30px;
    font-weight: 700  ;
`;

const NavbarItems = styled.ul`
    list-style: none;
    display: flex;
`;

const NavbarItem = styled.li`
    padding-left: 20px;
    padding-right: 20px;
    font-weight: 350;
    font-size: 15px;
`;

const TuderButton = styled(Link)`
    text-decoration: none;
    background-color: ${(props) => {
        if(props.type === 'red-button'){
            return 'rgb(255, 0, 0)'
        }else if(props.type === 'tudor-button'){
            return '#EB7B42'
        }
        return '#EB7B42'
    }};
    padding: 8px 13px;
    border: 0px solid;
    border-radius: 6px;
    color: white;
    &:hover{
        background-color: ${(props) => {
        if(props.type === 'red-button'){
            return 'rgb(255, 68, 68)'
        }else if(props.type === 'tudor-button'){
            return 'rgb(240, 123, 36)'
        }
        return 'rgb(240, 123, 36)'
    }};
    }
`;

const TuderLinkNav = styled(Link)`
    text-decoration: none;
    color: black;
    &:hover{
        color: #EB7B42;
    }
`;