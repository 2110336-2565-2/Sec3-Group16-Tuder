import React ,{useEffect} from 'react';
import toast, { useToasterStore } from 'react-hot-toast';
import styled from 'styled-components';
import Button from './Button';


export const confirm = (message, handleConfirm) => {
    toast(
        (t) => (
            <Container>
                <Text>{message}</Text>
                <ButtonSection>
                    <Button variance="submit" onClick={()=>{
                        handleConfirm()
                        // dismiss all toasts
                        toast.dismiss()}
                        }>
                        Confirm
                    </Button>
                    <Button variance="cancel" onClick={() => {
                        toast.dismiss(t.id)}
                        }>
                        Cancel
                    </Button>
                </ButtonSection>
            </Container>
        ),
        {
            duration : 5000,
        }
    );
}

const Container = styled.div`
    display: flex;
    flex-direction: column;
    padding: 10px;
    justify-content: space-between;
    align-items: center;
    gap: 20px;
`;

const Text = styled.span`
    text-align: center;
`;

const ButtonSection = styled.div`
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 20px;
`;