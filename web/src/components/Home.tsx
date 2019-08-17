import React, { useState, useEffect } from "react";
import { repositoryFactory } from "../api/repositoryFactory";
import { Link } from "../model/link";
import { Sidebar } from "./Sidebar";
import styled from "styled-components";
import { LinkIndex } from "./LinkIndex";
const linkRepository = repositoryFactory.get("links");

export const Home = () => {
  const [links, setLinks] = useState<Link[]>([]);
  useEffect(() => {
    linkRepository.getAllLinks().then(links => setLinks(links));
  }, []);
  return (
    <Wrapper>
      <Left>
        <Sidebar />
      </Left>
      <Right>
        <LinkIndex items={links} />
      </Right>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
`;

const Left = styled.div`
  flex: 1;
`;

const Right = styled.div`
  flex: 5;
`;

export default Home;
