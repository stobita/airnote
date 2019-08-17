import React from "react";
import styled from "styled-components";
import colors from "../colors";

export const Sidebar = () => {
  return (
    <Wrapper>
      <h2>menu</h2>
      <ul>
        <li>Links</li>
      </ul>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
  background: ${colors.mainBlack};
  color: ${colors.mainWhite};
  flex-direction: column;
  min-height: 100vh;
  padding: 16px;
`;
