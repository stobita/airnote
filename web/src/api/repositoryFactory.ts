import linksRepository from "./linksRepository";

const repositories = {
  links: linksRepository
};

type RepositoryName = keyof (typeof repositories);

export const repositoryFactory = {
  get: (name: RepositoryName) => repositories[name]
};
