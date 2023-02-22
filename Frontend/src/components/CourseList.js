import styled from 'styled-components';
import Course from '../components/Course';
import {useQuery} from 'react-query';
import { fetchCourseHandler } from '../handlers/fetchCourseHandler';




export default function CourseList(){
    const {data, isLoading, error} = useQuery(
        'courses',
        fetchCourseHandler
        );


    if (isLoading) {
        return <div>Loading...</div>
    }

    if (error) {
        return <div>Error: {error.message}</div>
    }

    return (
        <>

            <CourseListPage>
                { data.data.map(item => (  
                    <CourseListcontent key={item.courseid}>
                        <Course coursename={item.coursename} content={item.content} tutor={item.tutorname}/>       
                    </CourseListcontent>
                ))}
            </CourseListPage>
        </>
    )
}

const CourseListPage = styled.div`
    display: flex;
    flex-direction: row;
    gap: 20px;
    justify-content: center;
`

const CourseListcontent = styled.div`
    display: flex;
    border: 1px solid black;
    border-radius: 16px;
    padding: 10px;
`