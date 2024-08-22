'use client'

import { createContext, Dispatch, SetStateAction } from "react";
import { Link, Tag } from "../model/link";

type DataContextProps = {
  links: Link[];
  setLinks: Dispatch<SetStateAction<Link[]>>;

  tags: Tag[];
  setTags: Dispatch<SetStateAction<Tag[]>>;
};

const defaultProps = {
  links: [],
  setLinks: () => { },

  tags: [],
  setTags: () => { }
};

export const DataContext = createContext<DataContextProps>(defaultProps);
