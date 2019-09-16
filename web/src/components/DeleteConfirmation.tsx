import React from "react";
import styled from "styled-components";
import colors from "../colors";
import { Button } from "./Button";
import { ButtonPair } from "./ButtonPair";

interface Props {
  onSubmit: () => void;
  onCancel: () => void;
}

export const DeleteConfirmation = (props: Props) => {
  return (
    <Wrapper>
      <Message>Are you sure you want to delete it?</Message>
      <ButtonPair
        left={
          <Button danger type="button" onClick={props.onSubmit}>
            Delete
          </Button>
        }
        right={<Button onClick={props.onCancel}>Cancel</Button>}
      />
    </Wrapper>
  );
};

const Wrapper = styled.div`
  border: 1px solid ${props => props.theme.border};
  border-radius: 4px;
  padding: 16px;
  margin: 16px 0;
`;

const Message = styled.p`
  color: ${colors.danger};
`;
