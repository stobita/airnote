import React, { useState, useEffect, useCallback } from "react";
import styled from "styled-components";
import { Link, Tag } from "../model/link";
import { Sidebar } from "./Sidebar";
import { LinkIndex } from "./LinkIndex";
import { Header } from "./Header";
import { SlideMenu } from "./SlideMenu";
import { AddLinkForm } from "./AddLinkForm";
import { LinkDetail } from "./LinkDetail";
import linksRepository from "../api/linksRepository";
import tagsRepository from "../api/tagsRepository";

export const Home = () => {
  const [links, setLinks] = useState<Link[]>([]);
  const [tags, setTags] = useState<Tag[]>([]);
  const [formOpen, setFormOpen] = useState(false);
  const [detailOpen, setDetailOpen] = useState(false);
  const [selectedLink, setSelectedLink] = useState<Link>();

  useEffect(() => {
    linksRepository.getAllLinks().then(links => {
      setLinks(links);
    });
    tagsRepository.getAllTags().then(tags => {
      setTags(tags);
    });
  }, []);

  const refreshLinks = useCallback(() => {
    linksRepository.getAllLinks().then(links => {
      setLinks(links);
    });
  }, []);

  const onClickTag = useCallback(id => {
    tagsRepository.getLinks(id).then(links => {
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

  const handleAfterCreate = useCallback(async id => {
    const links = await linksRepository.getAllLinks();
    setLinks(links);
    const original = await linksRepository.getLinkOriginal(id);
    const newlinks = links.map(i =>
      i.id === id ? { ...i, title: original.title } : i
    );
    console.log(newlinks);
    setLinks(newlinks);
  }, []);

  const handleAfterUpdate = useCallback(() => {
    linksRepository.getAllLinks().then(links => {
      setLinks(links);
      if (selectedLink) {
        setSelectedLink(links.find(i => i.id === selectedLink.id));
      }
    });
  }, [selectedLink]);

  const handleAfterDelete = useCallback(() => {
    linksRepository.getAllLinks().then(links => {
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
        return <AddLinkForm afterSubmit={handleAfterCreate} tags={tags} />;
      case detailOpen:
        if (selectedLink) {
          return (
            <LinkDetail
              item={selectedLink}
              tags={tags}
              afterUpdate={handleAfterUpdate}
              afterDelete={handleAfterDelete}
              onClickTag={onClickTag}
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
        <Sidebar
          tags={tags}
          onClickTag={onClickTag}
          onClickTitle={refreshLinks}
        />
      </Left>
      <Right>
        <Header onClickAddButton={showForm} />
        <LinkIndex
          items={links}
          onSelectItem={selectItem}
          onClickTag={onClickTag}
        />
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
