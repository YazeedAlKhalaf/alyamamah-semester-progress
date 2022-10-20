# Al Yamamah Semester Progress

This project's purpose is to show how many days have passed from the current semester in Al Yamamah University.

## Getting Started

1. to generate access tokens, follow the insturctions here: https://developer.twitter.com/en/docs/authentication/oauth-1-0a/obtaining-user-access-tokens

2. put the accesss tokens you got and the consumer keys in the `app.env` by copying the `app.env.example` file.

> provide the environment variables in app.env.example to the docker container when you run it

## How this app works?

The app has a backend part and a frontend part.

#### It consists of two parts:

- API: uses go lang
- Client Website: made with Astro

When a PR is submitted and then merged to the `main` branch:

1. A trigger in Google Cloud triggers Cloud Build to build and deploy a new version to Cloud Run.
2. A trigger in Netlify triggers it to build the website and deploy it instantly.

The app uses the Twitter API to tweet everyday at "8:00 AM"(GMT+3).

Cloud Scheduler is what makes that magic happen!

## Contributing

You can contribute in multiple ways:

- Reporting bugs in the issues tab.
- Opening pull requets fixing bugs.
- Opening pull requests adding new semesters the `calendars/` folder.
  - The academic calendar we depend on is found here: https://yu.edu.sa/resources/academic-calendar/
  - We save a copy of the latest calendar as PDF in the repository under `resources/`
