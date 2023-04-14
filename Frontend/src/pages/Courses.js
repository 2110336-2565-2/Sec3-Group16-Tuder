import CourseList from "../components/CourseList";
import CourseSearchForm from "../components/form/FormCourseSearch";
import styled from "styled-components";
import { Outlet } from "react-router";
import React, { useState, createContext , useContext} from "react";
const DataContext = createContext({
        
        data: [],
        setData : () => {}
    });

export const useDataContext = () => useContext(DataContext);

export default function Courses(){
    const [data, setData] = useState([]);
    
    return (
        <DataContext.Provider value={{data, setData}}>
            <CourseListPage>
                <CenterTop><h1>Courses</h1></CenterTop>
                <CourseSearchForm />
                <HorizonContainer><hr></hr></HorizonContainer>
                <CourseList>
                </CourseList>
            </CourseListPage>
            <br></br>
        </DataContext.Provider>
        
    )
}

const CenterTop = styled.div`
    margin-top: 2rem;
    margin-bottom: 0.5rem;
    text-align: center;
`

const CourseListPage = styled.div`
    display: flex;
    flex-direction: column;
    gap: 20px;
    justify-content: center;
`
const HorizonContainer = styled.div`
    padding:3px 12% 3px 12%;
`