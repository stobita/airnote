import repository, { Collection } from "./repository";
import { Link } from "../model/link";

const resource = "/links";

export type LinkPayload = {
  url: string;
  description: string;
};

const linksRepository = {
  async getAllLinks(): Promise<Link[]> {
    const res = await repository.get<Collection<Link>>(`${resource}`);
    return res.data.items;
  },
  async createLink(payload: LinkPayload): Promise<Link> {
    const res = await repository.post<Link>(`${resource}`, payload);
    return res.data;
  },
  async updateLink(id: number, payload: LinkPayload): Promise<Link> {
    const res = await repository.put<Link>(`${resource}/${id}`, payload);
    return res.data;
  },
  async deleteLink(id: number) {
    await repository.delete(`${resource}/${id}`);
  }
};

export default linksRepository;
