import React, { ReactNode, useState, useEffect } from "react";
import { Link, Tag } from "./model/link";
import { DataContext } from "./context/dataContext";
import linksRepository from "./api/linksRepository";
import tagsRepository from "./api/tagsRepository";

interface Props {
  children: ReactNode;
}

export const DataContextProvider = (props: Props) => {
  const [links, setLinks] = useState<Link[]>([]);
  const [tags, setTags] = useState<Tag[]>([]);

  useEffect(() => {
    linksRepository.getAllLinks().then(links => {
      setLinks(links);
    });
    tagsRepository.getAllTags().then(tags => {
      setTags(tags);
    });
  }, []);

  return (
    <DataContext.Provider
      value={{
        links,
        tags,

        setLinks,
        setTags
      }}
    >
      {props.children}
    </DataContext.Provider>
  );
};
