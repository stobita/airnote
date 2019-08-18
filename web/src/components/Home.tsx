import React, { useState, useEffect, useCallback } from "react";
import styled from "styled-components";
import { Link } from "../model/link";
import { Sidebar } from "./Sidebar";
import { LinkIndex } from "./LinkIndex";
import { Header } from "./Header";
import { SlideMenu } from "./SlideMenu";
import { AddLinkForm } from "./AddLinkForm";
import { repositoryFactory } from "../api/repositoryFactory";

const linkRepository = repositoryFactory.get("links");

export const Home = () => {
  const [links, setLinks] = useState<Link[]>([]);
  const [slideOpen, setSlideOpen] = useState(false);

  useEffect(() => {
    linkRepository.getAllLinks().then(links => setLinks(links));
  }, []);

  const showForm = useCallback(() => {
    setSlideOpen(true);
  }, []);

  const hideForm = useCallback(() => {
    setSlideOpen(false);
  }, []);

  const updateLinks = useCallback(() => {
    linkRepository.getAllLinks().then(links => setLinks(links));
  }, []);

  return (
    <Wrapper>
      <Left>
        <Sidebar />
      </Left>
      <Right>
        <Header onClickAddButton={showForm} />
        <LinkIndex items={links} />
      </Right>
      <SlideMenu onClose={hideForm} open={slideOpen}>
        <AddLinkForm afterPost={updateLinks} />
      </SlideMenu>
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
