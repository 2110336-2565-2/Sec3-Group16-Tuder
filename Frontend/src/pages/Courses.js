import CourseList from "../components/CourseList";
import CourseSearchForm from "../components/FormCourseSearch";
import styled from "styled-components";

export default function Courses(){



    return (
        <CourseListPage>
            <h1>Courses</h1>
            <CourseSearchForm />
            <CourseList>
            </CourseList>
        </CourseListPage>
    )
}

const CourseListPage = styled.div`
    display: flex;
    flex-direction: column;
    gap: 20px;
    justify-content: center;
`
