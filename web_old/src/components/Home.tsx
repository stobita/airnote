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
        <Top>
          <Header
            onClickAddButton={showForm}
            onSubmitWordSearch={handleOnWordSearchSubmit}
          />
        </Top>
        <Bottom>
          <LinkIndex />
        </Bottom>
      </Right>
      <SlideMenu>
        <SlideMenuContent />
      </SlideMenu>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
  background: ${props => props.theme.bg};
  height: 100%;
  min-height: 100vh;
`;

const Left = styled.div`
  flex: 1;
  position: fixed;
  min-width: 240px;
`;

const Right = styled.div`
  flex: 5;
  margin-left: 240px;
`;

const Top = styled.div`
  height: 48px;
`;

const Bottom = styled.div`
  margin-top: 16px;
  heih
`;

export default Home;
