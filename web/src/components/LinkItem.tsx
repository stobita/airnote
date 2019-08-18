import React from "react";
import { Link } from "../model/link";
import { ReactComponent as DefaultImage } from "../assets/default.svg";
import styled from "styled-components";
import colors from "../colors";

interface Props {
  item: Link;
}

export const LinkItem = (props: Props) => {
  return (
    <Wrapper>
      <a href={props.item.url} target="_blank" rel="noopener noreferrer">
        <DefaultImage />
        <Bottom>
          <p>{props.item.description}</p>
        </Bottom>
      </a>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  flex: 0 1 25%;
  border: 1px solid ${colors.borderGray};
  border-radius: 4px;
  box-sizing: border-box;
  margin: 8px;
`;

const Bottom = styled.div`
  padding: 8px;
`;
