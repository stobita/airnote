import React from "react";
import { Link } from "../model/link";
import { LinkItem } from "./LinkItem";
import styled from "styled-components";

interface Props {
  items: Link[];
}

export const LinkIndex = (props: Props) => {
  return (
    <Wrapper>
      {props.items.map(link => (
        <Column>
          <LinkItem key={link.id} item={link} />
        </Column>
      ))}
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
  flex-wrap: wrap;
  padding: 16px;
`;

const Column = styled.div`
  flex: 0 1 25%;
  box-sizing: border-box;
`;
