import { defineConfig } from "astro/config";
import svelte from "@astrojs/svelte";
import partytown from "@astrojs/partytown";
import compress from "astro-compress";

// https://astro.build/config
export default defineConfig({
  vite: {
    ssr: {
      external: ["svgo"],
    },
  },
  integrations: [
    svelte(),
    partytown({
      config: {
        forward: ["dataLayer.push"],
      },
    }),
    compress(),
  ],
});
