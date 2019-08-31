import repository, { Collection } from "./repository";
import { Tag, Link } from "../model/link";

const resource = "/tags";

const tagsRepository = {
  async getAllTags(): Promise<Tag[]> {
    const res = await repository
      .get<Collection<Tag>>(`${resource}`)
      .catch(e => null);
    if (res === null || !res.data.items) {
      return [];
    }
    return res.data.items;
  },
  async getLinks(id: number): Promise<Link[]> {
    const res = await repository
      .get<Collection<Link>>(`${resource}/${id}/links`)
      .catch(e => null);
    if (res == null || !res.data.items) {
      return [];
    }
    return res.data.items;
  }
};

export default tagsRepository;
