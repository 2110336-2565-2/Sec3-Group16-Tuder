import React, { useEffect, useState } from "react";
import { useNavigate, useParams, useLocation } from "react-router-dom";
import styled from "styled-components";
import StudentInfo from "../components/profile/StudentInfo.js";
import TutorInfo from "../components/profile/TutorInfo.js";
import Error from "../components/global/Error.js";
import Footer from "../components/global/Footer.js";
// import { dummyStudent, dummyTutor } from "../datas/Profile.role.js";
import {
  getStudentByUsername,
  getTutorByUsername,
} from "../handlers/profile/getUserHandler.js";
import useRole from "../hooks/useRole.js";
import useUsername from "../hooks/useUsername.js";
import { IsUser } from "../components/IsAuth.js";

// icons
import { EditOutlined } from "@ant-design/icons";

export default function Profile() {
  const [error, setError] = useState(false);
  function capitalizeFirstLetter(string) {
    return string.charAt(0).toUpperCase() + string.slice(1);
  }
  // Handle owner and others' profile
  const othersUsername = useParams().username;
  // if path is /profile/:username, then othersUsername is not undefined
  // indicating that the user is viewing other's profile
  const isOwner = othersUsername === undefined;
  let othersRole = useLocation().pathname.split("/")[1];
  // remove last character 's' from role
  othersRole = othersRole.substring(0, othersRole.length - 1);
  // ------------------------------------

  const [user, setUser] = useState({});
  const [myRole, handleMyRole] = useRole();
  const role = isOwner ? myRole : othersRole;
  const [myUsername, handleMyUsername] = useUsername();
  const username = isOwner ? myUsername : othersUsername;

  useEffect(() => {
    if (role === "student") {
      getStudentByUsername(username)
        .then((res) => {
          setUser(res.data.data);
        })
        .catch((err) => {
          setError(true);
        });
    } else if (role === "tutor") {
      getTutorByUsername(username)
        .then((res) => {
          setUser(res.data.data);
        })
        .catch((err) => {
          setError(true);
        });
    }
  }, [role]);
  const navigate = useNavigate();
  // ------------------------------------

  if (error) {
    return (
      <>
        <Error error_msg="Sorry, this user does not exist." />
        <Footer />
      </>
    );
  } else {
    return (
      <Container>
        {isOwner ? <IsUser /> : null}
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
              {isOwner ? (
                <EditIcon onClick={() => navigate("edit-profile")} />
              ) : null}
            </TitleWrapper>
            {role === "student" ? (
              <StudentInfo
                user={user}
                isOwner={isOwner}
                othersUsername={othersUsername}
              />
            ) : (
              <TutorInfo
                user={user}
                isOwner={isOwner}
                othersUsername={othersUsername}
              />
            )}
          </MiddleSection>
        </Wrapper>
        <Footer />
      </Container>
    );
  }
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
