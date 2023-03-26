import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import styled from "styled-components";
import StudentInfo from "../components/profile/StudentInfo.js";
import TutorInfo from "../components/profile/TutorInfo.js";
import Footer from "../components/global/Footer.js";
// import { dummyStudent, dummyTutor } from "../datas/Profile.role.js";
import { getStudentByUsername, getTutorByUsername } from "../handlers/profile/getUserHandler.js";
import useRole from "../hooks/useRole.js";
import useUsername from "../hooks/useUsername.js";
import { IsUser } from "../components/IsAuth.js";

// icons
import { EditOutlined } from "@ant-design/icons";

export default function Profile() {
  function capitalizeFirstLetter(string) {
    return string.charAt(0).toUpperCase() + string.slice(1);
  }
  // CHANGE THIS TO GET USER FROM BACKEND
  const [user, setUser] = useState({});
  const [role, handleRole] = useRole();
  const [username, handleUsername] = useUsername();
  useEffect(() => {
    if (role === "student") {
      getStudentByUsername(username).then((res) => {
        setUser(res.data.data);
        console.log("res.data.data: ", res.data.data);
      });
    } else if (role === "tutor") {
      getTutorByUsername(username).then((res) => {
        setUser(res.data.data);
        console.log("res.data.data: ", res.data.data);
      });
    }
  }, [role]);
  // ------------------------------------

  const navigate = useNavigate();

  return (
    <Container>
      <IsUser>
      <TopSection>
        <Role>{capitalizeFirstLetter(role)}</Role>
        <ProfileImageWrapper>
          <ProfileImage src={user.profile_picture_URL} />
        </ProfileImageWrapper>
        <Name>{user.firstname + " " + user.lastname}</Name>
        <Email>{user.email}</Email>
      </TopSection>
      <Wrapper>
        <MiddleSection>
          <TitleWrapper>
            <Title>Information</Title>
            <EditIcon onClick={()=>navigate("/edit-profile")} />
          </TitleWrapper>
          {role === "student" ? (
            <StudentInfo user={user} />
          ) : (
            <TutorInfo user={user} />
          )}
        </MiddleSection>
      </Wrapper>
      <Footer />
    </IsUser>
    </Container>
    
  );
}

const Container = styled.div`
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
  background-color: #fdedeb;
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

const Wrapper = styled.div`
  border-radius: 50px 50px 0px 0px;
  background-color: white;
  margin-top: 10px;
  width: 100%;
`;

const MiddleSection = styled.div`
  display: flex;
  flex-direction: column;
  width: 100%;
  background-color: #ffffff;
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
