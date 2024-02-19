import { defineConfig } from "astro/config";
import svelte from "@astrojs/svelte";
import partytown from "@astrojs/partytown";
import compress from "astro-compress";
import netlify from "@astrojs/netlify";

// https://astro.build/config
export default defineConfig({
  output: "server",
  adapter: netlify(),
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
