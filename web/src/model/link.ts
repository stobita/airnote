export type Link = {
  id: number;
  url: string;
  title?: string;
  description: string;
  tags: Tag[];
};

export type LinkOriginal = {
  title: string;
};

type Tag = {
  id: number;
  text: string;
};
