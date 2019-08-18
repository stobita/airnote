import repository, { Collection } from "./repository";
import { Link } from "../model/link";

const resource = "/links";

export type CreateLinkPayload = {
  url: string;
  description: string;
};

const linksRepository = {
  async getAllLinks(): Promise<Link[]> {
    const res = await repository.get<Collection<Link>>(`${resource}`);
    return res.data.items;
  },
  async createLink(payload: CreateLinkPayload): Promise<Link> {
    const res = await repository.post<Link>(`${resource}`, payload);
    return res.data;
  }
};

export default linksRepository;
