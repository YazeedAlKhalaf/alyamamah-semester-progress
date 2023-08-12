import { defineCollection, z } from "astro:content";

const calendars = defineCollection({
  type: "data",
  schema: z.object({
    title: z.string(),
    events: z.array(
      z.object({
        name: z.string(),
        day: z.string(),
        start_date: z.string(),
        end_date: z.string().nullable(),
        week: z.string(),
      })
    ),
  }),
});

export const collections = { calendars };
