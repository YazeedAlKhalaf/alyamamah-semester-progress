# Al Yamamah Semester Progress

This project's purpose is to show how many days have passed from the current semester in Al Yamamah University.

#### It consists of two parts:

- API: uses go lang
- Client Website: made with Astro

## Getting Started

1. to generate access tokens, follow the insturctions here: https://developer.twitter.com/en/docs/authentication/oauth-1-0a/obtaining-user-access-tokens

2. put the accesss tokens you got and the consumer keys in the `app.env` by copying the `app.env.example` file.

> provide the environment variables in app.env.example to the docker container when you run it

## Contributing

You can contribute in multiple ways:

- Reporting bugs in the issues tab.
- Opening pull requets fixing bugs.
- Opening pull requests adding new semesters the `calendars/` folder.
  - The academic calendar we depend on is found here: https://yu.edu.sa/resources/academic-calendar/
  - We save a copy of the latest calendar as PDF in the repository under `resources/`
