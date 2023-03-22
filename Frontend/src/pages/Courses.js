import CourseList from "../components/CourseList";
import CourseSearchForm from "../components/form/FormCourseSearch";
import styled from "styled-components";
import { Outlet } from "react-router";
import React, { useState, createContext , useContext} from "react";
const DataContext = createContext({
        data : {
            data: []
        },
        setData : () => {}
    });

export const useDataContext = () => useContext(DataContext);

export default function Courses(){
    const [data, setData] = useState({data:[]});
    

    return (
        <DataContext.Provider value={{data, setData}}>
            <CourseListPage>
                <h1>Courses</h1>
            
                <CourseSearchForm />
                <CourseList>
                </CourseList>
            </CourseListPage>
        </DataContext.Provider>
        
    )
}

const CourseListPage = styled.div`
    display: flex;
    flex-direction: column;
    gap: 20px;
    justify-content: center;
`
