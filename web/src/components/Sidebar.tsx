import React from "react";
import styled from "styled-components";

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
  background: black;
  color: white;
  flex-direction: column;
  min-height: 100vh;
`;
