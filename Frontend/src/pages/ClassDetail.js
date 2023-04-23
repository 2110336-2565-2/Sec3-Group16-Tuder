import React, { useState } from "react";
import styled from "styled-components";
import Footer from "../components/global/Footer.js";
import { useNavigate, useParams } from "react-router-dom";
import { useQuery } from "react-query";
import { IsUser } from "../components/IsAuth";
import ClassDetails from "../components/ClassDetail";
import { fetchClassDetailById } from "../handlers/classesHandler";
import { toast } from "react-hot-toast";
export default function ClassDetail() {
  const { id } = useParams();
  const [data, setData] = useState({});
  const navigate = useNavigate();

  const { isLoading, error } = useQuery(
    "classDetail",
    () => {
      fetchClassDetailById(id)
        .then((res) => {
          if (res.data.success) {
            if (
              res.data.data === null ||
              res.data.data.status === "cancelled" ||
              res.data.data.status === "cancelling"
            ) {
              navigate("/classes");
            }
            setData(res.data.data);
          }
        })
        .catch((err) => {
          console.log(err);
          toast.error("Something went wrong");
        });
    },
    {
      refetchOnWindowFocus: false,
    }
  );

  if (isLoading) {
    return <div></div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  if (data === {}) {
    return <div>Loading</div>;
  }

  if (data === null) {
    return <div>Error</div>;
  }

  return (
    <IsUser>
      <Container>
        <ClassDetails classDetail={data} />
      </Container>
      <Footer />
    </IsUser>
  );
}

const Container = styled.div`
  margin-top: 3rem;
  margin-bottom: 3rem;
  padding-left: 7%;
  padding-right: 7%;
`;
