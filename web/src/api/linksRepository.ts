import repository, { Collection } from "./repository";
import { Link, LinkOriginal } from "../model/link";

const resource = "/links";

export type LinkPayload = {
  url: string;
  description: string;
  tags: string[];
};

const linksRepository = {
  async getAllLinks(): Promise<Link[]> {
    const res = await repository
      .get<Collection<Link>>(`${resource}`)
      .catch(e => {
        console.log(e);
        return null;
      });
    if (res == null || !res.data.items) {
      return [];
    }
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
  },
  async getLinkOriginal(id: number) {
    const res = await repository.get<LinkOriginal>(
      `${resource}/${id}/original`
    );
    return res.data;
  }
};

export default linksRepository;
