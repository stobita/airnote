import React, { useState, useEffect } from "react";
import { repositoryFactory } from "../api/repositoryFactory";
import { Link } from "../model/link";
import { Sidebar } from "./Sidebar";
import styled from "styled-components";
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
        <ul>
          {links.map(link => (
            <li key={link.url}>{link.url}</li>
          ))}
          <li>test</li>
        </ul>
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
