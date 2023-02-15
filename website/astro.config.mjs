import { defineConfig } from "astro/config";
import netlify from "@astrojs/netlify/functions";
import svelte from "@astrojs/svelte";
import partytown from "@astrojs/partytown";

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
  ],
});
