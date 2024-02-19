[![Netlify Status](https://api.netlify.com/api/v1/badges/c69d6bc5-ff38-49d9-bc51-87b2ae9e7368/deploy-status)](https://app.netlify.com/sites/alyamamah-semester-progress/deploys)

> [!IMPORTANT]
> After Twitter, currently known as X, stopped free access to their API, I decided to end the bot part, the code is still in the history if you want it, but it won't come back unless the price becomes free or someone sponsors the project.

# Al Yamamah Semester Progress

- This project shows the calendar for each semester in the university since I joined.
- I do use `pnpm`.

> [!NOTE]
> I didn't change the name since I want to keep its legacy.

## Contributing

You can contribute in multiple ways:

- Reporting bugs in the issues tab.
- Opening pull requets fixing bugs.
- Opening pull requests adding new semesters the `src/content/calendars/` folder.
  - The academic calendar we depend on is found here: https://yu.edu.sa/resources/academic-calendar/
  - We save a copy of the latest calendar as PDF in the repository under `resources/`
  - Then usually I use ChatGPT with GPT-4 to get the JSON done easily, first it started was with schedule 5, before it was manually.
