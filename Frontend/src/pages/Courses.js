import styled from 'styled-components';
import courseItems from '../datas/course.json';
import Course from '../components/Course';

export default function Courses(){

    



    return (
        <>
            <h1>Courses</h1>
            <CourseListPage>
                { courseItems.map(item => (  
                    <CourseListcontent>
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