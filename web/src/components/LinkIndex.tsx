import React, { useContext } from "react";
import { LinkItem } from "./LinkItem";
import styled, { css } from "styled-components";
import { DataContext } from "../context/dataContext";

export const LinkIndex = () => {
  const { links } = useContext(DataContext);

  return (
    <Wrapper>
      {links.map(link => (
        <Column key={link.id}>
          <LinkItem item={link} />
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
  display: flex;
  flex: 0 1 100%;
  box-sizing: border-box;
`;
