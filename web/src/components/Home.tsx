import React, { useState, useEffect, useCallback } from "react";
import styled from "styled-components";
import { Link } from "../model/link";
import { Sidebar } from "./Sidebar";
import { LinkIndex } from "./LinkIndex";
import { Header } from "./Header";
import { SlideMenu } from "./SlideMenu";
import { AddLinkForm } from "./AddLinkForm";
import { repositoryFactory } from "../api/repositoryFactory";
import { LinkDetail } from "./LinkDetail";

const linkRepository = repositoryFactory.get("links");

export const Home = () => {
  const [links, setLinks] = useState<Link[]>([]);
  const [formOpen, setFormOpen] = useState(false);
  const [detailOpen, setDetailOpen] = useState(false);
  const [selectedLink, setSelectedLink] = useState<Link>();

  useEffect(() => {
    linkRepository.getAllLinks().then(links => {
      setLinks(links);
    });
  }, []);

  const showForm = useCallback(() => {
    setFormOpen(true);
  }, []);

  const closeSlide = useCallback(() => {
    setFormOpen(false);
    setDetailOpen(false);
  }, []);

  const updateLinks = useCallback(() => {
    linkRepository.getAllLinks().then(links => {
      setLinks(links);
      if (selectedLink) {
        setSelectedLink(links.find(i => i.id === selectedLink.id));
      }
    });
  }, [selectedLink]);

  const handleAfterDelete = useCallback(() => {
    linkRepository.getAllLinks().then(links => {
      setLinks(links);
      setDetailOpen(false);
    });
  }, []);

  const selectItem = useCallback((l: Link) => {
    setDetailOpen(true);
    setSelectedLink(l);
  }, []);

  const SlideMenuContent = () => {
    switch (true) {
      case formOpen:
        return <AddLinkForm afterSubmit={updateLinks} />;
      case detailOpen:
        if (selectedLink) {
          return (
            <LinkDetail
              item={selectedLink}
              afterUpdate={updateLinks}
              afterDelete={handleAfterDelete}
            />
          );
        } else {
          return <></>;
        }
      default:
        return <></>;
    }
  };

  return (
    <Wrapper>
      <Left>
        <Sidebar />
      </Left>
      <Right>
        <Header onClickAddButton={showForm} />
        <LinkIndex items={links} onSelectItem={selectItem} />
      </Right>
      <SlideMenu onClose={closeSlide} open={formOpen || detailOpen}>
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
  flex: 5;
`;

export default Home;
