export type Link = {
  id: number;
  url: string;
  description: string;
  tags: Tag[];
};

type Tag = {
  id: number;
  text: string;
};
