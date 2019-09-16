import React, { useContext } from "react";
import styled from "styled-components";
import { Sidebar } from "./Sidebar";
import { LinkIndex } from "./LinkIndex";
import { Header } from "./Header";
import { SlideMenu } from "./SlideMenu";
import linksRepository from "../api/linksRepository";
import { DataContext } from "../context/dataContext";
import { ViewContext } from "../context/viewContext";
import { LinkForm } from "./LinkForm";

export const Home = () => {
  const { setLinks } = useContext(DataContext);
  const { setSlideOpen } = useContext(ViewContext);

  const refreshLinks = () => {
    linksRepository.getAllLinks().then(links => {
      setLinks(links);
    });
  };

  const showForm = () => {
    setSlideOpen(true);
  };

  const handleOnWordSearchSubmit = (word: string) => {
    linksRepository.searchLink(word).then(links => {
      setLinks(links);
    });
  };

  const SlideMenuContent = () => {
    return <LinkForm />;
  };

  return (
    <Wrapper>
      <Left>
        <Sidebar onClickTitle={refreshLinks} />
      </Left>
      <Right>
        <Header
          onClickAddButton={showForm}
          onSubmitWordSearch={handleOnWordSearchSubmit}
        />
        <LinkIndex />
      </Right>
      <SlideMenu>
        <SlideMenuContent />
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
  flex: 4;
`;

export default Home;
