import React from "react";
import styled from "styled-components";
import StudentInfo from "../components/profile/StudentInfo.js";
import TutorInfo from "../components/profile/TutorInfo.js";
import { dummyStudent, dummyTutor } from "../datas/Profile.role.js";

// icons
import { EditOutlined } from "@ant-design/icons";

export default function Profile() {
    function capitalizeFirstLetter(string) {
        return string.charAt(0).toUpperCase() + string.slice(1);
    }
    const user = dummyTutor;

    return (
        <Container>
        <TopSection>
            <Role>{capitalizeFirstLetter(user.role)}</Role>
            <ProfileImageWrapper>
            <ProfileImage src={user.profile_picture_URL} />
            </ProfileImageWrapper>
            <Name>{user.first_name + " " + user.last_name}</Name>
            <Email>{user.email}</Email>
        </TopSection>
        <MiddleSection>
            <TitleWrapper>
                <Title>Information</Title>
                <EditIcon />
            </TitleWrapper>
            {user.role === "student" ? (
            <StudentInfo user={user} />
            ) : (
            <TutorInfo user={user} />)}
        </MiddleSection>
        </Container>
    );
}

const Container = styled.div`
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    background-color: #FDEDEB;
`;

const TopSection = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
`;

const Role = styled.span`
    font-size: 36px;
    font-weight: bold;
    margin-top: 50px;
`;

const ProfileImageWrapper = styled.div`
    width: 200px;
    height: 200px;
    border-radius: 50%;
    overflow: hidden;
    margin-top: 18px;
`;

const ProfileImage = styled.img`
    width: 100%;
    height: 100%;
    object-fit: cover;
`;

const Name = styled.span`
    font-size: 30px;
    font-weight: bold;
    margin-top: 18px;
`;

const Email = styled.span`
    font-size: 16px;
    font-weight: medium;
    color: #ababab;
    margin-top: 10px;
`;

const MiddleSection = styled.div`
    display: flex;
    flex-direction: column;
    width: 100%;
    background-color: #FFFFFF;
    border-radius: 50px 50px 0px 0px;
    padding: 50px;
    margin-top: 20px;
`;

const TitleWrapper = styled.div`
    display: flex;
    align-items: center;
    gap: 10px;
`;

const Title = styled.h1`
    font-size: 36px;
    font-weight: bold;
`;

const EditIcon = styled(EditOutlined)`
    font-size: 30px;
    cursor: pointer;
`;

const InfoContainter = styled.div``;

const InfoTitle = styled.h2``;

const InfoContent = styled.span``;
