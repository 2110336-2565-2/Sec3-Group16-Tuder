import React, {useState, useEffect} from 'react'
import styled from 'styled-components'
import Reviews from '../components/review/Reviews.js'
import Footer from '../components/global/Footer.js'
import { IsTutor } from '../components/IsAuth.js'
import { tutorReviews } from '../datas/DummyReview'

export default function TutorReviews() {
  const [reviews, setReviews] = useState([])

  // NOTE to backend: Change this to fetch reviews from backend then it should work
  useEffect(() => {
    setReviews(tutorReviews)
  }, [])
  // ---------------------------------------------

  return (
    <Container>
      {/* <IsTutor /> */}
      <TitleWrapper>
        <Title>{reviews.tutor_firstname + " " + reviews.tutor_lastname}</Title>
      </TitleWrapper>
      {reviews.total_review > 0?
        <Reviews reviews={reviews} />:
        <p>No review yet.</p>}
      <Footer />
    </Container>
  )
}

const Container = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #EB7B42;
  width: 100%;
`;

const Title = styled.h1`
  color: white;
`;

const TitleWrapper = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 150px;
`

