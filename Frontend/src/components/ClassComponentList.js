import React, { useState } from "react";
import styled from "styled-components";
import MyClass from "./MyClass";
import { IsUser } from "./IsAuth";
import { useQuery } from "react-query";
import { fetchMatchesByIdHandler } from "../handlers/classesHandler";
import { getStudentID, getTutorID } from "../utils/jwtGet";
import { useDataContext } from "../pages/ClassList";

export default function ClassComponentList() {
  const { data, setData } = useDataContext();

  const { isLoading, error } = useQuery(
    "myClass",
    () => {
      fetchMatchesByIdHandler(getStudentID() || getTutorID())
        .then((res) => {
          if (res.data.success) {
            setData(res.data.data);
          }
        })
        .catch((err) => {
          console.log(err);
        });
    },
    {
      refetchOnWindowFocus: false,
    }
  );

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  if (data === null) {
    return <div>Error</div>;
  }

  if (data === []) {
    return <div>Empty</div>;
  }

  return (
    <Container>
      <IsUser>
        <h1>My Classes</h1>

        <ClassListPage>
          {/* use map for generate each class */}
          {data.map((item) => (
            <div key={item.match_id}>
              <MyClass data={item} />
            </div>
          ))}
        </ClassListPage>
      </IsUser>
    </Container>
  );
}

const Container = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 5px;
  margin-bottom: 25px;
`;

const ClassListPage = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: 10px;
  margin-top: 10px;
`;
