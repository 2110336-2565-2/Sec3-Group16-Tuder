import React, {useState, useEffect} from "react";
import styled from "styled-components";
import useRole from "../hooks/useRole";
import useUsername from "../hooks/useUsername.js";
import { getStudentByUsername, getTutorByUsername } from "../handlers/profile/getUserHandler.js";
import { dummyStudent, dummyTutor } from "../datas/Profile.role";

import FormEditProfile from "../components/profile/FormEditProfile";
import WaveFooter from "../components/global/WaveFooter";
import { IsUser } from "../components/IsAuth";

export default function EditProfile() {
  const [user, setUser] = useState();
  const [role, handleRole] = useRole();
  const [username, handleUsername] = useUsername();
  useEffect(() => {
    if (role === "student") {
      getStudentByUsername(username).then((res) => {
        setUser({...res.data.data, "role": role});
        console.log("res.data.data: ", res.data.data);
        console.log("student",user)
      });
    } else if (role === "tutor") {
      getTutorByUsername(username).then((res) => {
        setUser({...res.data.data, "role": role});
        console.log("res.data.data: ", res.data.data);
        console.log("tutor",user)
      });
    }
  }, []);

  return (
    <Container>
      <IsUser>
      <Wrapper>
        {user?<FormEditProfile user={user} />:<h1>Loading..</h1>}
      </Wrapper>
      <WaveFooter backgroundColor="#fdedeb" />
      </IsUser>
    </Container>
  );
}

const Container = styled.div`
  display: flex;
  flex-direction: column;
`;

const Wrapper = styled.div`
  display: flex;
  min-height: 95vh;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
  padding: 50px 0px;
  background-color: #fdedeb;
`;
