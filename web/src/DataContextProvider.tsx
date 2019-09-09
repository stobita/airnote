import React, { ReactNode, useState } from "react";
import { Link, Tag } from "./model/link";
import { DataContext } from "./context/dataContext";

interface Props {
  children: ReactNode;
}

export const DataContextProvider = (props: Props) => {
  const [links, setLinks] = useState<Link[]>([]);
  const [tags, setTags] = useState<Tag[]>([]);
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
